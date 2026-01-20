package main

import "fmt"

func main() {
	fmt.Println("Bienvenue sur Tic Tac Toe :)")
	fmt.Print("\ndebut du jeux\n\n")

	var tableau [3][3]uint8
	var utilisateur uint8 = 1
	victoire := false
	matchNul := false

	for !victoire && !matchNul {
		affichTab(tableau)

		var placer bool
		var caseX, caseY int

		if utilisateur == 2 {
			fmt.Println("L'IA", utilisateur, "réfléchit...")
			caseX, caseY = meilleurCoup(tableau, utilisateur)
			placer = placeChoix(&tableau, caseX, caseY, utilisateur)
		} else {
			fmt.Println("L'IA", utilisateur, "réfléchit...")
			caseX, caseY = meilleurCoup(tableau, utilisateur)
			placer = placeChoix(&tableau, caseX, caseY, utilisateur)
		}

		if placer {
			victoire = detectionsVictoire(tableau, utilisateur)
			matchNul = verifieMatchNul(tableau)

			if !victoire && !matchNul {
				switch utilisateur {
				case 1:
					utilisateur = 2
				case 2:
					utilisateur = 1
				}
			}
		}
	}

	affichTab(tableau)
	if victoire {
		fmt.Println("Félicitations à l'utilisateur", utilisateur, "! Vous avez gagné !")
	} else if matchNul {
		fmt.Println("Match nul ! La partie est terminée.")
	}
}