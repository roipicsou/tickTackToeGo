package main

import "testing"

func TestDetectionsVictoire(t *testing.T) {
	tests := []struct {
		name     string
		tab      [3][3]uint8
		symbole  uint8
		expected bool
	}{
		{
			name: "Victoire ligne",
			tab: [3][3]uint8{
				{1, 1, 1},
				{0, 2, 0},
				{2, 0, 0},
			},
			symbole:  1,
			expected: true,
		},
		{
			name: "Victoire colonne",
			tab: [3][3]uint8{
				{2, 1, 0},
				{2, 1, 0},
				{2, 0, 1},
			},
			symbole:  2,
			expected: true,
		},
		{
			name: "Victoire diagonale",
			tab: [3][3]uint8{
				{1, 0, 2},
				{0, 1, 0},
				{2, 0, 1},
			},
			symbole:  1,
			expected: true,
		},
		{
			name: "Pas de victoire",
			tab: [3][3]uint8{
				{1, 2, 1},
				{2, 1, 2},
				{2, 1, 2},
			},
			symbole:  1,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := detectionsVictoire(test.tab, test.symbole)
			if result != test.expected {
				t.Errorf("attendu %v, obtenu %v", test.expected, result)
			}
		})
	}
}

func TestVerifieMatchNul(t *testing.T) {
	tests := []struct {
		name     string
		tab      [3][3]uint8
		expected bool
	}{
		{
			name: "Grille pleine - match nul",
			tab: [3][3]uint8{
				{1, 2, 1},
				{2, 1, 2},
				{2, 1, 2},
			},
			expected: true,
		},
		{
			name: "Encore des coups possibles",
			tab: [3][3]uint8{
				{1, 0, 2},
				{0, 1, 0},
				{2, 0, 0},
			},
			expected: false,
		},
		{
			name: "Grille non pleine",
			tab: [3][3]uint8{
				{1, 2, 1},
				{2, 1, 2},
				{2, 1, 0},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := verifieMatchNul(test.tab)
			if result != test.expected {
				t.Errorf("attendu %v, obtenu %v", test.expected, result)
			}
		})
	}
}
