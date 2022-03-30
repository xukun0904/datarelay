// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: dataset.proto

package dataset

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	response "jhr.com/datarelay/api/gen/response"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataSourceType int32

const (
	DataSourceType_DST_NOT_SPECIFIED DataSourceType = 0
	DataSourceType_RDBMS             DataSourceType = 1
	DataSourceType_BIG_DATA          DataSourceType = 2
	DataSourceType_MQ                DataSourceType = 3
	DataSourceType_FILE_SERVER       DataSourceType = 4
)

// Enum value maps for DataSourceType.
var (
	DataSourceType_name = map[int32]string{
		0: "DST_NOT_SPECIFIED",
		1: "RDBMS",
		2: "BIG_DATA",
		3: "MQ",
		4: "FILE_SERVER",
	}
	DataSourceType_value = map[string]int32{
		"DST_NOT_SPECIFIED": 0,
		"RDBMS":             1,
		"BIG_DATA":          2,
		"MQ":                3,
		"FILE_SERVER":       4,
	}
)

func (x DataSourceType) Enum() *DataSourceType {
	p := new(DataSourceType)
	*p = x
	return p
}

func (x DataSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_dataset_proto_enumTypes[0].Descriptor()
}

func (DataSourceType) Type() protoreflect.EnumType {
	return &file_dataset_proto_enumTypes[0]
}

func (x DataSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceType.Descriptor instead.
func (DataSourceType) EnumDescriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{0}
}

type ConfFileType int32

const (
	ConfFileType_CFT_NOT_SPECIFIED ConfFileType = 0
	ConfFileType_KRB5_CONF         ConfFileType = 1
	ConfFileType_USER_KEYTAB       ConfFileType = 2
	ConfFileType_CORE_SITE_XML     ConfFileType = 3
	ConfFileType_HDFS_SITE_XML     ConfFileType = 4
	ConfFileType_HBASE_SITE_XML    ConfFileType = 5
)

// Enum value maps for ConfFileType.
var (
	ConfFileType_name = map[int32]string{
		0: "CFT_NOT_SPECIFIED",
		1: "KRB5_CONF",
		2: "USER_KEYTAB",
		3: "CORE_SITE_XML",
		4: "HDFS_SITE_XML",
		5: "HBASE_SITE_XML",
	}
	ConfFileType_value = map[string]int32{
		"CFT_NOT_SPECIFIED": 0,
		"KRB5_CONF":         1,
		"USER_KEYTAB":       2,
		"CORE_SITE_XML":     3,
		"HDFS_SITE_XML":     4,
		"HBASE_SITE_XML":    5,
	}
)

func (x ConfFileType) Enum() *ConfFileType {
	p := new(ConfFileType)
	*p = x
	return p
}

func (x ConfFileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfFileType) Descriptor() protoreflect.EnumDescriptor {
	return file_dataset_proto_enumTypes[1].Descriptor()
}

func (ConfFileType) Type() protoreflect.EnumType {
	return &file_dataset_proto_enumTypes[1]
}

func (x ConfFileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfFileType.Descriptor instead.
func (ConfFileType) EnumDescriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{1}
}

type DataConnType int32

const (
	DataConnType_DCT_NOT_SPECIFIED DataConnType = 0
	DataConnType_GAUSS_DB          DataConnType = 1
	DataConnType_MYSQL             DataConnType = 2
	DataConnType_ORACLE            DataConnType = 3
	DataConnType_HIVE              DataConnType = 21
	DataConnType_HBASE             DataConnType = 22
	DataConnType_KAFKA             DataConnType = 31
	DataConnType_FTP               DataConnType = 41
	DataConnType_SFTP              DataConnType = 42
	DataConnType_HDFS              DataConnType = 43
)

// Enum value maps for DataConnType.
var (
	DataConnType_name = map[int32]string{
		0:  "DCT_NOT_SPECIFIED",
		1:  "GAUSS_DB",
		2:  "MYSQL",
		3:  "ORACLE",
		21: "HIVE",
		22: "HBASE",
		31: "KAFKA",
		41: "FTP",
		42: "SFTP",
		43: "HDFS",
	}
	DataConnType_value = map[string]int32{
		"DCT_NOT_SPECIFIED": 0,
		"GAUSS_DB":          1,
		"MYSQL":             2,
		"ORACLE":            3,
		"HIVE":              21,
		"HBASE":             22,
		"KAFKA":             31,
		"FTP":               41,
		"SFTP":              42,
		"HDFS":              43,
	}
)

