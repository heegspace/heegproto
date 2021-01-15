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
    1:string                appid,
    2:string                refresh_token,
    3:string                access_token,
    4:map<string,string>    extra,
}

struct refresh_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct logout_wechat_req {
    1:string                appid,
    2:string                cookie,
    3:map<string,string>    extra,
}

struct logout_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct userinfo_wechat_req {
    1:i64                   uid,
    2:string                openid,
    3:string                access_token,
    4:string                refresh_token,  // 用于刷新access_token
    5:map<string,string>    extra,
}

struct userinfo_wechat_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:common.wechat_userinfo    userinfo,
    4:string                    cookie, // 是否更新cookie
    5:map<string,string>        extra,
}

service dartynode_service {
    // 登录微信
    login_wechat_res login_wechat(1:login_wechat_req req),

    // 刷新微信的token
    refresh_wechat_res refresh_wechat(1:refresh_wechat_req req),

    // 退出微信登录
    logout_wechat_res logout_wechat(1:logout_wechat_req req),

    // 获取用户信息
    userinfo_wechat_res userinfo_wechat(1:userinfo_wechat_req req),
}