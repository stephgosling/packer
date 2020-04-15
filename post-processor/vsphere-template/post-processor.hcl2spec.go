// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package vsphere_template

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Host                *string           `mapstructure:"host" cty:"host"`
	Insecure            *bool             `mapstructure:"insecure" cty:"insecure"`
	Username            *string           `mapstructure:"username" cty:"username"`
	Password            *string           `mapstructure:"password" cty:"password"`
	Datacenter          *string           `mapstructure:"datacenter" cty:"datacenter"`
	Folder              *string           `mapstructure:"folder" cty:"folder"`
	SnapshotEnable      *bool             `mapstructure:"snapshot_enable" cty:"snapshot_enable"`
	SnapshotName        *string           `mapstructure:"snapshot_name" cty:"snapshot_name"`
	SnapshotDescription *string           `mapstructure:"snapshot_description" cty:"snapshot_description"`
	ReregisterVM        *bool             `mapstructure:"reregister_vm" default:"true" cty:"reregister_vm"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"host":                       &hcldec.AttrSpec{Name: "host", Type: cty.String, Required: false},
		"insecure":                   &hcldec.AttrSpec{Name: "insecure", Type: cty.Bool, Required: false},
		"username":                   &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":                   &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"datacenter":                 &hcldec.AttrSpec{Name: "datacenter", Type: cty.String, Required: false},
		"folder":                     &hcldec.AttrSpec{Name: "folder", Type: cty.String, Required: false},
		"snapshot_enable":            &hcldec.AttrSpec{Name: "snapshot_enable", Type: cty.Bool, Required: false},
		"snapshot_name":              &hcldec.AttrSpec{Name: "snapshot_name", Type: cty.String, Required: false},
		"snapshot_description":       &hcldec.AttrSpec{Name: "snapshot_description", Type: cty.String, Required: false},
		"reregister_vm":              &hcldec.AttrSpec{Name: "reregister_vm", Type: cty.Bool, Required: false},
	}
	return s
}
