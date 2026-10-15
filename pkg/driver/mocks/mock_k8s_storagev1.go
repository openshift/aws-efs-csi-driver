// Source: k8s.io/client-go/kubernetes/typed/storage/v1 (interfaces: StorageV1Interface,CSINodeInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1 "k8s.io/client-go/applyconfigurations/storage/v1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	rest "k8s.io/client-go/rest"
)

// MockStorageV1Interface is a mock of StorageV1Interface interface.
type MockStorageV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageV1InterfaceMockRecorder
}

// MockStorageV1InterfaceMockRecorder is the mock recorder for MockStorageV1Interface.
type MockStorageV1InterfaceMockRecorder struct {
	mock *MockStorageV1Interface
}

// NewMockStorageV1Interface creates a new mock instance.
func NewMockStorageV1Interface(ctrl *gomock.Controller) *MockStorageV1Interface {
	mock := &MockStorageV1Interface{ctrl: ctrl}
	mock.recorder = &MockStorageV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageV1Interface) EXPECT() *MockStorageV1InterfaceMockRecorder {
	return m.recorder
}

// CSIDrivers mocks base method.
func (m *MockStorageV1Interface) CSIDrivers() storagev1.CSIDriverInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CSIDrivers")
	ret0, _ := ret[0].(storagev1.CSIDriverInterface)
	return ret0
}

// CSIDrivers indicates an expected call of CSIDrivers.
func (mr *MockStorageV1InterfaceMockRecorder) CSIDrivers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSIDrivers", reflect.TypeOf((*MockStorageV1Interface)(nil).CSIDrivers))
}

// CSINodes mocks base method.
func (m *MockStorageV1Interface) CSINodes() storagev1.CSINodeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CSINodes")
	ret0, _ := ret[0].(storagev1.CSINodeInterface)
	return ret0
}

// CSINodes indicates an expected call of CSINodes.
func (mr *MockStorageV1InterfaceMockRecorder) CSINodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSINodes", reflect.TypeOf((*MockStorageV1Interface)(nil).CSINodes))
}

// CSIStorageCapacities mocks base method.
func (m *MockStorageV1Interface) CSIStorageCapacities(arg0 string) storagev1.CSIStorageCapacityInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CSIStorageCapacities", arg0)
	ret0, _ := ret[0].(storagev1.CSIStorageCapacityInterface)
	return ret0
}

// CSIStorageCapacities indicates an expected call of CSIStorageCapacities.
func (mr *MockStorageV1InterfaceMockRecorder) CSIStorageCapacities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSIStorageCapacities", reflect.TypeOf((*MockStorageV1Interface)(nil).CSIStorageCapacities), arg0)
}

// RESTClient mocks base method.
func (m *MockStorageV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockStorageV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockStorageV1Interface)(nil).RESTClient))
}

// StorageClasses mocks base method.
func (m *MockStorageV1Interface) StorageClasses() storagev1.StorageClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageClasses")
	ret0, _ := ret[0].(storagev1.StorageClassInterface)
	return ret0
}

// StorageClasses indicates an expected call of StorageClasses.
func (mr *MockStorageV1InterfaceMockRecorder) StorageClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageClasses", reflect.TypeOf((*MockStorageV1Interface)(nil).StorageClasses))
}

// VolumeAttachments mocks base method.
func (m *MockStorageV1Interface) VolumeAttachments() storagev1.VolumeAttachmentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeAttachments")
	ret0, _ := ret[0].(storagev1.VolumeAttachmentInterface)
	return ret0
}

// VolumeAttachments indicates an expected call of VolumeAttachments.
func (mr *MockStorageV1InterfaceMockRecorder) VolumeAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeAttachments", reflect.TypeOf((*MockStorageV1Interface)(nil).VolumeAttachments))
}

// MockCSINodeInterface is a mock of CSINodeInterface interface.
type MockCSINodeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCSINodeInterfaceMockRecorder
}

// MockCSINodeInterfaceMockRecorder is the mock recorder for MockCSINodeInterface.
type MockCSINodeInterfaceMockRecorder struct {
	mock *MockCSINodeInterface
}

// NewMockCSINodeInterface creates a new mock instance.
func NewMockCSINodeInterface(ctrl *gomock.Controller) *MockCSINodeInterface {
	mock := &MockCSINodeInterface{ctrl: ctrl}
	mock.recorder = &MockCSINodeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCSINodeInterface) EXPECT() *MockCSINodeInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockCSINodeInterface) Apply(arg0 context.Context, arg1 *applyconfigurationsstoragev1.CSINodeApplyConfiguration, arg2 metav1.ApplyOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockCSINodeInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockCSINodeInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockCSINodeInterface) Create(arg0 context.Context, arg1 *v1.CSINode, arg2 metav1.CreateOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCSINodeInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCSINodeInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockCSINodeInterface) Delete(arg0 context.Context, arg1 string, arg2 metav1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCSINodeInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCSINodeInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockCSINodeInterface) DeleteCollection(arg0 context.Context, arg1 metav1.DeleteOptions, arg2 metav1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockCSINodeInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockCSINodeInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockCSINodeInterface) Get(arg0 context.Context, arg1 string, arg2 metav1.GetOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCSINodeInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCSINodeInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockCSINodeInterface) List(arg0 context.Context, arg1 metav1.ListOptions) (*v1.CSINodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.CSINodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCSINodeInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCSINodeInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockCSINodeInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 metav1.PatchOptions, arg5 ...string) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockCSINodeInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockCSINodeInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockCSINodeInterface) Update(arg0 context.Context, arg1 *v1.CSINode, arg2 metav1.UpdateOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCSINodeInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCSINodeInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockCSINodeInterface) Watch(arg0 context.Context, arg1 metav1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockCSINodeInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockCSINodeInterface)(nil).Watch), arg0, arg1)
}
