namespace go notenode

include "../rescode.thrift"

struct authorize {
    1:string    key,
    2:string    value,
    3:map<string,string> extra,
}

struct update_note_req {
    1:authorize             auth,
    2:string                nid,
    3:string                data,
    4:string                html,
    5:map<string,string>    extra,
}

service lognode_service {
    // ----------------- 笔记 ----------------// 
    //更新或者添加笔记信息
    note_meta_res updateNote(1:string uid, 2:update_note_req req),
    // 获取笔记列表
    note_meta_list_res noteMetaList(1:string uid, 2:i32 page, 3:i32 size),
    // 获取用户笔记数量
    note_list_count_res noteListCount(1:string uid),
    // 获取笔记数据列
    note_data_res noteData(1:string uid, 2:string nid),
    // 获取笔记html数据
    note_html_res noteHtml(1:string uid, 2:string nid),
}