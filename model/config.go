package model

import (
	"net/http"

	"jhr.com/datarelay/api/gen/dataset"
)

type ServerConfig struct {
	Port       uint16           `mapstructure:"port" json:"port" validate:"required"`
	Test       bool             `mapstructure:"test" json:"test"`
	Nacos      NacosConfig      `mapstructure:"nacos" json:"nacos"`
	Api        ApiConfig        `mapstructure:"api" json:"api"`
	DataSource DataSourceConfig `mapstructure:"data_source" json:"data_source"`
}

type ApiConfig struct {
	Protocol            Protocol            `mapstructure:"protocol" json:"protocol" validate:"required"`
	RequestMethod       RequestMethod       `mapstructure:"request_method" json:"request_method" validate:"required"`
	RequestContentType  RequestContentType  `mapstructure:"request_content_type" json:"request_content_type"`
	RequestParams       []RequestParam      `mapstructure:"request_params" json:"request_params"`
	Path                string              `mapstructure:"path" json:"path" validate:"required"`
	ResponseContentType ResponseContentType `mapstructure:"response_content_type" json:"response_content_type" validate:"required"`
}

type NacosConfig struct {
	ContextPath string `mapstructure:"context_path" json:"context_path"`
	Port        uint64 `mapstructure:"port" json:"port"`
	NamespaceId string `mapstructure:"namespace_id" json:"namespace_id"`
	IpAddr      string `mapstructure:"ip_addr" json:"ip_addr" validate:"ip4_addr"`
}

type RequestParam struct {
	Name         string           `mapstructure:"name" json:"name" validate:"required"`
	Type         RequestParamType `mapstructure:"type" json:"type" validate:"required"`
	Required     bool             `mapstructure:"required" json:"required"`
	DefaultValue string           `mapstructure:"default_value" json:"default_value"`
	Value        string           `mapstructure:"value" json:"value"`
	Pageable     bool             `mapstructure:"pageable" json:"pageable"`
	Condition    string           `mapstructure:"condition" json:"condition"`
}

type DataSourceConfig struct {
	Service DataSourceServerConfig `mapstructure:"service" json:"service"`
	Info    DataSourceInfo         `mapstructure:"info" json:"info"`
}

type DataSourceServerConfig struct {
	Name      string   `mapstructure:"name" json:"name" validate:"required"`
	GroupName string   `mapstructure:"group_name" json:"group_name"`
	Clusters  []string `mapstructure:"clusters" json:"clusters"`
}

type DataSourceInfo struct {
	Username             string                  `mapstructure:"username" json:"username"`
	IpAndPort            string                  `mapstructure:"ip_and_port" json:"ip_and_port"`
	UserPrincipal        string                  `mapstructure:"user_principal" json:"user_principal"`
	ConfFiles            []*dataset.ConfFileInfo `mapstructure:"conf_files" json:"conf_files"`
	DsourceType          dataset.DataConnType    `mapstructure:"dsource_type" json:"dsource_type" validate:"required"`
	ConnectType          dataset.ConnectType     `mapstructure:"connect_type" json:"connect_type"`
	Id                   string                  `mapstructure:"id" json:"id" validate:"required"`
	DbName               string                  `mapstructure:"db_name" json:"db_name" validate:"required"`
	AuthType             dataset.AuthType        `mapstructure:"auth_type" json:"auth_type"`
	ZookeeperNamespace   string                  `mapstructure:"zookeeper_namespace" json:"zookeeper_namespace"`
	Principal            string                  `mapstructure:"principal" json:"principal"`
	DsourceMfrs          dataset.DataSourceMfrs  `mapstructure:"dsource_mfrs" json:"dsource_mfrs"`
	Ip                   string                  `mapstructure:"ip" json:"ip"`
	Schema               string                  `mapstructure:"schema" json:"schema"`
	ServiceDiscoveryMode string                  `mapstructure:"service_discovery_mode" json:"service_discovery_mode"`
	AdvanceParamConf     string                  `mapstructure:"advance_param_conf" json:"advance_param_conf"`
	Port                 string                  `mapstructure:"port" json:"port"`
	Password             string                  `mapstructure:"password" json:"password"`
	SaslQop              string                  `mapstructure:"sasl_qop" json:"sasl_qop"`
	Sql                  string                  `mapstructure:"sql" json:"sql" validate:"required"`
	CountSql             string                  `mapstructure:"count_sql" json:"count_sql"`
	PageAndOrderSql      string                  `mapstructure:"page_and_order_sql" json:"page_and_order_sql"`
}

type RequestContentType uint8

const (
	REQUEST_CONTENT_TYPE_NOT_SPECIFIED RequestContentType = iota
	REQUEST_APPLICATION_JSON
	REQUEST_APPLICATION_XML
	REQUEST_APPLICATION_FORM
)

var (
	RequestContentTypeName = map[RequestContentType]string{
		REQUEST_CONTENT_TYPE_NOT_SPECIFIED: "",
		REQUEST_APPLICATION_JSON:           "application/json",
		REQUEST_APPLICATION_XML:            "application/xml",
		REQUEST_APPLICATION_FORM:           "application/x-www-form-urlencoded",
	}
)

type ResponseContentType uint8

const (
	RESPONSE_CONTENT_TYPE_NOT_SPECIFIED ResponseContentType = iota
	RESPONSE_APPLICATION_JSON
	RESPONSE_APPLICATION_XML
)

var (
	ResponseContentTypeName = map[ResponseContentType]string{
		RESPONSE_CONTENT_TYPE_NOT_SPECIFIED: "application/json",
		RESPONSE_APPLICATION_JSON:           "application/json",
		RESPONSE_APPLICATION_XML:            "application/xml",
	}
)

type Protocol uint8

const (
	PROTOCOL_NOT_SPECIFIED Protocol = iota
	PROTOCOL_HTTP
	PROTOCOL_HTTPS
)

type RequestMethod uint8

const (
	REQUEST_METHOD_NOT_SPECIFIED RequestMethod = iota
	REQUEST_METHOD_GET
	REQUEST_METHOD_POST
)

var (
	RequestMethodName = map[RequestMethod]string{
		REQUEST_METHOD_NOT_SPECIFIED: http.MethodGet,
		REQUEST_METHOD_GET:           http.MethodGet,
		REQUEST_METHOD_POST:          http.MethodPost,
	}
)

type RequestParamType uint8

const (
	REQUEST_PARAM_TYPE_NOT_SPECIFIED RequestParamType = iota
	REQUEST_PARAM_TYPE_QUERY
	REQUEST_PARAM_TYPE_BODY
	REQUEST_PARAM_TYPE_PATH
	REQUEST_PARAM_TYPE_HEADER
)
