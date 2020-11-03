package convert

import (
	"encoding/json"

	"github.com/adrianosiq/teste-backends/golang/domain"
)

type deps struct{}

//Event load is responsible for load depedencies
type Event struct{}

//NewEvent create new string
func NewEvent() Event {
	return Event{}
}

//CreateEvent is responsible for receive json and convert in event
func (event Event) CreateEvent(data [][]byte) []domain.Event {
	var events []domain.Event
	for _, item := range data {
		var event domain.Event
		json.Unmarshal(item, &event)
		events = append(events, event)
	}
	return events
}
