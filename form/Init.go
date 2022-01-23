// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TInit struct {
    *vcl.TForm
    Check    *vcl.TButton
    Group    *vcl.TGroupBox
    InitShow *vcl.TLabel
    License  *vcl.TMemo

    //::private::
    TInitFields
}

var Init *TInit




// vcl.Application.CreateForm(&Init)

func NewInit(owner vcl.IComponent) (root *TInit)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/Init.gfm
var initBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(Init, &initBytes)
