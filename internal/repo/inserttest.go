package repo

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"testproject/internal/entity"
	"time"
)

func (r *Repo) InsertData() error {
	batch, err := r.Conn.PrepareBatch(context.Background(), "INSERT INTO events (eventID, userID, eventTime) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	for i := 0; i < 1000; i++ {
		err := batch.AppendStruct(entity.Event{
			EventType: gofakeit.IntRange(1, 9999),
			UserID:    gofakeit.IntRange(1, 9999),
			EventTime: time.Now(),
		})
		if err != nil {
			return err
		}
	}
	return batch.Send()
}
func (r *Repo) LogData(request entity.Request) ([]entity.Event, error) {
	var event entity.Event
	err := r.Conn.Select(context.Background(), &event, "SELECT * FROM events WHERE eventID = ? AND eventTime BETWEEN ? AND ?", request.EventType, request.Before, request.After)
	if err != nil {
		return nil, err
	}
	return []entity.Event{event}, nil
}
