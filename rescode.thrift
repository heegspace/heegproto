namespace go heegproto.rescode

enum code {
    SUCCESS =   0,      // 成功
    ERROR   =   1,      // 失败
    DB_ERROR = 2,       // 数据异常
    AUTH_ERR = 3,       // 没有权限
    MOBILE_ERR = 4,     // 电话号码错误
    EMAIL_ERR = 5,      // 邮箱错误

    PARAM_ERR = 99,     // 参数错误
    EXISTS  =   100,    // 已经存在
    IS_SELF =   101,    // 自己


    NOT_EXISTS = 400,    // 不存在

    SEND_CODE_ERR = 10000,  // 发送验证码错误
    CODE_ERROR  = 10001,    // 验证码错误
    CODE_EXPIRE = 10002,    // 验证码过期
    CODE_RATE   = 10003,    // 请求频率太高
    CODE_LIMIT  = 10004,    // 验证码码达到限制
    CODE_TYPE_ERR = 10005,  // 不支持的类型
}