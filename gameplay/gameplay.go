package gameplay

import (
	"fmt"
	"time"

	"github.com/arispen/item-server/generator"
)

type player struct {
	life         int
	attackRating int
	defense      int
	damage       int
	exp          int
	blockChance  int
	inventory    []generator.Item
	gold         int
}

func (player *player) attack(monster *monster) (kill bool) {
	monsterDefense := 1 + monster.level/2
	// TODO
	monsterBlock := 5
	if isHit(player.attackRating, monsterDefense, monsterBlock) {
		fmt.Println("player hits")
		monster.life -= player.damage
		if monster.life <= 0 {
			return true
		}
	}
	return false
}

type monster struct {
	level int
	life  int
}

func (monster *monster) attack(player *player) (death bool) {
	monsterAttack := 1 + monster.level/2
	monsterDamage := 1 + monster.level/2
	if isHit(monsterAttack, player.defense, player.blockChance) {
		fmt.Println("monster hits")
		player.life -= monsterDamage
		if player.life <= 0 {
			return true
		}
		fmt.Print("life left")
		fmt.Print(player.life)
	}
	return false
}

func isHit(attackerAttack int, defenderDefense int, defenderBlock int) bool {
	// Chance To Hit = 200% * {AR / (AR + DR)} * {Alvl / (Alvl + Dlvl)}
	// TODO: levels?
	chanceToHit := 200 * (attackerAttack / (attackerAttack + defenderDefense))
	fmt.Println("chance to hit:", chanceToHit)
	// min 5% max 95%
	if chanceToHit < 5 {
		chanceToHit = 5
	}
	if chanceToHit > 95 {
		chanceToHit = 95
	}
	roll := generator.Roll(1, 100)
	return roll <= chanceToHit
}

func tick() {

}

func StartMainLoop() {
	// player attacks a monster
	//  - if hit then deal damage
	//	 - if killed then get exp, roll treasure, monster level ++
	// monster attacks the player character
	// 	- if hit then deal damage
	//	 - if player killed, go to monster level 1

	var player *player = &player{life: 20, attackRating: 1, defense: 1, damage: 1, blockChance: 5, inventory: make([]generator.Item, 0)}
	var monster *monster = &monster{level: 1, life: 1}

	ticker := time.NewTicker(1000 * time.Millisecond)
	tocks := 0
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				//fmt.Println("Tick at", t)
				tocks++
				if tocks%4 == 0 {
					fmt.Println("player.attack(monster)")
					if monsterKilled := player.attack(monster); monsterKilled {
						fmt.Println("monster killed!")
						player.exp += monster.level
						player.inventory = append(player.inventory, generator.GenerateItem(monster.level))
						fmt.Println(player.inventory)
						monster.level++
					}

				}
				if tocks%5 == 0 {
					fmt.Println("monster.attack(player)")
					if death := monster.attack(player); death {
						monster.level = 0
						fmt.Println("DEATH")
						ticker.Stop()
					}
					fmt.Println("")
				}
			}
		}
	}()

	time.Sleep(100 * 1000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
