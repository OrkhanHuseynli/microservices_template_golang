package models

import "encoding/json"

type PaymentCodec struct{}

func (c *PaymentCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *PaymentCodec) Decode(data []byte) (interface{}, error) {
	var p Payment
	return &p, json.Unmarshal(data, &p)
}


type PaymentListCodec struct{}

func (c *PaymentListCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *PaymentListCodec) Decode(data []byte) (interface{}, error) {
	var m []Payment
	err := json.Unmarshal(data, &m)
	return m, err
}


type ProcessedPaymentCodec struct{}

func (c *ProcessedPaymentCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *ProcessedPaymentCodec) Decode(data []byte) (interface{}, error) {
	var p ProcessedPayment
	return &p, json.Unmarshal(data, &p)
}
