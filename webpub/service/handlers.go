package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/sarrufat/natsk8s/pub/rand"
	"io"
	"log"
	"net/http"
)

type (
	MessageProducer struct {
		logger *log.Logger
		cli    *NatsClient
	}
	NatsTestRequest struct {
		Subject     string `json:"subject"`
		NumMsg      int    `json:"num_msg"`
		MessageSize int    `json:"message_size"`
	}
)

func (nt *NatsTestRequest) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(nt)
}
func (mp *MessageProducer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	mp.logger.Println("ServeHTTP ", request.Method)
	if request.Method == http.MethodPost {
		natsTest := &NatsTestRequest{}
		err := natsTest.FromJSON(request.Body)
		if err != nil {
			http.Error(writer, "Unable to marshal json", http.StatusInternalServerError)
			mp.logger.Print(err)
		}
		mp.cli.Do(natsTest.Prepare())
	}
}

func (nt *NatsTestRequest) Prepare() NatsAction {
	return func(conn *nats.Conn) {
		for cnt := 0; cnt < nt.NumMsg; cnt++ {
			data := rand.String(nt.MessageSize)
			conn.Publish(nt.Subject, []byte(data))
		}
	}
}
func NewMessageProducer(logger *log.Logger, url string) *MessageProducer {
	return &MessageProducer{
		logger: logger,
		cli:    &NatsClient{url: url},
	}
}
