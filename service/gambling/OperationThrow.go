package gambling

import (
	"math/rand"
	"time"
)

// 摔一下当前的骰子
func (d *Dice) RandNext() {
	rand.Seed(time.Now().UTC().UnixNano())
	d.ActualValue = 1 + rand.Intn(d.MaxNum)
}

// 摔一下当前的骰子
func (d *Play) ThrowNumOfDice() ([]Dice) {

	num := make([]int, d.DiceNum)

	var dices []Dice = make([]Dice, d.DiceNum)
	for i := range num {
		dices[i] = Dice{6, 0}
		dices[i].RandNext()
	}

	return dices
}
