package main

type Observer interface {
	Update(*Event)
}