func (x DataConnType) Enum() *DataConnType {
	p := new(DataConnType)
	*p = x
	return p
}

func (x DataConnType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataConnType) Descriptor() protoreflect.EnumDescriptor {
	return file_dataset_proto_enumTypes[2].Descriptor()
}

func (DataConnType) Type() protoreflect.EnumType {
	return &file_dataset_proto_enumTypes[2]
}

func (x DataConnType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataConnType.Descriptor instead.
func (DataConnType) EnumDescriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{2}
}

type AuthType int32

const (
	AuthType_AT_NOT_SPECIFIED AuthType = 0
	AuthType_NORMAL           AuthType = 1
	AuthType_KERBEROS         AuthType = 2
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "AT_NOT_SPECIFIED",
		1: "NORMAL",
		2: "KERBEROS",
	}
	AuthType_value = map[string]int32{
		"AT_NOT_SPECIFIED": 0,
		"NORMAL":           1,
		"KERBEROS":         2,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_dataset_proto_enumTypes[3].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_dataset_proto_enumTypes[3]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{3}
}

type DataSourceMfrs int32

const (
	DataSourceMfrs_DSM_NOT_SPECIFIED DataSourceMfrs = 0
	DataSourceMfrs_NL                DataSourceMfrs = 1
	DataSourceMfrs_HW                DataSourceMfrs = 7
)

// Enum value maps for DataSourceMfrs.
var (
	DataSourceMfrs_name = map[int32]string{
		0: "DSM_NOT_SPECIFIED",
		1: "NL",
		7: "HW",
	}
	DataSourceMfrs_value = map[string]int32{
		"DSM_NOT_SPECIFIED": 0,
		"NL":                1,
		"HW":                7,
	}
)

func (x DataSourceMfrs) Enum() *DataSourceMfrs {
	p := new(DataSourceMfrs)
	*p = x
	return p
}

func (x DataSourceMfrs) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceMfrs) Descriptor() protoreflect.EnumDescriptor {
	return file_dataset_proto_enumTypes[4].Descriptor()
}

func (DataSourceMfrs) Type() protoreflect.EnumType {
	return &file_dataset_proto_enumTypes[4]
}

func (x DataSourceMfrs) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceMfrs.Descriptor instead.
func (DataSourceMfrs) EnumDescriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{4}
}

type ConnectType int32

const (
	ConnectType_CT_NOT_SPECIFIED ConnectType = 0
	ConnectType_SID              ConnectType = 1
	ConnectType_SERVICE_NAME     ConnectType = 2
)

// Enum value maps for ConnectType.
var (
	ConnectType_name = map[int32]string{
		0: "CT_NOT_SPECIFIED",
		1: "SID",
		2: "SERVICE_NAME",
	}
	ConnectType_value = map[string]int32{
		"CT_NOT_SPECIFIED": 0,
		"SID":              1,
		"SERVICE_NAME":     2,
	}
)

func (x ConnectType) Enum() *ConnectType {
	p := new(ConnectType)
	*p = x
	return p
}

func (x ConnectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectType) Descriptor() protoreflect.EnumDescriptor {
	return file_dataset_proto_enumTypes[5].Descriptor()
}

func (ConnectType) Type() protoreflect.EnumType {
	return &file_dataset_proto_enumTypes[5]
}

func (x ConnectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectType.Descriptor instead.
func (ConnectType) EnumDescriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{5}
}

type ConfFileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DsourceType       DataSourceType `protobuf:"varint,1,opt,name=dsource_type,json=dsourceType,proto3,enum=dataset.DataSourceType" json:"dsource_type,omitempty"`
	ConfFilePath      string         `protobuf:"bytes,2,opt,name=conf_file_path,json=confFilePath,proto3" json:"conf_file_path,omitempty"`
	ConfFileType      ConfFileType   `protobuf:"varint,3,opt,name=conf_file_type,json=confFileType,proto3,enum=dataset.ConfFileType" json:"conf_file_type,omitempty"`
	LocalConfFilePath string         `protobuf:"bytes,4,opt,name=local_conf_file_path,json=localConfFilePath,proto3" json:"local_conf_file_path,omitempty"`
}

func (x *ConfFileInfo) Reset() {
	*x = ConfFileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfFileInfo) ProtoMessage() {}

