package handleFunc

import (
	"client/pkg/domain"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response := domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Server is started",
	}
	response.Send(w, "Health")
}
