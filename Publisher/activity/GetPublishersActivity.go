package activity

import (
	"Publisher/component"
	"encoding/json"
	"net/http"
)

func GetAllPublisherResponse(writer http.ResponseWriter, _ *http.Request) {
	publishers := component.GetAllPublisher()
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(publishers)
}
