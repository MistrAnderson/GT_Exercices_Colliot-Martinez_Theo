package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for {
		fmt.Println("Que souhaitez-vous faire ?")
		fmt.Println("1. Récupérer tout le texte contenu")
		fmt.Println("2. Ajouter du texte dans ce fichier")
		fmt.Println("3. Supprimer tout le contenu du fichier")
		fmt.Println("4. Remplacer le contenu")
		fmt.Println("5. Quitter")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			readFile()
		case "2":
			appendToFile()
		case "3":
			clearFile()
		case "4":
			replaceFileContents()
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func readFile() {
	data, err := ioutil.ReadFile("truc.txt")
	if err != nil {
		fmt.Println("Impossible de lire le fichier")
		return
	}
	fmt.Println(string(data))
}

func appendToFile() {
	fmt.Println("Entrez le texte à ajouter :")
	var text string
	fmt.Scanln(&text)
	file, err := os.OpenFile("truc.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, text)
	if err != nil {
		fmt.Println("Impossible d'écrire dans le fichier")
		return
	}
	fmt.Println("Texte ajouté avec succès")
}

func clearFile() {
	err := ioutil.WriteFile("truc.txt", []byte(""), 0644)
	if err != nil {
		fmt.Println("Impossible de supprimer le contenu du fichier")
		return
	}
	fmt.Println("Contenu du fichier supprimé avec succès")
}

func replaceFileContents() {
	fmt.Println("Entrez le texte à écrire dans le fichier :")
	var text string
	fmt.Scanln(&text)
	err := ioutil.WriteFile("truc.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Impossible d'écrire dans le fichier")
		return
	}
	fmt.Println("Contenu du fichier remplacé avec succès")
}
