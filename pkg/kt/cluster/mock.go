// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/cluster/types.go

// Package cluster is a generated GoMock package.
package cluster

import (
	context "context"
	reflect "reflect"

	options "github.com/alibaba/kt-connect/pkg/kt/options"
	util "github.com/alibaba/kt-connect/pkg/kt/util"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
)

// MockKubernetesInterface is a mock of KubernetesInterface interface.
type MockKubernetesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesInterfaceMockRecorder
}

// MockKubernetesInterfaceMockRecorder is the mock recorder for MockKubernetesInterface.
type MockKubernetesInterfaceMockRecorder struct {
	mock *MockKubernetesInterface
}

// NewMockKubernetesInterface creates a new mock instance.
func NewMockKubernetesInterface(ctrl *gomock.Controller) *MockKubernetesInterface {
	mock := &MockKubernetesInterface{ctrl: ctrl}
	mock.recorder = &MockKubernetesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesInterface) EXPECT() *MockKubernetesInterfaceMockRecorder {
	return m.recorder
}

// AddEphemeralContainer mocks base method.
func (m *MockKubernetesInterface) AddEphemeralContainer(ctx context.Context, containerName, podName string, options *options.DaemonOptions, envs map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEphemeralContainer", ctx, containerName, podName, options, envs)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddEphemeralContainer indicates an expected call of AddEphemeralContainer.
func (mr *MockKubernetesInterfaceMockRecorder) AddEphemeralContainer(ctx, containerName, podName, options, envs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEphemeralContainer", reflect.TypeOf((*MockKubernetesInterface)(nil).AddEphemeralContainer), ctx, containerName, podName, options, envs)
}

// ClusterCidrs mocks base method.
func (m *MockKubernetesInterface) ClusterCidrs(ctx context.Context, namespace string, connectOptions *options.ConnectOptions) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCidrs", ctx, namespace, connectOptions)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterCidrs indicates an expected call of ClusterCidrs.
func (mr *MockKubernetesInterfaceMockRecorder) ClusterCidrs(ctx, namespace, connectOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCidrs", reflect.TypeOf((*MockKubernetesInterface)(nil).ClusterCidrs), ctx, namespace, connectOptions)
}

// CreateService mocks base method.
func (m *MockKubernetesInterface) CreateService(ctx context.Context, name, namespace string, external bool, port int, labels map[string]string) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", ctx, name, namespace, external, port, labels)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockKubernetesInterfaceMockRecorder) CreateService(ctx, name, namespace, external, port, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockKubernetesInterface)(nil).CreateService), ctx, name, namespace, external, port, labels)
}

// DecreaseRef mocks base method.
func (m *MockKubernetesInterface) DecreaseRef(ctx context.Context, namespace, deployment string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseRef", ctx, namespace, deployment)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecreaseRef indicates an expected call of DecreaseRef.
func (mr *MockKubernetesInterfaceMockRecorder) DecreaseRef(ctx, namespace, deployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseRef", reflect.TypeOf((*MockKubernetesInterface)(nil).DecreaseRef), ctx, namespace, deployment)
}

// DeletePod mocks base method.
func (m *MockKubernetesInterface) DeletePod(ctx context.Context, podName, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePod", ctx, podName, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePod indicates an expected call of DeletePod.
func (mr *MockKubernetesInterfaceMockRecorder) DeletePod(ctx, podName, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePod", reflect.TypeOf((*MockKubernetesInterface)(nil).DeletePod), ctx, podName, namespace)
}

// Deployment mocks base method.
func (m *MockKubernetesInterface) Deployment(ctx context.Context, name, namespace string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployment", ctx, name, namespace)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deployment indicates an expected call of Deployment.
func (mr *MockKubernetesInterfaceMockRecorder) Deployment(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockKubernetesInterface)(nil).Deployment), ctx, name, namespace)
}

