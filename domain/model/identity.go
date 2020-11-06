package model

type UserID int64

//TODO: 既存のint64へのcastはここに統一する
//TODO:　テストも書いておく
func (id UserID) ConvertUserIdToInt64() int64 {
	return int64(id)
}
