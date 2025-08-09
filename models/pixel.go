package models

type Pixel struct {
	Index int `json:"index" bson:"index"` // 1..10000
	Color int `json:"color" bson:"color"` // 1..10
}
