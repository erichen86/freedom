//Package po generated by 'freedom new-po'
package po

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Cart .
type Cart struct {
	changes map[string]interface{}
	ID      int       `gorm:"primary_key;column:id"`
	UserID  int       `gorm:"column:user_id"`  // 用户ID
	GoodsID int       `gorm:"column:goods_id"` // 商品id
	Num     int       `gorm:"column:num"`      // 数量
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}

// TableName .
func (obj *Cart) TableName() string {
	return "cart"
}

// TakeChanges .
func (obj *Cart) TakeChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// updateChanges .
func (obj *Cart) setChanges(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}

// SetUserID .
func (obj *Cart) SetUserID(userID int) {
	obj.UserID = userID
	obj.setChanges("user_id", userID)
}

// SetGoodsID .
func (obj *Cart) SetGoodsID(goodsID int) {
	obj.GoodsID = goodsID
	obj.setChanges("goods_id", goodsID)
}

// SetNum .
func (obj *Cart) SetNum(num int) {
	obj.Num = num
	obj.setChanges("num", num)
}

// SetCreated .
func (obj *Cart) SetCreated(created time.Time) {
	obj.Created = created
	obj.setChanges("created", created)
}

// SetUpdated .
func (obj *Cart) SetUpdated(updated time.Time) {
	obj.Updated = updated
	obj.setChanges("updated", updated)
}

// AddUserID .
func (obj *Cart) AddUserID(userID int) {
	obj.UserID += userID
	obj.setChanges("user_id", gorm.Expr("user_id + ?", userID))
}

// AddGoodsID .
func (obj *Cart) AddGoodsID(goodsID int) {
	obj.GoodsID += goodsID
	obj.setChanges("goods_id", gorm.Expr("goods_id + ?", goodsID))
}

// AddNum .
func (obj *Cart) AddNum(num int) {
	obj.Num += num
	obj.setChanges("num", gorm.Expr("num + ?", num))
}
