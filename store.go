package main

type Store interface {
	GetMenu() (*Menu, error)
}
