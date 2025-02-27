package utils

import "time"

type HttpResponse struct {
	Status    uint        `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

func GetResponse(status uint, message string, data interface{}) HttpResponse {
	return HttpResponse{
		Status:    status,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC822),
	}
}
