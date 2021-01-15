namespace go dartynode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct login_wechat_req {
    1:string                appid,
    2:string                Code,
    3:string                State,
    14:map<string,string>   extra,
}

struct login_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct refresh_wechat_req {
    1:i64                   uid,
    2:string                appid,
    3:string                refresh_token,
    4:string                access_token,
    5:map<string,string>    extra,
}

struct refresh_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct logout_wechat_req {
    1:i64                   uid,
    2:string                appid,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct logout_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}


service dartynode_service {
    // 登录微信
    login_wechat_res login_wechat(1:login_wechat_req req),

    // 刷新微信的token
    refresh_wechat_res refresh_wechat(1:refresh_wechat_req req),

    // 退出微信登录
    logout_wechat_res logout_wechat(1:logout_wechat_req req),
}