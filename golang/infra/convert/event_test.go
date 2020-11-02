package convert

import (
	"testing"

	"github.com/adrianosiq/teste-backends/golang/infra"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	t.Run("Should CreateEvent return events on success", func(t *testing.T) {
		str := NewString()
		data, _ := str.StringToJSON(infra.MockStringToJSON)
		event := NewEvent()
		out := event.CreateEvent(data)
		var a = assert.New(t)
		a.Equal(infra.MockEventExp, out)
	})
}
