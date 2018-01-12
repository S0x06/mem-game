package game

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/wangxianzhuo/mem-game/server/pic"
)

// GameMap 游戏版面，保存所以实体，版面序列和对应图片资源
type GameMap struct {
	Entries []Entry          `json:"entries"`
	Map     map[int]*Entry   `json:"map"`
	Pics    map[int]*pic.Pic `json:"pics"`
}

func (m *GameMap) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for k, v := range m.Map {
		buffer.WriteString(fmt.Sprintf("[%v:%v]", k, v))
	}
	return fmt.Sprintf("entries: %v\nmap: %v\npics: %v", m.Entries, buffer.String(), m.Pics)
}

// ErrGameMapSize 版面上游戏实体的数量不是2的倍数返回这个异常
var ErrGameMapSize = errors.New("game map size is illegal")

// NewGameMap new a GameMap
func NewGameMap(len, wid int, picPath string) (*GameMap, error) {
	if (len*wid)%2 != 0 {
		return nil, ErrGameMapSize
	}
	var err error
	gameMap := GameMap{
		Entries: make([]Entry, len*wid),
	}

	uniqueEntryNum := len * wid / 2

	for index := 0; index < (len * wid); index++ {
		gameMap.Entries[index] = NewEntry(index, index%uniqueEntryNum, index%uniqueEntryNum)
	}

	gameMap.Pics, err = loadPics(picPath, uniqueEntryNum)
	if err != nil {
		return nil, fmt.Errorf("Create a new game map error: %v", err)
	}

	gameMap.Map, err = generateMap(len*wid, gameMap.Entries)
	if err != nil {
		return nil, fmt.Errorf("Create a new game map error: %v", err)
	}
	return &gameMap, nil
}

func loadPics(picPath string, picNum int) (map[int]*pic.Pic, error) {
	pics, err := pic.LoadPics(picPath, picNum)
	if err != nil {
		return nil, fmt.Errorf("load pics error: %v", err)
	}
	result := make(map[int]*pic.Pic)
	for index := 0; index < picNum; index++ {
		result[index] = &pics[index]
	}
	return result, nil
}

func generateMap(entryNum int, entries []Entry) (map[int]*Entry, error) {
	result := make(map[int]*Entry)

	eList := arraylist.New()
	for _, entry := range entries {
		eList.Add(entry)
	}
	for index := 0; index < entryNum; index++ {
		eIndex := getRandomIndex(eList.Size())
		v, _ := eList.Get(eIndex)
		e := v.(Entry)
		result[index] = &e
		eList.Remove(eIndex)
	}
	return result, nil
}

func getRandomIndex(rank int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(rank)
}