func (x *ConfFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfFileInfo.ProtoReflect.Descriptor instead.
func (*ConfFileInfo) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *ConfFileInfo) GetDsourceType() DataSourceType {
	if x != nil {
		return x.DsourceType
	}
	return DataSourceType_DST_NOT_SPECIFIED
}

func (x *ConfFileInfo) GetConfFilePath() string {
	if x != nil {
		return x.ConfFilePath
	}
	return ""
}

func (x *ConfFileInfo) GetConfFileType() ConfFileType {
	if x != nil {
		return x.ConfFileType
	}
	return ConfFileType_CFT_NOT_SPECIFIED
}

func (x *ConfFileInfo) GetLocalConfFilePath() string {
	if x != nil {
		return x.LocalConfFilePath
	}
	return ""
}

type ConnectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip                   string          `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 string          `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	DbName               string          `protobuf:"bytes,4,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	Username             string          `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password             string          `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Schema               string          `protobuf:"bytes,7,opt,name=schema,proto3" json:"schema,omitempty"`
	IpAndPort            string          `protobuf:"bytes,8,opt,name=ip_and_port,json=ipAndPort,proto3" json:"ip_and_port,omitempty"`
	AuthType             AuthType        `protobuf:"varint,9,opt,name=auth_type,json=authType,proto3,enum=dataset.AuthType" json:"auth_type,omitempty"`
	ServiceDiscoveryMode string          `protobuf:"bytes,10,opt,name=service_discovery_mode,json=serviceDiscoveryMode,proto3" json:"service_discovery_mode,omitempty"`
	ZookeeperNamespace   string          `protobuf:"bytes,11,opt,name=zookeeper_namespace,json=zookeeperNamespace,proto3" json:"zookeeper_namespace,omitempty"`
	SaslQop              string          `protobuf:"bytes,12,opt,name=sasl_qop,json=saslQop,proto3" json:"sasl_qop,omitempty"`
	Principal            string          `protobuf:"bytes,13,opt,name=principal,proto3" json:"principal,omitempty"`
	UserPrincipal        string          `protobuf:"bytes,14,opt,name=user_principal,json=userPrincipal,proto3" json:"user_principal,omitempty"`
	ConfFile             []*ConfFileInfo `protobuf:"bytes,15,rep,name=conf_file,json=confFile,proto3" json:"conf_file,omitempty"`
	DsourceType          DataConnType    `protobuf:"varint,16,opt,name=dsource_type,json=dsourceType,proto3,enum=dataset.DataConnType" json:"dsource_type,omitempty"`
	AdvanceParamConf     string          `protobuf:"bytes,17,opt,name=advance_param_conf,json=advanceParamConf,proto3" json:"advance_param_conf,omitempty"`
	DriverClassName      string          `protobuf:"bytes,18,opt,name=driver_class_name,json=driverClassName,proto3" json:"driver_class_name,omitempty"`
	TableName            string          `protobuf:"bytes,19,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	SqlMap               string          `protobuf:"bytes,20,opt,name=sql_map,json=sqlMap,proto3" json:"sql_map,omitempty"`
	DsourceMfrs          DataSourceMfrs  `protobuf:"varint,21,opt,name=dsource_mfrs,json=dsourceMfrs,proto3,enum=dataset.DataSourceMfrs" json:"dsource_mfrs,omitempty"`
	ConnectType          ConnectType     `protobuf:"varint,22,opt,name=connect_type,json=connectType,proto3,enum=dataset.ConnectType" json:"connect_type,omitempty"`
}

func (x *ConnectionInfo) Reset() {
	*x = ConnectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionInfo) ProtoMessage() {}

func (x *ConnectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionInfo.ProtoReflect.Descriptor instead.
func (*ConnectionInfo) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConnectionInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ConnectionInfo) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ConnectionInfo) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *ConnectionInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ConnectionInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConnectionInfo) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *ConnectionInfo) GetIpAndPort() string {
	if x != nil {
		return x.IpAndPort
	}
	return ""
}

func (x *ConnectionInfo) GetAuthType() AuthType {
	if x != nil {
		return x.AuthType
	}
	return AuthType_AT_NOT_SPECIFIED
}

func (x *ConnectionInfo) GetServiceDiscoveryMode() string {
	if x != nil {
		return x.ServiceDiscoveryMode
	}
	return ""
}

