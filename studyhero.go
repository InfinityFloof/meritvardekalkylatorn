package main

import "fmt"

var (
	personbetyg     = make(map[string]float32)
	steg            = make(map[string]float32)
	amnen           = []string{"Bild", "Biologi", "Engelska", "Fysik", "Geografi", "Hemkunskap", "Historia", "Idrott", "Kemi", "Matematik", "Moderna Språk", "Musik", "Religionkunskap", "Samhällskunskap", "Slöjd", "Svenska", "Teknik"}
	ezamnen         = []string{}
	hdamnen         = []string{}
	betyg           = make(map[string]float32)
	svar            string
	metod           int
	wishmerit       float32
	differencemerit float32
	totalmerit      float32
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
		personbetyg[amnen[i]] = betyg[svar]
		totalmerit += betyg[svar]
	}
	fmt.Printf("Ditt meritvärde är: %v\n", totalmerit)
	fmt.Printf("Om du inte når dina mål så är det inte kört.\n")
	fmt.Printf("Vi tittar vilka ämnen du tykcer är svåra\n")
	for i := 0; i < len(amnen); i++ {
		fmt.Printf("Tycker du %v är svårt? [J/N]: ", amnen[i])
		fmt.Scanf("%v", &svar)
		if svar == "J" {
			hdamnen = append(hdamnen, amnen[i])
		} else {
			ezamnen = append(ezamnen, amnen[i])
		}
	}
	fmt.Printf("Vad är ditt önskade meritvärde?: ")
	fmt.Scanf("%v", &wishmerit)
	differencemerit = wishmerit - totalmerit
	fmt.Printf("För att få ditt önskade meritvärde måste vi uppnå %v poäng till\n", differencemerit)
	fmt.Printf("Det finns flera sätt vi kan lösa det här på\n")
	fmt.Printf("1: Plugga flera ämnen men inte behöva höja mycket i dem\n")
	fmt.Printf("2: Plugge färre ämnen men behöva uppnå högt i dem\n")
	fmt.Printf("Vad vill du göra?: ")
	fmt.Scanf("%v", &metod)
	switch metod {
	case 1:
		for x := 0; x < 10; x++ {
			// Försök 1, lätta betyg
			for i := 0; i < cap(ezamnen); i++ {
				if differencemerit > 0 {
					if personbetyg[ezamnen[i]] != 20 {
						personbetyg[ezamnen[i]] += 2.5
						steg[ezamnen[i]] += 1
						differencemerit -= 2.5
					}
				}
			}
			// Försök 2, lätta betyg
			for i := 0; i < cap(ezamnen); i++ {
				if differencemerit > 0 {
					if personbetyg[ezamnen[i]] != 20 {
						personbetyg[ezamnen[i]] += 2.5
						steg[ezamnen[i]] += 1
						differencemerit -= 2.5
					}
				}
			}
			// Försök 3, lätta betyg
			for i := 0; i < cap(ezamnen); i++ {
				if differencemerit > 0 {
					if personbetyg[ezamnen[i]] != 20 {
						personbetyg[ezamnen[i]] += 2.5
						steg[ezamnen[i]] += 1
						differencemerit -= 2.5
					}
				}
			}
			// Försök 3, svåra betyg
			for i := 0; i < cap(hdamnen); i++ {
				if differencemerit > 0 {
					if personbetyg[hdamnen[i]] != 20 {
						personbetyg[hdamnen[i]] += 2.5
						steg[hdamnen[i]] += 1
						differencemerit -= 2.5
					}
				}
			}
			// Försök 4, svåra betyg
			for i := 0; i < cap(hdamnen); i++ {
				if differencemerit > 0 {
					if personbetyg[hdamnen[i]] != 20 {
						personbetyg[hdamnen[i]] += 2.5
						steg[hdamnen[i]] += 1
						differencemerit -= 2.5
					}
				}
			}
		}
	case 2:
		// Försök 1, Lätta betyg
		for x := 0; x <= 20; x++ {
			for i := 0; i < cap(ezamnen); i++ {
				if differencemerit > 0 {
					if personbetyg[ezamnen[i]] <= 20 {
						personbetyg[ezamnen[i]] += 2.5
						steg[ezamnen[i]] += 1
						differencemerit -= 2.5
					}
				}
			}
		}
		// Försök 2, Svåra betyg
		for x := 0; x <= 20; x++ {
			for i := 0; i < cap(hdamnen); i++ {
				if differencemerit > 0 {
					if personbetyg[hdamnen[i]] <= 20 {
						personbetyg[hdamnen[i]] += 2.5
						steg[hdamnen[i]] += 1
						differencemerit -= 2.5
					}
				}
			}
		}
	}
	fmt.Printf("Du kan öka i (Steg) för att nå ditt mål:\n")
	for i := 0; i < cap(amnen); i++ {
		fmt.Printf("%v: %v steg\n", amnen[i], steg[amnen[i]])
	}

}
