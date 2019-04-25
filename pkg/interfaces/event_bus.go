package interfaces

type EventBus interface {
	StartListen() error
	Shutdown()
}
