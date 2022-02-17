package models

import (
	"context"
	"demo1/pkg/setting"
	"demo1/pkg/types"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client

func Setup() {
	var err error
	var mysqlConfig = setting.Conf.MysqlConf
	var redisConfig = setting.Conf.RedisConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.DB)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("models.Setup err: %v\n", err)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	testRedis()
}

func Close() {

}

func testRedis() {
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}

type Member struct {
	//gorm.Model
	UserId   int64          `json:"user_id" gorm:"primary_key"`
	Nickname string         `json:"nickname"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	UserType types.UserType `json:"user_type"`
	Deleted  int            `json:"deleted"`
}

type Course struct {
	CourseId  int64  `json:"course_id" gorm:"primary_key"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	TeacherId string `json:"teacher_id"`
	Deleted   int    `json:"deleted"`
}

type CourseBook struct {
	StudentId int64 `json:"student_id"`
	CourseId  int64 `json:"course_id"`
	Deleted   int   `json:"deleted"`
}
