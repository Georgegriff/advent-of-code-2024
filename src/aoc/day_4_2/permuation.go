package main

type Permutation struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
}

func (p *Permutation) check() int {

	// S.S
	// .A.
	// M.M
	if p.TopLeft == "S" && p.TopRight == "S" && p.BottomLeft == "M" && p.BottomRight == "M" {
		return 1
	}

	// M.M
	// .A.
	// S.S
	if p.TopLeft == "M" && p.TopRight == "M" && p.BottomLeft == "S" && p.BottomRight == "S" {
		return 1
	}

	// S.M
	// .A.
	// S.M
	if p.TopLeft == "S" && p.TopRight == "M" && p.BottomLeft == "S" && p.BottomRight == "M" {
		return 1
	}

	// M.S
	// .A.
	// M.S
	if p.TopLeft == "M" && p.TopRight == "S" && p.BottomLeft == "M" && p.BottomRight == "S" {
		return 1
	}

	return 0
}
