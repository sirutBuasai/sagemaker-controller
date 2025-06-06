// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EndpointConfigSpec defines the desired state of EndpointConfig.
type EndpointConfigSpec struct {

	// Specifies configuration for how an endpoint performs asynchronous inference.
	// This is a required field in order for your Endpoint to be invoked using InvokeEndpointAsync
	// (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpointAsync.html).
	AsyncInferenceConfig *AsyncInferenceConfig `json:"asyncInferenceConfig,omitempty"`
	DataCaptureConfig    *DataCaptureConfig    `json:"dataCaptureConfig,omitempty"`
	// Sets whether all model containers deployed to the endpoint are isolated.
	// If they are, no inbound or outbound network calls can be made to or from
	// the model containers.
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation,omitempty"`
	// The name of the endpoint configuration. You specify this name in a CreateEndpoint
	// (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html)
	// request.
	//
	// Regex Pattern: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`
	// +kubebuilder:validation:Required
	EndpointConfigName *string `json:"endpointConfigName"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform actions on your behalf. For more information, see SageMaker Roles
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).
	//
	// To be able to pass this role to Amazon SageMaker, the caller of this action
	// must have the iam:PassRole permission.
	//
	// Regex Pattern: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
	ExecutionRoleARN *string `json:"executionRoleARN,omitempty"`
	// The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service
	// key that SageMaker uses to encrypt data on the storage volume attached to
	// the ML compute instance that hosts the endpoint.
	//
	// The KmsKeyId can be any of the following formats:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Alias name: alias/ExampleAlias
	//
	//   - Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias
	//
	// The KMS key policy must grant permission to the IAM role that you specify
	// in your CreateEndpoint, UpdateEndpoint requests. For more information, refer
	// to the Amazon Web Services Key Management Service section Using Key Policies
	// in Amazon Web Services KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)
	//
	// Certain Nitro-based instances include local storage, dependent on the instance
	// type. Local storage volumes are encrypted using a hardware module on the
	// instance. You can't request a KmsKeyId when using an instance type with local
	// storage. If any of the models that you specify in the ProductionVariants
	// parameter use nitro-based instances with local storage, do not specify a
	// value for the KmsKeyId parameter. If you specify a value for KmsKeyId when
	// using any nitro-based instances with local storage, the call to CreateEndpointConfig
	// fails.
	//
	// For a list of instance types that support local instance storage, see Instance
	// Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes).
	//
	// For more information about local instance storage encryption, see SSD Instance
	// Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html).
	//
	// Regex Pattern: `^[a-zA-Z0-9:/_-]*$`
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// An array of ProductionVariant objects, one for each model that you want to
	// host at this endpoint.
	// +kubebuilder:validation:Required
	ProductionVariants []*ProductionVariant `json:"productionVariants"`
	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags      []*Tag     `json:"tags,omitempty"`
	VPCConfig *VPCConfig `json:"vpcConfig,omitempty"`
}

// EndpointConfigStatus defines the observed state of EndpointConfig
type EndpointConfigStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// EndpointConfig is the Schema for the EndpointConfigs API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type EndpointConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointConfigSpec   `json:"spec,omitempty"`
	Status            EndpointConfigStatus `json:"status,omitempty"`
}

// EndpointConfigList contains a list of EndpointConfig
// +kubebuilder:object:root=true
type EndpointConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EndpointConfig{}, &EndpointConfigList{})
}
