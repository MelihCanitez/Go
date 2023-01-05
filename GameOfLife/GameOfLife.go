package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getSpielfeldGroesse() int{
	//Eingabe vom Nutzer
	fmt.Print("Wie gross soll das Spielfeld sein?: ")
	sizeReader := bufio.NewReader(os.Stdin)
	input, e := sizeReader.ReadString('\n')
	if e != nil {
		fmt.Println("Fehler beim einlesen der Spielfeldgroesse", e)
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r") // "\r" Wird zwar beim String nicht angezeigt, jedoch beim konvertieren wird dies mit beachtet
	
    //string to int
    groesse, e := strconv.Atoi(input)
    if e != nil {
        fmt.Println("Fehler bei der Konvertierung von String to Int", e)
    }
	return groesse
}

func erstelleSpielFeld(groesse int) [][]string{
	groesse = groesse +2 //Am Rand eine Reihe die nicht beachtet wird - später wichtig für die Regeln
	spielFeld := make([][]string, groesse) //Hier werden Slices genutzt - make()
	for i := range spielFeld {
		spielFeld[i] = make([]string, groesse)
	}
	
	for i := 0; i < groesse; i++{
		for j := 0; j < groesse; j++{
			if i == 0 || i == groesse-1 || j == 0 || j == groesse-1{
				spielFeld[i][j] = "Wird nicht beachtet"
			}else{
				spielFeld[i][j] = "X"
			}
		}
	}
	return spielFeld
}

func zeigeSpielFeldAn(spielFeld [][]string){
	fmt.Println()
	for i := 1; i < len(spielFeld)-1; i++{
		for j := 1; j < len(spielFeld[i])-1; j++{
			if(spielFeld[i][j] == "O"){ //Genau wie in Rust kann man auch in GoLang Ansi Escape Codes nutzen um die Ausgabe im Terminal zu verändern
				red := "\033[0;31m"
				reset := "\033[0m"
				fmt.Print(red + spielFeld[i][j] + reset + " ")
			}else{
				fmt.Print(spielFeld[i][j] + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func setzeSpielMarken(spielFeld [][]string){
		//Eingabe vom Nutzer
		fmt.Print("Wie viele Spielmarken möchten Sie setzen?: ")
		amountReader := bufio.NewReader(os.Stdin)
		input, e := amountReader.ReadString('\n')
		if e != nil {
			fmt.Println("Fehler beim einlesen der Spielmarken", e)
		}
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r")
		
		//string to int
		amount, e := strconv.Atoi(input)
		if e != nil {
			fmt.Println("Fehler bei der Konvertierung von String to Int", e)
		}
		
		for i := 0; i < amount; i++{
			outOfRange:
			fmt.Println("Geben Sie die Koordinaten der zu setzenden Spielmarke ein: ")
			xyReader := bufio.NewReader(os.Stdin)
			input, e := xyReader.ReadString('\n')
			if e != nil {
				fmt.Println("Fehler beim einlesen der Koordinaten", e)
			}
			input = strings.TrimSuffix(input, "\n")
			xStr := strings.Split(input, " ")[0]
			yStr := strings.Split(input, " ")[1]

			xStr = strings.TrimSuffix(xStr, "\r")
			yStr = strings.TrimSuffix(yStr, "\r")

			x, e := strconv.Atoi(xStr)
			if e != nil {
				fmt.Println("Fehler bei der Konvertierung von String to Int (Koordinaten)", e)
			}

			y, e := strconv.Atoi(yStr)
			if e != nil {
				fmt.Println("Fehler bei der Konvertierung von String to Int (Koordinaten)", e)
			}

			if x < 1 || x > len(spielFeld)-2 || y < 1 || y > len(spielFeld)-2{
				fmt.Println("Ausserhalb vom Spielfeld. Bitte bleiben Sie im Intervall!")
				goto outOfRange
			}

			spielFeld[y][x] = "O"
			zeigeSpielFeldAn(spielFeld)
		}
}

func starteGeneration(spielFeld [][]string){
		//Eingabe vom Nutzer
		fmt.Print("Wie viele Generationen soll es geben?: ")
		genReader := bufio.NewReader(os.Stdin)
		input, e := genReader.ReadString('\n')
		if e != nil {
			fmt.Println("Fehler beim einlesen der Anzahl an Generationen", e)
		}
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r") // "\r" Wird zwar beim String nicht angezeigt, jedoch beim konvertieren wird dies mit beachtet
		
		//string to int
		amountGen, e := strconv.Atoi(input)
		if e != nil {
			fmt.Println("Fehler bei der Konvertierung von String to Int", e)
		}
		fmt.Println("\n---------------")
		for i := 0; i < amountGen; i++{
			fmt.Printf("Generation %d:", (i+1))
			applyRules(spielFeld)
			zeigeSpielFeldAn(spielFeld)
		}
}

func applyRules(spielFeld [][]string){
	groesse := len(spielFeld) //Gleich grosser Slice (wie spielFeld), hier werden die Anzahl an lebenden Nachbarn gespeichert
	nachbarn := make([][]int, groesse)
	for i := range nachbarn {
		nachbarn[i] = make([]int, groesse)
	}
	
	for i := 1; i < len(spielFeld)-1; i++{
		counter := 0
		for j := 1; j < len(spielFeld)-1; j++{
			if spielFeld[i-1][j-1] == "O"{
				counter++
			}
            if spielFeld[i-1][j] == "O"{
				counter++
			}
            if spielFeld[i-1][j+1] == "O"{
				counter++
			}
            if spielFeld[i][j-1] == "O"{
				counter++
			}
            if spielFeld[i][j+1] == "O"{
				counter++
			}
            if spielFeld[i+1][j-1] == "O"{
				counter++
			}
            if spielFeld[i+1][j] == "O"{
				counter++
			}
            if spielFeld[i+1][j+1] == "O"{
				counter++
			}
			nachbarn[i-1][j-1] = counter
			counter = 0
		}
	}

	for i := 1; i < len(spielFeld)-1; i++{
		for j := 1; j < len(spielFeld)-1; j++{
			if spielFeld[i][j] == "X"{ //Regeln für tote Zellen
                if nachbarn[i-1][j-1] == 3{ //"Zum Leben erweckt" wenn eine tote Zelle genau 3 lebende Nachbarn hat
                    spielFeld[i][j] = "O";
                }
            }else{ //Regeln für lebende Zellen
                if nachbarn[i-1][j-1] < 2{ //"Sterben an Einsamkeit" wenn eine lebende Zelle weniger als 2 lebende Nachbarn hat
                    spielFeld[i][j] = "X";
                }else if (nachbarn[i-1][j-1] == 2) || (nachbarn[i-1][j-1] == 3){ //Eine lebende Zelle lebt in der Folgegeneration, wenn diese genau 2 oder 3 Nachbarn hat
                    continue;
                }else if nachbarn[i-1][j-1] > 3{ //"Sterben an Überbevölkerung" wenn eine lebende Zelle mehr als 3 lebende Nachbarn hat
                    spielFeld[i][j] = "X";
                }else{
                    fmt.Println("Test...Dieser Fall sollte nicht ausgelöst werden.");
                }
            }
		}
	}
}

func main(){
	var groesse int = getSpielfeldGroesse()
	spielFeld := erstelleSpielFeld(groesse)
	zeigeSpielFeldAn(spielFeld)
	setzeSpielMarken(spielFeld)
	starteGeneration(spielFeld)
}