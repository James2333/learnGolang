package db

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	type args struct {
		db  DB
		key string
	}
	ctrl:=gomock.NewController(t)
	m:=NewMockDB(ctrl)

	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFromDB(tt.args.db, tt.args.key); got != tt.want {
				t.Errorf("GetFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}