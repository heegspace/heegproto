namespace go loginnode

include "../rescode/rescode.thrift"

struct login_req {
    1:string                account;
    2:string                passwd;
    3:map<string,string>    extra;
}

struct login_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie;
    4:map<string,string>    extra;
}

struct login_by_code_req {
    1:string                account;
    2:string                code;
    3:map<string,string>    extra;
}

struct login_by_code_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie;
    4:map<string,string>    extra;
}


struct login_wechat_req {
    1:string                appid;
    2:string                Code;
    3:string                State;
    4:map<string,string>    extra;
}

struct login_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie;
    4:map<string,string>    extra;
}

struct logout_req {
    1:string                cookie;
    2:map<string,string>    extra;
}

struct logout_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra;
}

struct refresh_req {
    1:string                cookie;
    2:map<string,string>    extra;
}

struct refresh_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie;
    4:map<string,string>    extra;
}

struct login_alipay_req {
    1:string                app_id;
    2:string                source;
    3:string                scope;
    4:string                auth_code
    5:map<string,string>    extra;
}

struct login_alipay_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie;
    4:map<string,string>    extra;
}

service loginnode_service {
    // 登录
    login_res login(1:login_req req);
    login_by_code_res login_by_code(1:login_by_code_req req),
    login_wechat_res login_wechat(1:login_wechat_req req),
    login_alipay_res login_alipay(1:login_alipay_req req),

    // 退出登录
    logout_res logout(1:logout_req req);
    // 刷新
    refresh_res refresh(1:refresh_req req);
}