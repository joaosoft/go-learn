package main

import (
	"time"

	"go-app/app"
	"go-app/app/api"
	"go-app/app/connectors/mqueue"
	controllers "go-app/samples/api"
	"fmt"
)

// AppConfig ...
type AppConfig struct {
	NSQ     map[string]*mqueue.NSQConfig `mapstruct:"nsq"`
	Logging struct {
		Mode  string `mapstruct:"mode"`
		Level string `mapstruct:"level"`
	} `mapstruct:"logging"`
}

func main() {
	//var config AppConfig
	//_, err := app.LoadConfig("go-app", "app-1", &config)
	//if err != nil {
	//	panic(err)
	//}

	//app.ConfigureLogger(config.Logging.Mode, config.Logging.Level)

	//pman := app.NewProcessManager()

	// sample NSQ consumer
	//p0 := mqueue.NewNSQConsumer(config.NSQ["consumer-1"])
	// sample WebAPI
	//p1 := api.NewWebAPI("localhost:8080")
	// sample gRPC API
	p2 := api.NewGRPC("localhost:8081")

	// a sample implementation of an NSQ consumer controller
	//nsqController := &nsqSampleController{}
	//p0.SetHandler(nsqController, 1)

	// a sample implementation of a gRPC controller
	sample := controllers.NewSampleRPC(p2.Server)

	fmt.Println("START")
	//pman.AddProcess("API 0", p0)
	//pman.AddProcess("API 1", p1)
	//pman.AddProcess("API 3", p2)

	go p2.Start()

	response, err := sample.Ping(nil, nil)
	fmt.Println("MSG:", response, "ERROR:", err)

}

type nsqSampleController struct{}

func (api *nsqSampleController) HandleRequest(msg mqueue.IMessage) error {
	data := string(msg.GetContent())

	if data == "requeue" {
		switch msg.GetRetries() {
		case 1:
			app.Logger.Infof("requeuing message [type:%s][id:%s][retries:%d]", msg.GetContentType(), msg.GetID(), msg.GetRetries())
			msg.Requeue()
		case 2:
			app.Logger.Infof("requeuing message for 10 seconds [type:%s][id:%s][retries:%d]", msg.GetContentType(), msg.GetID(), msg.GetRetries())
			msg.RequeueWithDelay(10 * time.Second)
		default:
			app.Logger.Infof("requeued message processed [type:%s][id:%s][retries:%d]", msg.GetContentType(), msg.GetID(), msg.GetRetries())
			msg.Processed()
		}
	}

	if data == "skip processed" {
		return nil
	}

	app.Logger.Infof("requeued message processed [type:%s][id:%s]", msg.GetContentType(), msg.GetID())
	if err := msg.Processed(); err != nil {
		app.Logger.LogError(err, "error marking message as processed")
	}

	return nil
}
