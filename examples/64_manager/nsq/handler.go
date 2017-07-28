package nsq

import nsqlib "github.com/nsqio/go-nsq"

type IHandler interface {
	HandleMessage(message *nsqlib.Message) error
}
