package main

import (
	"github.com/golang/mock/gomock"
	"learn101/mock_interface_test/equipment"
	"testing"
)

func TestPerson_dayLife(t *testing.T) {
	type fields struct {
		name  string
		phone equipment.Phone
	}
	// 生成mockPhone对象
	mockCtl := gomock.NewController(t)
	mockPhone := equipment.NewMockPhone(mockCtl)
	// 设置mockPhone对象的接口方法返回值
	mockPhone.EXPECT().ZhiHu().Return(true)
	mockPhone.EXPECT().WeiXin().Return(true)
	mockPhone.EXPECT().WangZhe().Return(true)
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"case1", fields{"iphone6s", equipment.NewIphone6s()}, true},
		{"case2", fields{"mocked phone", mockPhone}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Person{
				name:  tt.fields.name,
				phone: tt.fields.phone,
			}
			if got := x.dayLife(); got != tt.want {
				t.Errorf("dayLife() = %v, want %v", got, tt.want)
			}
		})
	}
}
