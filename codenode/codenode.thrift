namespace go codenode

include "../rescode.thrift"

enum code_type {
    UID         = 0,
    ACCOUNT     = 1,
    MOBILE      = 2,
    EMAIL       = 3,
}

struct authorize {
    1:string    key,
    2:string    value,
    3:map<string,string> extra,
}

struct code_req {
    1:authorize             auth,
    2:string                desc,
    3:code_type             type,
    4:map<string,string>    extra,
}

struct code_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                code,
    4:map<string,string>    extra,
}

struct verify_code_req {
    1:authorize             auth,
    2:string                desc,
    4:string                code,
    5:map<string,string>    extra,
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