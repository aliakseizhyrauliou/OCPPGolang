package mapping

import (
	"encoding/json"
	"fmt"
)

type ChargerResponse struct {
	Rows [][]string `json:"0"`
}

type OCPPMessage struct {
	MessageTypeID float64                 `json:"0"`
	UniqueID      string                  `json:"1"`
	MessageType   string                  `json:"2"`
	Payload       BootNotificationPayload `json:"3"`
}

func Map(jsonData string) *OCPPMessage {
	var data []interface{}

	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return &OCPPMessage{}
	}

	ocpp := BuildOcppMessage(data)

	return ocpp
}

func BuildOcppMessage(data []interface{}) *OCPPMessage {

	id, err := data[0].(float64)
	if !err {
		fmt.Println("MessageTypeID is not a float64")
		return nil
	}

	payloadMap, ok := data[3].(map[string]interface{})
	if !ok {
		fmt.Println("Payload is not a map[string]interface{}")
		return nil
	}

	payload := BootNotificationPayload{
		ChargePointVendor:       payloadMap["chargePointVendor"].(string),
		ChargePointModel:        payloadMap["chargePointModel"].(string),
		ChargePointSerialNumber: payloadMap["chargePointSerialNumber"].(string),
		ChargeBoxSerialNumber:   payloadMap["chargeBoxSerialNumber"].(string),
		// Продолжайте добавлять другие параметры, если необходимо
	}

	return &OCPPMessage{
		MessageTypeID: id,
		UniqueID:      data[1].(string),
		MessageType:   data[2].(string),
		Payload:       payload,
	}
}
