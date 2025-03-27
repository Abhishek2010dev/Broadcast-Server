package types

import (
	"encoding/json"
)

type WebsocketError struct {
	Message string `json:"error_message"`
}

// For server side
func NewWebsocketError(message string) *WebsocketError {
	return &WebsocketError{
		Message: message,
	}
}

func (w *WebsocketError) ToJson() (string, error) {
	data, err := json.Marshal(w)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// For client side
func WebsocketErrorFromJson(v []byte) (*WebsocketError, error) {
	var websocketError WebsocketError
	err := json.Unmarshal(v, &websocketError)
	if err != nil {
		return nil, err
	}
	return &websocketError, nil
}
