// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	http "net/http"
	ent "aio/ent"
	entity "aio/internal/entity"
	pb "aio/internal/pb"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
	go_uuid "github.com/satori/go.uuid"
)

// MockBark is a mock of Bark interface.
type MockBark struct {
	ctrl     *gomock.Controller
	recorder *MockBarkMockRecorder
}

// MockBarkMockRecorder is the mock recorder for MockBark.
type MockBarkMockRecorder struct {
	mock *MockBark
}

// NewMockBark creates a new mock instance.
func NewMockBark(ctrl *gomock.Controller) *MockBark {
	mock := &MockBark{ctrl: ctrl}
	mock.recorder = &MockBarkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBark) EXPECT() *MockBarkMockRecorder {
	return m.recorder
}

// Pull mocks base method.
func (m *MockBark) Pull(ctx context.Context, device *entity.BarkDevice, offset, limit int) ([]*entity.BarkHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", ctx, device, offset, limit)
	ret0, _ := ret[0].([]*entity.BarkHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pull indicates an expected call of Pull.
func (mr *MockBarkMockRecorder) Pull(ctx, device, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockBark)(nil).Pull), ctx, device, offset, limit)
}

// Push mocks base method.
func (m *MockBark) Push(arg0 context.Context, arg1 string, arg2 *entity.APNsMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockBarkMockRecorder) Push(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockBark)(nil).Push), arg0, arg1, arg2)
}

// Register mocks base method.
func (m *MockBark) Register(arg0 context.Context, arg1 *entity.BarkDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockBarkMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockBark)(nil).Register), arg0, arg1)
}

// MockBarkRepo is a mock of BarkRepo interface.
type MockBarkRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBarkRepoMockRecorder
}

// MockBarkRepoMockRecorder is the mock recorder for MockBarkRepo.
type MockBarkRepoMockRecorder struct {
	mock *MockBarkRepo
}

// NewMockBarkRepo creates a new mock instance.
func NewMockBarkRepo(ctrl *gomock.Controller) *MockBarkRepo {
	mock := &MockBarkRepo{ctrl: ctrl}
	mock.recorder = &MockBarkRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBarkRepo) EXPECT() *MockBarkRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockBarkRepo) Get(arg0 context.Context, arg1 *entity.BarkDevice) (*entity.BarkDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.BarkDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBarkRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBarkRepo)(nil).Get), arg0, arg1)
}

// SaveMessage mocks base method.
func (m *MockBarkRepo) SaveMessage(ctx context.Context, message *entity.APNsMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveMessage", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveMessage indicates an expected call of SaveMessage.
func (mr *MockBarkRepoMockRecorder) SaveMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveMessage", reflect.TypeOf((*MockBarkRepo)(nil).SaveMessage), ctx, message)
}

// Store mocks base method.
func (m *MockBarkRepo) Store(arg0 context.Context, arg1 *entity.BarkDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockBarkRepoMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockBarkRepo)(nil).Store), arg0, arg1)
}

// MockBarkWebAPI is a mock of BarkWebAPI interface.
type MockBarkWebAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBarkWebAPIMockRecorder
}

// MockBarkWebAPIMockRecorder is the mock recorder for MockBarkWebAPI.
type MockBarkWebAPIMockRecorder struct {
	mock *MockBarkWebAPI
}

// NewMockBarkWebAPI creates a new mock instance.
func NewMockBarkWebAPI(ctrl *gomock.Controller) *MockBarkWebAPI {
	mock := &MockBarkWebAPI{ctrl: ctrl}
	mock.recorder = &MockBarkWebAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBarkWebAPI) EXPECT() *MockBarkWebAPIMockRecorder {
	return m.recorder
}

// Push mocks base method.
func (m *MockBarkWebAPI) Push(arg0 context.Context, arg1 *entity.APNsMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockBarkWebAPIMockRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockBarkWebAPI)(nil).Push), arg0, arg1)
}

// MockPushDeer is a mock of PushDeer interface.
type MockPushDeer struct {
	ctrl     *gomock.Controller
	recorder *MockPushDeerMockRecorder
}

// MockPushDeerMockRecorder is the mock recorder for MockPushDeer.
type MockPushDeerMockRecorder struct {
	mock *MockPushDeer
}

// NewMockPushDeer creates a new mock instance.
func NewMockPushDeer(ctrl *gomock.Controller) *MockPushDeer {
	mock := &MockPushDeer{ctrl: ctrl}
	mock.recorder = &MockPushDeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPushDeer) EXPECT() *MockPushDeerMockRecorder {
	return m.recorder
}