// GetAllExistingShadowPods mocks base method.
func (m *MockKubernetesInterface) GetAllExistingShadowPods(ctx context.Context, namespace string) ([]v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllExistingShadowPods", ctx, namespace)
	ret0, _ := ret[0].([]v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllExistingShadowPods indicates an expected call of GetAllExistingShadowPods.
func (mr *MockKubernetesInterfaceMockRecorder) GetAllExistingShadowPods(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllExistingShadowPods", reflect.TypeOf((*MockKubernetesInterface)(nil).GetAllExistingShadowPods), ctx, namespace)
}

// GetDeployment mocks base method.
func (m *MockKubernetesInterface) GetDeployment(ctx context.Context, name, namespace string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", ctx, name, namespace)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockKubernetesInterfaceMockRecorder) GetDeployment(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockKubernetesInterface)(nil).GetDeployment), ctx, name, namespace)
}

// GetOrCreateShadow mocks base method.
func (m *MockKubernetesInterface) GetOrCreateShadow(ctx context.Context, name string, options *options.DaemonOptions, labels, annotations, envs map[string]string) (string, string, string, *util.SSHCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateShadow", ctx, name, options, labels, annotations, envs)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(*util.SSHCredential)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// GetOrCreateShadow indicates an expected call of GetOrCreateShadow.
func (mr *MockKubernetesInterfaceMockRecorder) GetOrCreateShadow(ctx, name, options, labels, annotations, envs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateShadow", reflect.TypeOf((*MockKubernetesInterface)(nil).GetOrCreateShadow), ctx, name, options, labels, annotations, envs)
}

// GetPod mocks base method.
func (m *MockKubernetesInterface) GetPod(ctx context.Context, name, namespace string) (*v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", ctx, name, namespace)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPod indicates an expected call of GetPod.
func (mr *MockKubernetesInterfaceMockRecorder) GetPod(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockKubernetesInterface)(nil).GetPod), ctx, name, namespace)
}

// Pod mocks base method.
func (m *MockKubernetesInterface) Pod(ctx context.Context, name, namespace string) (*v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pod", ctx, name, namespace)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pod indicates an expected call of Pod.
func (mr *MockKubernetesInterfaceMockRecorder) Pod(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pod", reflect.TypeOf((*MockKubernetesInterface)(nil).Pod), ctx, name, namespace)
}

// Pods mocks base method.
func (m *MockKubernetesInterface) Pods(ctx context.Context, label, namespace string) (*v10.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pods", ctx, label, namespace)
	ret0, _ := ret[0].(*v10.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pods indicates an expected call of Pods.
func (mr *MockKubernetesInterfaceMockRecorder) Pods(ctx, label, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pods", reflect.TypeOf((*MockKubernetesInterface)(nil).Pods), ctx, label, namespace)
}

// RemoveConfigMap mocks base method.
func (m *MockKubernetesInterface) RemoveConfigMap(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveConfigMap", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveConfigMap indicates an expected call of RemoveConfigMap.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveConfigMap(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConfigMap", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveConfigMap), ctx, name, namespace)
}

// RemovePod mocks base method.
func (m *MockKubernetesInterface) RemovePod(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePod", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePod indicates an expected call of RemovePod.
func (mr *MockKubernetesInterfaceMockRecorder) RemovePod(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePod", reflect.TypeOf((*MockKubernetesInterface)(nil).RemovePod), ctx, name, namespace)
}

// RemoveService mocks base method.
func (m *MockKubernetesInterface) RemoveService(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveService", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveService indicates an expected call of RemoveService.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveService(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveService", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveService), ctx, name, namespace)
}

// Scale mocks base method.
func (m *MockKubernetesInterface) Scale(ctx context.Context, deployment *v1.Deployment, replicas *int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", ctx, deployment, replicas)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scale indicates an expected call of Scale.
func (mr *MockKubernetesInterfaceMockRecorder) Scale(ctx, deployment, replicas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockKubernetesInterface)(nil).Scale), ctx, deployment, replicas)
}

// ScaleTo mocks base method.
func (m *MockKubernetesInterface) ScaleTo(ctx context.Context, deployment, namespace string, replicas *int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleTo", ctx, deployment, namespace, replicas)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleTo indicates an expected call of ScaleTo.
func (mr *MockKubernetesInterfaceMockRecorder) ScaleTo(ctx, deployment, namespace, replicas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleTo", reflect.TypeOf((*MockKubernetesInterface)(nil).ScaleTo), ctx, deployment, namespace, replicas)
}

// ServiceHosts mocks base method.
func (m *MockKubernetesInterface) ServiceHosts(ctx context.Context, namespace string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceHosts", ctx, namespace)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// ServiceHosts indicates an expected call of ServiceHosts.
func (mr *MockKubernetesInterfaceMockRecorder) ServiceHosts(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceHosts", reflect.TypeOf((*MockKubernetesInterface)(nil).ServiceHosts), ctx, namespace)
}

// UpdatePod mocks base method.
func (m *MockKubernetesInterface) UpdatePod(ctx context.Context, namespace string, pod *v10.Pod) (*v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePod", ctx, namespace, pod)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePod indicates an expected call of UpdatePod.
func (mr *MockKubernetesInterfaceMockRecorder) UpdatePod(ctx, namespace, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePod", reflect.TypeOf((*MockKubernetesInterface)(nil).UpdatePod), ctx, namespace, pod)
}
