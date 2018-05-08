package main

type Subject interface {
	Register(*ConcreteFPMObserver)
	RemoveObserver(*ConcreteFPMObserver)
	NotifyAllObservers(*ConcreteFPMObserver)
}