func (x *ConnectionInfo) GetZookeeperNamespace() string {
	if x != nil {
		return x.ZookeeperNamespace
	}
	return ""
}

func (x *ConnectionInfo) GetSaslQop() string {
	if x != nil {
		return x.SaslQop
	}
	return ""
}

func (x *ConnectionInfo) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *ConnectionInfo) GetUserPrincipal() string {
	if x != nil {
		return x.UserPrincipal
	}
	return ""
}

func (x *ConnectionInfo) GetConfFile() []*ConfFileInfo {
	if x != nil {
		return x.ConfFile
	}
	return nil
}

func (x *ConnectionInfo) GetDsourceType() DataConnType {
	if x != nil {
		return x.DsourceType
	}
	return DataConnType_DCT_NOT_SPECIFIED
}

func (x *ConnectionInfo) GetAdvanceParamConf() string {
	if x != nil {
		return x.AdvanceParamConf
	}
	return ""
}

func (x *ConnectionInfo) GetDriverClassName() string {
	if x != nil {
		return x.DriverClassName
	}
	return ""
}

func (x *ConnectionInfo) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *ConnectionInfo) GetSqlMap() string {
	if x != nil {
		return x.SqlMap
	}
	return ""
}

func (x *ConnectionInfo) GetDsourceMfrs() DataSourceMfrs {
	if x != nil {
		return x.DsourceMfrs
	}
	return DataSourceMfrs_DSM_NOT_SPECIFIED
}

func (x *ConnectionInfo) GetConnectType() ConnectType {
	if x != nil {
		return x.ConnectType
	}
	return ConnectType_CT_NOT_SPECIFIED
}

