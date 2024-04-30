package Components

/*
	插件管理器对象
*/

import "github.com/dullgiulio/pingo"

type TPluginItem struct {
	Instance *pingo.Plugin
	Protol   string //协议
	Folder   string //路径
}

type TPlugins struct {
	Items []*TPluginItem
}

/*
	TPlugins方法：
	* AddItem()：增加项
	* SaveToFile()：保存到文件
	* LoadFromFile()：从文件中读取
*/
