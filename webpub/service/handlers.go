package service

import (
	"encoding/json"
	"fmt"
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
	NatsTestResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)

func (nt *NatsTestRequest) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(nt)
}
func (nr *NatsTestResponse) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(nr)
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
		outcome := mp.cli.Do(natsTest.Prepare())
		resp := &NatsTestResponse{
			Status:  http.StatusOK,
			Message: outcome,
		}
		resp.ToJSON(writer)
	}
}

func (nt *NatsTestRequest) Prepare() NatsAction {
	return func(conn *nats.Conn) string {
		for cnt := 0; cnt < nt.NumMsg; cnt++ {
			data := rand.String(nt.MessageSize)
			conn.Publish(nt.Subject, []byte(data))
		}
		return fmt.Sprintf("%d messages of %d sent", nt.NumMsg, nt.MessageSize)
	}
}
func NewMessageProducer(logger *log.Logger, url string) *MessageProducer {
	return &MessageProducer{
		logger: logger,
		cli:    &NatsClient{url: url},
	}
}
