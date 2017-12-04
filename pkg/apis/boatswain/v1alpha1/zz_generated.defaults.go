// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Cluster{}, func(obj interface{}) { SetObjectDefaults_Cluster(obj.(*Cluster)) })
	scheme.AddTypeDefaultingFunc(&ClusterList{}, func(obj interface{}) { SetObjectDefaults_ClusterList(obj.(*ClusterList)) })
	scheme.AddTypeDefaultingFunc(&Node{}, func(obj interface{}) { SetObjectDefaults_Node(obj.(*Node)) })
	scheme.AddTypeDefaultingFunc(&NodeGroup{}, func(obj interface{}) { SetObjectDefaults_NodeGroup(obj.(*NodeGroup)) })
	scheme.AddTypeDefaultingFunc(&NodeGroupList{}, func(obj interface{}) { SetObjectDefaults_NodeGroupList(obj.(*NodeGroupList)) })
	scheme.AddTypeDefaultingFunc(&NodeList{}, func(obj interface{}) { SetObjectDefaults_NodeList(obj.(*NodeList)) })
	return nil
}

func SetObjectDefaults_Cluster(in *Cluster) {
	SetDefaults_ClusterSpec(&in.Spec)
}

func SetObjectDefaults_ClusterList(in *ClusterList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Cluster(a)
	}
}

func SetObjectDefaults_Node(in *Node) {
	SetDefaults_NodeSpec(&in.Spec)
}

func SetObjectDefaults_NodeGroup(in *NodeGroup) {
	SetDefaults_NodeGroupSpec(&in.Spec)
}

func SetObjectDefaults_NodeGroupList(in *NodeGroupList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_NodeGroup(a)
	}
}

func SetObjectDefaults_NodeList(in *NodeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Node(a)
	}
}
