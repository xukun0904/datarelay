package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"jhr.com/datarelay/api/gen/dataset"
	"jhr.com/datarelay/api/gen/response"
	"jhr.com/datarelay/exception"
	"jhr.com/datarelay/global"
	"jhr.com/datarelay/model"
	"jhr.com/datarelay/util"
)

func HandleDataset(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	// 校验请求类型
	if err := checkRequestContentType(r); err != nil {
		return err
	}
	// 校验请求参数
	paramValues, countParamValues, whereSql, err := checkAndGetRequestParams(r, params)
	if err != nil {
		return err
	}
	// 设置数据源的sqlMap
	if err := setSqlMap(func(options *QueryOptions) {
		sqls := map[string]string{
			"count_sql": global.ServerConfig.DataSource.Info.CountSql + whereSql,
			"sql":       global.ServerConfig.DataSource.Info.Sql + whereSql + " " + global.ServerConfig.DataSource.Info.PageAndOrderSql,
		}
		options.Sqls = sqls
		options.paramValues = [][]interface{}{countParamValues, paramValues}
	}); err != nil {
		return err
	}
	// 执行查询并响应
	execQueryAndReturn(rw)
	return nil
}

func execQueryAndReturn(rw http.ResponseWriter) error {
	// 从nacos获取客户端ip和port
	grpcIp, grpcPort, err := getInstanceGrpcIpAndPort()
	if err != nil {
		return err
	}
	zap.S().Infof("[%s:%s] Execute sql: %v", grpcIp, grpcPort, global.DataSourceConnectionInfo.SqlMap)
	// 使用grpc执行查询数据源方法
	cc, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcIp, grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	tsc := dataset.NewDataSetServiceClient(cc)
	resp, err := tsc.QueryForList(context.Background(), &global.DataSourceConnectionInfo)
	if err != nil {
		return err
	}
	// 处理查询结果
	result, err := convertToResponse(resp)
	if err != nil {
		return err
	}
	if err = util.WriteResponse(rw, result); err != nil {
		return err
	}
	return nil
}

func convertToResponse(resp *response.ResponseResult) (model.Result, error) {
	// 处理查询结果， 反序列化成struct
	ds := dataset.DataSet{}
	if resp.Data != nil {
		if err := resp.Data.UnmarshalTo(&ds); err != nil {
			return model.Result{}, err
		}
	}
	data := rowDatasToStringMaps(&ds)
	// 转换成响应值
	result := model.Result{
		Success: resp.Success,
		Code:    resp.Code,
		Message: resp.Message,
		Data:    &data,
	}
	return result, nil
}

func rowDatasToStringMaps(ds *dataset.DataSet) model.PageResponse {
	var data []model.StringMap
	for _, rd := range ds.Dataset {
		var sm = make(map[string]string)
		for k, v := range rd.RowData {
			sm[k] = v
		}
		data = append(data, model.StringMap(sm))
	}
	return model.PageResponse{
		Rows:  &data,
		Total: ds.Total,
	}
}

func getInstanceGrpcIpAndPort() (string, string, error) {
	instance, err := global.NamingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: global.ServerConfig.DataSource.Service.Name,
		GroupName:   global.ServerConfig.DataSource.Service.GroupName,
		Clusters:    global.ServerConfig.DataSource.Service.Clusters,
	})
	if err != nil {
		return "", "", err
	}
	gRpcPort := instance.Metadata["gRPC_port"]
	if util.IsBlank(gRpcPort) {
		return "", "", exception.Cast(model.ResultMap[model.NOT_FOUND_HEALTHY_INSTANCE])
	}
	return instance.Ip, gRpcPort, nil
}

type QueryOption func(*QueryOptions)

type QueryOptions struct {
	Sqls        map[string]string
	paramValues [][]interface{}
}

func loadQueryOption(option ...QueryOption) *QueryOptions {
	options := new(QueryOptions)
	for _, e := range option {
		e(options)
	}
	return options
}

func setSqlMap(option ...QueryOption) error {
	op := loadQueryOption(option...)
	sqlMap := make(map[string][]interface{})
	i := 0
	for k, v := range op.Sqls {
		m := map[string]string{
			"type": k,
			k:      v,
		}
		b, err := json.Marshal(&m)
		if err != nil {
			return err
		}
		sqlMap[string(b)] = op.paramValues[i]
		i++
	}
	sqlJsonBytes, err := json.Marshal(&sqlMap)
	if err != nil {
		return err
	}
	global.DataSourceConnectionInfo.SqlMap = string(sqlJsonBytes)
	return nil
}

