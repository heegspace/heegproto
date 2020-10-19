namespace go lognode

include "../rescode.thrift"

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
    3:string    func,
    4:string    timestamp,
    5:string    status,
    6:string    req,
    7:string    res,
    8:map<string,string> extra,
}

struct log_req {
    1:log_level             level,
    2:log_desc              desc,
    3:string                timestamp,
    4:string                func,
    5:string                info,
    6:map<string,string>    extra, 
}

service lognode_service {
    void log(1:log_req req),
    void call_log(1:call_log_req req),
}
