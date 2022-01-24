package interceptor

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type MyInterceptor struct {
}

func NewMyInterceptor() MyInterceptor {
	return MyInterceptor{}
}

func (i MyInterceptor) OnSend(msg *sarama.ProducerMessage) {

	value, err := msg.Value.Encode()
	if err != nil {
		log.Printf("[Interceptor] encode error: %v", err)
		return
	}
	msg.Value = sarama.StringEncoder(fmt.Sprintf("[sarama] %s", string(value)))
}