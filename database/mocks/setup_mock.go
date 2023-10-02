package mocks

import "github.com/stretchr/testify/mock"

type CustomMock struct {
	*mock.Mock
}

func (s *Services) Close() error {
	return nil
}
