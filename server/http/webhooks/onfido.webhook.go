package webhooks

import (
	"context"
	"encoding/json"
	"ms.api/log"
	"ms.api/protos/pb/onfidoService"
	"net/http"
	"time"
)

type onfidoWebhook struct {
	Payload struct {
		ResourceType string `json:"resource_type"`
		Action       string `json:"action"`
		Object       struct {
			ID                 string    `json:"id"`
			Status             string    `json:"status"`
			CompletedAtIso8601 time.Time `json:"completed_at_iso8601"`
			Href               string    `json:"href"`
		} `json:"object"`
	} `json:"payload"`
}

func HandleOnfidoWebhook(rpc onfidoService.OnfidoServiceClient) http.HandlerFunc {
	const OnfidoCheckCompleted = "check.completed"

	// sample incoming payload
	// {
	//  "payload": {
	//    "resource_type": "check",
	//    "action": "check.completed",
	//    "object": {
	//      "id": "3a5a1033-f387-4be4-82a2-f3ea4bca927e",
	//      "status": "complete",
	//      "completed_at_iso8601": "2020-09-14T12:10:46Z",
	//      "href": "https://api.onfido.com/v3/checks/3a5a1033-f387-4be4-82a2-f3ea4bca927e"
	//    }
	//  }
	// }

	return func(writer http.ResponseWriter, request *http.Request) {
		// Get values from Onfido's result push
		decoder := json.NewDecoder(request.Body)
		var onfido onfidoWebhook
		err := decoder.Decode(&onfido)
		if err != nil {
			panic(err)
		}

		if onfido.Payload.Action != OnfidoCheckCompleted {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte("only onfido check webhooks is allowed."))
			log.Errorf("Failed to process onfido webhook push, Got the following payload: %v", onfido)
			return
		}

		if _, err := rpc.WebhookPush(context.Background(), &onfidoService.OnfidoCheckWebhookRequest{
			Id:          onfido.Payload.Object.ID,
			Status:      onfido.Payload.Object.Status,
			CompletedAt: onfido.Payload.Object.CompletedAtIso8601.String(),
		}); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte("failed to process onfido webhook push"))
			log.Errorf("Failed to process onfido webhook push via gRPC, Got the following error: %v", err)
			return
		}

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("successfully processed onfido's webhook push."))

		log.Debugf("Processed onfido webhook: %s", onfido.Payload.Object.ID)
	}
}
