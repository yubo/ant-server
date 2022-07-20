package services

import (
	"context"
)

type services struct{}

func New() *services {
	return &services{}
}

func (p *services) Start(ctx context.Context) error {

	// /api/v1/account
	if err := newAccount().install(ctx); err != nil {
		return err
	}

	// /api/v1/users
	if err := newUser().install(ctx); err != nil {
		return err
	}

	return nil
}
