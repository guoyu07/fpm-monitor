package main

type Observer interface {
	Update(*Event)
}

type FPMObserver struct {
	Observer
}

func (fsb *FPMObserver) Update(e *Event) {

}
