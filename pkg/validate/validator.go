package validate

// 自定义验证器
import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const (
	IdPattern          = "^22(999|006|180|008)([0-9][0-9][0-9])$"                                                                    // 前两位必须为22，第三位到第六位必须为999、006、008、180其中之一，后三位必须为纯数字
	ChineseNamePattern = "^[\\x{4e00}-\\x{9fa5}]{2,10}$"                                                                             // 中文姓名
	GenderPattern      = "^女|男$"                                                                                                     // 性别
	MajorPattern       = "^计算机科学与技术|计科师范|软件工程|网络工程$"                                                                                 // 专业信息必须为其中之一
	QQPattern          = "^[1-9][0-9]{4,10}$"                                                                                        // QQ号匹配
	PhonePattern       = "^(?:(?:\\+|00)86)?1\\d{10}$"                                                                               // 手机号码匹配
	ProvincePattern    = "^浙江|上海|北京|天津|重庆|黑龙江|吉林|辽宁|内蒙古|河北|新疆|甘肃|青海|陕西|宁夏|河南|山东|山西|安徽|湖北|湖南|江苏|四川|贵州|云南|广西|西藏|江西|广东|福建|台湾|海南|香港|澳门$" // 省份匹配
)

var Validate *validator.Validate
var rex *regexp.Regexp

func InitValidator() (err error) {

	Validate = binding.Validator.Engine().(*validator.Validate)

	if err = Validate.RegisterValidation("valid-id", checkId); err != nil {
		fmt.Printf("检测中文姓名的验证器初始化错误！错误值：%v\n", err)
	}
	if err = Validate.RegisterValidation("is-chineseName", checkName); err != nil {
		fmt.Printf("检测中文姓名的验证器初始化错误！错误值：%v\n", err)
	}
	if err = Validate.RegisterValidation("valid-gender", checkGender); err != nil {
		fmt.Printf("检测性别的验证器初始化错误！错误值：%v\n", err)
	}
	if err = Validate.RegisterValidation("valid-major", checkMajor); err != nil {
		fmt.Printf("检测专业信息的验证器初始化错误！错误值：%v\n", err)
	}
	if err = Validate.RegisterValidation("is-qq", checkQQ); err != nil {
		fmt.Printf("检测QQ号的验证器初始化错误！错误值：%v\n", err)
	}
	if err = Validate.RegisterValidation("is-phone", checkPhone); err != nil {
		fmt.Printf("检测QQ号的验证器初始化错误！错误值：%v\n", err)
	}
	if err = Validate.RegisterValidation("valid-province", checkProvince); err != nil {
		fmt.Printf("检测省份信息的验证器初始化错误！错误值：%v\n", err)
	}
	return
}
func checkId(fl validator.FieldLevel) bool {
	rex = regexp.MustCompile(IdPattern)
	result := rex.MatchString(fl.Field().String())
	return result
}

func checkName(fl validator.FieldLevel) bool {
	rex = regexp.MustCompile(ChineseNamePattern)
	result := rex.MatchString(fl.Field().String())
	return result
}

func checkGender(fl validator.FieldLevel) bool {
	rex = regexp.MustCompile(GenderPattern)
	result := rex.MatchString(fl.Field().String())
	return result
}

func checkMajor(fl validator.FieldLevel) bool {
	rex = regexp.MustCompile(MajorPattern)
	result := rex.MatchString(fl.Field().String())
	return result
}

func checkQQ(fl validator.FieldLevel) bool {
	rex = regexp.MustCompile(QQPattern)
	result := rex.MatchString(fl.Field().String())
	return result
}

func checkPhone(fl validator.FieldLevel) bool {
	rex = regexp.MustCompile(PhonePattern)
	result := rex.MatchString(fl.Field().String())
	return result
}

func checkProvince(fl validator.FieldLevel) bool {
	rex = regexp.MustCompile(ProvincePattern)
	result := rex.MatchString(fl.Field().String())
	return result
}
