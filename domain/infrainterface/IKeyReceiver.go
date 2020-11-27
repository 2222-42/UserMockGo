package infrainterface

type IKeyReceiver interface {
	ReceiveSecretKey() interface{}
	ReceivePublicKey() interface{}
}
