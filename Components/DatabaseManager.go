package Components

import (
	"errors"
	"github.com/gogf/gcache-adapter/adapter"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

/*
	数据库配置对象
*/

// TRedisItem Redis配置项
type TRedisItem struct {
	Name string //名称
	Link string //连接字符串
}

// TRedises Redis配置管理主对象
type TRedises struct {
	Items []*TRedisItem
}

/*
	TRedises对象方法：
	* AddItem()：添加项
	* SaveToFile()：保存到文件
	* LoadFromFile()：从文件中读取
	* Apply()：应用配置到覆盖主设置中
*/

func (p *TRedises) AddItem(AItem TRedisItem) {
	mItem := TRedisItem{
		Name: AItem.Name,
		Link: AItem.Link,
	}
	p.Items = append(p.Items, &mItem)
}

func (p *TRedises) SaveToFile(AFilename string) {
	var mItems []g.Map
	for _, v := range p.Items {
		mItem := g.Map{
			"Name": v.Name,
			"Link": v.Link,
		}
		mItems = append(mItems, mItem)
		mJson := gjson.New(mItems)
		_ = gfile.PutContents(AFilename, mJson.Export())
	}
}

func (p *TRedises) LoadFromFile(AFilename string) error {
	if gfile.Exists(AFilename) == true {
		mContents := gfile.GetContents(AFilename)
		mJson := gjson.New(mContents)
		mJsonArray := mJson.Array()
		for _, v := range mJsonArray {
			mJsonItem := gjson.New(v)
			var mItem TRedisItem = TRedisItem{
				Name: mJsonItem.GetString("Name"),
				Link: mJsonItem.GetString("Link"),
			}
			p.Items = append(p.Items, &mItem)
		}
		return nil
	} else {
		return errors.New(AFilename + "文件不存在")
	}
}

func (p *TRedises) Apply() {
	for _, v := range p.Items {
		_ = g.Cfg().Set("redis."+v.Name+"Link", v.Link)
	}
}

// TDatabaseItem 数据项对象
type TDatabaseItem struct {
	Name           string //名称
	Type           string //类型
	Link           string //连接字符串
	EnableCache    bool   //是否缓存
	UseMemoryCache bool   //是否使用内存缓存，默认为false
	RedisName      string //使用哪个Redis，如果为空，则使用default
}

// TDatabases 数据库管理器对象
type TDatabases struct {
	Items []*TDatabaseItem
}

/*
	数据库管理对象方法：
	* AddItem()：增加项
	* SaveToFile()：保存到设置文件
	* LoadFromFile()：从设置文件中载入
	* Apply()：覆盖到系统设置中
	* SetCache()：为数据库连接设置缓存
*/

func (p *TDatabases) AddItem(AItem TDatabaseItem) {
	mItem := TDatabaseItem{
		Name:           AItem.Name,
		Type:           AItem.Type,
		Link:           AItem.Link,
		EnableCache:    AItem.EnableCache,
		UseMemoryCache: AItem.UseMemoryCache,
		RedisName:      AItem.RedisName,
	}
	p.Items = append(p.Items, &mItem)
}

// SavetoFile 保存到设置文件中
func (p *TDatabases) SavetoFile(AFilename string) {
	var mItems []g.Map
	for _, v := range p.Items {
		mItem := g.Map{
			"Name":           v.Name,
			"Type":           v.Type,
			"Link":           v.Link,
			"EnableCache":    v.EnableCache,
			"UseMemoryCache": v.UseMemoryCache,
			"RedisName":      v.RedisName,
		}
		mItems = append(mItems, mItem)
	}
	mJson := gjson.New(mItems)
	_ = gfile.PutContents(AFilename, mJson.Export())
}

// LoadFromFile 从文件中读取设置
func (p *TDatabases) LoadFromFile(AFilename string) error {
	if gfile.Exists(AFilename) == true {
		mContents := gfile.GetContents(AFilename)
		mJsons := gjson.New(mContents)
		mJsonArray := mJsons.Array()
		for _, v := range mJsonArray {
			mJsonItem := gjson.New(v)
			var mItem TDatabaseItem = TDatabaseItem{
				Name:           mJsonItem.GetString("Name"),
				Type:           mJsonItem.GetString("Type"),
				Link:           mJsonItem.GetString("Link"),
				EnableCache:    mJsonItem.GetBool("EnableCache"),
				UseMemoryCache: mJsonItem.GetBool("UseMemoryCache"),
				RedisName:      mJsonItem.GetString("RedisName"),
			}
			p.Items = append(p.Items, &mItem)
		}
		return nil
	} else {
		return errors.New(AFilename + "文件不存在")
	}
}

// Apply 应用，即覆盖到系统设置中
func (p *TDatabases) Apply() {
	for _, v := range p.Items {
		_ = g.Cfg().Set("database."+v.Name+".link", v.Link)
	}
}

// SetCache 循环设置数据库缓存
func (p *TDatabases) SetCache() {
	for _, v := range p.Items {
		if v.EnableCache == true {
			if v.UseMemoryCache == false {
				//设置Redis缓存
				mAdapter := adapter.NewRedis(g.Redis(v.RedisName))
				g.DB(v.Name).GetCache().SetAdapter(mAdapter)
			}
		}
	}
}
