namespace go notenode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct note_item {
    1:i64                   nid,
    2:string                data,
    3:string                html,
    4:list<string>          tag,
    5:string                bgcolor,
    6:i64                   note_type,
}

struct update_note_req {
    1:common.authorize             auth,
    2:i64                       uid,
    3:note_item             note,
    5:map<string,string>    extra,
}

struct note_meta_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:common.note_meta  meta,
    4:map<string,string> extra,
}
 
struct note_meta_list_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:i64                   userid,
    4:string                tag,
    5:i64                   status,
    6:i32                   page,
    7:i32                   size,
    8:i64                   note_type,
    9:map<string,string>    extra,
}

struct note_meta_list_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<common.note_meta>      meta,
    4:map<string,string> extra,
}

struct note_list_count_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:i64                   userid,
    4:string                tag,
    5:i64                   status,
    6:i64                   note_type,
    7:map<string,string>    extra,
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


struct note_cooper_req {
    1:common.authorize             auth,
    2:i64                   nid,
    3:i64                   userid,
    4:list<string>          user,
    5:map<string,string>    extra,
}

struct note_cooper_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<common.user_info>         user,
    4:map<string,string> extra,
}

struct note_tag_req {
    1:common.authorize             auth,
    2:i64                   nid,
    3:i64                   userid,
    4:list<string>          tags,
    5:map<string,string>    extra,
}

struct note_tag_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<string>      tags,
    4:map<string,string> extra,
}

struct note_bgcolor_req {
    1:common.authorize          auth,
    2:i64                   nid,
    3:i64                   userid,
    4:string                bgcolor,
    5:map<string,string>    extra,
} 

struct note_bgcolor_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:string             bgcolor,
    4:map<string,string> extra,
}


struct note_tag_add_req {
    1:common.authorize          auth,
    2:i64                   userid,
    3:string                cn,
    4:string                eg,
    5:map<string,string>    extra,
}

struct note_tag_add_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct note_tag_list_req {
    1:common.authorize          auth,
    2:i64                   user_id,
    3:string                lang,
    4:i64                   status,
    5:i32                   page,
    6:i32                   size,
    7:map<string,string>    extra,
}

struct note_tag_list_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.note_tag>     tags,
    4:map<string,string>        extra,
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


    // 更新笔记的协作者、标签、颜色
    note_cooper_res note_cooper(1:note_cooper_req req),
    note_tag_res note_tag(1:note_tag_req req),
    note_bgcolor_res note_bgcolor(1:note_bgcolor_req req),

    note_tag_add_res note_tag_add(1:note_tag_add_req req),
    note_tag_list_res note_tag_list(1:note_tag_list_req req),
}