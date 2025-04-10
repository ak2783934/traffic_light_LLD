package trafficlightslld

import (
	"sync"
	"time"
)

var instance *TrafficController
var once sync.Once

type TrafficController struct {
	roads map[string]*Road
	mu    sync.Mutex
}

func NewTrafficController() *TrafficController {
	once.Do(func() {
		instance = &TrafficController{
			roads: make(map[string]*Road),
		}
	})
	return instance
}

func (tc *TrafficController) AddRoad(road *Road) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.roads[road.ID] = road
}

func (tc *TrafficController) StartTrafficController() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	for _, road := range tc.roads {
		trafficLight := road.TrafficLight
		go func(tl *TrafficLight) {
			for {
				time.Sleep(time.Duration(tl.RedDuration) * time.Millisecond)
				tl.ChangeSignal(Green)
				time.Sleep(time.Duration(tl.GreenDuration) * time.Millisecond)
				tl.ChangeSignal(Yellow)
				time.Sleep(time.Duration(tl.YellowDuration) * time.Millisecond)
				tl.ChangeSignal(Red)
			}
		}(trafficLight)
	}
}

func (tc *TrafficController) HandleEmergencySignal(roadID string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	road, exists := tc.roads[roadID]
	if exists {
		trafficLight := road.TrafficLight
		if trafficLight != nil {
			trafficLight.ChangeSignal(Green)
		}
	} else {
		panic("Road not found")
	}
}
