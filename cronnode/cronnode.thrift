namespace go cronnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct cron_call_req {
    1:string                func,
    2:i64                   call_time,
    3:map<string,string>    extra,
}

struct cron_call_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

service cronnode_service {
    // 提交实名
    cron_call_res cron_call(1:cron_call_req req),
}