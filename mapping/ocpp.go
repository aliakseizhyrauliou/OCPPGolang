package mapping

import (
	"encoding/json"
	"errors"
	"fmt"
)

type OCPPMessage struct {
	MessageTypeID float64                `json:"0"`
	UniqueID      string                 `json:"1"`
	MessageType   string                 `json:"2"`
	Payload       map[string]interface{} `json:"3"`
}

func MapToOcpp(jsonData string) (*OCPPMessage, error) {
	var data []interface{}

	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return &OCPPMessage{}, err
	}

	id, err := data[0].(float64)
	if !err {
		fmt.Println("MessageTypeID is not a float64")
		return &OCPPMessage{}, errors.New("MessageTypeID is not a float64")
	}

	payloadMap, ok := data[3].(map[string]interface{})
	if !ok {
		fmt.Println("Payload is not a map[string]interface{}")
		return &OCPPMessage{}, errors.New("Error mapping payload")
	}

	return &OCPPMessage{
		MessageTypeID: id,
		UniqueID:      data[1].(string),
		MessageType:   data[2].(string),
		Payload:       payloadMap,
	}, nil

}
