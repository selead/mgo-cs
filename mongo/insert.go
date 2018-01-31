package mongo

import (
	"fmt"

	"github.com/spiderorg/mgo-cs/pool"
)

// 插入新数据
type Insert struct {
	Database   string        // 数据库
	Collection string        // 集合
	Docs       []interface{} // 文档
}

const (
	MaxLen int = 5000 //分配插入
)

func (self *Insert) Exec(resultPtr interface{}) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = fmt.Errorf("%v", re)
		}
	}()
	count := len(self.Docs)

	return Call(func(src pool.Src) error {
		c := src.(*MgoSrc).DB(self.Database).C(self.Collection)
		loop := count / MaxLen
		for i := 0; i < loop; i++ {
			err := c.Insert(self.Docs[i*MaxLen : (i+1)*MaxLen]...)
			if err != nil {
				return err
			}
		}
		if count%MaxLen == 0 {
			return nil
		}
		return c.Insert(self.Docs[loop*MaxLen:]...)
	})
}
