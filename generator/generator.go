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
	Helm   Kind = 0
	Armor  Kind = 1
	Sword  Kind = 2
	Axe    Kind = 3
	Mace   Kind = 4
	Shield Kind = 5
)

type Item struct {
	name        string
	Tier        Tier
	kind        Kind
	defense     int
	blockChance int
	attack      int
	damage      int
	level       int
	mods        map[Mod]int
}

func roll(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	return rnd.Intn(max-min) + min
}

func rollTier() Tier {
	var tier Tier
	d100 := roll(1, 100)
	if d100 <= 1 {
		tier = Unique
	} else if d100 <= 5 {
		tier = Rare
	} else if d100 <= 25 {
		tier = Magic
	} else {
		tier = Normal
	}
	return tier
}

func rollKind() Kind {
	n := roll(0, 5)
	return Kind(n)
}

func GenerateItem(monsterLevel int) Item {

	tier := rollTier()
	kind := rollKind()
	level := monsterLevel

	return Item{Tier: tier, kind: kind, defense: roll(5, 10), level: level}
}
