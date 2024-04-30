package Components

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/guid"
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
	Name      string  //应用名称
	Version   float32 //版本编号
	Serial    string  //唯一序列号
	ConfigDir string  //保存自定义设置文件的文件夹路径，默认为./Config
	StaticDir string  //保存静态资源文件的文件夹路径，默认为./Static
	CertDir   string  //保存各种证书的文件夹路径，默认为./Cert
	DataDir   string  //保存本地数据的文件夹路径，默认为./Data
	LuaDir    string  //保存Lua脚本的子文件夹名，默认为Lua
	PluginDir string  //保存插件的子文件夹名，默认为Plugin

	Databases    TDatabases    //数据库管理器
	Redises      TRedises      //缓存Redis管理器
	SystemConfig TSystemConfig //系统设置项
	MainServer   *ghttp.Server //主Web服务
}

/*
	TApp对象方法：
	* Init()：初始化方法
	* Start()：启动方法
*/

// Init 初始化主程序对象
func (p *TApp) Init() error {
	//初始化默认值
	p.Name = "GoBlocks"
	p.Version = 0.01
	p.Serial = guid.S()
	p.ConfigDir = Const_Config_Dir
	p.DataDir = Const_Data_Dir
	p.PluginDir = Const_Plugin_Dir
	p.CertDir = Const_Cert_Dir
	p.LuaDir = Const_Lua_Dir
	p.StaticDir = Const_Static_Dir
	p.MainServer = g.Server()
	//检查目录是否存在，如果不存在则创建之
	if gfile.Exists(Const_Config_Dir) == false {
		_ = gfile.Mkdir(Const_Config_Dir)
	}
	if gfile.Exists(Const_Data_Dir) == false {
		_ = gfile.Mkdir(Const_Data_Dir)
	}
	if gfile.Exists(Const_Plugin_Dir) == false {
		_ = gfile.Mkdir(Const_Plugin_Dir)
	}
	if gfile.Exists(Const_Lua_Dir) == false {
		_ = gfile.Mkdir(Const_Lua_Dir)
	}
	if gfile.Exists(Const_Cert_Dir) == false {
		_ = gfile.Mkdir(Const_Cert_Dir)
	}
	if gfile.Exists(Const_Static_Dir) == false {
		_ = gfile.Mkdir(Const_Static_Dir)
	}
	//读取自定义的系统设置
	e1 := p.SystemConfig.LoadFromFile(gfile.Join(p.ConfigDir, Const_Default_System_Config_File))
	if e1 != nil {
		return e1
	}
	//初始化数据库连接和缓存，并绑定缓存到数据库连接，Redis必须先于数据库执行
	p.Redises = TRedises{}
	e3 := p.Redises.LoadFromFile(gfile.Join(p.ConfigDir, Const_Default_Redis_Config_File))
	if e3 != nil {
		return e3
	}
	p.Redises.Apply()
	p.Databases = TDatabases{}
	e2 := p.Databases.LoadFromFile(gfile.Join(p.ConfigDir, Const_Default_Db_Config_File))
	if e2 != nil {
		return e2
	}
	p.Databases.Apply()
	p.Databases.SetCache()
	//初始化插件管理器

	//初始化路由，包括脚本、插件和反向代理，为简化起见一律使用固定的路由前缀

	//加载默认环境，包括Session和全局对象

	return nil
}

// Start 启动主服务
func (p *TApp) Start() {

	p.MainServer.Run()
}
