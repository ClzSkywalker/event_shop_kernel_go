package constx

type RecoverFunc func()
type RecoverReadChanFunc func(ch <-chan DbInitStateType)
type RecoverWriteChanFunc func(ch chan<- DbInitStateType)

type CreateTableFunc func() error
type DropTableFunc func() error
type AutoMigrateFunc func() (err error)
