package main
import (
	"fmt"
)

type city struct{
	climate string
	name string

}



func main(){
	// User includes input either via CLI or via the webfrontend to select their region
	var defaultCities = []city {
		{climate: "costal", name: "Pisa"},
		{climate: "cool", name: "Aberdeen"},
		{climate: "arctic", name: "Narvik"},
		{climate: "temperate", name: "Colmar"},
		{climate: "warm", name: "Athens"},
		{climate: "desert", name: "Algiers"},
		{climate: "jungle", name: "Cayenne"},
		{climate: "nothern coast", name: "Hamburg"},
		{climate: "southern coast", name: "Marseille"},
		{climate: "mountainous", name: "Innsbruck"},
		{climate: "desert mountains", name: "Almería"},
	}
	for i, s := range defaultCities {
		fmt.Println(i, s.climate)
	}
	var selection int
	fmt.Println("Enter the index number of the location you wish to use")
	fmt.Scanln(&selection)
	var stringArgument = defaultCities[selection]
	fmt.Println(stringArgument)
	//fmt.Printf("Starting up data collection!")
	//fmt.Println(defaultCities)
}

func connectToApiService() {
	//key= f0fa4704d8ba2c556ca8238a82f9aa75


}