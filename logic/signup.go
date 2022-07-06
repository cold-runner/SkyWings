package logic

import (
	"SkyWings/dao/mysql"
	"SkyWings/models"
	"SkyWings/pkg/snowflake"
	"SkyWings/settings"
	"encoding/base64"
	"io/ioutil"
	"time"
)

func SignUp(student *models.Student) error {

	// 根据手机号码、学号、QQ查询数据库中是否已存在报名者信息
	if err := mysql.IsExit(student); err != nil {
		return err
	}

	// 图片处理
	bytes, _ := base64.StdEncoding.DecodeString(student.Photo)
	photoStorePath := settings.Conf.PhotoPath + student.Id + student.Name + ".jpg"
	_ = ioutil.WriteFile(photoStorePath, bytes, 0666)
	// = ioutil.WriteFile(path+"/dao/photos/"+student.Name+".jpg", bytes, 0666)

	// 将student结构体中各个字段的格式变成数据库所需要的格式
	genId, _ := snowflake.GetID() // 得到自然主键
	stuBir, _ := time.Parse("2006-01-02", student.Birth)
	stuAge := time.Now().Year() - stuBir.Year() // 通过出生年得到年龄
	// 插入数据到数据库
	if err := mysql.Insert(student, genId, uint8(stuAge), photoStorePath); err != nil {
		return err
	}

	return nil
}
