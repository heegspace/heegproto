namespace go friendnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct add_friend_req {
    1:common.authorize auth,
    2:i64               uid,
    3:string            phone,
    4:map<string,string> extra,
}

struct add_friend_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct agree_friend_req {
    1:common.authorize     auth,
    2:i64                   suid, // 用户id
    3:i64                   uid,
    4:map<string,string> extra,
}

struct agree_friend_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct friend_list_req {
    1:common.authorize     auth,
    2:i64                   uid,
}

struct friend_item {
    1:i64       uid,
    2:string    note,
    3:string    account,
    4:string    nick_name,
    5:string    avatar,
    6:string    brithday,
    7:i64       join_at,
}

struct group {
    1:i64               sort,
    2:string            name,
    3:list<friend_item> lists,
}

struct friend {
    1:i64                   total,
    2:list<group>           data,
    3:list<friend_item>     invite,
    4:list<friend_item>     blacks,
}

struct friend_list_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:friend                friends,
    4:map<string,string>    extra,
}

struct create_group_req {
    1:common.authorize     auth,
    2:i64                   uid,
    3:string                name,
    4:map<string,string> extra,
}

struct create_group_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct rename_group_req {
    1:common.authorize     auth,
    2:i64                   uid,
    3:string                old_name,
    4:string                new_name,
    5:map<string,string> extra,
}

struct rename_group_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct add_friend_note_req {
    1:common.authorize         auth,
    2:i64                   suid,
    3:i64                   uid,
    4:string                note,
    5:map<string,string> extra,
}

struct add_friend_note_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct move_group_req {
    1:common.authorize             auth,
    2:i64                       suid,
    3:i64                       uid,
    4:string                    group,
    5:map<string,string>        extra,
}

struct move_group_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct remove_friend_req {
    1:common.authorize         auth,
    2:i64                   suid,
    3:i64                   uid,
    4:map<string,string> extra,
}

struct remove_friend_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

service friendnode_service {
    // --------------- 好友接口 ----------- //
    // 添加好友
    add_friend_res addFriends(1:add_friend_req req),

    // 同意好友
    agree_friend_res agreeFriends(1:agree_friend_req req),

    //  请求好友列表
    friend_list_res friendsList(1:friend_list_req req),

    // 添加组
    create_group_res createGroup(1:create_group_req req),

    // 重命名组
    rename_group_res renameGroup(1:rename_group_req req),

    // 添加好友备注
    add_friend_note_res addNoteFriend(1:add_friend_note_req req),
    
    // 移动到新的组
    move_group_res moveToNewGroup(1:move_group_req req),

    // 删除好友
    remove_friend_res removeFriend(1:remove_friend_req req),
}