namespace go authnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"


struct verify_token_req {
    1:string                token,
    2:common.authorize      auth,
    3:map<string,string>    extra,
}

struct verify_token_res {
    1:rescode.code          rescode,
    2:map<string,string>    extra,
}

struct admin_role_req {
    1:string                token,
    2:common.authorize      auth,
    3:map<string,string>    extra,
}

struct admin_role_res {
    1:rescode.code          rescode,
    2:map<string,string>    extra,
}

struct cooper_role_req {
    1:string                token,
    2:common.authorize      auth,
    3:map<string,string>    extra,
}

struct cooper_role_res {
    1:rescode.code          rescode,
    2:map<string,string>    extra,
}

service authnode_service {
    // 验证token
    verify_token_res verify_token(1:verify_token_req req);

    // 验证是否是管理用户
    admin_role_res admin_role(1:admin_role_req req);

    // 验证是否是合作用户
    cooper_role_res cooper_role(1:cooper_role_req req);
}