// 由res2go IDE插件自动生成，不要编辑。
package main

import (
	"github.com/ying32/govcl/vcl"
)

type TForm1 struct {
	*vcl.TForm
	Label1        *vcl.TLabel
	Label2        *vcl.TLabel
	InFile        *vcl.TEdit
	OutFile       *vcl.TEdit
	InFileButton  *vcl.TButton
	OutFileButton *vcl.TButton
	ComRadio      *vcl.TRadioButton
	UnComRadio    *vcl.TRadioButton
	KeyFile       *vcl.TEdit
	Label3        *vcl.TLabel
	KeyFileButton *vcl.TButton
	EnButton      *vcl.TCheckBox
	RunButton     *vcl.TToggleBox
	InFileDialog  *vcl.TOpenDialog
	KeyFileDialog *vcl.TOpenDialog

	//::private::
	TForm1Fields
}

var Form1 *TForm1

// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1) {
	vcl.CreateResForm(owner, &root)
	return
}

var form1Bytes = []byte("\x54\x50\x46\x30\x06\x54\x46\x6F\x72\x6D\x31\x05\x46\x6F\x72\x6D\x31\x04\x4C\x65\x66\x74\x03\x87\x01\x06\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x03\x54\x6F\x70\x03\x0B\x01\x05\x57\x69\x64\x74\x68\x03\xF5\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\x41\x72\x63\x68\x69\x76\x65\x20\x2D\x20\x52\x65\x68\x74\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xF5\x01\x0D\x44\x65\x73\x69\x67\x6E\x54\x69\x6D\x65\x50\x50\x49\x02\x78\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x32\x2E\x32\x2E\x30\x2E\x34\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x0C\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x11\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBE\x93\xE5\x85\xA5\xE6\x96\x87\xE4\xBB\xB6\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x0C\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x40\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBE\x93\xE5\x87\xBA\xE6\x96\x87\xE4\xBB\xB6\x00\x00\x05\x54\x45\x64\x69\x74\x06\x49\x6E\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x03\x20\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x12\xE8\xBE\x93\xE5\x85\xA5\xE6\x96\x87\xE4\xBB\xB6\xE8\xB7\xAF\xE5\xBE\x84\x00\x00\x05\x54\x45\x64\x69\x74\x07\x4F\x75\x74\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x38\x05\x57\x69\x64\x74\x68\x03\x20\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x12\xE8\xBE\x93\xE5\x87\xBA\xE6\x96\x87\xE4\xBB\xB6\xE8\xB7\xAF\xE5\xBE\x84\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x49\x6E\x46\x69\x6C\x65\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x03\x80\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE9\x80\x89\xE6\x8B\xA9\xE6\x96\x87\xE4\xBB\xB6\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0D\x4F\x75\x74\x46\x69\x6C\x65\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x03\x80\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x38\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE9\x80\x89\xE6\x8B\xA9\xE6\x96\x87\xE4\xBB\xB6\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x0C\x54\x52\x61\x64\x69\x6F\x42\x75\x74\x74\x6F\x6E\x08\x43\x6F\x6D\x52\x61\x64\x69\x6F\x04\x4C\x65\x66\x74\x02\x12\x06\x48\x65\x69\x67\x68\x74\x02\x18\x03\x54\x6F\x70\x03\xC0\x00\x05\x57\x69\x64\x74\x68\x02\x36\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xA7\xA3\xE5\x8E\x8B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x0C\x54\x52\x61\x64\x69\x6F\x42\x75\x74\x74\x6F\x6E\x0A\x55\x6E\x43\x6F\x6D\x52\x61\x64\x69\x6F\x04\x4C\x65\x66\x74\x02\x12\x06\x48\x65\x69\x67\x68\x74\x02\x18\x03\x54\x6F\x70\x03\x98\x00\x05\x57\x69\x64\x74\x68\x02\x36\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8E\x8B\xE7\xBC\xA9\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x05\x54\x45\x64\x69\x74\x07\x4B\x65\x79\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x60\x05\x57\x69\x64\x74\x68\x03\x20\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x12\xE5\x8A\xA0\xE5\xAF\x86\xE8\xAF\x81\xE4\xB9\xA6\xE8\xB7\xAF\xE5\xBE\x84\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x0C\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x68\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xAF\x81\xE4\xB9\xA6\xE6\x96\x87\xE4\xBB\xB6\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0D\x4B\x65\x79\x46\x69\x6C\x65\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x03\x80\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x60\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE9\x80\x89\xE6\x8B\xA9\xE6\x96\x87\xE4\xBB\xB6\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x07\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x08\x45\x6E\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x02\x68\x06\x48\x65\x69\x67\x68\x74\x02\x18\x03\x54\x6F\x70\x03\x98\x00\x05\x57\x69\x64\x74\x68\x02\x36\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8A\xA0\xE5\xAF\x86\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x08\x00\x00\x0A\x54\x54\x6F\x67\x67\x6C\x65\x42\x6F\x78\x09\x52\x75\x6E\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x03\x80\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x03\xB0\x00\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xBF\x90\xE8\xA1\x8C\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x09\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x0C\x49\x6E\x46\x69\x6C\x65\x44\x69\x61\x6C\x6F\x67\x04\x4C\x65\x66\x74\x03\xD0\x00\x03\x54\x6F\x70\x03\x90\x00\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x0D\x4B\x65\x79\x46\x69\x6C\x65\x44\x69\x61\x6C\x6F\x67\x04\x4C\x65\x66\x74\x03\x18\x01\x03\x54\x6F\x70\x03\x90\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(Form1, &form1Bytes)