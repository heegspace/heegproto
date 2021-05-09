namespace go rescode

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
    
    NOT_DATA = 404,      // 没有数据

    SEND_CODE_ERR = 10000,  // 发送验证码错误
    CODE_ERROR  = 10001,    // 验证码错误
    CODE_EXPIRE = 10002,    // 验证码过期
    CODE_RATE   = 10003,    // 请求频率太高
    CODE_LIMIT  = 10004,    // 验证码码达到限制
    CODE_TYPE_ERR = 10005,  // 不支持的类型
    JSON_MAR_ERR = 10006,   // json编码错误
    JSON_UNMAR_ERR = 10007, // json解码错误

    CODE_NODE_ERROR   = 9000, // 验证码服务异常
    CODE_NODE_NOINIT  = 9001, // 验证码服务没有初始化
    DATA_NODE_ERROR   = 9002, // 数据服务异常
    DATA_NODE_NOINIT  = 9003, // 数据服务没有初始化
    DARTY_NODE_ERROR  = 9004, // 第三方服务异常
    DARTY_NODE_NOINIT = 9005, // 第三方服务没有初始化

    WECHAT_LOGIN_ERROR = 30001,     // 微信登录错误
    WECHAT_TOKEN_TIMEOUT = 30002,   // 微信token过期
    WECHAT_REFRESH_ERROR = 30003,   // 微信token刷新失败
    WECHAT_USERINFO_ERR = 30004,    // 微信信息获取错误

    ALIPAY_LOGIN_ERROR = 30031,     // 微信登录错误
    ALIPAY_TOKEN_TIMEOUT = 30032,   // 微信token过期
    ALIPAY_REFRESH_ERROR = 30033,   // 微信token刷新失败
    ALIPAY_USERINFO_ERR = 30034,    // 微信信息获取错误

    BAIDU_ACCESS_TOKEN_ERROR = 30061,   // 百度accesstoken错误
    BAIDU_ENTITY_ERROR = 30062,         // 百度entity请求错误
    BAIDU_IDIDENT_ERROR = 30063,        // 百度身份证识别请求错误
    BAIDU_IDIDENT_FAIL = 30084,         // 百度身份证识别失败
}