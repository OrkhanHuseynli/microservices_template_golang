package collector

import (
	"context"
	"encoding/json"
	"github.com/lovoo/goka"
	"github.com/microservices_template_golang/payment_history/models"
)

const maxMessages = 5

var (
	group goka.Group = "collector"
	Table goka.Table = goka.GroupTable(group)
)

type MessageListCodec struct{}

func (c *MessageListCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *MessageListCodec) Decode(data []byte) (interface{}, error) {
	var m []models.Message
	err := json.Unmarshal(data, &m)
	return m, err
}

func collect(ctx goka.Context, msg interface{}) {
	var ml []models.Message
	if v := ctx.Value(); v != nil {
		ml = v.([]models.Message)
	}

	m := msg.(*models.Message)
	ml = append(ml, *m)

	if len(ml) > maxMessages {
		ml = ml[len(ml)-maxMessages:]
	}
	ctx.SetValue(ml)
}

func Run(ctx context.Context, brokers []string) func() error {
	return func() error {
		g := goka.DefineGroup(group,
			goka.Input(models.ReceivedStream, new(models.MessageCodec), collect),
			goka.Persist(new(MessageListCodec)),
		)
		p, err := goka.NewProcessor(brokers, g)
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}