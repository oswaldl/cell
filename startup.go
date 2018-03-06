package main

import (
	"github.com/go-martini/martini"
	"github.com/oswaldl/cell/service/gambling"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
)

const MAX_DICE_NUM  = 5

func main() {
	m := martini.Classic()

	registerWebService(m)

	m.RunOnAddr(":8887")
}

func registerWebService(m *martini.ClassicMartini) {

	m.Get("/help", func() string {
		return "/gambling/:diceNum"
	})

	m.Get("/gambling/:diceNum", func(params martini.Params) (int, string) {

		if len(params) == 0 {
			// No params. Return entire collection encoded as JSON.

			// Failed encoding collection.
			return http.StatusInternalServerError, "no param: diceNum"
		}

		// Convert id to integer.
		diceNum, err := strconv.Atoi(params["diceNum"])
		if err != nil {
			// Id was not a number.
			return http.StatusBadRequest, "invalid diceNum"
		}

		num := make([]int, diceNum)

		if diceNum > MAX_DICE_NUM {
			return http.StatusBadRequest, fmt.Sprintf("%s%d", "diceNum is over " , MAX_DICE_NUM)
		}

		var dices [MAX_DICE_NUM]gambling.Dice
		for i := range num {
			dices[i] = gambling.Dice{6, 0}
			dices[i].RandNext()
		}

		jsonStr, _ := json.Marshal(dices)
		return http.StatusOK, string(jsonStr)
	})

}
