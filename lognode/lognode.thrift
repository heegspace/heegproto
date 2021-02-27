namespace go lognode

include "../rescode/rescode.thrift"

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

service lognode_service {
    void log(1:log_req req),
    void call_log(1:call_log_req req),
}
