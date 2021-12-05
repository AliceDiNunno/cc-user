package usecases

type Logger interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type UserRepo interface {
}

type interactor struct {
	userRepo UserRepo
}

func NewInteractor(u UserRepo) interactor {
	return interactor{
		userRepo: u,
	}
}
