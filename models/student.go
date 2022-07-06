package models

type Student struct {
	Name     string `json:"name" binding:"required,is-chineseName" db:"student_name"`         // 学生姓名
	Id       string `json:"id" binding:"required,number,len=8,valid-id"`                      // 8位固定学号
	Gender   string `json:"gender" binding:"required,valid-gender" db:"student_gender"`       // 性别
	Birth    string `json:"birth" binding:"required,datetime=2006-01-02"`                     // 时间
	Province string `json:"province" binding:"required,valid-province" db:"student_province"` // 省份
	Major    string `json:"major" binding:"required,valid-major" db:"student_major"`          // 专业
	Intro    string `json:"intro" binding:"required,min=100" db:"student_introduce"`          // 防止注入攻击
	Phone    string `json:"phone" binding:"required,is-phone" db:"student_phone"`             // 手机号码
	QQ       string `json:"qq" binding:"required,is-qq" db:"student_qq"`                      // QQ号
	Photo    string `json:"photo" binding:"required,base64"`                                  // 照片
}
