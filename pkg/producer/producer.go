package producer

import (
	"encoding/json"
	"log"

	"github.com/caarlos0/env/v11"
	rhiaqey_env "github.com/rhiaqey/common-go/pkg/env"
	"github.com/rhiaqey/common-go/pkg/executor"
	"github.com/rhiaqey/common-go/pkg/rpc"
	"github.com/rhiaqey/sdk-go/pkg/sdk"
)

func Run[T sdk.Producer, S any](producer T, settings S) {
	log.Println("running producer")

	// read env
	cfg := rhiaqey_env.Env{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err)
	}

	// create executor
	exec, err := executor.Setup(&cfg)
	if err != nil {
		log.Panic(err)
	}

	log.Printf(
		"producer [id=%s,name=%s] is ready",
		exec.GetID(),
		exec.GetName(),
	)

	// prepare producer
	comms, err := producer.Setup(settings)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: read settings here

	// TODO: prepare publisher_registration_message
	message := rpc.RegisterPublisherRPCMessageData{
		Id:        exec.GetID(),
		Name:      exec.GetName(),
		Namespace: exec.GetNamespace(),
		Schema:    []byte(producer.Schema()),
	}
	err = exec.RPC(message)
	if err != nil {
		log.Fatal(err)
	}

	go producer.Start()

	// TODO: private http server start to serve /alive /ready /metrics /version
	// TODO: read channels and set global counter TOTAL_CHANNELS.set(channel_count);

	log.Println("ready, set, go...")

	for {
		message := <-comms
		_, err := json.Marshal(message)
		if err != nil {
			log.Fatalln(err)
		}

		// publish message

		log.Println(message)
	}
}