type RowData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowData map[string]string `protobuf:"bytes,1,rep,name=row_data,json=rowData,proto3" json:"row_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RowData) Reset() {
	*x = RowData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowData) ProtoMessage() {}

func (x *RowData) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowData.ProtoReflect.Descriptor instead.
func (*RowData) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{2}
}

func (x *RowData) GetRowData() map[string]string {
	if x != nil {
		return x.RowData
	}
	return nil
}

type DataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataset []*RowData `protobuf:"bytes,1,rep,name=dataset,proto3" json:"dataset,omitempty"`
	Total   int64      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *DataSet) Reset() {
	*x = DataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSet) ProtoMessage() {}

func (x *DataSet) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSet.ProtoReflect.Descriptor instead.
func (*DataSet) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{3}
}

func (x *DataSet) GetDataset() []*RowData {
	if x != nil {
		return x.Dataset
	}
	return nil
}

func (x *DataSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_dataset_proto protoreflect.FileDescriptor

var file_dataset_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x1a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde,
	0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x3a, 0x0a, 0x0c, 0x64, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x64, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x66, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x3b, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f,
	0x0a, 0x14, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0xb9, 0x06, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x1e, 0x0a, 0x0b, 0x69, 0x70, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x2e, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x34, 0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x73, 0x6c, 0x5f, 0x71,
	0x6f, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x73, 0x6c, 0x51, 0x6f,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x64, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43,
	0x6f, 0x6e, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x64, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43, 0x6f,
	0x6e, 0x66, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x71, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x71, 0x6c, 0x4d, 0x61, 0x70, 0x12, 0x3a, 0x0a, 0x0c, 0x64, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6d, 0x66, 0x72, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x66, 0x72, 0x73, 0x52, 0x0b, 0x64, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x66,
	0x72, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x7f, 0x0a, 0x07, 0x52,
	0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x6f, 0x77, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x6f, 0x77, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x3a, 0x0a, 0x0c, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x07,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x59, 0x0a, 0x0e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44,
	0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x44, 0x42, 0x4d, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x42, 0x49, 0x47, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x4d,
	0x51, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x10, 0x04, 0x2a, 0x7f, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x46, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4b,
	0x52, 0x42, 0x35, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x54, 0x41, 0x42, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x4f, 0x52, 0x45, 0x5f, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x58, 0x4d, 0x4c, 0x10, 0x03, 0x12, 0x11,
	0x0a, 0x0d, 0x48, 0x44, 0x46, 0x53, 0x5f, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x58, 0x4d, 0x4c, 0x10,
	0x04, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x49, 0x54, 0x45, 0x5f,
	0x58, 0x4d, 0x4c, 0x10, 0x05, 0x2a, 0x87, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f,
	0x6e, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x43, 0x54, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x47, 0x41, 0x55, 0x53, 0x53, 0x5f, 0x44, 0x42, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x59, 0x53, 0x51, 0x4c, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x52, 0x41, 0x43, 0x4c, 0x45,
	0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x56, 0x45, 0x10, 0x15, 0x12, 0x09, 0x0a, 0x05,
	0x48, 0x42, 0x41, 0x53, 0x45, 0x10, 0x16, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x41, 0x46, 0x4b, 0x41,
	0x10, 0x1f, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x54, 0x50, 0x10, 0x29, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x46, 0x54, 0x50, 0x10, 0x2a, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x44, 0x46, 0x53, 0x10, 0x2b, 0x2a,
	0x3a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x41,
	0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x4b, 0x45, 0x52, 0x42, 0x45, 0x52, 0x4f, 0x53, 0x10, 0x02, 0x2a, 0x37, 0x0a, 0x0e, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x66, 0x72, 0x73, 0x12, 0x15, 0x0a,
	0x11, 0x44, 0x53, 0x4d, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4c, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02,
	0x48, 0x57, 0x10, 0x07, 0x2a, 0x3e, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x49, 0x44,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x02, 0x32, 0x55, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x6a,
	0x68, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x3b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dataset_proto_rawDescOnce sync.Once
	file_dataset_proto_rawDescData = file_dataset_proto_rawDesc
)

func file_dataset_proto_rawDescGZIP() []byte {
	file_dataset_proto_rawDescOnce.Do(func() {
		file_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_dataset_proto_rawDescData)
	})
	return file_dataset_proto_rawDescData
}

var file_dataset_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dataset_proto_goTypes = []interface{}{
	(DataSourceType)(0),             // 0: dataset.DataSourceType
	(ConfFileType)(0),               // 1: dataset.ConfFileType
	(DataConnType)(0),               // 2: dataset.DataConnType
	(AuthType)(0),                   // 3: dataset.AuthType
	(DataSourceMfrs)(0),             // 4: dataset.DataSourceMfrs
	(ConnectType)(0),                // 5: dataset.ConnectType
	(*ConfFileInfo)(nil),            // 6: dataset.ConfFileInfo
	(*ConnectionInfo)(nil),          // 7: dataset.ConnectionInfo
	(*RowData)(nil),                 // 8: dataset.RowData
	(*DataSet)(nil),                 // 9: dataset.DataSet
	nil,                             // 10: dataset.RowData.RowDataEntry
	(*response.ResponseResult)(nil), // 11: response.ResponseResult
}
var file_dataset_proto_depIdxs = []int32{
	0,  // 0: dataset.ConfFileInfo.dsource_type:type_name -> dataset.DataSourceType
	1,  // 1: dataset.ConfFileInfo.conf_file_type:type_name -> dataset.ConfFileType
	3,  // 2: dataset.ConnectionInfo.auth_type:type_name -> dataset.AuthType
	6,  // 3: dataset.ConnectionInfo.conf_file:type_name -> dataset.ConfFileInfo
	2,  // 4: dataset.ConnectionInfo.dsource_type:type_name -> dataset.DataConnType
	4,  // 5: dataset.ConnectionInfo.dsource_mfrs:type_name -> dataset.DataSourceMfrs
	5,  // 6: dataset.ConnectionInfo.connect_type:type_name -> dataset.ConnectType
	10, // 7: dataset.RowData.row_data:type_name -> dataset.RowData.RowDataEntry
	8,  // 8: dataset.DataSet.dataset:type_name -> dataset.RowData
	7,  // 9: dataset.DataSetService.QueryForList:input_type -> dataset.ConnectionInfo
	11, // 10: dataset.DataSetService.QueryForList:output_type -> response.ResponseResult
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_dataset_proto_init() }
func file_dataset_proto_init() {
	if File_dataset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dataset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfFileInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dataset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dataset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dataset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dataset_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dataset_proto_goTypes,
		DependencyIndexes: file_dataset_proto_depIdxs,
		EnumInfos:         file_dataset_proto_enumTypes,
		MessageInfos:      file_dataset_proto_msgTypes,
	}.Build()
	File_dataset_proto = out.File
	file_dataset_proto_rawDesc = nil
	file_dataset_proto_goTypes = nil
	file_dataset_proto_depIdxs = nil
}
