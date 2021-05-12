namespace go certnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct submit_cert_req {
    1:common.authorize      auth,
    2:common.person_cert    cert,
    3:map<string,string>    extra,
}

struct submit_cert_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct cert_approved_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:i64                   approve_uid,
    4:string                approve_name,
    5:string                info,
    6:map<string,string>    extra,
}

struct cert_approved_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:common.person_cert    cert,
    4:map<string,string>    extra,
}

struct cert_refuse_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:i64                   refuse_uid,
    4:string                refuse_name
    5:string                info,
    6:map<string,string>    extra,
}

struct cert_refuse_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:common.person_cert    cert,
    4:map<string,string>    extra,
}


struct cert_info_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:string                status,
    4:string                idname,
    5:string                idnum,
    6:i32                   page,
    7:i32                   size,
    8:map<string,string>    extra,
}

struct cert_info_res {
    1:rescode.code                  rescode,
    2:string                        resmsg,
    3:list<common.person_cert>      cert,
    4:map<string,string>            extra,
}

struct cert_flow_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:i32                   page,
    4:i32                   size,
    5:map<string,string>    extra,
}

struct cert_flow_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.person_cert>  cert,
    4:map<string,string>        extra,
}

struct cert_cancel_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:string                info,
    4:map<string,string>    extra,
}

struct cert_cancel_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct cert_cache_req {
    1:common.authorize      auth,
    2:string                key,
    3:string                value,
    4:i64                  expire,
    5:map<string,string>    extra,
}

struct cert_cache_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

service certnode_service {
    // 提交实名
    submit_cert_res submit_cert(1:submit_cert_req req),

    // 获取实名信息
    cert_info_res cert_info(1:cert_info_req req),

    // 实名日志记录
    cert_flow_res cert_flow(1:cert_flow_req req),

    // 取消实名
    cert_cache_res cert_cancel(1:cert_cancel_req req),


    //--------------需要授权--------------//
    // 审核通过
    cert_approved_res cert_approved(1:cert_approved_req req),

    // 实名失败 
    cert_refuse_res cert_refuse(1:cert_refuse_req req),

    // 实名缓存
    cert_cache_res cert_cache(1:cert_cache_req req),
}