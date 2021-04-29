package main

type Tier int
type Slot int

const (
	Normal Tier = 0
	Magic  Tier = 1
	Rare   Tier = 2
	Unique Tier = 3
)

const (
	Head      Slot = 0
	Chest     Slot = 1
	LeftHand  Slot = 2
	RightHand Slot = 3
)

type Item struct {
	tier    Tier
	slot    Slot
	defense int
	attack  int
	damage  int
	//mods map[Mod]int
}

func GenerateItem() Item {
	return Item{tier: Normal, slot: Head, defense: 1}
}
