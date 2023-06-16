package errmsg

const (
	SUCCSE                 = 200
	ERROR                  = 500
	INVALIDDATAFORMAT      = 999 // 非法数据格式
	VALIDATEERROR          = 998 // 数据字段验证失败
	FAILED_TO_QUERY_DATA   = 997 // 查询数据失败
	ERROE_ADD_DATA         = 996 // 添加数据失败
	ERROE_UPDATE_DATA      = 995 // 修改数据失败
	CONFIGLOADERR          = 994 // 配置加载失败
	FILE_TYPE_IS_INVALID   = 993 // 文件类型不合法
	MKDIR_FAIL             = 992 // MkdirAll失败
	FILE_GET_FAIL          = 991 // 文件获取失败
	ERROE_DELETE_DATA      = 990 // 删除数据失败
	MARSHAL_JSON_FAIL      = 989 // json序列化失败
	UNMARSHAL_JSON_FAIL    = 988 //  json反序列化失败
	ERROR_DATA_DUPLICATION = 987 //  数据重复

	// code= 1000... 管理员模块的错误
	ERROR_USERNAME_USED       = 1001 // 用户名已存在
	ERROR_PASSWORD_WRONG      = 1002 // 密码错误
	ERROR_USER_NOT_EXIST      = 1003 // 用户不存在
	ERROR_TOKEN_EXIST         = 1004 // TOKEN不存在,请重新登陆
	ERROR_TOKEN_RUNTIME       = 1005 // TOKEN已过期,请重新登陆
	ERROR_TOKEN_WRONG         = 1006 // TOKEN不正确,请重新登陆
	ERROR_TOKEN_TYPE_WRONG    = 1007 // TOKEN格式错误,请重新登陆
	ERROR_USER_NO_RIGHT       = 1008 // 该用户无权限
	ERROR_PASSWORD_ENCRYPTION = 1009 // 密码加密失败
	ERROR_TOKEN_GENERATE      = 1010 // token生成失败
	ERROR_USER_DISABLE        = 1011 // 管理员账户被禁用

	// code= 2000... 文章模块的错误

	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	INVALIDDATAFORMAT:      "非法数据格式",
	VALIDATEERROR:          "数据字段验证失败",
	FAILED_TO_QUERY_DATA:   "查询数据失败",
	ERROE_ADD_DATA:         "添加数据失败",
	ERROE_UPDATE_DATA:      "修改数据失败",
	CONFIGLOADERR:          "配置加载失败",
	FILE_TYPE_IS_INVALID:   "文件类型不合法",
	MKDIR_FAIL:             "MkdirAll失败",
	FILE_GET_FAIL:          "文件获取失败",
	MARSHAL_JSON_FAIL:      "json序列化失败",
	UNMARSHAL_JSON_FAIL:    "json反序列反失败",
	ERROR_DATA_DUPLICATION: "数据重复",

	ERROR_USERNAME_USED:       "用户名已存在！",
	ERROR_PASSWORD_WRONG:      "密码错误",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_TOKEN_EXIST:         "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:       "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:         "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG:    "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:       "该用户无权限",
	ERROR_PASSWORD_ENCRYPTION: "密码加密失败",
	ERROR_TOKEN_GENERATE:      "token生成失败",
	ERROR_USER_DISABLE:        "管理员账户被禁用",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
