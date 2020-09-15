package webhooks

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"ms.api/fakes"
	"ms.api/protos/pb/onfidoService"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleOnfidoWebhook(t *testing.T) {
	const testPayload = `{
	"payload": {
	    "resource_type": "check",
	    "action": "check.completed",
	    "object": {
	      "id": "3a5a1033-f387-4be4-82a2-f3ea4bca927e",
	      "status": "complete",
	      "completed_at_iso8601": "2020-09-14T12:10:46Z",
	      "href": "https://api.onfido.com/v3/checks/3a5a1033-f387-4be4-82a2-f3ea4bca927e"
	    }
	  }
	}`

	onfidoClient := fakes.NewFakeOnFidoClient(&onfidoService.Void{}, nil, nil)
	webHookHandler := HandleOnfidoWebhook(onfidoClient)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(testPayload)))
	webHookHandler.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestHandleOnfidoWebhook_WebhookPush_Error(t *testing.T) {
	const testPayload = `{
	"payload": {
	    "resource_type": "check",
	    "action": "check.completed",
	    "object": {
	      "id": "3a5a1033-f387-4be4-82a2-f3ea4bca927e",
	      "status": "complete",
	      "completed_at_iso8601": "2020-09-14T12:10:46Z",
	      "href": "https://api.onfido.com/v3/checks/3a5a1033-f387-4be4-82a2-f3ea4bca927e"
	    }
	  }
	}`

	onfidoClient := fakes.NewFakeOnFidoClient(nil, nil, errors.New("failed to call client"))
	webHookHandler := HandleOnfidoWebhook(onfidoClient)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(testPayload)))
	webHookHandler.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
}
