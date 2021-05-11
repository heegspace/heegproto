namespace go lognode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

enum log_level {
    INFO  = 0,
    DEBUG = 1,
    WARN  = 2,
    ERROR = 3,
}

enum log_desc {
    CONSOLE = 0,
    FILE = 1,
    DB = 2,
}

struct call_log_req {
    1:log_level level,
    2:log_desc  desc,
    3:string    server_name,
    4:string    ip,
    5:string    func,
    6:string    timestamp,
    7:string    status,
    8:string    req,
    9:string    res,
    10:map<string,string> extra,
}

struct log_req {
    1:log_level             level,
    2:log_desc              desc,
    3:string                server_name,
    4:string                ip,
    5:string                timestamp,
    6:string                func,
    7:string                info,
    8:map<string,string>    extra, 
}

struct query_user_log_req {
    1:common.authorize      auth,
    2:i64                   id,
    3:i64                   uid,
    4:i32                   log_type,
    5:i32                   page,
    6:i32                   size,
    7:map<string,string>    extra,
}

struct query_user_log_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.user_log>     logs,
    4:map<string,string>        extra,
}

service lognode_service {
    void log(1:log_req req),
    void call_log(1:call_log_req req),

    // 获取日志
    query_user_log_res query_user_log(1:query_user_log_req req),
}
