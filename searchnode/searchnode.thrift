namespace go searchnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct search_question_req {
    1:common.authorize auth,
    2:i64   uid,
    3:string keyword,
    4:i32    page,
    5:i32   size,
    6:map<string,string> extra,
}

struct search_question_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:double                            timestamp,
    4:common.search_hits_total          total,
    5:list<common.search_hits_item>     hits,
    6:map<string,string>                extra, 
}

struct search_item_req {
    1:common.authorize auth,
    2:string            keyword,
    3:i32               page,
    4:i32                size,
    5:map<string,string> extra,
}

struct search_item_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:list<string>                      lists,
    4:map<string,string>                extra,
}

struct search_history_req {
    1:common.authorize auth,
    2:i64           uid,
    3:i32           page,
    4:i32           size,
    5:map<string,string> extra,
}

struct search_history_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:list<common.search_history_item>  lists,
    4:map<string,string>                extra,
}

service searchnode_service {
    // 搜索试题
    search_question_res search_question(1:search_question_req req),

    // 搜索关键字补全
    search_item_res search_item(1:search_item_req req),

    // 获取搜索记录
    search_history_res search_history(1:search_history_req req),
}