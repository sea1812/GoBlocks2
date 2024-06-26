package main

import (
	App "GoBlocks2/Components"
)

func main() {
	//读取自定义设置，覆盖原始Config
	//初始化数据库连接和缓存，并绑定缓存到数据库连接
	//初始化插件管理器
	//初始化路由，包括脚本、插件和反向代理，为简化起见一律使用固定的路由前缀
	//加载默认环境，包括Session和全局对象
	//启动主服务器
	TApp := App.TApp{}
	er := TApp.Init()
	if er == nil {
		TApp.Start()
	} else {
		panic(er)
	}
}
