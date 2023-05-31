package main

import (
	"fmt"
)

func main() {
	laura := Engine{}
	laura.Learn("greeting", "hello hi ok-lol ok-lol")
	laura.Learn("weather", "weather how is the weather like")
	laura.Learn("weather", "hvordan er været")
	fmt.Println(laura.GetLikelyIntent("været ok været"))
}
