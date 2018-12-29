package model

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	ID           uint
	Name         string
	CategoryId   uint
	ItemProperty []ItemProperty `gorm:"ForeignKey:ItemId"`
}

type ItemProperty struct {
	ID              uint
	ItemId          uint
	PropertyId      uint
	PropertyValueId uint
}

type Property struct {
	ID            uint
	Name          string
	PropertyValue []PropertyValue `gorm:"ForeignKey:PropertyId"`
}

type PropertyValue struct {
	ID         uint
	PropertyId uint
	Name       string
}

// GetItemsByCategoryId 获取分类商品列表
func GetItemsByCategoryId(id uint) ([]*Item, error) {
	var items []*Item
	err := DB.Find(&items, Item{CategoryId: uint(id)}).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemById(id uint) (*Item, error) {
	var item Item
	err := DB.First(&item, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var propertys []ItemProperty
	// err = DB.Model(&item).Related(&propertys).Error
	// if err != nil {
	// 	return nil, err
	// }
	rows, err := DB.Model(&item).Joins("left join item_property on item_property.item_id = item.id").Rows()
	for rows.Next() {
		var 
	}
	// item.ItemProperty = propertys
	// return &item, nil

	// rows, err := DB.Model(&Item{})
}
