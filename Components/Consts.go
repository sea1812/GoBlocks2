package Components

/*
	常量定义
*/

const (
	Const_Config_Dir                 = "./Config"
	Const_Static_Dir                 = "./Static"
	Const_Cert_Dir                   = "./Cert"
	Const_Data_Dir                   = "./Data"
	Const_Lua_Dir                    = "./Lua"
	Const_Plugin_Dir                 = "./Plugin"
	Const_Default_System_Config_File = "sys.conf"
	Const_Default_Db_Config_File     = "db.conf"
	Const_Default_Redis_Config_File  = "redis.conf"
	Const_Plugins_Config_File        = "plugins.conf"
)

type TRouteType int

const (
	Route_Static TRouteType = 1
	Route_Lua    TRouteType = 2
	Route_Plugin TRouteType = 3
	Route_Proxy  TRouteType = 4
)
