// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package common

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*ISOConfig) HCL2Schema() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"ISOChecksum": &hcldec.AttrSpec{Name:"ISOChecksum", Type:cty.String, Required:false},
		"ISOChecksumURL": &hcldec.AttrSpec{Name:"ISOChecksumURL", Type:cty.String, Required:false},
		"ISOChecksumType": &hcldec.AttrSpec{Name:"ISOChecksumType", Type:cty.String, Required:false},
		"RawSingleISOUrl": &hcldec.AttrSpec{Name:"RawSingleISOUrl", Type:cty.String, Required:false},
		"ISOUrls": &hcldec.AttrSpec{Name:"ISOUrls", Type:cty.List(cty.String), Required:false},
		"TargetPath": &hcldec.AttrSpec{Name:"TargetPath", Type:cty.String, Required:false},
		"TargetExtension": &hcldec.AttrSpec{Name:"TargetExtension", Type:cty.String, Required:false},
	}
	return hcldec.ObjectSpec(s)
}
