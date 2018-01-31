package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/spiderorg/mgo-cs/pool"
)

// 在指定集合进行条件查询
type Find struct {
	Database   string      // 数据库
	Collection string      // 集合
	Query      bson.M      // 查询语句
	Sort       []string    // 排序，用法如Sort("firstname", "-lastname")，优先按firstname正序排列，其次按lastname倒序排列
	Skip       int         // 跳过前n个文档
	Limit      int         // 返回最多n个文档
	Select     interface{} // 只查询、返回指定字段，如{"name":1}
}

func (f *Find) Exec(resultPtr interface{}) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = fmt.Errorf("%v", re)
		}
	}()
	resultPtr2 := resultPtr.(bson.M)

	err = Call(func(src pool.Src) error {
		c := src.(*MgoSrc).DB(f.Database).C(f.Collection)

		if id, ok := f.Query["_id"]; ok {
			if idStr, ok2 := id.(string); ok2 {
				f.Query["_id"] = bson.ObjectIdHex(idStr)
			}
		}

		q := c.Find(f.Query)

		resultPtr2["Total"], _ = q.Count()

		if len(f.Sort) > 0 {
			q.Sort(f.Sort...)
		}

		if f.Skip > 0 {
			q.Skip(f.Skip)
		}

		if f.Limit > 0 {
			q.Limit(f.Limit)
		}

		if f.Select != nil {
			q.Select(f.Select)
		}
		r := []interface{}{}
		err = q.All(&r)

		resultPtr2["Docs"] = r

		return err
	})
	return
}
