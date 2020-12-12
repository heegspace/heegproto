namespace go registernode

include "../rescode/rescode.thrift"
include "../common/common.thrift"


struct normal_user_req {
    1:common.authorize auth,
    2:string    account,
    3:string    passwd,
    4:string    code,
    5:bool      policy,
    6:string    source,
    7:string    invitor,
    8:map<string,string> extra,
}

struct normal_user_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct cooperator_user_req {
    1:common.authorize auth,
    2:string    account,
    3:string    passwd,
    4:string    code,  
    5:string    contactor,  // 联系人
    6:string    email,
    7:bool      policy,
    8:string    source,
    9:string    invitor,
    10:map<string,string> extra,
}

struct cooperator_user_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

service registernode_service {
    // 创建普通用户
    normal_user_res   normal_user(1:normal_user_req req), 

    // 创建合作用户
    cooperator_user_res cooperator_user(1:cooperator_user_req req),
}