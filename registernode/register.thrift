namespace go registernode

include "../rescode.thrift"

struct authorize {
    1:string    key,
    2:string    value,
    3:map<string,string> extra,
}

struct normal_user_req {
    1:authorize auth,
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
    1:authorize auth,
    2:string    account,
    3:string    passwd,
    4:string    code,  
    5:string    contactor,  // 联系人
    6:string    email,
    7:string    brand_name, 
    8:string    company_name,
    9:bool      policy,
    10:string    source,
    11:string    invitor,
    12:map<string,string> extra,
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