// GenNewKey mocks base method.
func (m *MockPushDeer) GenNewKey(ctx context.Context, token, name string) ([]*entity.PushKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenNewKey", ctx, token, name)
	ret0, _ := ret[0].([]*entity.PushKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenNewKey indicates an expected call of GenNewKey.
func (mr *MockPushDeerMockRecorder) GenNewKey(ctx, token, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenNewKey", reflect.TypeOf((*MockPushDeer)(nil).GenNewKey), ctx, token, name)
}

// GetAllDevice mocks base method.
func (m *MockPushDeer) GetAllDevice(ctx context.Context, token string) ([]*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDevice", ctx, token)
	ret0, _ := ret[0].([]*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDevice indicates an expected call of GetAllDevice.
func (mr *MockPushDeerMockRecorder) GetAllDevice(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevice", reflect.TypeOf((*MockPushDeer)(nil).GetAllDevice), ctx, token)
}

// GetAllKey mocks base method.
func (m *MockPushDeer) GetAllKey(ctx context.Context, token string) ([]*entity.PushKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllKey", ctx, token)
	ret0, _ := ret[0].([]*entity.PushKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllKey indicates an expected call of GetAllKey.
func (mr *MockPushDeerMockRecorder) GetAllKey(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllKey", reflect.TypeOf((*MockPushDeer)(nil).GetAllKey), ctx, token)
}

// GetMessage mocks base method.
func (m *MockPushDeer) GetMessage(ctx context.Context, token string, offset, limit uint64) ([]*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", ctx, token, offset, limit)
	ret0, _ := ret[0].([]*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockPushDeerMockRecorder) GetMessage(ctx, token, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockPushDeer)(nil).GetMessage), ctx, token, offset, limit)
}

// GetUser mocks base method.
func (m *MockPushDeer) GetUser(ctx context.Context, token string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, token)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockPushDeerMockRecorder) GetUser(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockPushDeer)(nil).GetUser), ctx, token)
}

// PushMessage mocks base method.
func (m *MockPushDeer) PushMessage(ctx context.Context, key string, msg *entity.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushMessage", ctx, key, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushMessage indicates an expected call of PushMessage.
func (mr *MockPushDeerMockRecorder) PushMessage(ctx, key, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMessage", reflect.TypeOf((*MockPushDeer)(nil).PushMessage), ctx, key, msg)
}

// RegenKey mocks base method.
func (m *MockPushDeer) RegenKey(ctx context.Context, token, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegenKey", ctx, token, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegenKey indicates an expected call of RegenKey.
func (mr *MockPushDeerMockRecorder) RegenKey(ctx, token, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegenKey", reflect.TypeOf((*MockPushDeer)(nil).RegenKey), ctx, token, id)
}

// Register mocks base method.
func (m *MockPushDeer) Register(ctx context.Context, appleID, email, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, appleID, email, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockPushDeerMockRecorder) Register(ctx, appleID, email, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPushDeer)(nil).Register), ctx, appleID, email, name)
}

