package asciiart

import (
	"bufio"
	"log"
	"os"
)

func GetTheSymboles(banner string) [][]string {
	file, err := os.Open("./Files/" + banner + ".txt")
	if err != nil {
		log.Fatal("Erreur:", err)
	}
	defer file.Close()                // Assure la fermeture du fichier après usage
	scanner := bufio.NewScanner(file) // Initialise un scanner pour lire le fichier ligne par ligne
	count := 0
	symbole := []string{}
	symboles := [][]string{}

	// Boucle à travers chaque ligne du fichier texte
	for scanner.Scan() {
		line := scanner.Text() + " "
		symbole = append(symbole, line)
		count++

		// Chaque symbole ASCII dans ce format est constitué de 9 lignes
		if count == 9 {
			symboles = append(symboles, symbole) // Ajoute le symbole complet à la liste de symboles
			symbole = []string{}
			count = 0
		}
	}
	// Vérifie si le fichier contient au moins 95 symboles.
	if len(symboles) < 95 {
		log.Fatal("Please make sure all characters are present in the file.")
	}
	return symboles
}
