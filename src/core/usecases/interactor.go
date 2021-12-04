package usecases

type Logger interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type interactor struct {
}

func NewInteractor() interactor {
	return interactor{}
}
