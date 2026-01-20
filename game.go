package main

import "fmt"

// affiche le contenue de la liste sous forme d'un tableu
func affichTab(tab [3][3]uint8) {
	for x := 2; x >= 0; x-- {
		fmt.Print(x + 1, "|")
		for y := range tab[x] {
			fmt.Print(tab[x][y], " ")
		}
		fmt.Println("")
	}

	fmt.Println(" -------")
	fmt.Println("  1 2 3")
}

// permet de connaitre le choix de l'utilisateur retourne les cordonée choisi si il sont corecte et que la place est vide
func choixUtile(tab [3][3]uint8, utilsateur uint8) (int, int) {
	var (
		ligne   int
		colonne int
	)

	fmt.Println("c'est a l'utilsateur", utilsateur, "de jouer")
	fmt.Print("Entre la ligne choisie (1-3) : ")
	fmt.Scanln(&ligne)

	fmt.Print("Entre la colonne choisie (1-3) : ")
	fmt.Scanln(&colonne)

	dansLesLimites := ligne >= 1 && ligne <= 3 && colonne >= 1 && colonne <= 3

	ligne = ligne - 1
	colonne = colonne - 1

	if dansLesLimites && tab[ligne][colonne] == 0 {
		return ligne, colonne
	} else {
		fmt.Println("Coordonnées invalides ou case déjà prise !")
		return -1, -1
	}
}

// permet de modifier un nombre a des cordonée choisie
func placeChoix(tab *[3][3]uint8, x int, y int, symbole uint8) bool {
	if x != -1 && y != -1 {
		tab[x][y] = symbole
		return true
	} else {
		return false
	}
}

// detectionsVictoire vérifie s'il y a une condition de victoire pour le symbole donné.
func detectionsVictoire(tab [3][3]uint8, symbole uint8) bool {
	// Vérification des lignes
	for i := range 3 {
		if tab[i][0] == symbole && tab[i][1] == symbole && tab[i][2] == symbole {
			return true
		}
	}

	// Vérification des colonnes
	for i := range 3 {
		if tab[0][i] == symbole && tab[1][i] == symbole && tab[2][i] == symbole {
			return true
		}
	}

	// Vérification des diagonales
	if tab[0][0] == symbole && tab[1][1] == symbole && tab[2][2] == symbole {
		return true
	}
	if tab[0][2] == symbole && tab[1][1] == symbole && tab[2][0] == symbole {
		return true
	}

	return false
}

// cette fonction remet a zero le tableau
func resetTableau(tab *[3][3]uint8) {
	for x := 2; x >= 0; x-- {
		for y := range tab[x] {
			tab[x][y] = 0
		}
	}
}

func verifieMatchNul(tab [3][3]uint8) bool {
	// D'abord, vérifier si le plateau est plein.
	for _, ligne := range tab {
		for _, caseVal := range ligne {
			if caseVal == 0 {
				return false
			}
		}
	}

	// Si le plateau est plein, vérifier si un joueur a gagné.
	if detectionsVictoire(tab, 1) || detectionsVictoire(tab, 2) {
		// Si un joueur a gagné, ce n'est pas un match nul.
		return false
	}

	// Si le plateau est plein et que personne n'a gagné, c'est un match nul.
	return true
}
