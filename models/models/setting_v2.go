package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SettingV2 struct {
	any                    `collection:"settings"`
	BaseModelV2[SettingV2] `bson:",inline"`
	Key                    string `json:"key" bson:"key"`
	Value                  bson.M `json:"value" bson:"value"`
}
