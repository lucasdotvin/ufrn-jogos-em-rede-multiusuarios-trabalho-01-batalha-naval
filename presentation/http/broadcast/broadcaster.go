package broadcast

import "net/http"

type Broadcaster interface {
	Subscribe(channelKey string, subscriberKey string, w http.ResponseWriter, r *http.Request, callbackers ...*Callbacker) error
	Broadcast(channelKey string, message []byte)
}
