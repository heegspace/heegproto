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

service loginnode_service {
    // 登录
    login_res login(1:login_req req);
    // 退出登录
    logout_res logout(1:logout_req req);
    // 刷新
    refresh_res refresh(1:refresh_req req);
}