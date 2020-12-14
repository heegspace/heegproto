namespace go codenode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

enum code_type {
    UID         = 0,
    ACCOUNT     = 1,
    MOBILE      = 2,
    EMAIL       = 3,
}

// 请求验证码的操作类型
enum operate_type {
    REGISTER    = 101,  // 注册请求
    LOGIN       = 101,  // 登录请求
}

struct code_req {
    1:common.authorize      auth,
    2:string                desc,
    3:code_type             type,
    4:operate_type         optype,
    5:map<string,string>    extra,
}

struct code_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                code,
    4:map<string,string>    extra,
}

struct verify_code_req {
    1:common.authorize      auth,
    2:string                desc,
    4:string                code,
    5:operate_type         optype,
    6:map<string,string>    extra,
}

struct verify_code_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    4:map<string,string>    extra,
}

service codenode_service {
    // 发送验证码
    code_res   send_code(1:code_req req), 
    // 请求验证码
    verify_code_res verify_code(1:verify_code_req req),
}