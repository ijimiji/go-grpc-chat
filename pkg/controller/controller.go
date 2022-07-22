package controller

type Sendable interface {
	Send(string) error
	Run(string)
}

type Updatable interface {
	Update(string, string) error
}
