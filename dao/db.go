package dao

import (
	"WinterDemo/configs"
	"WinterDemo/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 这里我先通过尝试连接数据库来判断数据库是否存在
func InitDB() error {
	conf := configs.GlobalConfig.Database //简化代码

	// fmt.Printf("正在使用以下配置连接数据库：\n用户名：%s\n主机：%s\n端口：%d\n数据库：%s\n",
	// conf.Username, conf.Host, conf.Port, conf.DBName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		conf.Username, conf.Password, conf.Host, conf.Port,
	) //这个dsn是/,没有连接到指定数据库

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("数据库连接失败: %s", err)
	} //连接mysql但是没有连接到指定数据库

	//row在mysql中检查数据库是否存在,不存在会返回sql.ErrNoRows存储在row中，必须读取才能知道
	//row是一个*sql.Row类型的指针，所以需要Scan()把结果传给dbName来获取数据判断是否存在
	row := DB.Raw("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", conf.DBName).Row()
	var dbName string
	err = row.Scan(&dbName)
	if err != nil {
		err := DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", conf.DBName)).Error
		if err != nil {
			return fmt.Errorf("创建数据库失败: %s", err)
		}
		fmt.Printf("数据库%s创建成功\n", conf.DBName)
	}

	fulldsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		conf.Username, conf.Password, conf.Host, conf.Port,
		conf.DBName, conf.Charset, conf.ParseTime, conf.Loc,
	) //这个dsn是连接到指定数据库

	DB, err = gorm.Open(mysql.Open(fulldsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("数据库连接失败: %s", err)
	} //连接到指定数据库

	fmt.Println("数据库连接成功")

	//接下来依据models中的结构体创建表
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("创建表失败: %s", err)
	}

	return nil
}

func CloseDB() error {
	if DB != nil { //检测DB是否为空
		db, err := DB.DB() //DB是*gorm.DB类型，而我需要的是DB.DB()返回的*sql.DB类型
		if err != nil {
			return fmt.Errorf("获取数据库实例(即*sql.DB)失败: %s", err)
		}
		db.Close()

		fmt.Println("数据库连接已关闭")
	}
	return nil
}
