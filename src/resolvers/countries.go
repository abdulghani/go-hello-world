package resolvers

import (
	"context"
	"encoding/json"
	"hello_world/src/utils"

	"github.com/nats-io/nats.go"
)

var inbox string = "_INBOX.72HMCBZXALI2BQ949SGKXA" + utils.CreateUlid()

type NatsPayload struct {
	Pattern string      `json:"pattern"`
	Data    interface{} `json:"data"`
	Id      string      `json:"id"`
}

func Countries(ctx context.Context) (string, error) {
	client := utils.GetNatsClient()

	client.Subscribe(inbox, func(m *nats.Msg) {
		utils.Inspect("GET MESSAGE FROM INBOX", m, m.Data)
		var reply interface{}
		json.Unmarshal(m.Data, &reply)
		utils.Inspect("DATA", reply)
	})

	address := "general/country/get-all"
	encoded, _ := json.Marshal(&NatsPayload{
		Pattern: address,
		Data:    nil,
		Id:      utils.CreateUlid(),
	})
	client.PublishMsg(&nats.Msg{
		Subject: address,
		Reply:   inbox,
		Data:    encoded,
	})

	return "hello world", nil
}
