package mysql

import (
	"SkyWings/models"
	"go.uber.org/zap"
)

func Insert(stu *models.Student, genId uint64, stuAge uint8, photoPath string) (err error) {
	// 将报名者信息插入数据库
	_, err = db.Exec("insert into applicant_information(id, student_name, student_id, student_gender, student_age, student_province, student_major, student_introduce, student_phone, student_qq, student_photo) VALUES (?,?,?,?,?,?,?,?,?,?,?)", genId, stu.Name, stu.Id, stu.Gender, stuAge, stu.Province, stu.Major, stu.Intro, stu.Phone, stu.QQ, photoPath)
	if err != nil {
		zap.L().Error("插入数据失败！err:", zap.Error(err))
		return ErrorDao
	}
	return nil
}

func IsExit(std *models.Student) error {
	// 根据手机号码、学号、QQ查询数据库中是否已存在报名者信息

	// 学号查重
	var stuId string
	_ = db.Get(&stuId, "select student_id from applicant_information where student_id=?", std.Id)
	if std.Id == stuId {
		return ErrorUserExit
	}

	// 手机号查重
	var stuPhoneEx string
	_ = db.Get(&stuPhoneEx, "select student_phone from applicant_information where student_phone=?", std.Phone)
	if stuPhoneEx == std.Phone {
		return ErrorUserExit
	}

	// QQ号查重
	var stuQQEx string
	_ = db.Get(&stuQQEx, "select student_qq from applicant_information where student_qq=?", std.QQ)
	if std.QQ == stuQQEx {
		return ErrorUserExit
	}

	return nil
}
