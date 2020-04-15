// Code generated by "mapstructure-to-hcl2 -type Config,DtlArtifact,ArtifactParameter"; DO NOT EDIT.
package devtestlabsartifacts

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatArtifactParameter is an auto-generated flat version of ArtifactParameter.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatArtifactParameter struct {
	Name  *string `mapstructure:"name" cty:"name"`
	Value *string `mapstructure:"value" cty:"value"`
	Type  *string `mapstructure:"type" cty:"type"`
}

// FlatMapstructure returns a new FlatArtifactParameter.
// FlatArtifactParameter is an auto-generated flat version of ArtifactParameter.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ArtifactParameter) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatArtifactParameter)
}

// HCL2Spec returns the hcl spec of a ArtifactParameter.
// This spec is used by HCL to read the fields of ArtifactParameter.
// The decoded values from this spec will then be applied to a FlatArtifactParameter.
func (*FlatArtifactParameter) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name":  &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"value": &hcldec.AttrSpec{Name: "value", Type: cty.String, Required: false},
		"type":  &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
	}
	return s
}

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName        *string                `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType      *string                `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug            *bool                  `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce            *bool                  `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError          *string                `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars         map[string]string      `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars    []string               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	CloudEnvironmentName   *string                `mapstructure:"cloud_environment_name" required:"false" cty:"cloud_environment_name"`
	ClientID               *string                `mapstructure:"client_id" cty:"client_id"`
	ClientSecret           *string                `mapstructure:"client_secret" cty:"client_secret"`
	ClientCertPath         *string                `mapstructure:"client_cert_path" cty:"client_cert_path"`
	ClientJWT              *string                `mapstructure:"client_jwt" cty:"client_jwt"`
	ObjectID               *string                `mapstructure:"object_id" cty:"object_id"`
	TenantID               *string                `mapstructure:"tenant_id" required:"false" cty:"tenant_id"`
	SubscriptionID         *string                `mapstructure:"subscription_id" cty:"subscription_id"`
	DtlArtifacts           []FlatDtlArtifact      `mapstructure:"dtl_artifacts" cty:"dtl_artifacts"`
	LabName                *string                `mapstructure:"lab_name" cty:"lab_name"`
	ResourceGroupName      *string                `mapstructure:"lab_resource_group_name" cty:"lab_resource_group_name"`
	VMName                 *string                `mapstructure:"vm_name" cty:"vm_name"`
	PollingDurationTimeout *string                `mapstructure:"polling_duration_timeout" required:"false" cty:"polling_duration_timeout"`
	AzureTags              map[string]*string     `mapstructure:"azure_tags" cty:"azure_tags"`
	Json                   map[string]interface{} `cty:"json"`
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
		"cloud_environment_name":     &hcldec.AttrSpec{Name: "cloud_environment_name", Type: cty.String, Required: false},
		"client_id":                  &hcldec.AttrSpec{Name: "client_id", Type: cty.String, Required: false},
		"client_secret":              &hcldec.AttrSpec{Name: "client_secret", Type: cty.String, Required: false},
		"client_cert_path":           &hcldec.AttrSpec{Name: "client_cert_path", Type: cty.String, Required: false},
		"client_jwt":                 &hcldec.AttrSpec{Name: "client_jwt", Type: cty.String, Required: false},
		"object_id":                  &hcldec.AttrSpec{Name: "object_id", Type: cty.String, Required: false},
		"tenant_id":                  &hcldec.AttrSpec{Name: "tenant_id", Type: cty.String, Required: false},
		"subscription_id":            &hcldec.AttrSpec{Name: "subscription_id", Type: cty.String, Required: false},
		"dtl_artifacts":              &hcldec.BlockListSpec{TypeName: "dtl_artifacts", Nested: hcldec.ObjectSpec((*FlatDtlArtifact)(nil).HCL2Spec())},
		"lab_name":                   &hcldec.AttrSpec{Name: "lab_name", Type: cty.String, Required: false},
		"lab_resource_group_name":    &hcldec.AttrSpec{Name: "lab_resource_group_name", Type: cty.String, Required: false},
		"vm_name":                    &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"polling_duration_timeout":   &hcldec.AttrSpec{Name: "polling_duration_timeout", Type: cty.String, Required: false},
		"azure_tags":                 &hcldec.AttrSpec{Name: "azure_tags", Type: cty.Map(cty.String), Required: false},
		"json":                       &hcldec.AttrSpec{Name: "json", Type: cty.Map(cty.String), Required: false},
	}
	return s
}

// FlatDtlArtifact is an auto-generated flat version of DtlArtifact.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDtlArtifact struct {
	ArtifactName *string                 `mapstructure:"artifact_name" cty:"artifact_name"`
	ArtifactId   *string                 `mapstructure:"artifact_id" cty:"artifact_id"`
	Parameters   []FlatArtifactParameter `mapstructure:"parameters" cty:"parameters"`
}

// FlatMapstructure returns a new FlatDtlArtifact.
// FlatDtlArtifact is an auto-generated flat version of DtlArtifact.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DtlArtifact) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDtlArtifact)
}

// HCL2Spec returns the hcl spec of a DtlArtifact.
// This spec is used by HCL to read the fields of DtlArtifact.
// The decoded values from this spec will then be applied to a FlatDtlArtifact.
func (*FlatDtlArtifact) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"artifact_name": &hcldec.AttrSpec{Name: "artifact_name", Type: cty.String, Required: false},
		"artifact_id":   &hcldec.AttrSpec{Name: "artifact_id", Type: cty.String, Required: false},
		"parameters":    &hcldec.BlockListSpec{TypeName: "parameters", Nested: hcldec.ObjectSpec((*FlatArtifactParameter)(nil).HCL2Spec())},
	}
	return s
}
