namespace go notenode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct note_item {
    1:i64                   nid,
    2:string                data,
    3:string                html,
}

struct update_note_req {
    1:common.authorize             auth,
    2:i64                   uid,
    3:note_item             note,
    5:map<string,string>    extra,
}

struct note_meta_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:common.note_meta         meta,
    4:map<string,string> extra,
}
 
struct note_meta_list_req {
    1:common.authorize             auth,
    2:i64                   uid,
    3:i32                   page,
    4:i32                   size,
    5:map<string,string> extra,
}

struct note_meta_list_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<common.note_meta>      meta,
    4:map<string,string> extra,
}

struct note_list_count_req {
    1:common.authorize             auth,
    2:i64                   uid,
    3:map<string,string>    extra,
}

struct note_list_count_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:i32      count,
    4:map<string,string> extra,
}

struct note_data_req {
    1:common.authorize             auth,
    2:i64                uid,
    3:i64                nid,
    4:map<string,string>    extra,
}

struct note_data_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:string            data,
    4:map<string,string> extra,
}

struct note_html_req {
    1:common.authorize             auth,
    2:i64                uid,
    3:i64                nid,
    4:map<string,string>    extra,
}

struct note_html_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:string            data,
    4:map<string,string> extra,
}

service notenode_service {
    //更新或者添加笔记信息
    note_meta_res updateNote(1:update_note_req req),

    // 获取笔记列表
    note_meta_list_res noteMetaList(1:note_meta_list_req req),

    // 获取用户笔记数量
    note_list_count_res noteListCount(1:note_list_count_req req),

    // 获取笔记数据列
    note_data_res noteData(1:note_data_req req),

    // 获取笔记html数据
    note_html_res noteHtml(1:note_html_req req),
}