package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	t.Run("Car builder", func(t *testing.T) {
		carBuilder := &CarBuilder{}
		new(ManufacturingDirector).SetBuilder(carBuilder).Construct()

		car := carBuilder.GetVehicle()

		assertWheels(t, "Car", 4, car.Wheels)
		assertStructure(t, "Car", "Car", car.Structure)
		assertSeats(t, "Car", 5, car.Seats)
	})

	t.Run("Bike Builder", func(t *testing.T) {
		bikeBuilder := &BikeBuilder{}
		new(ManufacturingDirector).SetBuilder(bikeBuilder).Construct()

		motorbike := bikeBuilder.GetVehicle()

		assertWheels(t, "Motorbike", 2, motorbike.Wheels)
		assertStructure(t, "Motorbike", "Motorbike", motorbike.Structure)
		assertSeats(t, "Motorbike", 2, motorbike.Seats)
	})

	t.Run("Bus Builder", func(t *testing.T) {
		busBuilder := &BusBuilder{}
		new(ManufacturingDirector).SetBuilder(busBuilder).Construct()

		bus := busBuilder.GetVehicle()

		assertWheels(t, "Bus", 8, bus.Wheels)
		assertStructure(t, "Bus", "Bus", bus.Structure)
		assertSeats(t, "Bus", 30, bus.Seats)
	})
}

func assertWheels(t *testing.T, vehicle string, expected int, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("Wheels on a %s must be %d and they were %d.\n", vehicle, expected, got)
	}
}

func assertStructure(t *testing.T, vehicle string, expected string, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("Structure on a %s must be %s and they were %s.\n", vehicle, expected, got)
	}
}

func assertSeats(t *testing.T, vehicle string, expected int, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("Seats on a %s must be %d and they were %d.\n", vehicle, expected, got)
	}
}
