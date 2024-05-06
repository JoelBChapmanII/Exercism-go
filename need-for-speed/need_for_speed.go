package speed

// TODO: define the 'Car' type struct
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	battery_tmp := car.battery - car.batteryDrain
	// If battery can't accommodate the drain return the car as is
	if battery_tmp < 0 {
		return car
	}
	car.battery = battery_tmp
	car.distance = car.distance + car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// Get how far the car can travel + the current distance the car has already traveled
	// Check if the total amount the car can travel is greater or equal to the track length
	return (car.battery/car.batteryDrain*car.speed + car.distance) >= track.distance
}
