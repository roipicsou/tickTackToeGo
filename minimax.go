package main

// meilleurCoup analyse le plateau actuel et retourne la meilleure position (ligne, colonne)
// pour le joueur donné (ia).
// Il retourne (-1, -1) si aucun coup n'est possible.
func meilleurCoup(tab [3][3]uint8, ia uint8) (int, int) {
	bestScore := -1000
	moveX, moveY := -1, -1

	// Détermine l'ID de l'adversaire
	adversaire := (ia % 2) + 1

	// On parcourt toutes les cases pour trouver les coups possibles
	for x := range 3 {
		for y := range 3 {
			// Si la case est vide
			if tab[x][y] == 0 {
				// 1. On joue le coup
				tab[x][y] = ia

				// 2. On appelle minimax pour évaluer ce coup
				// depth = 0, isMaximizing = false (c'est au tour de l'adversaire)
				score := minimax(tab, 0, false, ia, adversaire)

				// 3. On annule le coup (backtracking)
				tab[x][y] = 0

				// 4. Si ce coup est meilleur que le précédent, on le garde
				if score > bestScore {
					bestScore = score
					moveX = x
					moveY = y
				}
			}
		}
	}

	return moveX, moveY
}

// minimax est la fonction récursive qui évalue l'arbre des coups possibles.
func minimax(tab [3][3]uint8, depth int, isMaximizing bool, ia uint8, adversaire uint8) int {
	// --- Conditions d'arrêt (Terminal states) ---

	// Si l'IA gagne
	if detectionsVictoire(tab, ia) {
		return 10 - depth // On favorise les victoires rapides
	}
	// Si l'adversaire gagne
	if detectionsVictoire(tab, adversaire) {
		return depth - 10 // On favorise les défaites lentes (résistance)
	}
	// Si match nul (plateau plein sans vainqueur)
	if verifieMatchNul(tab) {
		return 0
	}

	// --- Récursion ---

	if isMaximizing {
		// Tour de l'IA : on cherche à maximiser le score
		maxScore := -1000
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if tab[x][y] == 0 {
					tab[x][y] = ia
					score := minimax(tab, depth+1, false, ia, adversaire)
					tab[x][y] = 0
					if score > maxScore {
						maxScore = score
					}
				}
			}
		}
		return maxScore
	} else {
		// Tour de l'adversaire : il joue pour minimiser notre score
		minScore := 1000
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if tab[x][y] == 0 {
					tab[x][y] = adversaire
					score := minimax(tab, depth+1, true, ia, adversaire)
					tab[x][y] = 0
					if score < minScore {
						minScore = score
					}
				}
			}
		}
		return minScore
	}
}