func checkRequestContentType(r *http.Request) error {
	apiConfig := global.ServerConfig.Api
	// 校验请求Content-Type是否正确
	var requestContentTypeConfig = model.RequestContentTypeName[apiConfig.RequestContentType]
	if util.IsNotBlank(requestContentTypeConfig) {
		requestContentType := r.Header.Get("Content-Type")
		if requestContentTypeConfig != requestContentType {
			return exception.Cast(model.ResultMap[model.REQUEST_CONTENT_INCORRECT])
		}
	}
	return nil
}

func checkAndGetRequestParams(r *http.Request, params httprouter.Params) ([]interface{}, []interface{}, string, error) {
	// 获取body内容
	m, err := getRequestBody(r)
	if err != nil {
		return nil, nil, "", err
	}
	apiConfig := global.ServerConfig.Api
	// 校验参数是否正确, 并封装成数组
	requestParams := apiConfig.RequestParams
	var paramValues []interface{}
	var countParamValues []interface{}
	var conditions []string
	for _, rp := range requestParams {
		var paramValue string
		switch {
		case rp.Type == model.REQUEST_PARAM_TYPE_QUERY:
			paramValue = r.FormValue(rp.Name)
		case rp.Type == model.REQUEST_PARAM_TYPE_BODY:
			if apiConfig.RequestContentType == model.REQUEST_APPLICATION_FORM {
				paramValue = r.PostFormValue(rp.Name)
			} else {
				paramValue = m[rp.Name]
			}
		case rp.Type == model.REQUEST_PARAM_TYPE_PATH:
			paramValue = params.ByName(rp.Name)
		case rp.Type == model.REQUEST_PARAM_TYPE_HEADER:
			if len(r.Header[rp.Name]) > 0 {
				paramValue = r.Header[rp.Name][0]
			}
		default:
			return nil, nil, "", exception.Cast(model.ResultMap[model.REQUEST_PARAMETER_TYPE_INCORRECT])
		}
		if util.IsBlank(paramValue) && util.IsNotBlank(rp.DefaultValue) {
			paramValue = rp.DefaultValue
		}
		if rp.Required && util.IsBlank(paramValue) {
			return nil, nil, "", exception.Cast(model.ResultMap[model.REQUEST_PARAMETER_BLANK])
		}
		if util.IsNotBlank(paramValue) {
			arg, condition := getArgumentAndCondition(paramValue, rp)
			paramValues = append(paramValues, arg)
			// 判断是否是分页参数
			if !rp.Pageable {
				countParamValues = append(countParamValues, arg)
				conditions = append(conditions, condition)
			}
		}
	}
	return paramValues, countParamValues, spliceWhereSql(conditions), nil
}

func getArgumentAndCondition(paramValue string, rp model.RequestParam) (interface{}, string) {
	var arg interface{}
	condition := rp.Condition
	if strings.Contains(condition, "IN") || strings.Contains(condition, "NOT IN") {
		args := strings.Split(paramValue, ",")
		placeholder := strings.Repeat("?,", len(args))
		placeholder = strings.TrimRight(placeholder, ",")
		condition = strings.Replace(condition, "?", placeholder, 1)
		arg = args
	} else if strings.Contains(condition, "LIKE") || strings.Contains(condition, "NOT LIKE") {
		arg = "%" + paramValue + "%"
	} else {
		arg = paramValue
	}
	return arg, condition
}

func spliceWhereSql(conditions []string) string {
	// 拼接Where条件
	var whereSql string
	if len(conditions) > 0 {
		whereSql = " where "
		for i, condition := range conditions {
			if i != 0 {
				whereSql += " and "
			}
			whereSql += condition
		}
	}
	return whereSql
}

func getRequestBody(r *http.Request) (map[string]string, error) {
	apiConfig := global.ServerConfig.Api
	var m map[string]string
	if apiConfig.RequestContentType == model.REQUEST_APPLICATION_JSON || apiConfig.RequestContentType == model.REQUEST_APPLICATION_XML {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return m, err
		}
		if len(b) == 0 {
			return m, nil
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		if apiConfig.RequestContentType == model.REQUEST_APPLICATION_JSON {
			if err := json.Unmarshal(b, &m); err != nil {
				return m, err
			}
		} else if apiConfig.RequestContentType == model.REQUEST_APPLICATION_XML {
			var sm model.StringMap
			if err := xml.Unmarshal(b, &sm); err != nil {
				return m, err
			}
			m = sm
		}
	}
	return m, nil
}
