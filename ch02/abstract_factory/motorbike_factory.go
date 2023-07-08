package abstract_factory

import (
	"errors"
	"fmt"
)

type MotorbikeType int

const (
	SportMotorbikeType MotorbikeType = iota
	CruiseMotorbikeType
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) GetVehicle(mt MotorbikeType) (Vehicle, error) {
	switch mt {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", mt))
	}
}
