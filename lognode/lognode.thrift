namespace go lognode

include "../rescode.thrift"

struct call_log_req {
    1:i32       level,
    2:string    func,
    3:string    timestamp,
    4:string    status,
    5:string    req,
    6:string    res,
    7:map<string,string> extra,
}

struct log_req {
    1:i32                   level,
    2:string                timestamp,
    3:string                func,
    4:string                info,
    5:map<string,string>    extra, 
}

service lognode_service {
    void log(1:log_req req),
    void call_log(1:call_log_req req),
}
