package App

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
)

/*
	根应用程序的对象
*/

// TSystemConfig 系统设置对象
type TSystemConfig struct {
	BindAddress        string //监听IP地址，如果为空则为0.0.0.0
	BindDomain         string //绑定域名，如果为空则不绑定
	BindHttpPort       int    //绑定的Http端口，默认80
	EnableTLS          bool   //是否启用TLS，默认为False
	BindHttpsPort      int    //绑定的Https端口，默认443
	TLSCrtFile         string //绑定Https端口的CRT证书文件路径
	TLSKeyFile         string //绑定Https端口的KEY证书文件路径
	EnableCerbot       bool   //是否启用Cerbot申请TLS证书，服务商为Let's Encrypt
	EnableAutoRenewTLS bool   //是否启用自动更新TLS证书
	EnableLua          bool   //是否启用Lua脚本
	EnablePlugin       bool   //是否启用插件
	EnableProxy        bool   //是否启动反向代理
}

/*
	TSystemConfig方法
	* LoadFromFile()：从文件中加载属性配置
	* SaveToFiles()：保存属性配置到文件
	* ApplyCustomCfg()：加载客制配置文件并应用到全部CFG上
*/

// LoadFromFile 从文件中加载属性配置
func (p *TSystemConfig) LoadFromFile(AFilename string) error {
	if gfile.Exists(AFilename) == true {
		mContents := gfile.GetContents(AFilename)
		mJson := gjson.New(mContents)
		p.BindAddress = mJson.GetString("BindAddress")
		p.BindDomain = mJson.GetString("BindDomain")
		p.BindHttpPort = mJson.GetInt("BindHttpPort")
		p.EnableTLS = mJson.GetBool("EnableTLS")
		p.BindHttpsPort = mJson.GetInt("BindHttpsPort")
		p.TLSCrtFile = mJson.GetString("TLSCrtFile")
		p.TLSKeyFile = mJson.GetString("TLSKeyFile")
		p.EnableCerbot = mJson.GetBool("EnableCerbot")
		p.EnableAutoRenewTLS = mJson.GetBool("EnableAutoRenewTLS")
		p.EnableLua = mJson.GetBool("EnableLua")
		p.EnablePlugin = mJson.GetBool("EnablePlugin")
		p.EnableProxy = mJson.GetBool("EnableProxy")
		return nil
	} else {
		return errors.New(AFilename + "文件不存在")
	}
}

// SaveToFile 保存设置属性到文件
func (p *TSystemConfig) SaveToFile(AFilename string) {
	mJson := gjson.New(p)
	_ = gfile.PutContents(AFilename, mJson.Export())
}

// TApp 应用成勋对象
type TApp struct {
	Name         string  //应用名称
	Version      float32 //版本编号
	ConfigDir    string  //保存自定义设置文件的文件夹路径
	StaticDir    string  //保存静态资源文件的文件夹路径
	LuaSubDir    string  //保存Lua脚本的子文件夹名
	PluginSubDir string  //保存插件的子文件夹名

	SystemConfig TSystemConfig //系统设置项
}
