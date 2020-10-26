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

struct update_user_req {
    1:authorize auth,
    2:string uid,
    3:string user_name,
    4:string brithday,
    5:string card_id,
    6:string address,
    7:string nick_name,
    8:string avatar,
    9:string attention,
    10:string   login_at,
    11:string   login_ip,
    12:string   last_at,
    13:i16      status,
    14:string   email,
    20:string   update_at,
    16:string   contact_name,
    17:string   brand_name,
    18:string   company_name,
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
    update_user_res update_user(1:update_user_req req),

    user_info_res user_info(1:user_info_req req), 
}