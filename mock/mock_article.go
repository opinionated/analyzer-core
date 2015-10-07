// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/opinionated/scraper-core/scraper (interfaces: Article)

package mock_scraper

import (
	gomock "github.com/golang/mock/gomock"
	html "golang.org/x/net/html"
)

// Mock of Article interface
type MockArticle struct {
	ctrl     *gomock.Controller
	recorder *_MockArticleRecorder
}

// Recorder for MockArticle (not exported)
type _MockArticleRecorder struct {
	mock *MockArticle
}

func NewMockArticle(ctrl *gomock.Controller) *MockArticle {
	mock := &MockArticle{ctrl: ctrl}
	mock.recorder = &_MockArticleRecorder{mock}
	return mock
}

func (_m *MockArticle) EXPECT() *_MockArticleRecorder {
	return _m.recorder
}

func (_m *MockArticle) DoParse(_param0 *html.Tokenizer) error {
	ret := _m.ctrl.Call(_m, "DoParse", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockArticleRecorder) DoParse(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoParse", arg0)
}

func (_m *MockArticle) GetData() string {
	ret := _m.ctrl.Call(_m, "GetData")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockArticleRecorder) GetData() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetData")
}

func (_m *MockArticle) GetDescription() string {
	ret := _m.ctrl.Call(_m, "GetDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockArticleRecorder) GetDescription() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDescription")
}

func (_m *MockArticle) GetLink() string {
	ret := _m.ctrl.Call(_m, "GetLink")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockArticleRecorder) GetLink() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLink")
}

func (_m *MockArticle) GetTitle() string {
	ret := _m.ctrl.Call(_m, "GetTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockArticleRecorder) GetTitle() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTitle")
}

func (_m *MockArticle) SetData(_param0 string) {
	_m.ctrl.Call(_m, "SetData", _param0)
}

func (_mr *_MockArticleRecorder) SetData(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetData", arg0)
}