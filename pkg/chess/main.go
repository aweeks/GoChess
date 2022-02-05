package chess

import (
	"fmt"
)

func Main() {
	var state = StartingGameState
	fmt.Println(state)
	for _, s := range state.Board.PlayerSquares(state.Turn) {
		fmt.Println(state.Board.PrettyBoard(ToSquares(state.MovesFrom(s))))
	}
}
