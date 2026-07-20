// Package testingtooling provides small units suitable for table, fake, benchmark, and fuzz tests.
package testingtooling

import (
	"context"
	"errors"
	"strings"
)

var ErrEmailTaken = errors.New("email already registered")

func Sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

type UserStore interface {
	Exists(ctx context.Context, email string) (bool, error)
}

type RegistrationService struct {
	store UserStore
}

func NewRegistrationService(store UserStore) RegistrationService {
	return RegistrationService{store: store}
}

func (s RegistrationService) Register(ctx context.Context, email string) error {
	exists, err := s.store.Exists(ctx, NormalizeEmail(email))
	if err != nil {
		return err
	}
	if exists {
		return ErrEmailTaken
	}
	return nil
}
