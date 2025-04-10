package trafficlightslld

type Road struct {
	ID           string
	Name         string
	TrafficLight *TrafficLight
}

func NewRoad(id, name string) *Road {
	return &Road{
		ID:   id,
		Name: name,
	}
}

func (rd *Road) SetTrafficLight(trafficLight *TrafficLight) {
	rd.TrafficLight = trafficLight
}
