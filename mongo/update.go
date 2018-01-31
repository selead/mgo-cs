package mongo

import (
	"github.com/spiderorg/mgo-cs/pool"
)

// 更新第一个匹配的数据
type Update struct {
	Database   string      // 数据库
	Collection string      // 集合
	Type       string      // Type
	Selector   interface{} // 文档选择器
	Change     interface{} // 文档更新内容
}

func (u *Update) Exec(_ interface{}) error {
	return Call(func(src pool.Src) error {
		c := src.(*MgoSrc).DB(u.Database).C(u.Collection)
		if u.Type == "Upsert" {
			_, err := c.Upsert(u.Selector, u.Change)
			return err
		}
		return c.Update(u.Selector, u.Change)
	})
}
