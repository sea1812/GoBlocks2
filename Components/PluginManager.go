package Components

/*
	插件管理器对象
*/

import (
	"errors"
	"github.com/dullgiulio/pingo"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

type TPluginItem struct {
	Instance *pingo.Plugin
	Name     string  //名称
	Version  float32 //版本号
	Protol   string  //协议
	CMD      string  //路径
}

// Start 启动插件，创建实例
func (p *TPluginItem) Start() error {
	mFile := gfile.Join(Const_Plugin_Dir, p.CMD)
	if p.Instance != nil {
		if gfile.Exists(mFile) == true {
			p.Instance = pingo.NewPlugin("tcp", mFile)
			p.Instance.Start()
			return nil
		} else {
			return errors.New(mFile + "插件文件不存在")
		}
	} else {
		return errors.New("实例已存在")
	}
}

// Stop 停止插件，释放实例
func (p *TPluginItem) Stop() {
	if p.Instance != nil {
		p.Instance.Stop()
		p.Instance = nil
	}
}

type TPlugins struct {
	Items []*TPluginItem
}

/*
	TPlugins方法：
	* AddItem()：增加项
	* SaveToFile()：保存到文件
	* LoadFromFile()：从文件中读取
	* StartAll()：启动全部插件
	* Items()：用名字查找插件
*/

func (p *TPlugins) AddItem(AItem TPluginItem) {
	mItem := TPluginItem{
		Instance: nil,
		Name:     AItem.Name,
		Version:  AItem.Version,
		Protol:   AItem.Protol,
		CMD:      AItem.CMD,
	}
	p.Items = append(p.Items, &mItem)
}

// ItemByName 按名字查找插件
func (p *TPlugins) ItemByName(AName string) *TPluginItem {
	var mR *TPluginItem = nil
	for _, v := range p.Items {
		if v.Name == AName {

		}
	}
	return mR
}

// StartAll 启动全部插件
func (p *TPlugins) StartAll() {
	for _, v := range p.Items {
		_ = v.Start()
	}
}

// StopAll 停止全部插件
func (p *TPlugins) StopAll() {
	for _, v := range p.Items {
		v.Stop()
	}
}

// SaveToFile 保存插件设置到文件
func (p *TPlugins) SaveToFile(AFilename string) {
	var mMaps []g.Map
	for _, v := range p.Items {
		mMap := g.Map{
			"Name":    v.Name,
			"Version": v.Version,
			"Cmd":     v.CMD,
		}
		mMaps = append(mMaps, mMap)
	}
	mJson := gjson.New(mMaps)
	mFilename := gfile.Join(Const_Plugin_Dir, AFilename)
	_ = gfile.PutContents(mFilename, mJson.Export())
}

// LoadFromFile 从文件中载入全部插件
func (p *TPlugins) LoadFromFile(AFilename string) error {
	if gfile.Exists(AFilename) == true {
		mFilename := gfile.Join(Const_Plugin_Dir, AFilename)
		mContents := gfile.GetContents(mFilename)
		mJson := gjson.New(mContents)
		mJsonArray := mJson.Array()
		for _, v := range mJsonArray {
			mV := gjson.New(v)
			var mItem TPluginItem
			mItem.Name = mV.GetString("Name")
			mItem.Version = mV.GetFloat32("Version")
			mItem.CMD = mV.GetString("Cmd")
			p.Items = append(p.Items, &mItem)
		}
		return nil
	} else {
		return errors.New(AFilename + "插件设置文件不存在")
	}
}
