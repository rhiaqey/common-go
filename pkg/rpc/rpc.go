package rpc

type RPCMessage interface {
	Data() any
}

type RegisterPublisherRPCMessageData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Schema    []byte `json:"schema"`
}

func (msg RegisterPublisherRPCMessageData) Data() any {
	return msg
}
