package uaa

type Command interface {
	Execute() error
}
