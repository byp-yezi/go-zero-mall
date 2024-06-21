package errx

var message map[uint32]string

func init() {
   message = make(map[uint32]string)

   // 全局模块
   message[OK] = "SUCCESS"
   message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
   message[REUQEST_PARAM_ERROR] = "参数错误"
   message[TOKEN_EXPIRE_ERROR] = "token失效, 请重新登陆"
   message[TOKEN_GENERATE_ERROR] = "生成token失败"
   message[DB_ERROR] = "数据库繁忙,请稍后再试"

   // 用户模块
   message[USER_EXIST_ERROR] = "该用户已注册-rpc"
   message[USER_NOTEXIST_ERROR] = "该用户不存在-rpc"
   message[PASSWORD_ERROR] = "密码错误-rpc"

   //产品模块
   message[CREATE_PRODUCT_ERROR] = "产品创建失败"
   message[GET_PRODUCT_BY_ID_ERROR] = "通过id获取产品失败"
}

func MapErrMsg(errcode uint32) string {
   if msg, ok := message[errcode]; ok {
      return msg
   } else {
      return "服务器开小差啦,稍后再来试一试"
   }
}

func IsCodeErr(errcode uint32) bool {
   if _, ok := message[errcode]; ok /* || errcode >= 1000 */{
      return true
   } else {
      return false
   }
   // return errcode >= 1000 && errcode <= 2000
}