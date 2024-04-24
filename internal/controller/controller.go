package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"testproject/internal/entity"
	"testproject/internal/repo"
	"testproject/internal/service"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.3 --name=Controller --output=../mocks/controller
type Controller interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEvent(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type Gateway struct {
	s service.Service
}

func NewController() Controller {
	return &Gateway{s: service.NewService(repo.NewRepo())}
}

// CreateEvent handles HTTP POST requests to insert an event into ClickHouse.
// @Summary Insert Event
// @Description Insert event into ClickHouse.
// @Tags Create
// @Accept json
// @Produce json
// @Param user body Event true "Event object to insert"
// @Success 200 {string} string "event inserted successfully"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /api/event [post]
func (c *Gateway) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.s.ProcessEvent(context.Background(), event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetEvent handles HTTP GET requests to retrieve events from ClickHouse.
// @Summary Get Events
// @Description Retrieve events from ClickHouse.
// @Tags Get
// @Produce json
// @Success 200 {array} Event "events retrieved successfully"
// @Failure 500 {string} string "error getting events from ClickHouse"
// @Router /api/event [get]
func (c *Gateway) GetEvent(w http.ResponseWriter, r *http.Request) {
	events, err := c.s.GetEvents(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Test handles HTTP GET requests to insert random test data.
// @Summary Test
// @Description Insert random test data.
// @Tags Get
// @Produce json
// @Success 200 {string} string "test data inserted successfully"
// @Failure 500 {string} string "error"
// @Router /api/test [get]
func (c *Gateway) Insert(w http.ResponseWriter, r *http.Request) {
	err := c.s.Insert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Search handles HTTP GET requests to search for events in ClickHouse.
// @Summary Search Events
// @Description Search for events in ClickHouse based on criteria.
// @Tags Get
// @Produce json
// @Param user body Request true "Search criteria"
// @Success 200 {array} Event "events matching search criteria"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "error parsing events from ClickHouse"
// @Router /api/search [get]
func (c *Gateway) Search(w http.ResponseWriter, r *http.Request) {
	var request entity.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := c.s.Search(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
