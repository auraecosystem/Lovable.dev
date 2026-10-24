// go get github.com/gorilla/websocket
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/websocket"
)

func main() {
	headers := http.Header{"Authorization": {"Bearer " + os.Getenv("XAI_API_KEY")}}
	ws, _, err := websocket.DefaultDialer.Dial("wss://api.x.ai/v1/realtime?agent_id=agent_LhOy7jXCV7vuxLMq", headers)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	hello := map[string]any{
		"type": "conversation.item.create",
		"item": map[string]any{"type": "message", "role": "user",
			"content": []map[string]any{{"type": "input_text", "text": "Hello!"}}},
	}
	ws.WriteJSON(hello)
	ws.WriteJSON(map[string]any{"type": "response.create"})

	for {
		_, raw, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		var event struct {
			Type  string `json:"type"`
			Delta string `json:"delta"`
		}
		json.Unmarshal(raw, &event)
		if event.Type == "response.output_audio_transcript.delta" {
			fmt.Print(event.Delta)
		}
	}
}
