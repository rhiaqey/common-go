package executor

import (
	"log"

	"github.com/rhiaqey/common-go/pkg/env"
	"github.com/rhiaqey/common-go/pkg/rpc"
	"github.com/rhiaqey/sdk-go/pkg/sdk"
)

type Executor struct {
	env env.Env
}

func Setup(env *env.Env) (Executor, error) {
	return Executor{
		env: *env,
	}, nil
}

func (exe *Executor) GetID() string {
	return exe.env.Id
}

func (exe *Executor) GetName() string {
	return exe.env.Name
}

func (exe *Executor) GetNamespace() string {
	return exe.env.Namespace
}

func (exe *Executor) GetPrivatePort() uint16 {
	return exe.env.PrivatePort
}

func (exe *Executor) GetPublicPort() uint16 {
	return exe.env.PublicPort
}

func (exe *Executor) RPC(message rpc.RPCMessage) error {
	log.Print("rpc ")
	log.Println(message.Data())
	return nil
}

func (exe *Executor) Publish(message sdk.ProducerMessage) error {
	log.Print("pub ")
	log.Println(message)
	return nil
}
