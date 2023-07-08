package abstract_factory

type CruiseMotorbike struct{}

// Vehicle
func (c *CruiseMotorbike) GetWheels() int {
	return 2
}

func (c *CruiseMotorbike) GetSeats() int {
	return 2
}

// Motorbike
func (c *CruiseMotorbike) GetType() MotorbikeType {
	return CruiseMotorbikeType
}
