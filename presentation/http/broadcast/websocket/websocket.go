package websocket

import (
	"net/http"
	"nhooyr.io/websocket"
	"trabalho-01-batalha-naval/config"
	"trabalho-01-batalha-naval/presentation/http/broadcast"
)

type Subscriber struct {
	Key          string
	InputChannel chan []byte
}

type Broadcaster struct {
	allowedOrigins string
	subscribers    map[string][]*Subscriber
}

func NewWebSocketBroadcaster(cfg config.Config) *Broadcaster {
	return &Broadcaster{
		allowedOrigins: cfg.WebSocketAllowedOrigins,
		subscribers:    make(map[string][]*Subscriber),
	}
}

func (b *Broadcaster) Subscribe(channelKey string, subscriberKey string, w http.ResponseWriter, r *http.Request, callbackers ...*broadcast.Callbacker) error {
	subscriber := &Subscriber{
		Key:          subscriberKey,
		InputChannel: make(chan []byte),
	}

	b.registerSubscriber(channelKey, subscriber)

	connection, err := websocket.Accept(w, r, &websocket.AcceptOptions{OriginPatterns: []string{b.allowedOrigins}})

	if err != nil {
		return err
	}

	defer func() {
		for _, callbacker := range callbackers {
			callbacker.OnDisconnect(channelKey, subscriberKey)
		}

		_ = connection.CloseNow()
	}()

	defer b.unregisterSubscriber(channelKey, subscriber)

	ctx := connection.CloseRead(r.Context())

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case message := <-subscriber.InputChannel:
			err := connection.Write(ctx, websocket.MessageText, message)

			if err != nil {
				return err
			}
		}
	}
}

func (b *Broadcaster) Broadcast(channelKey string, message []byte) {
	if _, ok := b.subscribers[channelKey]; !ok {
		return
	}

	for _, subscriber := range b.subscribers[channelKey] {
		subscriber.InputChannel <- message
	}
}

func (b *Broadcaster) registerSubscriber(channelKey string, subscriber *Subscriber) {

	if _, ok := b.subscribers[channelKey]; !ok {
		b.subscribers[channelKey] = make([]*Subscriber, 0)
	}

	b.subscribers[channelKey] = append(b.subscribers[channelKey], subscriber)
}

func (b *Broadcaster) unregisterSubscriber(channelKey string, unregisteringSubscriber *Subscriber) {
	if _, ok := b.subscribers[channelKey]; !ok {
		return
	}

	for i, registeringSubscriber := range b.subscribers[channelKey] {
		if registeringSubscriber == unregisteringSubscriber {
			b.subscribers[channelKey] = append(b.subscribers[channelKey][:i], b.subscribers[channelKey][i+1:]...)
			return
		}
	}
}
