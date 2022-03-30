package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"jhr.com/datarelay/api/gen/dataset"
	"jhr.com/datarelay/global"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {
	zap.S().Info("Start initializing configuration")
	v := viper.New()
	// 配置文件名称及类型
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	// 获取配置文件所在路径
	if datarelayConfigPath := GetEnvInfo("CONFIG_PATH"); datarelayConfigPath != "" {
		zap.S().Info("The path of the configuration file is: ", datarelayConfigPath)
		v.AddConfigPath(datarelayConfigPath)
	} else {
		// 默认在当前目录查找
		v.AddConfigPath(".")
	}
	// 加载配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic("Failed to load configuration file: ", err.Error())
	}
	// 序列化成struct
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panic("Serialization configuration file to struct failed: ", err.Error())
	}
	// 验证配置文件是否正确
	validate := validator.New()
	if err := validate.Struct(&global.ServerConfig); err != nil {
		zap.S().Panic("Configuration verification failed: ", err.Error())
	}
	zap.S().Info("Initial configuration succeeded")
}

func InitDataSourceConnectionInfo() {
	dataSourceInfo := global.ServerConfig.DataSource.Info
	global.DataSourceConnectionInfo = dataset.ConnectionInfo{
		Id:                   dataSourceInfo.Id,
		Ip:                   dataSourceInfo.Ip,
		Port:                 dataSourceInfo.Port,
		DbName:               dataSourceInfo.DbName,
		Username:             dataSourceInfo.Username,
		Password:             dataSourceInfo.Password,
		Schema:               dataSourceInfo.Schema,
		IpAndPort:            dataSourceInfo.IpAndPort,
		AuthType:             dataSourceInfo.AuthType,
		ServiceDiscoveryMode: dataSourceInfo.ServiceDiscoveryMode,
		ZookeeperNamespace:   dataSourceInfo.ZookeeperNamespace,
		SaslQop:              dataSourceInfo.SaslQop,
		Principal:            dataSourceInfo.Principal,
		UserPrincipal:        dataSourceInfo.UserPrincipal,
		ConfFile:             dataSourceInfo.ConfFiles,
		DsourceType:          dataSourceInfo.DsourceType,
		AdvanceParamConf:     dataSourceInfo.AdvanceParamConf,
		ConnectType:          dataSourceInfo.ConnectType,
	}
	zap.S().Infof("Convert to data source configuration struct successfully. Data source configuration is: {%+v}", &global.DataSourceConnectionInfo)
}
