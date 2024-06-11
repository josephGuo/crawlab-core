package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserV2 struct {
	any                 `collection:"users"`
	BaseModelV2[UserV2] `bson:",inline"`
	Username            string `json:"username" bson:"username"`
	Password            string `json:"-,omitempty" bson:"password"`
	Role                string `json:"role" bson:"role"`
	Email               string `json:"email" bson:"email"`
}
