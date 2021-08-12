package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_counter(t *testing.T) {
	type args struct {
		out chan<- int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"negative + negative",
			args{-1, -1},
			-2,
		},
		{
			"negative + positive",
			args{-1, 1},
			0,
		},
		{
			"positive + positive",
			args{1, 1},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Add(tt.args.x, tt.args.y), tt.want, tt.name)
		})
	}
}