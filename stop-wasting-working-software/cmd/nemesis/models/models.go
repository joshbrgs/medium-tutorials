package models

import "github.com/joshbrgs/mongorm/cmd/mongorm"

type Nemesis struct {
	mongorm.Model `bson:",inline"`
	Name          string `json:"name" bson:"name"`
	Power         string `json:"power" bson:"power"`
}
