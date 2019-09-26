// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package chroot

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"AMIMappings":       nil,
		"ChrootMounts":      nil,
		"CommandWrapper":    &hcldec.AttrSpec{Name: "CommandWrapper", Type: cty.String, Required: false},
		"CopyFiles":         &hcldec.AttrSpec{Name: "CopyFiles", Type: cty.List(cty.String), Required: false},
		"DevicePath":        &hcldec.AttrSpec{Name: "DevicePath", Type: cty.String, Required: false},
		"NVMEDevicePath":    &hcldec.AttrSpec{Name: "NVMEDevicePath", Type: cty.String, Required: false},
		"FromScratch":       &hcldec.AttrSpec{Name: "FromScratch", Type: cty.Bool, Required: false},
		"MountOptions":      &hcldec.AttrSpec{Name: "MountOptions", Type: cty.List(cty.String), Required: false},
		"MountPartition":    &hcldec.AttrSpec{Name: "MountPartition", Type: cty.String, Required: false},
		"MountPath":         &hcldec.AttrSpec{Name: "MountPath", Type: cty.String, Required: false},
		"PostMountCommands": &hcldec.AttrSpec{Name: "PostMountCommands", Type: cty.List(cty.String), Required: false},
		"PreMountCommands":  &hcldec.AttrSpec{Name: "PreMountCommands", Type: cty.List(cty.String), Required: false},
		"RootDeviceName":    &hcldec.AttrSpec{Name: "RootDeviceName", Type: cty.String, Required: false},
		"RootVolumeSize":    nil,
		"RootVolumeType":    &hcldec.AttrSpec{Name: "RootVolumeType", Type: cty.String, Required: false},
		"SourceAmi":         &hcldec.AttrSpec{Name: "SourceAmi", Type: cty.String, Required: false},
		"SourceAmiFilter":   nil,
		"RootVolumeTags":    nil,
		"Architecture":      &hcldec.AttrSpec{Name: "Architecture", Type: cty.String, Required: false},
	}
	return hcldec.ObjectSpec(s)
}