// RegisterDevice mocks base method.
func (m *MockPushDeer) RegisterDevice(ctx context.Context, info *entity.RegInfo) ([]*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterDevice", ctx, info)
	ret0, _ := ret[0].([]*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterDevice indicates an expected call of RegisterDevice.
func (mr *MockPushDeerMockRecorder) RegisterDevice(ctx, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterDevice", reflect.TypeOf((*MockPushDeer)(nil).RegisterDevice), ctx, info)
}

// RemoveDevice mocks base method.
func (m *MockPushDeer) RemoveDevice(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDevice", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDevice indicates an expected call of RemoveDevice.
func (mr *MockPushDeerMockRecorder) RemoveDevice(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDevice", reflect.TypeOf((*MockPushDeer)(nil).RemoveDevice), ctx, id)
}

// RemoveKey mocks base method.
func (m *MockPushDeer) RemoveKey(ctx context.Context, token, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveKey", ctx, token, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveKey indicates an expected call of RemoveKey.
func (mr *MockPushDeerMockRecorder) RemoveKey(ctx, token, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveKey", reflect.TypeOf((*MockPushDeer)(nil).RemoveKey), ctx, token, id)
}

// RemoveMessage mocks base method.
func (m *MockPushDeer) RemoveMessage(ctx context.Context, token, msgID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMessage", ctx, token, msgID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMessage indicates an expected call of RemoveMessage.
func (mr *MockPushDeerMockRecorder) RemoveMessage(ctx, token, msgID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMessage", reflect.TypeOf((*MockPushDeer)(nil).RemoveMessage), ctx, token, msgID)
}

// RenameDevice mocks base method.
func (m *MockPushDeer) RenameDevice(ctx context.Context, id, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameDevice", ctx, id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameDevice indicates an expected call of RenameDevice.
func (mr *MockPushDeerMockRecorder) RenameDevice(ctx, id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameDevice", reflect.TypeOf((*MockPushDeer)(nil).RenameDevice), ctx, id, name)
}

// RenameKey mocks base method.
func (m *MockPushDeer) RenameKey(ctx context.Context, token, id, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameKey", ctx, token, id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameKey indicates an expected call of RenameKey.
func (mr *MockPushDeerMockRecorder) RenameKey(ctx, token, id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameKey", reflect.TypeOf((*MockPushDeer)(nil).RenameKey), ctx, token, id, name)
}

// Upgrade mocks base method.
func (m *MockPushDeer) Upgrade(ctx context.Context, deviceID string, w http.ResponseWriter, r *http.Request, h http.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", ctx, deviceID, w, r, h)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockPushDeerMockRecorder) Upgrade(ctx, deviceID, w, r, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockPushDeer)(nil).Upgrade), ctx, deviceID, w, r, h)
}

// ValidateToken mocks base method.
func (m *MockPushDeer) ValidateToken(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockPushDeerMockRecorder) ValidateToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockPushDeer)(nil).ValidateToken), ctx, token)
}

// MockPushDeerRepo is a mock of PushDeerRepo interface.
type MockPushDeerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPushDeerRepoMockRecorder
}

// MockPushDeerRepoMockRecorder is the mock recorder for MockPushDeerRepo.
type MockPushDeerRepoMockRecorder struct {
	mock *MockPushDeerRepo
}

// NewMockPushDeerRepo creates a new mock instance.
func NewMockPushDeerRepo(ctrl *gomock.Controller) *MockPushDeerRepo {
	mock := &MockPushDeerRepo{ctrl: ctrl}
	mock.recorder = &MockPushDeerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPushDeerRepo) EXPECT() *MockPushDeerRepoMockRecorder {
	return m.recorder
}

// GetAllDevice mocks base method.
func (m *MockPushDeerRepo) GetAllDevice(arg0 context.Context, arg1 string) ([]*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDevice", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDevice indicates an expected call of GetAllDevice.
func (mr *MockPushDeerRepoMockRecorder) GetAllDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevice", reflect.TypeOf((*MockPushDeerRepo)(nil).GetAllDevice), arg0, arg1)
}

// GetAllPushKey mocks base method.
func (m *MockPushDeerRepo) GetAllPushKey(arg0 context.Context, arg1 string) ([]*entity.PushKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPushKey", arg0, arg1)
	ret0, _ := ret[0].([]*entity.PushKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPushKey indicates an expected call of GetAllPushKey.
func (mr *MockPushDeerRepoMockRecorder) GetAllPushKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPushKey", reflect.TypeOf((*MockPushDeerRepo)(nil).GetAllPushKey), arg0, arg1)
}

// GetDevice mocks base method.
func (m *MockPushDeerRepo) GetDevice(arg0 context.Context, arg1 string) (*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevice", arg0, arg1)
	ret0, _ := ret[0].(*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice.
func (mr *MockPushDeerRepoMockRecorder) GetDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockPushDeerRepo)(nil).GetDevice), arg0, arg1)
}

// GetMessage mocks base method.
func (m *MockPushDeerRepo) GetMessage(ctx context.Context, userID string, offset, count uint64) ([]*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", ctx, userID, offset, count)
	ret0, _ := ret[0].([]*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockPushDeerRepoMockRecorder) GetMessage(ctx, userID, offset, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockPushDeerRepo)(nil).GetMessage), ctx, userID, offset, count)
}

// GetPushKey mocks base method.
func (m *MockPushDeerRepo) GetPushKey(arg0 context.Context, arg1 int64, arg2, arg3 string) (*entity.PushKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPushKey", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*entity.PushKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPushKey indicates an expected call of GetPushKey.
func (mr *MockPushDeerRepoMockRecorder) GetPushKey(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPushKey", reflect.TypeOf((*MockPushDeerRepo)(nil).GetPushKey), arg0, arg1, arg2, arg3)
}

// GetUser mocks base method.
func (m *MockPushDeerRepo) GetUser(ctx context.Context, uuid, appleID string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, uuid, appleID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockPushDeerRepoMockRecorder) GetUser(ctx, uuid, appleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockPushDeerRepo)(nil).GetUser), ctx, uuid, appleID)
}

