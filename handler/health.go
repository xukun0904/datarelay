package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"jhr.com/datarelay/model"
	"jhr.com/datarelay/util"
)

func HandleHealth(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	// 查看是否有可用客户端
	_, _, err := getInstanceGrpcIpAndPort()
	if err != nil {
		return err
	}
	util.WriteResponse(rw, model.ResultMap[model.SUCCESS])
	return nil
}
