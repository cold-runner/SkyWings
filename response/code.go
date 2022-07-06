package response

type RespCode uint32

const (
	CodeSuccess            RespCode = 2022
	CodeNotExist           RespCode = 1011
	CodeSignupValidateFail RespCode = 9999
	CodeStudentExist       RespCode = 9900
	CodeDaoError           RespCode = 7777
	CodeServerBusy         RespCode = 1234
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:            "成功",
	CodeStudentExist:       "用户已存在",
	CodeDaoError:           "数据库错误",
	CodeSignupValidateFail: "参数校验失败",
	CodeNotExist:           "网页不存在",
}

func (rc RespCode) msg() string {
	msg, ok := codeMsgMap[rc]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
