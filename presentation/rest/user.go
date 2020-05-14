package rest

// プレゼンテーション層の実装
// RESTful API用

import (
	"context"

	goa "github.com/yyh-gl/goa-tutorial/gen/user"
	"github.com/yyh-gl/goa-tutorial/usecase"
)

type user struct {
	u usecase.User
}

func NewUser(u usecase.User) goa.Service {
	return &user{u: u}
}

func (u user) Create(ctx context.Context, p *goa.CreatePayload) (*goa.CreateResult, error) {
	// ユーザ作成に関するビジネスロジックは
	// usecase層以下にあるため（今回で言えば Register()内にメインのロジックがある）
	// gRPCに乗り換えるときはpresentation/restだけを修正すればOK
	ru := u.u.Register(p.Name, p.Age)
	return &goa.CreateResult{
		User: &goa.User{
			Name: ru.Name,
			Age:  ru.Age,
		},
	}, nil
}

func (u user) Index(ctx context.Context) (*goa.IndexResult, error) {
	return nil, nil
}
