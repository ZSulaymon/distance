package main

func main() {

}
func distance(fuelConsumed int, amountOfFuel int) int {
  const distanceInKm int = 100
  carCanRideFor := distanceInKm * amountOfFuel / fuelConsumed
  return carCanRideFor
}
