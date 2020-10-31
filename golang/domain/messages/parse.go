package messages

import (
	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/adrianosiq/teste-backends/golang/infra"
)

//Parse is responsible for load services
type Parse struct {
	convertString infra.ConvertStringToJSON
	convertEvent  infra.ConvertEvent
}

//NewParse create a new parse
func NewParse(convertString infra.ConvertStringToJSON, convertEvent infra.ConvertEvent) Parse {
	return Parse{
		convertString: convertString,
		convertEvent:  convertEvent,
	}
}

//StringToEvent is responsible for convert string to event
func (parse Parse) StringToEvent(inputs []string) ([]domain.Event, error) {
	dataJSON, err := parse.convertString.StringToJSON(inputs)
	if err != nil {
		return nil, err
	}
	events := parse.convertEvent.CreateEvent(dataJSON)
	return events, nil
}
