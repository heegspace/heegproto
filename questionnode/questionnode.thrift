namespace go questionnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct grade_cate_add_req {
    1:common.authorize auth,
    2:i64               uid,
    3:common.grade_cate grade,
    4:map<string,string> extra,
}

struct grade_cate_add_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct grade_cate_count_req {
    1:common.authorize auth,
    2:i64               uid,
    3:map<string,string> extra,
}

struct grade_cate_count_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:i32                   count,
    4:map<string,string>    extra,
}

struct grade_cate_list_req {
    1:common.authorize auth,
    2:i64   uid,
    3:i32   page,
    4:i32   size,
    5:map<string,string> extra,
}

struct grade_cate_list_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.grade_cate>  lists,
    4:map<string,string>    extra,
}

struct home_black_data_req {
    1:common.authorize auth,
    2:i64               uid,
    3:i16               school;
    4:map<string,string> extra,
}

struct home_black_data_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.school_static>  lists,
    4:map<string,string>    extra,
}

struct grade_subject_req {
    1:common.authorize auth,
    2:i64   uid,
    3:i32    index,
    4:map<string,string> extra,
}

struct grade_subject_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.grade_subject>  lists,
    4:map<string,string>    extra,
}

struct school_roll_add_req {
    1:common.authorize auth,
    2:i64 uid,
    3:common.school_roll roll,
    4:map<string,string> extra,
}

struct school_roll_add_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct school_roll_count_req {
    1:common.authorize auth,
    2:i64 uid,
    3:map<string,string> extra,
}

struct school_roll_count_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:i32                   count,
    4:map<string,string>    extra,
}

struct school_list_req {
    1:common.authorize auth,
    2:i64       uid,
    3:i32       page,
    4:i32       size,
    5:map<string,string> extra,
}

struct school_list_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.school_roll> lists,
    4:map<string,string>    extra,
}

struct subject_cate_add_req {
    1:common.authorize auth,
    2:i64    uid,
    3:common.subject_cate subject,
    4:map<string,string> extra,
}

struct subject_cate_add_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct subject_cate_count_req {
    1:common.authorize auth,
    2:i64    uid,
    3:map<string,string> extra,
}

struct subject_cate_count_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:i32                   count,
    4:map<string,string>    extra,
}

struct subejct_cate_list_req {
    1:common.authorize auth,
    2:i64       uid,
    3:i32       page,
    4:i32       size,
    5:map<string,string> extra,
}

struct subject_cate_list_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.subject_cate> lists,
    4:map<string,string>    extra,
}

struct question_version_req {
    1:common.authorize auth,
    2:i64    uid,
    3:common.question_query query,
    4:map<string,string> extra,
}

struct question_version_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.version_item> lists,
    4:map<string,string>    extra,
}

struct question_chapter_req {
    1:common.authorize auth,
    2:i64    uid,
    3:common.question_query query,
    4:map<string,string> extra,
}

struct chapter_res_item {
    1:string       uid,
    2:string    title,
    3:list<common.chapter_item> childs,
}

struct question_chapter_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<chapter_res_item> lists,
    4:map<string,string>    extra,
}

struct subject_name_req {
    1:common.authorize auth,
    2:i64       uid,
    3:string    param,
    4:map<string,string> extra,
}

struct subject_name_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                name,
    4:map<string,string>    extra,
}

struct question_source_req {
    1:common.authorize auth,
    2:i64    uid,
    3:common.question_query query,
    4:map<string,string> extra,
}

struct question_source_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.source_item> sources,
    4:map<string,string>    extra,
}

struct question_tixing_req {
    1:common.authorize auth,
    2:i64    uid,
    3:common.question_query query,
    4:map<string,string> extra,
}

struct question_tixing_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.tixing_item> tixings,
    4:map<string,string>    extra,
}

struct question_timu_count_req {
    1:common.authorize auth,
    2:i64    uid,
    3:common.question_query query,
    4:map<string,string> extra,
}

struct question_timu_count_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:i32                   count,
    4:map<string,string>    extra,
}

struct question_timu_req {
    1:common.authorize auth,
    2:i64    uid,
    3:common.question_query query,
    4:map<string,string> extra,
}

struct question_timu_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.timu_item> timus,
    4:map<string,string>    extra,
}

struct collect_timu_req {
    1:common.authorize auth,
    2:i64       uid,
    3:string    cid,
    4:i64       op,
    5:map<string,string> extra,
}

struct collect_timu_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct query_collect_timu_req {
    1:common.authorize auth,
    2:i64       uid,
    3:i32       page,
    4:i32       size,
    5:map<string,string> extra,
}

struct query_collect_timu_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.timu_item> timus,
    4:map<string,string>    extra,
}

struct modify_question_req {
    1:common.authorize auth,
    2:i64       uid,
    3:string    tid,
    4:common.question timu,
    5:map<string,string> extra,
}

struct modify_question_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct modify_list_req {
    1:common.authorize auth,
    2:i64       uid,
    3:i32       page,
    4:i32       size,
    5:string    sorted,
    6:string    status,
    7:map<string,string> extra,
}

struct modify_list_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<common.modify_item> timus,
    4:map<string,string>    extra,
}

struct modify_count_req {
    1:common.authorize auth,
    2:i64       uid,
    3:string    status,
    4:map<string,string> extra,
}

struct modify_count_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:i32                   count,
    4:map<string,string>    extra,
}

struct approve_req {
    1:common.authorize auth,
    2:i64                   id,
    3:i64                   uid,
    4:string                status,
    5:string                info,
    6:double                reward,
    7:common.approve_dest   dest,
    8:map<string,string>    extra, 
}

