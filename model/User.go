package model

import "ums/utils/result"

type User struct {
	ID int `json:"id"`
	NickName string `gorm:"type:varchar(100);not null" json:"nickName"`
	Email string `gorm:"type:varchar(100);not null" json:"email"`
	Phone string `gorm:"type:varchar(20);not null" json:"phone"`
	Address string `gorm:"type:varchar(20);not null" json:"address"`
}


func CreateUser(user *User) int {
	err := DB.Create(user).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func ModifyUser(user *User) int {
	err := DB.Save(user).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func RemoveUser(ids []int) int {
	var users []User
	err := DB.Delete(&users, ids).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func PageUser(pageNum, pageSize int, name string) ([]User, int64) {
	var userList []User
	var total int64
	if name == "" {
		err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList).Count(&total).Error
	}else {
		err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Where("nick_name like ?", "%"+name+"%").Find(&userList).Count(&total).Error
	}
	if err != nil {
		return nil, total
	}
	return userList, total
}