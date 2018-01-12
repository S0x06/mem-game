package game

import (
	"fmt"
	"testing"
)

func Test_NewMap(t *testing.T) {
	gameMap, err := NewGameMap(4, 4, "")
	if err != nil {
		t.Error(err)
	}
	_, err = NewGameMap(5, 3, "")
	if err != nil && err == GameMapSizeErr {
		t.Logf("game map size error")
	} else if err != nil {
		t.Error(err)
	}
	fmt.Println(gameMap)
}
