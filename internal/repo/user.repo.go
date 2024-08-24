package repo

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) GetInfoUser() string {
	return "user"
}
