package usecase

// ユースケース（アプリケーション）層の実装

type User struct{}

func NewUser() User {
	return User{}
}

type RegisteredUser struct {
	Name string
	Age  int
}

func (u User) Register(name string, age int) *RegisteredUser {
	return &RegisteredUser{
		Name: name,
		Age:  age,
	}
}
