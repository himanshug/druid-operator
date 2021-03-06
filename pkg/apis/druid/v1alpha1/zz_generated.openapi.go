// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/druid/v1alpha1.Druid":       schema_pkg_apis_druid_v1alpha1_Druid(ref),
		"./pkg/apis/druid/v1alpha1.DruidSpec":   schema_pkg_apis_druid_v1alpha1_DruidSpec(ref),
		"./pkg/apis/druid/v1alpha1.DruidStatus": schema_pkg_apis_druid_v1alpha1_DruidStatus(ref),
	}
}

func schema_pkg_apis_druid_v1alpha1_Druid(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Druid is the Schema for the druids API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/druid/v1alpha1.DruidSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/druid/v1alpha1.DruidStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/druid/v1alpha1.DruidSpec", "./pkg/apis/druid/v1alpha1.DruidStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_druid_v1alpha1_DruidSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DruidSpec defines the desired state of Druid",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_druid_v1alpha1_DruidStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DruidStatus defines the observed state of Druid",
				Type:        []string{"object"},
			},
		},
	}
}