// GetUserID mocks base method.
func (m *MockPushDeerRepo) GetUserID(ctx context.Context, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", ctx, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockPushDeerRepoMockRecorder) GetUserID(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockPushDeerRepo)(nil).GetUserID), ctx, token)
}

// RemoveDevice mocks base method.
func (m *MockPushDeerRepo) RemoveDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDevice indicates an expected call of RemoveDevice.
func (mr *MockPushDeerRepoMockRecorder) RemoveDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDevice", reflect.TypeOf((*MockPushDeerRepo)(nil).RemoveDevice), arg0, arg1)
}

// RemoveMessage mocks base method.
func (m *MockPushDeerRepo) RemoveMessage(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMessage indicates an expected call of RemoveMessage.
func (mr *MockPushDeerRepoMockRecorder) RemoveMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMessage", reflect.TypeOf((*MockPushDeerRepo)(nil).RemoveMessage), arg0, arg1)
}

// RemovePushKey mocks base method.
func (m *MockPushDeerRepo) RemovePushKey(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePushKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePushKey indicates an expected call of RemovePushKey.
func (mr *MockPushDeerRepoMockRecorder) RemovePushKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePushKey", reflect.TypeOf((*MockPushDeerRepo)(nil).RemovePushKey), arg0, arg1)
}

// StoreDevice mocks base method.
func (m *MockPushDeerRepo) StoreDevice(arg0 context.Context, arg1 *entity.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreDevice indicates an expected call of StoreDevice.
func (mr *MockPushDeerRepoMockRecorder) StoreDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreDevice", reflect.TypeOf((*MockPushDeerRepo)(nil).StoreDevice), arg0, arg1)
}

// StoreMessage mocks base method.
func (m *MockPushDeerRepo) StoreMessage(arg0 context.Context, arg1 *entity.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreMessage indicates an expected call of StoreMessage.
func (mr *MockPushDeerRepoMockRecorder) StoreMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMessage", reflect.TypeOf((*MockPushDeerRepo)(nil).StoreMessage), arg0, arg1)
}

// StorePushKey mocks base method.
func (m *MockPushDeerRepo) StorePushKey(arg0 context.Context, arg1 *entity.PushKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorePushKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorePushKey indicates an expected call of StorePushKey.
func (mr *MockPushDeerRepoMockRecorder) StorePushKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorePushKey", reflect.TypeOf((*MockPushDeerRepo)(nil).StorePushKey), arg0, arg1)
}

// StoreUser mocks base method.
func (m *MockPushDeerRepo) StoreUser(ctx context.Context, appleID, email, name, uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUser", ctx, appleID, email, name, uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUser indicates an expected call of StoreUser.
func (mr *MockPushDeerRepoMockRecorder) StoreUser(ctx, appleID, email, name, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUser", reflect.TypeOf((*MockPushDeerRepo)(nil).StoreUser), ctx, appleID, email, name, uuid)
}

// UpdateDeviceName mocks base method.
func (m *MockPushDeerRepo) UpdateDeviceName(ctx context.Context, deviceID, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeviceName", ctx, deviceID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeviceName indicates an expected call of UpdateDeviceName.
func (mr *MockPushDeerRepoMockRecorder) UpdateDeviceName(ctx, deviceID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeviceName", reflect.TypeOf((*MockPushDeerRepo)(nil).UpdateDeviceName), ctx, deviceID, name)
}

// UpdatePushKey mocks base method.
func (m *MockPushDeerRepo) UpdatePushKey(arg0 context.Context, arg1 *entity.PushKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePushKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePushKey indicates an expected call of UpdatePushKey.
func (mr *MockPushDeerRepoMockRecorder) UpdatePushKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePushKey", reflect.TypeOf((*MockPushDeerRepo)(nil).UpdatePushKey), arg0, arg1)
}

// MockPushDeerWebAPI is a mock of PushDeerWebAPI interface.
type MockPushDeerWebAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPushDeerWebAPIMockRecorder
}

// MockPushDeerWebAPIMockRecorder is the mock recorder for MockPushDeerWebAPI.
type MockPushDeerWebAPIMockRecorder struct {
	mock *MockPushDeerWebAPI
}

// NewMockPushDeerWebAPI creates a new mock instance.
func NewMockPushDeerWebAPI(ctrl *gomock.Controller) *MockPushDeerWebAPI {
	mock := &MockPushDeerWebAPI{ctrl: ctrl}
	mock.recorder = &MockPushDeerWebAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPushDeerWebAPI) EXPECT() *MockPushDeerWebAPIMockRecorder {
	return m.recorder
}

// Push mocks base method.
func (m *MockPushDeerWebAPI) Push(arg0 context.Context, arg1 []*entity.Device, arg2 *entity.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockPushDeerWebAPIMockRecorder) Push(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockPushDeerWebAPI)(nil).Push), arg0, arg1, arg2)
}

