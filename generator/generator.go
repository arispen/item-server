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
	reqLvl      int
	reqStr      int
	reqDex      int
}

func roll(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	return rnd.Intn(max-min) + min
}

func rollTier() Tier {
	var tier Tier
	dk := roll(1, 1000)
	if dk <= 10 {
		tier = Unique
	} else if dk <= 50 {
		tier = Rare
	} else if dk <= 250 {
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
	var name string
	var kind Kind
	var level, reqLvl, defense, reqStr int
	var mods map[Mod]int

	if tier == Unique {
		// TODO: roll unique item
		name = "Shako"
		kind = Helm
		reqLvl = 62
		reqStr = 50
		defense = roll(98, 141)
		mods = map[Mod]int{
			AllSkills:     2,
			AllAttributes: 2,
			MagicFind:     50,
			DamageReduced: 10,
			Life:          100,
		}

	} else {
		// TODO: roll mods
		kind = rollKind()
		level = monsterLevel
	}

	return Item{name: name, Tier: tier, kind: kind, defense: defense,
		level: level, mods: mods, reqLvl: reqLvl, reqStr: reqStr}
}
