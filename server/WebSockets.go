package server


import (
	secrets "ms.api/config"
	"ms.api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var WSUpgradeManager = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSockets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if secrets.GetSecrets().Environment == secrets.Production {
			// Do origin check.
		}
		//
		if r.Header.Get("Sec-WebSocket-Extensions") != "" {
			r.Header.Del("Sec-WebSocket-Extensions")
		}

		conn, e := WSUpgradeManager.Upgrade(w, r, r.Header)
		if e != nil {
			fmt.Print("ERROR: ", e.Error(), "\n")
			return
		}
		go func(conn *websocket.Conn) {
			for {
				_, msg, connectionError := conn.ReadMessage()
				if connectionError != nil {
					_ = conn.Close()
					return
				}

				var message WSMessageInput
				if err := json.Unmarshal(msg, &message); err != nil {
					fmt.Print(message, "m.")
					// Skip it and move on.
					_ = conn.WriteJSON(WSMessageOutput{
						Channel: "connection",
						Action:  "previous-input",
						Result: models.Result{
							Success:      false,
							Message:      "Invalid JSON Provided",
							//ReturnStatus: models.ReturnStatusNotOk,
						},
					})
					continue
				}

				// Pushing message back to client.
				_ = conn.WriteJSON(message)
			}
		}(conn)
	}
}

type WSMessageInput struct {
	Channel string      `json:"channel"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

type WSMessageOutput struct {
	Channel string        `json:"channel"`
	Action  string        `json:"action"`
	Result  models.Result `json:"result"`
}
