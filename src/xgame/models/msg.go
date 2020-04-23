package models

const (
	POS_LOBBY = iota
	POS_LANLORD
)

//
type Msg struct {
	pos int
	msg struct{ Name int }
}
