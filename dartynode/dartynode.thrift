namespace go dartynode

include "../rescode/rescode.thrift"

struct login_req {
    1:string                account;
    2:string                passwd;
    3:map<string,string>    extra;
}

struct login_wechat_req {
    1:string                appid;
    2:string                Code;
    3:string                State;
    14:map<string,string>   extra;
}

struct login_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie;
    4:map<string,string>    extra;
}

service dartynode_service {
    // 登录
    login_wechat_res login_wechat(1:login_wechat_req req),
}