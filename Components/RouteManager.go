package Components

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

/*
	路由管理器对象
*/

type TRouteItem struct {
	Name   string     //名称
	Type   TRouteType //路由类型
	Target string     //路由目标
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
