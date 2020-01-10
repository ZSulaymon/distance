package main

import "testing"

func Test_distance(t *testing.T) {
  tests := []struct {
    name         string
    fuelConsumed int
    amountOfFuel int
    want         int
  }{
    {"More than hundred km", 7, 14, 200},
    {"Exactly hundred km", 9, 9, 100},
    {"Less than hundred km", 10, 1, 10},
  }
  for _, test := range tests {
    got := distance(test.fuelConsumed, test.amountOfFuel)
    if got != test.want {
      t.Error("Test fail:", test.name, "got:", got, "want:", test.want)
    }
  }
}
