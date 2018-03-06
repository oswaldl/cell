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
