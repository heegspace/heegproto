namespace go usernode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

enum attention_op_type {
    ADD_ATTENTION = 1;
    DEL_ATTENTION = 2;
}

struct user_obj {
    1:string    phone,
    2:string    email,
    3:string    account,
    4:i16       status,
    5:i64       role,
    6:i64       vip,
    7:double    coin
    8:string    reg_at,
    9:string    nick_name,
    10:string   avatar,
    11:i32      sex,
    12:string   province,
    13:string   city,
    14:string   country,
    15:bool     vip_expire,
}

struct update_cardid_req {
    1:common.authorize auth,
    2:i64 uid,
    3:string card_id
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
    1:common.authorize      auth,
    2:i64                   uid,
    3:common.from_platom    platom, // 帐号来源平台
    4:map<string,string>    extra,
}

struct user_info_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:user_obj          user,
    4:map<string,string> extra,
}

struct update_userinfo_req {
    1:common.authorize                  auth,
    2:common.update_userinfo_item       info,
    3:map<string,string>                extra,
}

struct update_userinfo_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:string                    value,
    4:map<string,string>        extra,
}

struct user_cache_req {
    1:common.authorize      auth,
    2:string                key,
    3:string                value,
    4:i64                  expire,
    5:map<string,string>    extra,
}

struct user_cache_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

service usernode_service {
    // 更新用户信息
    update_userinfo_res updateUserInfo(1:update_userinfo_req req),

    // 更新身份证号
    update_user_res update_cardid(1:update_cardid_req req),

    // 更新关注对象
    update_user_res update_attention(1:update_attention_req req),

    // 获取用户信息
    user_info_res user_info(1:user_info_req req), 

    // 用户缓存
    user_cache_res user_cache(1:user_cache_req req),
}