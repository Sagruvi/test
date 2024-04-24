package tests

import (
	"context"
	"reflect"
	"testing"
	"testproject/internal/entity"
	"testproject/internal/repo"
	"testproject/internal/service"
)

func TestNewService(t *testing.T) {
	type args struct {
		repo repo.Repository
	}
	tests := []struct {
		name string
		args args
		want service.Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetEvents(t *testing.T) {
	type fields struct {
		Repository repo.Repository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.service{
				Repository: tt.fields.Repository,
			}
			got, err := s.GetEvents(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvents() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Insert(t *testing.T) {
	type fields struct {
		Repository repo.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.service{
				Repository: tt.fields.Repository,
			}
			if err := s.Insert(); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ProcessEvent(t *testing.T) {
	type fields struct {
		Repository repo.Repository
	}
	type args struct {
		ctx   context.Context
		event entity.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.service{
				Repository: tt.fields.Repository,
			}
			if err := s.ProcessEvent(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ProcessEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Search(t *testing.T) {
	type fields struct {
		Repository repo.Repository
	}
	type args struct {
		request entity.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.service{
				Repository: tt.fields.Repository,
			}
			got, err := s.Search(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
