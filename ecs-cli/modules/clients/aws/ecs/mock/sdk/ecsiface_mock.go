// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/aws-sdk-go/service/ecs/ecsiface (interfaces: ECSAPI)

package mock_ecsiface

import (
	request "github.com/aws/aws-sdk-go/aws/request"
	ecs "github.com/aws/aws-sdk-go/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ECSAPI interface
type MockECSAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockECSAPIRecorder
}

// Recorder for MockECSAPI (not exported)
type _MockECSAPIRecorder struct {
	mock *MockECSAPI
}

func NewMockECSAPI(ctrl *gomock.Controller) *MockECSAPI {
	mock := &MockECSAPI{ctrl: ctrl}
	mock.recorder = &_MockECSAPIRecorder{mock}
	return mock
}

func (_m *MockECSAPI) EXPECT() *_MockECSAPIRecorder {
	return _m.recorder
}

func (_m *MockECSAPI) CreateCluster(_param0 *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateCluster", _param0)
	ret0, _ := ret[0].(*ecs.CreateClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateCluster", arg0)
}

func (_m *MockECSAPI) CreateClusterRequest(_param0 *ecs.CreateClusterInput) (*request.Request, *ecs.CreateClusterOutput) {
	ret := _m.ctrl.Call(_m, "CreateClusterRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.CreateClusterOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) CreateClusterRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateClusterRequest", arg0)
}

func (_m *MockECSAPI) CreateService(_param0 *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateService", _param0)
	ret0, _ := ret[0].(*ecs.CreateServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) CreateService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateService", arg0)
}

func (_m *MockECSAPI) CreateServiceRequest(_param0 *ecs.CreateServiceInput) (*request.Request, *ecs.CreateServiceOutput) {
	ret := _m.ctrl.Call(_m, "CreateServiceRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.CreateServiceOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) CreateServiceRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateServiceRequest", arg0)
}

func (_m *MockECSAPI) DeleteAttributes(_param0 *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteAttributes", _param0)
	ret0, _ := ret[0].(*ecs.DeleteAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteAttributes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAttributes", arg0)
}

func (_m *MockECSAPI) DeleteAttributesRequest(_param0 *ecs.DeleteAttributesInput) (*request.Request, *ecs.DeleteAttributesOutput) {
	ret := _m.ctrl.Call(_m, "DeleteAttributesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DeleteAttributesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteAttributesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAttributesRequest", arg0)
}

func (_m *MockECSAPI) DeleteCluster(_param0 *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteCluster", _param0)
	ret0, _ := ret[0].(*ecs.DeleteClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCluster", arg0)
}

func (_m *MockECSAPI) DeleteClusterRequest(_param0 *ecs.DeleteClusterInput) (*request.Request, *ecs.DeleteClusterOutput) {
	ret := _m.ctrl.Call(_m, "DeleteClusterRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DeleteClusterOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteClusterRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteClusterRequest", arg0)
}

func (_m *MockECSAPI) DeleteService(_param0 *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteService", _param0)
	ret0, _ := ret[0].(*ecs.DeleteServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteService", arg0)
}

func (_m *MockECSAPI) DeleteServiceRequest(_param0 *ecs.DeleteServiceInput) (*request.Request, *ecs.DeleteServiceOutput) {
	ret := _m.ctrl.Call(_m, "DeleteServiceRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DeleteServiceOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteServiceRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteServiceRequest", arg0)
}

func (_m *MockECSAPI) DeregisterContainerInstance(_param0 *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error) {
	ret := _m.ctrl.Call(_m, "DeregisterContainerInstance", _param0)
	ret0, _ := ret[0].(*ecs.DeregisterContainerInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeregisterContainerInstance(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterContainerInstance", arg0)
}

func (_m *MockECSAPI) DeregisterContainerInstanceRequest(_param0 *ecs.DeregisterContainerInstanceInput) (*request.Request, *ecs.DeregisterContainerInstanceOutput) {
	ret := _m.ctrl.Call(_m, "DeregisterContainerInstanceRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DeregisterContainerInstanceOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeregisterContainerInstanceRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterContainerInstanceRequest", arg0)
}

func (_m *MockECSAPI) DeregisterTaskDefinition(_param0 *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error) {
	ret := _m.ctrl.Call(_m, "DeregisterTaskDefinition", _param0)
	ret0, _ := ret[0].(*ecs.DeregisterTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeregisterTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterTaskDefinition", arg0)
}

func (_m *MockECSAPI) DeregisterTaskDefinitionRequest(_param0 *ecs.DeregisterTaskDefinitionInput) (*request.Request, *ecs.DeregisterTaskDefinitionOutput) {
	ret := _m.ctrl.Call(_m, "DeregisterTaskDefinitionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DeregisterTaskDefinitionOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeregisterTaskDefinitionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterTaskDefinitionRequest", arg0)
}

func (_m *MockECSAPI) DescribeClusters(_param0 *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeClusters", _param0)
	ret0, _ := ret[0].(*ecs.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeClusters(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeClusters", arg0)
}

func (_m *MockECSAPI) DescribeClustersRequest(_param0 *ecs.DescribeClustersInput) (*request.Request, *ecs.DescribeClustersOutput) {
	ret := _m.ctrl.Call(_m, "DescribeClustersRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DescribeClustersOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeClustersRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeClustersRequest", arg0)
}

func (_m *MockECSAPI) DescribeContainerInstances(_param0 *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeContainerInstances", _param0)
	ret0, _ := ret[0].(*ecs.DescribeContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeContainerInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeContainerInstances", arg0)
}

func (_m *MockECSAPI) DescribeContainerInstancesRequest(_param0 *ecs.DescribeContainerInstancesInput) (*request.Request, *ecs.DescribeContainerInstancesOutput) {
	ret := _m.ctrl.Call(_m, "DescribeContainerInstancesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DescribeContainerInstancesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeContainerInstancesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeContainerInstancesRequest", arg0)
}

func (_m *MockECSAPI) DescribeServices(_param0 *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeServices", _param0)
	ret0, _ := ret[0].(*ecs.DescribeServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeServices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeServices", arg0)
}

func (_m *MockECSAPI) DescribeServicesRequest(_param0 *ecs.DescribeServicesInput) (*request.Request, *ecs.DescribeServicesOutput) {
	ret := _m.ctrl.Call(_m, "DescribeServicesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DescribeServicesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeServicesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeServicesRequest", arg0)
}

func (_m *MockECSAPI) DescribeTaskDefinition(_param0 *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeTaskDefinition", _param0)
	ret0, _ := ret[0].(*ecs.DescribeTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTaskDefinition", arg0)
}

func (_m *MockECSAPI) DescribeTaskDefinitionRequest(_param0 *ecs.DescribeTaskDefinitionInput) (*request.Request, *ecs.DescribeTaskDefinitionOutput) {
	ret := _m.ctrl.Call(_m, "DescribeTaskDefinitionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DescribeTaskDefinitionOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeTaskDefinitionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTaskDefinitionRequest", arg0)
}

func (_m *MockECSAPI) DescribeTasks(_param0 *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeTasks", _param0)
	ret0, _ := ret[0].(*ecs.DescribeTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeTasks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTasks", arg0)
}

func (_m *MockECSAPI) DescribeTasksRequest(_param0 *ecs.DescribeTasksInput) (*request.Request, *ecs.DescribeTasksOutput) {
	ret := _m.ctrl.Call(_m, "DescribeTasksRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DescribeTasksOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeTasksRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTasksRequest", arg0)
}

func (_m *MockECSAPI) DiscoverPollEndpoint(_param0 *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
	ret := _m.ctrl.Call(_m, "DiscoverPollEndpoint", _param0)
	ret0, _ := ret[0].(*ecs.DiscoverPollEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DiscoverPollEndpoint(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiscoverPollEndpoint", arg0)
}

func (_m *MockECSAPI) DiscoverPollEndpointRequest(_param0 *ecs.DiscoverPollEndpointInput) (*request.Request, *ecs.DiscoverPollEndpointOutput) {
	ret := _m.ctrl.Call(_m, "DiscoverPollEndpointRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.DiscoverPollEndpointOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DiscoverPollEndpointRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiscoverPollEndpointRequest", arg0)
}

func (_m *MockECSAPI) ListAttributes(_param0 *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListAttributes", _param0)
	ret0, _ := ret[0].(*ecs.ListAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListAttributes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAttributes", arg0)
}

func (_m *MockECSAPI) ListAttributesRequest(_param0 *ecs.ListAttributesInput) (*request.Request, *ecs.ListAttributesOutput) {
	ret := _m.ctrl.Call(_m, "ListAttributesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.ListAttributesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListAttributesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAttributesRequest", arg0)
}

func (_m *MockECSAPI) ListClusters(_param0 *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	ret := _m.ctrl.Call(_m, "ListClusters", _param0)
	ret0, _ := ret[0].(*ecs.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListClusters(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListClusters", arg0)
}

func (_m *MockECSAPI) ListClustersPages(_param0 *ecs.ListClustersInput, _param1 func(*ecs.ListClustersOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListClustersPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) ListClustersPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListClustersPages", arg0, arg1)
}

func (_m *MockECSAPI) ListClustersRequest(_param0 *ecs.ListClustersInput) (*request.Request, *ecs.ListClustersOutput) {
	ret := _m.ctrl.Call(_m, "ListClustersRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.ListClustersOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListClustersRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListClustersRequest", arg0)
}

func (_m *MockECSAPI) ListContainerInstances(_param0 *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListContainerInstances", _param0)
	ret0, _ := ret[0].(*ecs.ListContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListContainerInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainerInstances", arg0)
}

func (_m *MockECSAPI) ListContainerInstancesPages(_param0 *ecs.ListContainerInstancesInput, _param1 func(*ecs.ListContainerInstancesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListContainerInstancesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) ListContainerInstancesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainerInstancesPages", arg0, arg1)
}

func (_m *MockECSAPI) ListContainerInstancesRequest(_param0 *ecs.ListContainerInstancesInput) (*request.Request, *ecs.ListContainerInstancesOutput) {
	ret := _m.ctrl.Call(_m, "ListContainerInstancesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.ListContainerInstancesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListContainerInstancesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainerInstancesRequest", arg0)
}

func (_m *MockECSAPI) ListServices(_param0 *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListServices", _param0)
	ret0, _ := ret[0].(*ecs.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListServices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListServices", arg0)
}

func (_m *MockECSAPI) ListServicesPages(_param0 *ecs.ListServicesInput, _param1 func(*ecs.ListServicesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListServicesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) ListServicesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListServicesPages", arg0, arg1)
}

func (_m *MockECSAPI) ListServicesRequest(_param0 *ecs.ListServicesInput) (*request.Request, *ecs.ListServicesOutput) {
	ret := _m.ctrl.Call(_m, "ListServicesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.ListServicesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListServicesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListServicesRequest", arg0)
}

func (_m *MockECSAPI) ListTaskDefinitionFamilies(_param0 *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionFamilies", _param0)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionFamiliesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitionFamilies(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionFamilies", arg0)
}

func (_m *MockECSAPI) ListTaskDefinitionFamiliesPages(_param0 *ecs.ListTaskDefinitionFamiliesInput, _param1 func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionFamiliesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitionFamiliesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionFamiliesPages", arg0, arg1)
}

func (_m *MockECSAPI) ListTaskDefinitionFamiliesRequest(_param0 *ecs.ListTaskDefinitionFamiliesInput) (*request.Request, *ecs.ListTaskDefinitionFamiliesOutput) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionFamiliesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.ListTaskDefinitionFamiliesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitionFamiliesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionFamiliesRequest", arg0)
}

func (_m *MockECSAPI) ListTaskDefinitions(_param0 *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitions", _param0)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitions", arg0)
}

func (_m *MockECSAPI) ListTaskDefinitionsPages(_param0 *ecs.ListTaskDefinitionsInput, _param1 func(*ecs.ListTaskDefinitionsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitionsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionsPages", arg0, arg1)
}

func (_m *MockECSAPI) ListTaskDefinitionsRequest(_param0 *ecs.ListTaskDefinitionsInput) (*request.Request, *ecs.ListTaskDefinitionsOutput) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.ListTaskDefinitionsOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitionsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionsRequest", arg0)
}

func (_m *MockECSAPI) ListTasks(_param0 *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
	ret := _m.ctrl.Call(_m, "ListTasks", _param0)
	ret0, _ := ret[0].(*ecs.ListTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTasks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasks", arg0)
}

func (_m *MockECSAPI) ListTasksPages(_param0 *ecs.ListTasksInput, _param1 func(*ecs.ListTasksOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListTasksPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) ListTasksPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasksPages", arg0, arg1)
}

func (_m *MockECSAPI) ListTasksRequest(_param0 *ecs.ListTasksInput) (*request.Request, *ecs.ListTasksOutput) {
	ret := _m.ctrl.Call(_m, "ListTasksRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.ListTasksOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTasksRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasksRequest", arg0)
}

func (_m *MockECSAPI) PutAttributes(_param0 *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error) {
	ret := _m.ctrl.Call(_m, "PutAttributes", _param0)
	ret0, _ := ret[0].(*ecs.PutAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) PutAttributes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutAttributes", arg0)
}

func (_m *MockECSAPI) PutAttributesRequest(_param0 *ecs.PutAttributesInput) (*request.Request, *ecs.PutAttributesOutput) {
	ret := _m.ctrl.Call(_m, "PutAttributesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.PutAttributesOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) PutAttributesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutAttributesRequest", arg0)
}

func (_m *MockECSAPI) RegisterContainerInstance(_param0 *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
	ret := _m.ctrl.Call(_m, "RegisterContainerInstance", _param0)
	ret0, _ := ret[0].(*ecs.RegisterContainerInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RegisterContainerInstance(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterContainerInstance", arg0)
}

func (_m *MockECSAPI) RegisterContainerInstanceRequest(_param0 *ecs.RegisterContainerInstanceInput) (*request.Request, *ecs.RegisterContainerInstanceOutput) {
	ret := _m.ctrl.Call(_m, "RegisterContainerInstanceRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.RegisterContainerInstanceOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RegisterContainerInstanceRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterContainerInstanceRequest", arg0)
}

func (_m *MockECSAPI) RegisterTaskDefinition(_param0 *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
	ret := _m.ctrl.Call(_m, "RegisterTaskDefinition", _param0)
	ret0, _ := ret[0].(*ecs.RegisterTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RegisterTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterTaskDefinition", arg0)
}

func (_m *MockECSAPI) RegisterTaskDefinitionRequest(_param0 *ecs.RegisterTaskDefinitionInput) (*request.Request, *ecs.RegisterTaskDefinitionOutput) {
	ret := _m.ctrl.Call(_m, "RegisterTaskDefinitionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.RegisterTaskDefinitionOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RegisterTaskDefinitionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterTaskDefinitionRequest", arg0)
}

func (_m *MockECSAPI) RunTask(_param0 *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "RunTask", _param0)
	ret0, _ := ret[0].(*ecs.RunTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RunTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RunTask", arg0)
}

func (_m *MockECSAPI) RunTaskRequest(_param0 *ecs.RunTaskInput) (*request.Request, *ecs.RunTaskOutput) {
	ret := _m.ctrl.Call(_m, "RunTaskRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.RunTaskOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RunTaskRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RunTaskRequest", arg0)
}

func (_m *MockECSAPI) StartTask(_param0 *ecs.StartTaskInput) (*ecs.StartTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "StartTask", _param0)
	ret0, _ := ret[0].(*ecs.StartTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) StartTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTask", arg0)
}

func (_m *MockECSAPI) StartTaskRequest(_param0 *ecs.StartTaskInput) (*request.Request, *ecs.StartTaskOutput) {
	ret := _m.ctrl.Call(_m, "StartTaskRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.StartTaskOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) StartTaskRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTaskRequest", arg0)
}

func (_m *MockECSAPI) StopTask(_param0 *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "StopTask", _param0)
	ret0, _ := ret[0].(*ecs.StopTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) StopTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopTask", arg0)
}

func (_m *MockECSAPI) StopTaskRequest(_param0 *ecs.StopTaskInput) (*request.Request, *ecs.StopTaskOutput) {
	ret := _m.ctrl.Call(_m, "StopTaskRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.StopTaskOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) StopTaskRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopTaskRequest", arg0)
}

func (_m *MockECSAPI) SubmitContainerStateChange(_param0 *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
	ret := _m.ctrl.Call(_m, "SubmitContainerStateChange", _param0)
	ret0, _ := ret[0].(*ecs.SubmitContainerStateChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) SubmitContainerStateChange(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubmitContainerStateChange", arg0)
}

func (_m *MockECSAPI) SubmitContainerStateChangeRequest(_param0 *ecs.SubmitContainerStateChangeInput) (*request.Request, *ecs.SubmitContainerStateChangeOutput) {
	ret := _m.ctrl.Call(_m, "SubmitContainerStateChangeRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.SubmitContainerStateChangeOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) SubmitContainerStateChangeRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubmitContainerStateChangeRequest", arg0)
}

func (_m *MockECSAPI) SubmitTaskStateChange(_param0 *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
	ret := _m.ctrl.Call(_m, "SubmitTaskStateChange", _param0)
	ret0, _ := ret[0].(*ecs.SubmitTaskStateChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) SubmitTaskStateChange(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubmitTaskStateChange", arg0)
}

func (_m *MockECSAPI) SubmitTaskStateChangeRequest(_param0 *ecs.SubmitTaskStateChangeInput) (*request.Request, *ecs.SubmitTaskStateChangeOutput) {
	ret := _m.ctrl.Call(_m, "SubmitTaskStateChangeRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.SubmitTaskStateChangeOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) SubmitTaskStateChangeRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubmitTaskStateChangeRequest", arg0)
}

func (_m *MockECSAPI) UpdateContainerAgent(_param0 *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateContainerAgent", _param0)
	ret0, _ := ret[0].(*ecs.UpdateContainerAgentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateContainerAgent(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateContainerAgent", arg0)
}

func (_m *MockECSAPI) UpdateContainerAgentRequest(_param0 *ecs.UpdateContainerAgentInput) (*request.Request, *ecs.UpdateContainerAgentOutput) {
	ret := _m.ctrl.Call(_m, "UpdateContainerAgentRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.UpdateContainerAgentOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateContainerAgentRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateContainerAgentRequest", arg0)
}

func (_m *MockECSAPI) UpdateContainerInstancesState(_param0 *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateContainerInstancesState", _param0)
	ret0, _ := ret[0].(*ecs.UpdateContainerInstancesStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateContainerInstancesState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateContainerInstancesState", arg0)
}

func (_m *MockECSAPI) UpdateContainerInstancesStateRequest(_param0 *ecs.UpdateContainerInstancesStateInput) (*request.Request, *ecs.UpdateContainerInstancesStateOutput) {
	ret := _m.ctrl.Call(_m, "UpdateContainerInstancesStateRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.UpdateContainerInstancesStateOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateContainerInstancesStateRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateContainerInstancesStateRequest", arg0)
}

func (_m *MockECSAPI) UpdateService(_param0 *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateService", _param0)
	ret0, _ := ret[0].(*ecs.UpdateServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateService", arg0)
}

func (_m *MockECSAPI) UpdateServiceRequest(_param0 *ecs.UpdateServiceInput) (*request.Request, *ecs.UpdateServiceOutput) {
	ret := _m.ctrl.Call(_m, "UpdateServiceRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecs.UpdateServiceOutput)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateServiceRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateServiceRequest", arg0)
}

func (_m *MockECSAPI) WaitUntilServicesInactive(_param0 *ecs.DescribeServicesInput) error {
	ret := _m.ctrl.Call(_m, "WaitUntilServicesInactive", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) WaitUntilServicesInactive(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitUntilServicesInactive", arg0)
}

func (_m *MockECSAPI) WaitUntilServicesStable(_param0 *ecs.DescribeServicesInput) error {
	ret := _m.ctrl.Call(_m, "WaitUntilServicesStable", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) WaitUntilServicesStable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitUntilServicesStable", arg0)
}

func (_m *MockECSAPI) WaitUntilTasksRunning(_param0 *ecs.DescribeTasksInput) error {
	ret := _m.ctrl.Call(_m, "WaitUntilTasksRunning", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) WaitUntilTasksRunning(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitUntilTasksRunning", arg0)
}

func (_m *MockECSAPI) WaitUntilTasksStopped(_param0 *ecs.DescribeTasksInput) error {
	ret := _m.ctrl.Call(_m, "WaitUntilTasksStopped", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSAPIRecorder) WaitUntilTasksStopped(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitUntilTasksStopped", arg0)
}
