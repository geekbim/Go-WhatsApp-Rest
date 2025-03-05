// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "go_wa_rest/domain/entity"

	mock "github.com/stretchr/testify/mock"

	types "go.mau.fi/whatsmeow/types"

	whatsmeow "go.mau.fi/whatsmeow"
)

// WhatsAppService is an autogenerated mock type for the WhatsAppService type
type WhatsAppService struct {
	mock.Mock
}

// WhatsAppComposeJID provides a mock function with given fields: jid
func (_m *WhatsAppService) WhatsAppComposeJID(jid string) types.JID {
	ret := _m.Called(jid)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppComposeJID")
	}

	var r0 types.JID
	if rf, ok := ret.Get(0).(func(string) types.JID); ok {
		r0 = rf(jid)
	} else {
		r0 = ret.Get(0).(types.JID)
	}

	return r0
}

// WhatsAppComposeStatus provides a mock function with given fields: jid, rjid, isComposing, isAudio
func (_m *WhatsAppService) WhatsAppComposeStatus(jid string, rjid types.JID, isComposing bool, isAudio bool) {
	_m.Called(jid, rjid, isComposing, isAudio)
}

// WhatsAppDecomposeJID provides a mock function with given fields: jid
func (_m *WhatsAppService) WhatsAppDecomposeJID(jid string) string {
	ret := _m.Called(jid)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppDecomposeJID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(jid)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WhatsAppGenerateQR provides a mock function with given fields: qrChan
func (_m *WhatsAppService) WhatsAppGenerateQR(qrChan <-chan whatsmeow.QRChannelItem) (string, int) {
	ret := _m.Called(qrChan)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppGenerateQR")
	}

	var r0 string
	var r1 int
	if rf, ok := ret.Get(0).(func(<-chan whatsmeow.QRChannelItem) (string, int)); ok {
		return rf(qrChan)
	}
	if rf, ok := ret.Get(0).(func(<-chan whatsmeow.QRChannelItem) string); ok {
		r0 = rf(qrChan)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(<-chan whatsmeow.QRChannelItem) int); ok {
		r1 = rf(qrChan)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// WhatsAppGroup provides a mock function with given fields: jid
func (_m *WhatsAppService) WhatsAppGroup(jid string) ([]*types.GroupInfo, error) {
	ret := _m.Called(jid)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppGroup")
	}

	var r0 []*types.GroupInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*types.GroupInfo, error)); ok {
		return rf(jid)
	}
	if rf, ok := ret.Get(0).(func(string) []*types.GroupInfo); ok {
		r0 = rf(jid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.GroupInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WhatsAppIsClientOK provides a mock function with given fields: jid
func (_m *WhatsAppService) WhatsAppIsClientOK(jid string) error {
	ret := _m.Called(jid)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppIsClientOK")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(jid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhatsAppLogin provides a mock function with given fields: jid
func (_m *WhatsAppService) WhatsAppLogin(jid string) (string, int, error) {
	ret := _m.Called(jid)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppLogin")
	}

	var r0 string
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (string, int, error)); ok {
		return rf(jid)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(jid)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(jid)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(jid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// WhatsAppLogout provides a mock function with given fields: jid
func (_m *WhatsAppService) WhatsAppLogout(jid string) error {
	ret := _m.Called(jid)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppLogout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(jid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhatsAppReconnect provides a mock function with given fields: jid
func (_m *WhatsAppService) WhatsAppReconnect(jid string) error {
	ret := _m.Called(jid)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppReconnect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(jid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhatsAppSendDocument provides a mock function with given fields: ctx, jid, rjid, whatsAppDocument
func (_m *WhatsAppService) WhatsAppSendDocument(ctx context.Context, jid string, rjid types.JID, whatsAppDocument *entity.WhatsAppDocument) (string, error) {
	ret := _m.Called(ctx, jid, rjid, whatsAppDocument)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppSendDocument")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.JID, *entity.WhatsAppDocument) (string, error)); ok {
		return rf(ctx, jid, rjid, whatsAppDocument)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.JID, *entity.WhatsAppDocument) string); ok {
		r0 = rf(ctx, jid, rjid, whatsAppDocument)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.JID, *entity.WhatsAppDocument) error); ok {
		r1 = rf(ctx, jid, rjid, whatsAppDocument)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WhatsAppSendImage provides a mock function with given fields: ctx, jid, rjid, whatsAppImage
func (_m *WhatsAppService) WhatsAppSendImage(ctx context.Context, jid string, rjid types.JID, whatsAppImage *entity.WhatsAppImage) (string, error) {
	ret := _m.Called(ctx, jid, rjid, whatsAppImage)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppSendImage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.JID, *entity.WhatsAppImage) (string, error)); ok {
		return rf(ctx, jid, rjid, whatsAppImage)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.JID, *entity.WhatsAppImage) string); ok {
		r0 = rf(ctx, jid, rjid, whatsAppImage)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.JID, *entity.WhatsAppImage) error); ok {
		r1 = rf(ctx, jid, rjid, whatsAppImage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WhatsAppSendText provides a mock function with given fields: ctx, jid, rjid, message
func (_m *WhatsAppService) WhatsAppSendText(ctx context.Context, jid string, rjid types.JID, message string) (string, error) {
	ret := _m.Called(ctx, jid, rjid, message)

	if len(ret) == 0 {
		panic("no return value specified for WhatsAppSendText")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.JID, string) (string, error)); ok {
		return rf(ctx, jid, rjid, message)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.JID, string) string); ok {
		r0 = rf(ctx, jid, rjid, message)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.JID, string) error); ok {
		r1 = rf(ctx, jid, rjid, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWhatsAppService creates a new instance of WhatsAppService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWhatsAppService(t interface {
	mock.TestingT
	Cleanup(func())
}) *WhatsAppService {
	mock := &WhatsAppService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
