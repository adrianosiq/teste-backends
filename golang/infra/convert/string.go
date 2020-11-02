package convert

import (
	"fmt"
	"strings"
)

type dependecies interface {
	Fields() []string
	Split(input string) []string
	SelectType(schema string, action string, fields []string) ([]string, error)
	CreateJSON(fields []string, values []string) ([]byte, error)
}

type deps struct{}

func (d deps) Fields() []string {
	return []string{
		"event_id", "event_schema", "event_action", "event_timestamp", "proposal_id",
		"proposal_loan_value", "proposal_number_of_monthly_installments",
		"warranty_id", "warranty_value", "warranty_province",
		"proponent_id", "proponent_name", "proponent_age", "proponent_monthly_income", "proponent_is_main",
	}
}

func (d deps) Split(input string) []string {
	return strings.Split(input, ",")
}

func (d deps) SelectType(schema string, action string, fields []string) ([]string, error) {
	selected := fields[:5]
	switch {
	case schema == "proposal" && action != "deleted":
		selected = append(selected, fields[5:7]...)
		return selected, nil
	case schema == "warranty":
		selected = append(selected, fields[7:8]...)
		if action != "removed" {
			selected = append(selected, fields[8:10]...)
		}
		return selected, nil
	case schema == "proponent":
		selected = append(selected, fields[10:11]...)
		if action != "removed" {
			selected = append(selected, fields[11:]...)
		}
		return selected, nil
	}
	return nil, fmt.Errorf("unmapped schema")
}

func (d deps) CreateJSON(fields []string, values []string) ([]byte, error) {
	var str string
	for i, field := range fields {
		if i > len(values) {
			return nil, fmt.Errorf("malformed values in the input")
		}
		str += fmt.Sprintf("\"%s\":\"%s\",", field, values[i])
	}
	str = fmt.Sprintf(`{%s}`, str)
	return []byte(str), nil
}

//String load is responsible for load depedencies
type String struct {
	deps dependecies
}

//NewString create new string
func NewString() String {
	return String{
		deps: deps{},
	}
}

//StringToJSON is responsible for receive string and convert in json
func (str String) StringToJSON(inputs []string) ([][]byte, error) {
	fields := str.deps.Fields()
	var dataJSON [][]byte
	for _, input := range inputs {
		strSplit := str.deps.Split(input)
		if len(strSplit) > 1 {
			selected, err := str.deps.SelectType(strSplit[1], strSplit[2], fields)
			if err != nil {
				return [][]byte{}, err
			}
			createdJSON, err := str.deps.CreateJSON(selected, strSplit)
			if err != nil {
				return [][]byte{}, err
			}
			dataJSON = append(dataJSON, createdJSON)
		}
	}
	return dataJSON, nil
}
