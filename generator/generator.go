package generator

import (
	"math/rand"
	"time"
)

type Tier int

const (
	Normal Tier = 0
	Magic  Tier = 1
	Rare   Tier = 2
	Unique Tier = 3
)

type Kind int

const (
	Helm  Kind = 0
	Chest Kind = 1
	Sword Kind = 2
	Axe   Kind = 3
	Mace  Kind = 4
)

type Item struct {
	tier        Tier
	kind        Kind
	defense     int
	blockChance int
	attack      int
	damage      int
	mods        map[Mod]int
}

func roll(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	return rnd.Intn(max-min) + min
}

func GenerateItem() Item {
	return Item{tier: Normal, kind: Helm, defense: roll(5, 10)}
}
