package main

import "fmt"

var (
	amnen = []string{"Bild", "Biologi", "Engelska", "Fysik", "Geografi", "Hemkunskap", "Historia", "Idrott", "Kemi", "Matematik", "Moderna Språk", "Musik", "Religionkunskap", "Samhällskunskap", "Slöjd", "Svenska", "Teknik"}
	betyg = make(map[string]float32)
	svar  string
	total float32
)

func main() {
	betyg["A"] = 20
	betyg["B"] = 17.5
	betyg["C"] = 15
	betyg["D"] = 12.5
	betyg["E"] = 10
	betyg["F"] = 0
	fmt.Printf("OBS: Stora bokstäver!\n")
	for i := 0; i < cap(amnen); i++ {
		fmt.Printf("%v betyg: ", amnen[i])
		fmt.Scanf("%v", &svar)
		total += betyg[svar]
	}
	fmt.Printf("Ditt meritvärde är: %v\nKämpa på!\n", total)
}
