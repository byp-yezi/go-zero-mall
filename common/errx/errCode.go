package errx

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005

//用户模块
const (
	USER_EXIST_ERROR uint32 = iota + 10001
	USER_NOTEXIST_ERROR
	PASSWORD_ERROR
)

//产品模块
const (
	CREATE_PRODUCT_ERROR uint32 = iota + 11001
	GET_PRODUCT_BY_ID_ERROR
)