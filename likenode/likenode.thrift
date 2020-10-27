namespace go likenode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct likes_count_req {
    1:common.authorize     auth,
    2:string        uid,
    3:string        mid,
}

struct likes_count_res {
    1:rescode.code      rescode,
    2:i32               count,
    3:string            resmsg,
    4:map<string,string> extra,
}

struct likes_add_req {
    1:common.authorize     auth,
    2:string        uid,
    3:string        mid,
}

struct likes_add_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct likes_list_req {
    1:common.authorize     auth,
    2:string        uid,
    3:string        mid,
    4:i32           page,
    5:i32           size,
}

struct likes {
    1:string            uid,
    2:i64               create_at,
}

struct likes_list_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<likes>       data,
    4:map<string,string> extra,
}

service likenode_service {
    // 点赞数量
    likes_count_res likesCount(1:likes_count_req req),

    // 添加点赞
    likes_add_res likesAdd(1:likes_add_req req),

    // 获取点赞列表
    likes_list_res likesList(1:likes_list_req req),
}