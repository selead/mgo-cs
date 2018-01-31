package mongo

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/spiderorg/logrus"
	mgo "gopkg.in/mgo.v2"

	"github.com/spiderorg/mgo-cs/pool"
)

type MgoSrc struct {
	*mgo.Session
}

var (
	connGcSecond = time.Duration(viper.GetInt("database.mongodb.gcSeconds")) * 1e9
	session      *mgo.Session
	err          error
	MgoPool      = pool.ClassicPool(
		viper.GetInt("database.mongodb.connCAP"),
		viper.GetInt("database.mongodb.connCAP")/5,
		func() (pool.Src, error) {
			// if err != nil || session.Ping() != nil {
			// 	session, err = newSession()
			// }
			return &MgoSrc{session.Clone()}, err
		},
		connGcSecond)
)

func Refresh() {
	url := fmt.Sprintf("mongodb://%s:%s@%s", viper.GetString("database.mongodb.user"),
		viper.GetString("database.mongodb.password"),
		viper.GetString("database.mongodb.connect"),
	)

	session, err = mgo.Dial(url)
	if err != nil {
		logrus.Fatalln("MongoDB", err, "|", viper.GetString("database.mongodb.connect"))
	} else if err = session.Ping(); err != nil {
		logrus.Fatalln("MongoDB", err, "|", viper.GetString("database.mongodb.connect"))
	} else {
		session.SetPoolLimit(viper.GetInt("database.mongodb.connCAP"))
	}
	logrus.Infoln("To open mongo is ok.")
}

// 判断资源是否可用
func (self *MgoSrc) Usable() bool {
	if self.Session == nil || self.Session.Ping() != nil {
		return false
	}
	return true
}

// 使用后的重置方法
func (*MgoSrc) Reset() {}

// 被资源池删除前的自毁方法
func (self *MgoSrc) Close() {
	if self.Session == nil {
		return
	}
	self.Session.Close()
}

func Error() error {
	return err
}

// 调用资源池中的资源
func Call(fn func(pool.Src) error) error {
	return MgoPool.Call(fn)
}

// 销毁资源池
func Close() {
	MgoPool.Close()
}

// 返回当前资源数量
func Len() int {
	return MgoPool.Len()
}

// 获取所有数据
func DatabaseNames() (names []string, err error) {
	err = MgoPool.Call(func(src pool.Src) error {
		names, err = src.(*MgoSrc).DatabaseNames()
		return err
	})
	return
}

// 获取数据库集合列表
func CollectionNames(dbname string) (names []string, err error) {
	MgoPool.Call(func(src pool.Src) error {
		names, err = src.(*MgoSrc).DB(dbname).CollectionNames()
		return err
	})
	return
}
