package models

type MatchType int

const (
	One_x_One MatchType = iota
	Two_x_Two
	Four_x_Four
	Only_One_Win
)
