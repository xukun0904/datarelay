package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"jhr.com/datarelay/api/gen/dataset"
	"jhr.com/datarelay/global"
)

func HandleReady(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	validationQuery := "select 1"
	if global.ServerConfig.DataSource.Info.DsourceType == dataset.DataConnType_ORACLE {
		validationQuery = "select 1 from dual"
	}
	if err := setSqlMap(func(options *QueryOptions) {
		sqls := map[string]string{
			"sql": validationQuery,
		}
		options.Sqls = sqls
		options.paramValues = make([][]interface{}, 1)
	}); err != nil {
		return err
	}
	execQueryAndReturn(rw)
	return nil
}