// Register mocks base method.
func (m *MockPushDeerWebAPI) Register(ctx context.Context, key string, c *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", ctx, key, c)
}

// Register indicates an expected call of Register.
func (mr *MockPushDeerWebAPIMockRecorder) Register(ctx, key, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPushDeerWebAPI)(nil).Register), ctx, key, c)
}

// MockExtension is a mock of Extension interface.
type MockExtension struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionMockRecorder
}

// MockExtensionMockRecorder is the mock recorder for MockExtension.
type MockExtensionMockRecorder struct {
	mock *MockExtension
}

// NewMockExtension creates a new mock instance.
func NewMockExtension(ctrl *gomock.Controller) *MockExtension {
	mock := &MockExtension{ctrl: ctrl}
	mock.recorder = &MockExtensionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtension) EXPECT() *MockExtensionMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockExtension) Add(ctx context.Context, uid go_uuid.UUID, group ...*pb.Group) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, uid}
	for _, a := range group {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockExtensionMockRecorder) Add(ctx, uid interface{}, group ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, uid}, group...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockExtension)(nil).Add), varargs...)
}

// Connect mocks base method.
func (m *MockExtension) Connect(ctx context.Context, uid go_uuid.UUID, wsConn *websocket.Conn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx, uid, wsConn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockExtensionMockRecorder) Connect(ctx, uid, wsConn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockExtension)(nil).Connect), ctx, uid, wsConn)
}

// Pull mocks base method.
func (m *MockExtension) Pull(ctx context.Context, uid go_uuid.UUID) (*ent.ExtensionClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", ctx, uid)
	ret0, _ := ret[0].(*ent.ExtensionClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pull indicates an expected call of Pull.
func (mr *MockExtensionMockRecorder) Pull(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockExtension)(nil).Pull), ctx, uid)
}

// Register mocks base method.
func (m *MockExtension) Register(ctx context.Context, name, extensionID string) (*ent.ExtensionClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, name, extensionID)
	ret0, _ := ret[0].(*ent.ExtensionClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockExtensionMockRecorder) Register(ctx, name, extensionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockExtension)(nil).Register), ctx, name, extensionID)
}

// RemoveGroup mocks base method.
func (m *MockExtension) RemoveGroup(ctx context.Context, uid, groupUid go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveGroup", ctx, uid, groupUid)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveGroup indicates an expected call of RemoveGroup.
func (mr *MockExtensionMockRecorder) RemoveGroup(ctx, uid, groupUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGroup", reflect.TypeOf((*MockExtension)(nil).RemoveGroup), ctx, uid, groupUid)
}

// RemoveTab mocks base method.
func (m *MockExtension) RemoveTab(ctx context.Context, uid, groupUid, tabUid go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTab", ctx, uid, groupUid, tabUid)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTab indicates an expected call of RemoveTab.
func (mr *MockExtensionMockRecorder) RemoveTab(ctx, uid, groupUid, tabUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTab", reflect.TypeOf((*MockExtension)(nil).RemoveTab), ctx, uid, groupUid, tabUid)
}

// SwapTab mocks base method.
func (m *MockExtension) SwapTab(ctx context.Context, uid, firstGroupUid, firstGroupTabUid, secondGroupUid, secondGroupTabUid go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapTab", ctx, uid, firstGroupUid, firstGroupTabUid, secondGroupUid, secondGroupTabUid)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwapTab indicates an expected call of SwapTab.
func (mr *MockExtensionMockRecorder) SwapTab(ctx, uid, firstGroupUid, firstGroupTabUid, secondGroupUid, secondGroupTabUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapTab", reflect.TypeOf((*MockExtension)(nil).SwapTab), ctx, uid, firstGroupUid, firstGroupTabUid, secondGroupUid, secondGroupTabUid)
}
