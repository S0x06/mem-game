package game

// Entry 游戏选择实体，通俗讲就是所选的牌
type Entry struct {
	ID        int `json:"id"`
	CompareID int `json:"compareId"`
	PicID     int `json:"picId"`
}

// NewEntry new an Entry
func NewEntry(id, compareID, picID int) Entry {
	return Entry{
		ID:        id,
		CompareID: compareID,
		PicID:     picID,
	}
}
