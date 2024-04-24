package repo

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"net"
	"testproject/internal/entity"
	"time"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.3 --name=Repository --output=../mocks/repo
type Repository interface {
	CreateEvent(ctx context.Context, event entity.Event) error
	GetEvent(ctx context.Context) ([]entity.Event, error)
	InsertData() error
	LogData(request entity.Request) ([]entity.Event, error)
}

type Repo struct {
	driver.Conn
}

func NewRepo() *Repo {
	driver, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	return &Repo{
		Conn: driver,
	}
}

func (r *Repo) CreateEvent(ctx context.Context, event entity.Event) error {
	err := r.Conn.Exec(ctx, "INSERT INTO events (eventID, userID, eventTime) VALUES (?, ?, ?)", event.EventType, event.UserID, event.EventTime)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repo) GetEvent(ctx context.Context) ([]entity.Event, error) {
	var event []entity.Event
	err := r.Conn.Select(ctx, &event, "SELECT * FROM events LIMIT 100")
	if err != nil {
		return nil, err
	}
	return event, nil
}
func connect() (driver.Conn, error) {
	dialCount := 0
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			dialCount++
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (r *Repo) Close() error {
	return r.Conn.Close()
}