struct approve_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:map<string,string> extra, 
}

struct timu_by_id_req {
    1:common.authorize      auth,
    2:i64                   uid,
    3:string                tid,
    4:map<string,string>    extra,
}

struct timu_by_id_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:common.timu_item   timu,
    4:map<string,string> extra,
}

struct timu_add_req {
    1:common.authorize          auth,
    2:i64                       uid,
    3:common.add_timu_item      timu,
    4:map<string,string>        extra,
}

struct timu_add_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:map<string,string> extra,
}


struct add_list_req {
    1:common.authorize   auth,
    2:i64               uid,
    3:i32               page,
    4:i32               size,
    5:string            sorted,
    6:string            status,
}

struct add_list_res {
    1:rescode.code                  rescode,
    2:string                        resmsg,
    3:list<common.add_timu_item>    timus,
    4:map<string,string>            extra,
}

struct add_count_req {
    1:common.authorize          auth,
    2:i64                       uid,
    3:string                    status,
    4:map<string,string>        extra, 
}

struct add_count_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:i32                count,
    4:map<string,string> extra, 
}

struct tixing_by_vid_req {
    1:common.authorize          auth,
    2:string                   vid,
    3:map<string,string>    extra,
}

struct tixing_by_vid_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<string>          tixings,
    4:map<string,string>    extra,
}

struct source_by_vid_req {
    1:common.authorize          auth,
    2:string                   vid,
    3:map<string,string>    extra,
}

struct source_by_vid_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<string>          sources,
    4:map<string,string>    extra,
}


struct ident_by_iid_req {
    1:common.authorize       auth,
    2:i64                    uid,
    3:string                 iid, 
    4:map<string,string>     extra,
}

struct ident_by_iid_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:common.baidu_ident_item  result,
    4:map<string,string>        extra,
}

struct ident_list_req {
    1:common.authorize          auth,
    2:i64       uid,
    3:i32       page,
    4:i32       size,
    5:string    sorted,
    6:string    status,
    7:string    iid,
    8:string    sign,
}

struct ident_list_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:list<common.baidu_ident_item>    results,
    4:map<string,string>                extra,
}

struct ident_count_req {
    1:common.authorize          auth,
    2:i64       uid,
    3:string    status,
    4:map<string,string> extra, 
}

struct ident_count_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:i32                count,
    4:map<string,string> extra, 
}

service questionnode_service {
    // 添加年纪信息
    grade_cate_add_res grade_cate_add(1:grade_cate_add_req req),
    // 获取年纪总数
    grade_cate_count_res grade_cate_count(1:grade_cate_count_req req),
    // 获取年级信息
    grade_cate_list_res grade_cate_list(1:grade_cate_list_req req),
    // 获取主页最新动态统计信息
    home_black_data_res home_black_data(1:home_black_data_req req),
    // 获取学籍对应的年级+科目
    grade_subject_res grade_subject(1:grade_subject_req req),
    // 添加学级信息
    school_roll_add_res school_roll_add(1:school_roll_add_req req),
    // 获取年级总数
    school_roll_count_res school_roll_count(1:school_roll_count_req req),
    // 获取年级列表信息
    school_list_res school_list(1:school_list_req req),
    // 添加科目信息
    subject_cate_add_res subject_cate_add(1:subject_cate_add_req req),
    // 获取科目总数
    subject_cate_count_res subject_cate_count(1:subject_cate_count_req req),
    // 获取科目列表
    subject_cate_list_res subject_cate_list(1:subejct_cate_list_req req),


    // 获取版本
    question_version_res question_version(1:question_version_req req),
    // 获取章节
    question_chapter_res question_chapter(1:question_chapter_req req),
    // 获取科目名
    subject_name_res subject_name(1:subject_name_req req),
    // 获取来源
    question_source_res question_source(1:question_source_req req),
    // 获取题型
    question_tixing_res question_tixing(1:question_tixing_req req),
    // 获取题目数量
    question_timu_count_res question_timu_count(1:question_timu_count_req req),
    // 获取题目
    question_timu_res question_timu(1:question_timu_req req),
    // 收藏题目
    collect_timu_res collect_timu(1:collect_timu_req req),
    // 获取收藏的题目
    query_collect_timu_res query_collect_timu(1:query_collect_timu_req req),
    // 纠错试题
    modify_question_res modify_question(1:modify_question_req req),
    // 获取纠错列表
    modify_list_res modify_list(1:modify_list_req req),
    // 获取纠错数量
    modify_count_res modify_count(1:modify_count_req req),
    // 添加试题
    timu_add_res question_timu_add(1:timu_add_req req),
    // 获取添加的试题
    add_list_res add_list(1:add_list_req req),
    // 获取添加的数量
    add_count_res add_count(1:add_count_req req),
    // 审核修改的试题
    approve_res approve(1:approve_req req),

    // 根据试题id请求题目
    timu_by_id_res question_timu_by_id(1:timu_by_id_req req),

    // 通过版本id获取题型列表
    tixing_by_vid_res tixing_by_vid(1:tixing_by_vid_req req),

    // 通过版本id获取来源
    source_by_vid_res source_by_vid(1:source_by_vid_req req),

    // 根据识别id获取识别记录
    ident_by_iid_res ident_by_iid(1:ident_by_iid_req req),
    // 获取识别列表
    ident_list_res ident_list(1:ident_list_req req),
    // 获取识别列表熟数量
    ident_count_res ident_count(1:ident_count_req req),
}