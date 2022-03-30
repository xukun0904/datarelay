package global

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"jhr.com/datarelay/api/gen/dataset"
	"jhr.com/datarelay/model"
)

var (
	ServerConfig             model.ServerConfig
	NamingClient             naming_client.INamingClient
	DataSourceConnectionInfo dataset.ConnectionInfo
)
