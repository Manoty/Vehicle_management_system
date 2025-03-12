package main

import "fmt"

// Define a Vehicle interface
type Vehicle interface{
	Drive()

}
//define common structs
type Whip struct{
	Brand string
	Colour string
	Age string
	Speed int
}
//method for the motor (base behaviour)
func (w Whip)Drive(){
	fmt.Println("A", w.Colour, w.Brand , "that looks", w.Age, "just passed at a freaking speed of",w.Speed, "mph")


}

type Lorry struct{
	Whip
	Transmission string
}

type Offroader struct{
	Whip
	Hp int
}
type Sedan struct{
	Whip
	Boot string
}
type Truck struct{
	Whip
	Body string

}
func main(){
	//create instances
	lorry := Lorry{Whip: Whip{Brand:"FTR", Colour: "brown", Age: "New", Speed: 180}, Transmission: "manual"}
	sedan := Sedan{Whip: Whip{Brand:"Subaru", Colour: "black", Age: "New", Speed: 240}, Boot:"hatchback"}
	offroader := Offroader{Whip: Whip{Brand:"Ford Ranger", Colour: "White", Age: "second hand", Speed: 180}, Hp:250}
	truck := Truck{Whip: Whip{Brand:"Jeep Wrangler", Colour: "Green", Age: "New", Speed: 200}, Body:"hard body"}


	//add slices to store vehicles
	vehicles := []Vehicle{lorry, sedan, offroader, truck}

	//iterate over the slice and call the Drive method
	for _, vehicle := range vehicles{
		vehicle.Drive()
	}
}


