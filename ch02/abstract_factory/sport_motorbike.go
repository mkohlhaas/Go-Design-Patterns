package abstract_factory

type SportMotorbike struct{}

// Vehicle
func (s *SportMotorbike) GetWheels() int {
	return 2
}

func (s *SportMotorbike) GetSeats() int {
	return 1
}

// Motorbike
func (s *SportMotorbike) GetType() MotorbikeType {
	return SportMotorbikeType
}
