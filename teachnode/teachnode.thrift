namespace go teachnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct teaccher_subject {
    1:string        name,
    2:i32           score,
}

struct teacher_item {
    1:i64                       uid,
    2:string                    tag,
    3:string                    resume,
    4:list<teaccher_subject>    subject,
    5:i32                       status,
    6:i64                       tmlong,
}

struct update_teacher_req {
    1:common.authorize      auth,
    2:teacher_item          teacher,
    3:map<string,string>    extra,
}

struct update_teacher_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:map<string,string> extra, 
}

struct find_teacher_item {
    1:teacher_item      teacher,
    2:common.user_info  info,
}

struct find_teacher_req {
    1:common.authorize      auth,
    2:i64                   tuid,
    3:i64                   uuid,
    4:string                subject,
    5:i32                   page,
    6:i32                   size,
    7:string                sorted,
    8:map<string,string>    extra,
}

struct find_teacher_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<find_teacher_item>   teachers,
    4:map<string,string>        extra,
}

struct focus_teacher_req {
    1:common.authorize      auth,
    2:i64                   tuid,
    3:i64                   uuid,
    4:i32                   status,
    5:map<string,string>    extra,
}

struct focus_teacher_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:map<string,string> extra,
}

service teachnode_service {
    // 添加或者根据教师信息
    update_teacher_res update_teacher(1:update_teacher_req req),

    // 查询教师信息
    find_teacher_res get_teacher(1:find_teacher_req req),

    // 关注/取消关注教师
    focus_teacher_res focus_teacher(1:focus_teacher_req req),
}