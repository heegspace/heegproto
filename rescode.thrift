namespace go heegproto.rescode

enum code {
    SUCCESS =   0,      // 成功
    ERROR   =   1,      // 失败
    
    PARAM_ERR = 99,     // 参数错误

    EXISTS  =   100,    // 已经存在
    IS_SELF =   101,    // 自己

    NOT_EXISTS = 400,    // 不存在
    SEND_CODE_ERR = 401, // 发送验证码错误
}