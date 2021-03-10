namespace go sensinode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct refresh_modify_reward_req {
    1:common.authorize      auth,
    2:i64                   id,
    3:double                reward,
    4:map<string,string>    extra,
} 

struct refresh_modify_reward_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                code,
    4:map<string,string>    extra,
}

struct refresh_user_coin_req {
    1:common.authorize      auth,
    2:i64                   userid,
    3:double                coin,
    4:map<string,string>    extra,
}

struct refresh_user_coin_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                code,
    4:map<string,string>    extra,
}



struct refresh_add_reward_req {
    1:common.authorize      auth,
    2:i64                   id,
    3:double                reward,
    4:map<string,string>    extra,
} 

struct refresh_add_reward_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                code,
    4:map<string,string>    extra,
}


service sensinode_service {
    // 更新修改试题的奖励积分
    refresh_modify_reward_res refresh_modify_reward(1:refresh_modify_reward_req req),

    // 更新添加试题的奖励积分
    refresh_add_reward_res refresh_add_reward(1:refresh_add_reward_req req),

    // 更新用户的coin数值
    refresh_user_coin_res refresh_user_coin(1:refresh_user_coin_req req),
}