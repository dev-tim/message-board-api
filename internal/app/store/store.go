package store

type IStore interface {
	Messages() IMessagesRepository
}
