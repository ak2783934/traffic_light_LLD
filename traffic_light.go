package trafficlightslld

import "fmt"

type TrafficLight struct {
	ID             string
	CurrentSignal  Signal
	RedDuration    int
	YellowDuration int
	GreenDuration  int
}

func NewTrafficLight(id string, redTime int, yellowTime int, greenTime int) *TrafficLight {
	return &TrafficLight{
		ID:             id,
		CurrentSignal:  Red,
		RedDuration:    redTime,
		YellowDuration: yellowTime,
		GreenDuration:  greenTime,
	}
}

func (TL *TrafficLight) ChangeSignal(newSignal Signal) {
	TL.CurrentSignal = newSignal
	TL.notifySignalChange()
}

func (TL *TrafficLight) notifySignalChange() {
	fmt.Printf(fmt.Sprintf("TrafficLigth id: %s Signal Changed to %s\n", TL.ID, TL.CurrentSignal))
}
