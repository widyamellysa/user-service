package model

import (

	//"entities"

	"digileaps/user/entities"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	Db         *mgo.Database
	Collection string
}

func (userModel UserModel) FindAll() (user []entities.User, err error) {
	err = userModel.Db.C(userModel.Collection).Find(bson.M{}).All(&user)
	return
}

func (userModel UserModel) Find(id string) (user entities.User, err error) {
	err = userModel.Db.C(userModel.Collection).FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

func (userModel UserModel) Create(id string) (user *entities.User, err error) {
	err = userModel.Db.C(userModel.Collection).Insert(&user)
	return
}

func (userModel UserModel) Update(id string) (user *entities.User, err error) {
	err = userModel.Db.C(userModel.Collection).UpdateId(user.ID, &user)
	return
}

func (userModel UserModel) Delete(user entities.User) error {
	err := userModel.Db.C(userModel.Collection).RemoveId(user)
	return err
}
