namespace go searchnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

enum Search_Style {
    PUBLIC_SEARCH           = 0,        // 公共搜索
    PRIVATE_SEARCH          = 1,        // 私有搜索
}

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

struct search_history_req {
    1:common.authorize auth,
    2:i64   uid,
    3:Search_Style  style,
    4:i32    page,
    5:i32   size,
    6:map<string,string> extra,
}

struct search_history_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:list<common.search_history_item>  lists,
    4:i32                               page,
    5:i32                               size,
    6:map<string,string>                extra,
}

service searchnode_service {
    // 搜索试题
    search_question_res search_question(1:search_question_req req),

    // 获取搜索记录
    search_history_res search_history(1:search_history_req req),
}