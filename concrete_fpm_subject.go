package main

type ConcreteFPMSubject struct {
	Subject
}

func (fs *ConcreteFPMSubject) Register(*ConcreteFPMObserver) {

}

func (fs *ConcreteFPMSubject) RemoveObserver(*ConcreteFPMObserver) {

}

func (fs *ConcreteFPMSubject) NotifyAllObservers(*Event) {

}
