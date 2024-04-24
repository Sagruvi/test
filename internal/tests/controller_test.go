package tests

import (
	"net/http"
	"reflect"
	"testing"
	"testproject/internal/controller"
	"testproject/internal/service"
)

func TestGateway_CreateEvent(t *testing.T) {
	type fields struct {
		s service.Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller.Gateway{
				s: tt.fields.s,
			}
			c.CreateEvent(tt.args.w, tt.args.r)
		})
	}
}

func TestGateway_GetEvent(t *testing.T) {
	type fields struct {
		s service.Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller.Gateway{
				s: tt.fields.s,
			}
			c.GetEvent(tt.args.w, tt.args.r)
		})
	}
}

func TestGateway_Insert(t *testing.T) {
	type fields struct {
		s service.Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller.Gateway{
				s: tt.fields.s,
			}
			c.Insert(tt.args.w, tt.args.r)
		})
	}
}

func TestGateway_Search(t *testing.T) {
	type fields struct {
		s service.Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller.Gateway{
				s: tt.fields.s,
			}
			c.Search(tt.args.w, tt.args.r)
		})
	}
}

func TestNewController(t *testing.T) {
	tests := []struct {
		name string
		want controller.Controller
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controller.NewController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
