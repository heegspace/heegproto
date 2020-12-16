namespace go usernode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

enum attention_op_type {
    ADD_ATTENTION = 1;
    DEL_ATTENTION = 2;
}

struct user_obj {
    1:i64       uid,
    2:string   phone,
    3:string   email,
    4:string    account,
    5:string    pass_wd,
    6:i16      status,
    7:i64      role,
    8:i64      vip,
    9:double   coin
    10:string  reg_at;
}

struct update_username_req {
    1:common.authorize auth,
    2:i64 uid,
    3:string user_name
    4:map<string,string> extra,
}

struct update_brithday_req {
    1:common.authorize auth,
    2:i64 uid,
    3:string brithday
    4:map<string,string> extra,
}

struct update_cardid_req {
    1:common.authorize auth,
    2:i64 uid,
    3:string card_id
    4:map<string,string> extra,
}

struct update_avatar_req {
    1:common.authorize auth,
    2:i64 uid,
    3:string avatar
    4:map<string,string> extra,
}

struct update_attention_req {
    1:common.authorize auth,
    2:i64               uid,
    3:attention_op_type op, // 0:添加;1:删除
    4:list<i64>         aid,
    5:map<string,string> extra,
}

struct update_user_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct user_info_req {
    1:common.authorize auth,
    2:i64    uid,
    3:map<string,string> extra,
}

struct user_info_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:user_obj          user,
    4:map<string,string> extra,
}

service usernode_service {
    // 更新用户名
    update_user_res update_username(1:update_username_req req),

    // 更新生日
    update_user_res update_brithday(1:update_brithday_req req),

    // 更新身份证号
    update_user_res update_cardid(1:update_cardid_req req),

    // 更新头像信息
    update_user_res update_avatar(1:update_avatar_req req),

    // 更新关注对象
    update_user_res update_attention(1:update_attention_req req),

    // 获取用户信息
    user_info_res user_info(1:user_info_req req), 
}