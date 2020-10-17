namespace go registernode

include "../rescode.thrift"

struct normal_user_req {
    1:string    account,
    2:string    passwd,
    3:string    code,
    4:bool      policy,
    5:string    source,
    6:string    invitor,
    7:map<string,string> extra,
}

struct normal_user_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct cooperator_user_req {
    1:string    account,
    2:string    passwd,
    3:string    code,  
    4:string    contactor,  // 联系人
    5:string    email,
    6:string    brand_name, 
    7:string    company_name,
    8:bool      policy,
    9:string    invitor,
    10:map<string,string> extra,
}

struct cooperator_user_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

service codenode_service {
    // 创建普通用户
    normal_user_res   normal_user(1:normal_user_req req), 

    // 创建合作用户
    cooperator_user_res cooperator_user(1:cooperator_user_req req),
}