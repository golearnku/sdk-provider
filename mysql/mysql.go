package mysql

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	internalMysql *Mysql
	onceMysql     sync.Once

	// ErrRecordNotFound returns a "record not found error". Occurs only when attempting to query the database with a struct; querying with a slice won't return this error
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// IsRecordNotFoundError returns true if error contains a RecordNotFound error
func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

type Config struct {
	DBName          string //数据库名称
	User            string //用户名
	Password        string //密码
	Adds            string //连接地址
	Debug           bool   //是否开启 mysql debug
	Charset         string //字符集 （默认 utf8）
	ConnMaxLifetime int    //设置连接可重用的最大时间量
	MaxIdleConns    int    //设置连接池空闲时的最大连接数
	MaxOpenConns    int    //设置数据库的最大打开连接数
	SingularTable   bool   //关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
}

type Mysql struct {
	client *gorm.DB //gorm  实例
}

//mysql ping
func (mysql *Mysql) Ping() error {
	if mysql.client != nil {
		return mysql.client.DB().Ping()
	}
	return nil
}

//mysql close
func (mysql *Mysql) Close() error {
	if mysql.client != nil {
		return mysql.client.DB().Close()
	}
	return nil
}

//初始化单例客户端
func NewClient(config Config) (*Mysql, error) {
	var (
		db  *gorm.DB
		err error
	)
	onceMysql.Do(func() {
		db, err = NewInstance(config)
		if err != nil {
			return
		}
		internalMysql = &Mysql{client: db}
	})
	if internalMysql == nil {
		return nil, err
	}
	return internalMysql, nil
}

//初始化实例
func NewInstance(config Config) (db *gorm.DB, err error) {
	charset := "utf8mb4"
	if config.Charset != "" {
		charset = config.Charset
	}

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", config.User, config.Password, config.Adds, config.DBName, charset)
	db, err = gorm.Open("mysql", url)
	if err != nil {
		return
	}

	// 启用Logger，显示详细日志
	db.LogMode(config.Debug)
	if config.ConnMaxLifetime > 0 {
		db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.ConnMaxLifetime))
	}
	if config.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxOpenConns > 0 {
		db.DB().SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.SingularTable {
		db.SingularTable(true)
	}
	return
}

//获取mysql单例对象
func GetConnDB() *gorm.DB {
	if internalMysql == nil || internalMysql.client == nil {
		panic("Mysql Client is not initialized")
	}
	return internalMysql.client
}
