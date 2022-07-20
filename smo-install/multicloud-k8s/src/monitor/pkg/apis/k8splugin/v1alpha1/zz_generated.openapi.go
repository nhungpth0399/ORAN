// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/k8splugin/v1alpha1.PodStatus":               schema_pkg_apis_k8splugin_v1alpha1_PodStatus(ref),
		"./pkg/apis/k8splugin/v1alpha1.ResourceBundleState":     schema_pkg_apis_k8splugin_v1alpha1_ResourceBundleState(ref),
		"./pkg/apis/k8splugin/v1alpha1.ResourceBundleStateSpec": schema_pkg_apis_k8splugin_v1alpha1_ResourceBundleStateSpec(ref),
		"./pkg/apis/k8splugin/v1alpha1.ResourceBundleStatus":    schema_pkg_apis_k8splugin_v1alpha1_ResourceBundleStatus(ref),
	}
}

func schema_pkg_apis_k8splugin_v1alpha1_PodStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodStatus defines the observed state of ResourceBundleState",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.PodStatus"),
						},
					},
				},
				Required: []string{"ready"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.PodStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_k8splugin_v1alpha1_ResourceBundleState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceBundleState is the Schema for the ResourceBundleStatees API",
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
							Ref: ref("./pkg/apis/k8splugin/v1alpha1.ResourceBundleStateSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/k8splugin/v1alpha1.ResourceBundleStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/k8splugin/v1alpha1.ResourceBundleStateSpec", "./pkg/apis/k8splugin/v1alpha1.ResourceBundleStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_k8splugin_v1alpha1_ResourceBundleStateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceBundleStateSpec defines the desired state of ResourceBundleState",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"selector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
				},
				Required: []string{"selector"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_k8splugin_v1alpha1_ResourceBundleStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceBundleStatus defines the observed state of ResourceBundleState",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"ready": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"resourceCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"podStatuses": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/k8splugin/v1alpha1.PodStatus"),
									},
								},
							},
						},
					},
					"serviceStatuses": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Service"),
									},
								},
							},
						},
					},
					"configMapStatuses": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ConfigMap"),
									},
								},
							},
						},
					},
				},
				Required: []string{"ready", "resourceCount", "podStatuses", "serviceStatuses"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/k8splugin/v1alpha1.PodStatus", "k8s.io/api/core/v1.ConfigMap", "k8s.io/api/core/v1.Service"},
	}
}
