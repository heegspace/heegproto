namespace go s2sname

include "../rescode/rescode.thrift"

enum Const {
    Expired  = 5, // 超时时间
}

struct s2sname {
    1:string    host,
    2:i32   port,
    3:i32   prority,
    4:string name,
    5:i64   expired;
}

struct register_req {
    1:string    name,
    2:s2sname   s2s,
}

struct register_res {
    1:rescode.code    rescode,
    2:string    resmsg,
}

struct update_req {
    1:string    name,
    2:s2sname   s2s,
}

struct update_res {
    1:rescode.code    rescode,
    2:string    resmsg,
}

struct fetch_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<s2sname>     s2ss,
}

struct heart_req {
    1:string    name,
    2:s2sname   s2s,
}

struct heart_res {
    1:rescode.code    rescode,
    2:string    resmsg,
}

service s2sname_service {
    // 注册s2s服务
    register_res registerS2sname(1:register_req req),

    // 更新s2sname信息
    update_res updateS2sname(1:update_req req),

    // 根据参数获取对应的s2s列表
    fetch_res fetchS2sname(1:string req),

    // 获取所有的s2sname 列表信息
    fetch_res fetchS2snames(),

    // 服务心跳包，用于更新服务的状态
    heart_res heart(1:heart_req req),
}