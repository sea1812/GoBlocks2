package Components

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"net/http/httputil"
	"net/url"
)

/*
	路由管理器对象
*/

type TRouteItem struct {
	Name   string     //名称
	Type   TRouteType //路由类型
	Source string     //路由来源
	Target string     //路由目标，只对反代和静态目录有效
}

type TRoutes struct {
	Items []*TRouteItem
}

/*
	路由管理器方法：
	* AddItem()：增加项
	* SaveToFile()：保存到文件
	* LoadFromFile()：从文件中读取
*/

func (p *TRoutes) AddItem(AItem TRouteItem) {
	mItem := TRouteItem{
		Name:   AItem.Name,
		Type:   AItem.Type,
		Target: AItem.Target,
	}
	p.Items = append(p.Items, &mItem)
}

func (p *TRoutes) LoadFromFile(AFilename string) error {
	p.Items = nil
	if gfile.Exists(AFilename) == true {
		mContents := gfile.GetContents(AFilename)
		mJson := gjson.New(mContents)
		mJsonArray := mJson.Array()
		for _, v := range mJsonArray {
			mItem := gjson.New(v)
			mItem_i := TRouteItem{
				Name:   mItem.GetString("Name"),
				Type:   TRouteType(mItem.GetInt("Type")),
				Target: mItem.GetString("Target"),
			}
			p.Items = append(p.Items, &mItem_i)
		}
		return nil
	} else {
		return errors.New(AFilename + "文件不存在")
	}
}

func (p *TRoutes) SaveToFile(AFilename string) {
	var mMaps []g.Map
	for _, v := range p.Items {
		mMap := g.Map{
			"Target": v.Target,
			"Name":   v.Name,
			"Type":   v.Type,
		}
		mMaps = append(mMaps, mMap)
	}
	mJson := gjson.New(mMaps)
	_ = gfile.PutContents(AFilename, mJson.Export())
}

// Apply 应用路由
func (p *TRoutes) Apply(AServer *ghttp.Server) {
	for _, v := range p.Items {
		switch v.Type {
		case Route_Static:
			//处理静态目录
			AServer.AddStaticPath(v.Source, v.Target)
			return
		case Route_Plugin:
			//处理插件路由
			AServer.BindHandler("/plugin/{plugin_name}", HandlerPlugin)
			return
		case Route_Lua:
			//处理Lua脚本路由
			AServer.BindHandler("/lua/{lua_name}", HandlerLua)
			return
		case Route_Proxy:
			//处理反向代理路由
			targetUrl, _ := url.Parse(v.Target)
			proxy := httputil.NewSingleHostReverseProxy(targetUrl)
			AServer.BindHandler(v.Source, func(r *ghttp.Request) {
				proxy.ServeHTTP(r.Response.Writer, r.Request)
			})
			return
		}
		//处理Admin后台的路由

		//处理内置功能的路由

	}
}
