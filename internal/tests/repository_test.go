package tests

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"reflect"
	"testing"
	"testproject/internal/entity"
	"testproject/internal/repo"
)

func TestNewRepo(t *testing.T) {
	tests := []struct {
		name string
		want *repo.Repo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repo.NewRepo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_Close(t *testing.T) {
	type fields struct {
		Conn driver.Conn
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
			r := &repo.Repo{
				Conn: tt.fields.Conn,
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_CreateEvent(t *testing.T) {
	type fields struct {
		Conn driver.Conn
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
			r := &repo.Repo{
				Conn: tt.fields.Conn,
			}
			if err := r.CreateEvent(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GetEvent(t *testing.T) {
	type fields struct {
		Conn driver.Conn
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
			r := &repo.Repo{
				Conn: tt.fields.Conn,
			}
			got, err := r.GetEvent(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_InsertData(t *testing.T) {
	type fields struct {
		Conn driver.Conn
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
			r := &repo.Repo{
				Conn: tt.fields.Conn,
			}
			if err := r.InsertData(); (err != nil) != tt.wantErr {
				t.Errorf("InsertData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_LogData(t *testing.T) {
	type fields struct {
		Conn driver.Conn
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
			r := &repo.Repo{
				Conn: tt.fields.Conn,
			}
			got, err := r.LogData(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connect(t *testing.T) {
	tests := []struct {
		name    string
		want    driver.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() got = %v, want %v", got, tt.want)
			}
		})
	}
}
