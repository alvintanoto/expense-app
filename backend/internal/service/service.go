package service

import "go.uber.org/dig"

type Holder struct {
	dig.In
}

func Register(container *dig.Container) error {
	return nil
}
