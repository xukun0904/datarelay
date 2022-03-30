package initialize

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"jhr.com/datarelay/global"
	"jhr.com/datarelay/util"
)

func InitNamingClient() {
	nacos := global.ServerConfig.Nacos
	zap.S().Infof("Start initializing nacos naming client. Nacos config is: %+v", &nacos)
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacos.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "error",
	}
	// 创建ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacos.IpAddr,
			ContextPath: util.BlankToDefault(nacos.ContextPath, "/nacos"),
			Port:        util.ZeroToDefault(nacos.Port, 8848),
			Scheme:      "http",
		},
	}
	// 创建服务发现客户端
	var err error
	global.NamingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		zap.S().Panic("Failed to initialize nacos naming client: ", err)
	}
	zap.S().Info("Nacos naming client initialized successfully")
}
