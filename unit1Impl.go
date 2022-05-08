package main

import (
	"flag"
	"github.com/Rehtt/archive/utils"
	"github.com/ying32/govcl/vcl"
	"io/ioutil"
	"log"
	"path/filepath"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnInFileChange(sender vcl.IObject) {
}

func (f *TForm1) OnOutFileChange(sender vcl.IObject) {

}

func (f *TForm1) OnInFileButtonClick(sender vcl.IObject) {
	if ok, path := vcl.SelectDirectory2("选择输入路径", "", true); ok {
		f.InFile.SetText(path)
	}
}

func (f *TForm1) OnOutFileButtonClick(sender vcl.IObject) {
	if ok, path := vcl.SelectDirectory2("选择输出路径", "", true); ok {
		f.OutFile.SetText(filepath.Join(path, "out.tar.xz"))
	}
}

func (f *TForm1) OnComRadioChange(sender vcl.IObject) {
	if f.ComRadio.Checked() {
		f.UnComRadio.SetChecked(false)
	}
}

func (f *TForm1) OnUnComRadioChange(sender vcl.IObject) {
	if f.UnComRadio.Checked() {
		f.ComRadio.SetChecked(false)
		f.EnButton.SetChecked(false)
	}
}

func (f *TForm1) OnKeyFileChange(sender vcl.IObject) {

}

func (f *TForm1) OnKeyFileButtonClick(sender vcl.IObject) {
	if f.KeyFileDialog.Execute() {
		f.KeyFile.SetText(f.KeyFileDialog.FileName())
	}
}

func (f *TForm1) OnEnButtonChange(sender vcl.IObject) {

}

func (f *TForm1) OnRunButtonChange(sender vcl.IObject) {
	if f.EnButton.Checked() {
		if f.KeyFile.GetTextLen() == 0 {
			vcl.ShowMessage("需要指定公钥")
			return
		}
		data, err := ioutil.ReadFile(f.KeyFile.Text())
		if err != nil {
			log.Fatalln(err)
			flag.Usage()
			return
		}
		utils.InitEncrypt(nil, data)
	}
	if f.InFile.GetTextLen() == 0 {
		vcl.ShowMessage("输入路径为空")
		return
	}
	if f.OutFile.GetTextLen() == 0 {
		vcl.ShowMessage("输出路径为空")
		return
	}

	if !f.ComRadio.Checked() {
		err := utils.Compress(f.InFile.Text(), f.OutFile.Text(), f.EnButton.Checked())
		if err != nil {
			vcl.ShowMessage("错误：" + err.Error())
			return
		}
		vcl.ShowMessage("完成")
		return
	}

	if !f.UnComRadio.Checked() {
		if f.KeyFile.GetTextLen() != 0 {
			data, err := ioutil.ReadFile(f.KeyFile.Text())
			if err == nil {
				utils.InitEncrypt(data, nil)
			}
		}
		err := utils.Uncompress(f.InFile.Text(), f.OutFile.Text())
		if err != nil {
			vcl.ShowMessage("错误：" + err.Error())
			return
		}
		vcl.ShowMessage("完成")
		return
	}

}
