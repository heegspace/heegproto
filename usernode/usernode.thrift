namespace go usernode

include "../rescode.thrift"

struct authorize {
    1:string    key,
    2:string    value,
    3:map<string,string> extra,
}

struct user_obj {
    1:string    uid,
    2:string    account,
    4:string    user_name,
    5:string    brithday,
    6:string    card_id,
    7:string    address,
    8:string    nick_name,
    9:string    avatar,
    10:string   phone,
    11:string   login_at,
    12:string   login_ip,
    13:string   last_at,
    14:i16      status,
    15:i64      role,
    16:string   email,
    17:string   contact_name,
    18:string   brand_name,
    19:string   company_name,
    20:string   attention,
    21:string   update_at,
    22:i64      vip,
    23:double   coin
}

struct update_username_req {
    1:authorize auth,
    2:string uid,
    3:string user_name
    19:map<string,string> extra,
}

struct update_brithday_req {
    1:authorize auth,
    2:string uid,
    3:string brithday
    19:map<string,string> extra,
}

struct update_cardid_req {
    1:authorize auth,
    2:string uid,
    3:string card_id
    19:map<string,string> extra,
}

struct update_avatar_req {
    1:authorize auth,
    2:string uid,
    3:string card_id
    19:map<string,string> extra,
}

struct update_attention_req {
    1:authorize auth,
    2:string uid,
    3:string card_id
    19:map<string,string> extra,
}

struct update_user_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct user_info_req {
    1:authorize auth,
    2:string    uid,
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