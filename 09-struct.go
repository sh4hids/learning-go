package main

import "fmt"

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_km   float64
}

func main() {
	car_a := car{gas_pedal: 22341, brake_pedal: 23423, steering_wheel: 344, top_speed_km: 123.456}
	car_b := car{14354, 2354, 54, 2114.2}

	fmt.Println(car_a)
	fmt.Println(car_b)
}
