namespace go datanode

include "../rescode.thrift"
struct user{
    1:string    uid,
    2:string    account,
    3:string    pass_wd,
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
    23:double  coin
}

struct user_info_req {
    1:string    uid
}

struct new_user_req {
    1:string    account,
    2:string    code,
    3:string    pass_wd,
    4:string    client_ip,
    5:i32       role,        
}

struct search_user_req {
    1:list<string>      uids
    2:string            user_name
    3:string            phone
    4:string            email
}

struct user_res {
    1:rescode.code    rescode,
    2:string    resmsg,
    3:user      user,
}

service datanode_service {
    // ---------- 用户接口 ------- //
    // 创建新用户
    user_res createUser(1:new_user_req req),
    // 获取用户信息
    user_res userInfo(1:user_info_req req),
    // 搜索用户信息
    list<user> searchUser(1:search_user_req req),
    // 更新用户信息
    user_res updateUser(1:user req),
}
