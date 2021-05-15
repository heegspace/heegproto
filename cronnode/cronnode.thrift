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

struct add_task_req {
    1:common.authorize      auth,
    2:common.cron_item      cron,
    3:string                admin,
    4:map<string,string>    extra,
}

struct add_task_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct get_task_req {
    1:common.authorize      auth,
    2:bool                  repeated,
    3:bool                  mutitask,
    4:string                start_at,
    5:map<string,string>    extra,
}

struct get_task_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.cron_item>    crons,
    4:map<string,string>        extra,
}

struct get_task_count_req {
    1:common.authorize      auth,
    2:bool                  repeated,
    3:bool                  mutitask,
    4:string                start_at,
    5:map<string,string>    extra,
}

struct get_task_count_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:i64                       count,
    4:map<string,string>        extra,
}

service cronnode_service {
    // 提交实名
    cron_call_res cron_call(1:cron_call_req req),

    // 添加任务
    add_task_res add_task(1:add_task_req req),

    // 获取任务
    get_task_res get_task(1:get_task_req req),

    // 获取任务数量
    get_task_count_res get_task_count(1:get_task_count_req req),
}