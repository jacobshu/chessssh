package game

import "fmt"

type Rank int

const (
	Rank1 Rank = iota + 1
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

func (r Rank) Index() int {
	return int(r)
}

func (r Rank) String() string {
	return fmt.Sprintf("%d", int(r))
}

type File int

const (
	FileA File = iota + 1
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) Index() int {
	return int(f)
}

func (f File) String() string {
	return [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}[f-1]
}

type Position struct {
	Rank Rank
	File File
}

func (p Position) String() string {
	return fmt.Sprintf("%s%s", p.Rank.String(), p.File.String())
}

func (p Position) ToBlack(step int) Position {
	return Position{
		Rank: Rank(p.Rank.Index() + step),
		File: p.File,
	}
}

func (p Position) ToWhite(step int) Position {
	return Position{
		Rank: Rank(p.Rank.Index() - step),
		File: p.File,
	}
}

func (p Position) ToA(step int) Position {
	return Position{
		Rank: p.Rank,
		File: File(p.File.Index() - step),
	}
}

func (p Position) ToH(step int) Position {
	return Position{
		Rank: p.Rank,
		File: File(p.File.Index() + step),
	}
}
