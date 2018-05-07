package main

type Subject interface {
	Register(*FPMObserver)
	RemoveObserver(*FPMObserver)
	NotifyAllObservers(*Event)
}

type FPMSubject struct {
	Subject
}

func (fs *FPMSubject) Register(*FPMObserver) {

}

func (fs *FPMSubject) RemoveObserver(*FPMObserver) {

}

func (fs *FPMSubject) NotifyAllObservers(*Event) {

}
