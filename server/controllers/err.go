package controllers

import (
	"encoding/json"
)

type ErrMessage struct {
	Message string
}

func NewErrMsg(msg string) ErrMessage {
	return ErrMessage{Message: msg}
}

func (m ErrMessage) String() string {
	msg, _ := json.Marshal(m)
	return string(msg)
}
