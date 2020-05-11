package config

var(
	// 数据库地址
	// root：用户名，123456：密码，localhost:3306：地址，gomicro：对应的库，去mysql建一个
	MysqlDNS = "root:123456@(localhost:3306)/gomicro"
	// 服务命名空间
	Namespace = "com.class.cinema."
)
