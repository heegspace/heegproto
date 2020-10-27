namespace go momnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct moments_count_req {
    1:common.authorize     auth,
    2:string        uid,
}

struct moments_count_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:i32               count,
    4:map<string,string> extra,
}

struct extra {
    1:string path,
    2:i32 type,
}

struct moments {
    1:i32   id,
    2:string mid,
    3:string text,
    4:list<extra> extra,
    5:i32 create_at,
}

struct add_moments_req {
    1:common.authorize     auth,
    2:string        uid,
    3:moments       moments,
    4:map<string,string> extra,
}

struct add_moments_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct moments_list_req {
    1:common.authorize     auth,
    2:string        uid,
    3:i32           page,
    4:i32           size,
}

struct moments_list_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<moments>     data,
}

service momnode_service {
    // 获取动态的数量
    moments_count_res momentsCount(1:moments_count_req req),

    // 添加动态
    add_moments_res momentsAdd(1:add_moments_req req),

    // 获取动态列表
    moments_list_res momentsList(1:moments_list_req req),
}