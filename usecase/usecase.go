package usecase

type Cmd interface {
	Execute() error
}
