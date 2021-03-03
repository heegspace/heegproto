namespace go datanode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct user{
    1:i64       uid,
    2:string   phone,
    3:string   email,
    4:string    account,
    5:string    pass_wd,
    6:i16      status,
    7:i64      role,
    8:i64      vip,
    9:double   coin,
    10:i64      reg_ip,
    11:string  reg_at,
}

struct user_info_req {
    1:i64    uid,
    2:map<string,string> extra,
}

struct new_user_req {
    1:string    account,
    2:string    pass_wd,
    3:i32       role,     
    4:string    invitor,
    5:string    source,
    6:string    contactor,
    7:string    email,
    8:i64       reg_ip;
    9:i64       status;
    10:map<string,string> extra,
}

struct search_user_req {
    1:list<i64>         uids,
    2:string            account,
    3:string            phone,
    4:string            email,
    5:map<string,string> extra,
}

struct user_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:user              user,
    4:map<string,string> extra,
}

struct update_req {
    1:i64        uid,
    2:string   phone,
    3:string   email,
    4:string    account,
    5:string    pass_wd,
    6:i16      status,
    7:i64      role,
    8:i64      vip,
    9:double   coin
    23:map<string,string> extra,
}

struct add_friend_req {
    1:string   phone,
    2:map<string,string> extra,
}

struct add_friend_res {
    1:rescode.code    rescode,
    2:string          resmsg,
    3:map<string,string> extra,
}

struct agree_friend_req {
    1:i64               uid,
    2:map<string,string> extra,
}

struct agree_friend_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
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

struct friend_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:friend                friends,
    4:map<string,string>    extra,
}

struct create_group_req {
    1:string    name,
    2:map<string,string> extra,
}

struct create_group_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct rename_group_req {
    1:string    old_name,
    2:string    new_name,
    3:map<string,string> extra,
}

struct rename_group_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct add_friend_note_req {
    1:i64                uid,
    2:string            note,
    3:map<string,string> extra,
}

struct add_friend_note_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct move_group_req {
    1:i64                uid,
    2:string            group,
    3:map<string,string> extra,
}

struct move_group_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct remove_friend_req {
    1:i64               uid,
    2:map<string,string> extra,
}

struct remove_friend_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct likes_count_res {
    1:rescode.code      rescode,
    2:i32               count,
    3:string            resmsg,
    4:map<string,string> extra,
}

struct likes_add_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct likes {
    1:i64               uid,
    2:i64               create_at,
}

struct likes_list_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<likes>       data,
    4:map<string,string> extra,
}

struct note_meta {
    1:i64       uid,
    2:i64   user_id,
    3:string title,
    4:string thumb,
    5:string desc,
    6:string sign,
    7:string create_at,
    8:string update_at,
}

struct update_note_req {
    1:i64       uid,
    2:string    data,
    3:string    html,
    4:map<string,string> extra,
}

struct note_meta_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:note_meta         meta,
    4:map<string,string> extra,
}

struct note_meta_list_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<note_meta>      meta,
    4:map<string,string> extra,
}

struct note_list_count_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:i32               count,
    4:map<string,string> extra,
}

struct note_data_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:string            data,
    4:map<string,string> extra,
}

struct note_html_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:string            data,
    4:map<string,string> extra,
}

struct moments_count_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:i32               count,
    4:map<string,string> extra,
}

struct extra {
    1:string path,
    2:i32 type,
}

struct moments {
    1:i64       id,
    2:string    mid,
    3:string    text,
    4:list<extra> extra,
    5:i32 create_at,
}

struct add_moments_res {
    1:rescode.code      rescode,
    2:string            resmsg,
}

struct moments_list_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<moments>     data,
}

struct FileItem {
    1:string    name,
    2:i32       type,
    3:string    path,
    4:i64       size,
    5:string    url,
    6:i64       create_at,
    7:i64       update_at,
    8:string    extra,
}

struct file_dir_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:list<FileItem>     dirs,
}

struct add_dir_req {
    1:string    parent,
    2:string    name,
}

struct add_dir_res {
    1:rescode.code  rescode,
    2:string        resmsg,
}

struct add_file_req {
    1:i64    size,
    2:string path,
    3:string name,
    4:string url,
    5:string thumb,
    6:string content_type,
}

struct add_file_res {
    1:rescode.code  rescode,
    2:string        resmsg,
}

struct file_attr_res {
    1:rescode.code  rescode,
    2:string        resmsg,
    3:i64           capacity,
    4:i32           number,
}

struct thumbnail_res {
    1:rescode.code  rescode,
    2:string        resmsg,
    3:string        thumb,
}

struct set_keyvalue_req {
    1:string        key,
    2:string        value,
    3:i64           expire,
}

struct set_keyvalue_res {
    1:rescode.code  rescode,
    2:string        resmsg,
}

struct get_keyvalue_req {
    1:string        key,
}

struct get_keyvalue_res {
    1:rescode.code  rescode,
    2:string        resmsg,
    3:string        value,
}

struct del_keyvalue_req {
    1:string        key,
}

struct del_keyvalue_res {
    1:rescode.code  rescode,
    2:string        resmsg,
}

struct version_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:list<common.version_item> versions,
    4:map<string,string> extra,
}

struct subject_name_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:string             name,
    4:map<string,string> extra,
}

struct chapter_res_item {
    1:string                    uid,
    2:string                    title,
    3:list<common.chapter_item> childs,
}

struct chapter_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<chapter_res_item> chapters,
    4:map<string,string> extra,
}

struct source_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:list<common.source_item>  sources,
    4:map<string,string> extra,
}

struct tixing_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:list<common.tixing_item>  tixings,
    4:map<string,string> extra,
}

struct timu_count_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:i32                count,
    4:map<string,string> extra,
}

struct timu_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:list<common.timu_item>    timus,
    4:map<string,string> extra,
}

struct timu_by_id_req {
    1:i64                   uid,    // 用户id
    2:string                tid,    // 题目id
    3:map<string,string>    extra,
}

struct timu_by_id_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:common.timu_item   timus,
    4:map<string,string> extra,
}

struct collect_req {
    1:string    cid,
    2:i64       op,
}

struct collect_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:map<string,string> extra,
}

struct collect_list_req {
    1:i64       uid,
    2:i32       page,
    3:i32       size,
}

struct collect_list_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:list<common.timu_item>    timus,
    4:map<string,string> extra,
}

struct modify_req {
    1:i64               uid,
    2:string            tid,
    3:common.question   timu,
}

struct modify_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:map<string,string> extra, 
}

struct modify_list_req {
    1:i64       uid,
    2:i32       page,
    3:i32       size,
    4:string    sorted,
    5:string    status,
}

struct modify_list_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:list<common.modify_item>    timus,
    4:map<string,string> extra,
}

struct modify_count_req {
    1:i64       uid,
    2:string    status,
    3:map<string,string> extra, 
}

struct modify_count_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:i32                count,
    4:map<string,string> extra, 
}

struct approve_modify_req {
    1:i64                   id,
    2:i64                   uid,
    3:string                status,
    4:string                info,
    5:map<string,string>    extra, 
}

struct approve_modify_res {
    1:rescode.code       rescode,
    2:string             resmsg,
    3:map<string,string> extra, 
}

struct search_req {
    1:string        keyword,
    2:i64           uid,
    3:common.search_tyle   style,
    4:i32           page,
    5:i32           size,
}

struct search_res {
    1:rescode.code                  rescode,
    2:string                        resmsg,
    3:double                        timestamp,
    4:common.search_hits_total         total,
    5:list<common.search_hits_item>    hits,
    6:map<string,string> extra, 
}

struct add_grade_cate_res {
    1:rescode.code      rescode,
    2:string            resmsg,
}

struct grade_cate_count_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:i64               count,
    4:map<string,string> extra,
}

struct grade_cate_list_req {
    1:i32       page,
    2:i32       size,
    3:map<string,string> extra,
}

struct grade_cate_list_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.grade_cate>   lists,
    4:map<string,string>        extra,
}

struct home_black_data_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.school_static> lists,
    4:map<string,string>        extra,
}

struct grade_subject_req {
    1:i32               index,
    2:map<string,string> extra,
}

struct grade_subject_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.grade_subject> lists,
    4:map<string,string>        extra,
}

struct school_roll_add_res {
    1:rescode.code      rescode,
    2:string            resmsg,
}

struct school_roll_count_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:i64               count,
}

struct school_roll_list_req {
    1:i32   page,
    2:i32   size,
    3:map<string,string> extra,
}

struct school_roll_list_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.school_roll>  lists,
    4:map<string,string>        extra,
}

struct subject_cate_add_res {
    1:rescode.code      rescode,
    2:string            resmsg,
}

struct subject_cate_count_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:i64               count,
}

struct subject_cate_list_req {
    1:i32   page,
    2:i32   size,
    3:map<string,string> extra,
}

struct subject_cate_list_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<common.subject_cate>  lists,
    4:map<string,string>        extra,
}

struct attention_item {
    1:i64       id,
    2:string    name,
    3:string    desc,
    4:i64       create_at,
}

struct add_attention_req {
    1:attention_item        item,
    2:map<string,string>    extra,
}

struct add_attention_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct del_attention_req {
    1:i64                   aid,
    2:map<string,string>    extra,
}

struct del_attention_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct get_attention_req {
    1:i64           aid,
    2:i32           page,
    3:i32           size,
    4:map<string,string> extra,
}

struct get_attention_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<attention_item>      lists,
    4:map<string,string>        extra,
}

struct add_user_attention_req {
    1:i64           aid,
    2:i64           uid,
    3:map<string,string> extra,
}

struct add_user_attention_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct get_user_attention_req {
    1:i64                   uid,
    2:map<string,string>    extra,
}

struct get_user_attention_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:list<attention_item>      lists,
    4:map<string,string>        extra,
}

struct del_user_attention_req {
    1:i64           aid,
    2:i64           uid,
    3:map<string,string> extra,
}

struct del_user_attention_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct set_user_info_req {
    1:i64                       uid,
    2:common.user_info          info,
    3:map<string,string>        extra,
}

struct set_user_info_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct get_user_info_req {
    1:i64                       uid,
    2:map<string,string>        extra,
}

struct get_user_info_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:common.user_info          info,
    4:map<string,string>        extra,
}

struct set_wechat_uinfo_req {
    1:i64                       uid,
    2:common.wechat_userinfo    info,
    3:map<string,string>        extra,
}

struct set_wechat_uinfo_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct get_wechat_uinfo_req {
    1:i64                       uid,
    2:map<string,string>        extra,
}

struct get_wechat_uinfo_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:common.wechat_userinfo    info,
    4:map<string,string>        extra,
}

struct set_alipay_uinfo_req {
    1:i64                       uid,
    2:common.alipay_userinfo    info,
    3:map<string,string>        extra,
}

struct set_alipay_uinfo_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct get_alipay_uinfo_req {
    1:i64                       uid,
    2:map<string,string>        extra,
}

struct get_alipay_uinfo_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:common.alipay_userinfo    info,
    4:map<string,string>        extra,
}

struct set_baidu_entity_req {
    1:string                    bid,
    2:list<common.baidu_entity> entitys,
    3:map<string,string>        extra,
}

struct set_baidu_entity_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:map<string,string>        extra,
}

struct baidu_entity_req {
    1:string                    bid,
    2:i64                       page,
    3:i64                       size,
    4:map<string,string>        extra,
}

struct baidu_entity_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:i64                       page,
    4:i64                       size,
    5:list<common.baidu_entity> entitys,
    6:map<string,string>        extra,
}

struct search_item_req {
    1:string            keyword,
    2:i32               page,
    3:i32                size,
    4:map<string,string> extra,
}

struct search_item_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:list<string>                      lists,
    4:map<string,string>                extra,
}

struct search_history_req {
    1:i64           uid,
    2:i32           page,
    3:i32           size,
    4:map<string,string> extra,
}

struct search_history_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:list<common.search_history_item>  lists,
    4:map<string,string>                extra,
}

struct refresh_modify_reward_req {
    1:string                auth,
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
    1:string                auth,
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

service datanode_service {
    // ---------- 用户接口 ------- //
    // 创建新用户
    user_res createUser(1:new_user_req req),
    // 获取用户信息
    user_res userInfo(1:user_info_req req),
    // 搜索用户信息
    list<user> searchUser(1:search_user_req req),
    // 更新用户信息
    user_res updateUser(1:update_req req),

    // 关注对象操作
    add_attention_res add_attention(1:add_attention_req req),
    del_attention_res del_attention(1:del_attention_req req),
    get_attention_res get_attention(1:get_attention_req req),

    // 用户
    add_user_attention_res add_user_attention(1:add_user_attention_req req),
    del_user_attention_res del_user_attention(1:del_user_attention_req req),
    get_user_attention_res get_user_attention(1:get_user_attention_req req),


    // --------------- 好友接口 ----------- //
    // 添加好友
    add_friend_res addFriends(1:i64 uid,2:add_friend_req req),
    // 同意好友
    agree_friend_res agreeFriends(1:i64 uid,2:agree_friend_req req),
    //  请求好友列表
    friend_res friendsList(1:i64 uid, 2:i64 uid2),
    // 添加组
    create_group_res createGroup(1:i64 uid,2:create_group_req req),
    // 重命名组
    rename_group_res renameGroup(1:i64 uid,2:rename_group_req req),
    // 添加好友备注
    add_friend_note_res addNoteFriend(1:i64 uid,2:add_friend_note_req req),
    // 移动到新的组
    move_group_res moveToNewGroup(1:i64 uid,2:move_group_req req),
    // 删除好友
    remove_friend_res removeFriend(1:i64 uid,2:remove_friend_req req),

    // ----------------- 点赞 --------------- //
    // 获取点赞的数量
    likes_count_res likesCount(1:string mid),
    // 添加点赞
    likes_add_res likesAdd(1:string mid,2:i64 uid),
    // 获取点赞列表
    likes_list_res likesList(1:string mid,2:i32 page, 3:i32 size),

    // ----------------- 笔记 ----------------// 
    //更新或者添加笔记信息
    note_meta_res updateNote(1:i64 uid, 2:update_note_req req),
    // 获取笔记列表
    note_meta_list_res noteMetaList(1:i64 uid, 2:i32 page, 3:i32 size),
    // 获取用户笔记数量
    note_list_count_res noteListCount(1:i64 uid),
    // 获取笔记数据列
    note_data_res noteData(1:i64 uid, 2:i64 nid),
    // 获取笔记html数据
    note_html_res noteHtml(1:i64 uid, 2:i64 nid),

    // ---------------- 动态 --------------- //
    // 获取动态的数量
    moments_count_res momentsCount(1:i64 uid),
    // 添加动态
    add_moments_res momentsAdd(1:i64 uid, 2:moments moments),
    // 获取动态列表
    moments_list_res momentsList(1:i64 uid, 2:i32 page, 3:i32 size),

    // ---------------- 文件 ------------------//
    // 获取目录,目录之间的/替换成-
    file_dir_res fileDirOne(1:i64 uid, 2:string path),
    // 添加目录
    add_dir_res fileDirAdd(1:i64 uid, 2:add_dir_req req),
    // 添加文件
    add_file_res addFile(1:i64 uid, 2:add_file_req req),
    // 获取云盘属性，容量和文件数量
    file_attr_res yunSaveAttr(1:i64 uid),

    // 获取图像缩略图
    thumbnail_res thumbnail(1:i64 uid, 2:string path)

    // ------------- key-value ----------------//
    // 设置缓存
    set_keyvalue_res setKeyValue(1:set_keyvalue_req req),
    // 获取缓存
    get_keyvalue_res getKeyValue(1:get_keyvalue_req req),
    // 删除
    del_keyvalue_res delKeyValue(1:del_keyvalue_req req),


    // --------------- 试题相关 -------------------//
    // 获取版本
    version_res question_version(1:common.question_query req),
    // 获取章节
    chapter_res question_chapter(1:common.question_query req),
    // 科目名
    subject_name_res  subject_name(1:string req),
    // 获取来源
    source_res question_source(1:common.question_query req),
    // 获取题型
    tixing_res question_tixing(1:common.question_query req),
    // 获取题目数量
    timu_count_res question_timu_count(1:common.question_query req),
    // 请求题目
    timu_res question_timu(1:i64 uid, 2:common.question_query req),
    // 根据试题id请求题目
    timu_by_id_res question_timu_by_id(1:timu_by_id_req req),
    // 收藏试题
    collect_res collect_timu(1:i64 uid, 2:collect_req req),
    // 获取收藏列表
    collect_list_res query_collect_timu(1:collect_list_req req),
    // 纠错试题（包含更新或添加）
    modify_res modify_question(1:modify_req req),
    // 获取个人纠错的试题
    modify_list_res modify_list(1:modify_list_req req),
    // 获取个人纠错的试题的数量
    modify_count_res modify_count(1:modify_count_req req),
    // 审核修改的试题
    approve_modify_res approve_modify(1:approve_modify_req req),
    // 搜索试题
    search_res on_search(1:search_req req),

    // 添加年纪信息
    add_grade_cate_res grade_cate_add(1:common.grade_cate req),
    // 获取年纪总数
    grade_cate_count_res grade_cate_count(),
    // 获取年级信息
    grade_cate_list_res grade_cate_list(1:grade_cate_list_req req),
    // 获取主页最新动态统计信息
    home_black_data_res home_black_data(),
    // 获取学籍对应的年级+科目
    grade_subject_res grade_subject(1:grade_subject_req req),
    // 添加学级信息
    school_roll_add_res school_roll_add(1:common.school_roll req),
    // 获取年级总数
    school_roll_count_res school_roll_count(),
    // 获取年级列表信息
    school_roll_list_res school_list(1:school_roll_list_req req),
    // 添加科目信息
    subject_cate_add_res subject_cate_add(1:common.subject_cate req),
    // 获取科目总数
    subject_cate_count_res subject_cate_count(),
    // 获取科目列表
    subject_cate_list_res subject_cate_list(1:subject_cate_list_req req),

    // 用户信息
    set_user_info_res set_user_info(1:set_user_info_req req),
    get_user_info_res get_user_info(1:get_user_info_req req),

    // 微信用户信息操作
    set_wechat_uinfo_res set_wechat_uinfo(1:set_wechat_uinfo_req req),
    get_wechat_uinfo_res get_wechat_uinfo(1:get_wechat_uinfo_req req),

    // 支付宝用户信息操作
    set_alipay_uinfo_res set_alipay_uinfo(1:set_alipay_uinfo_req req),
    get_alipay_uinfo_res get_alipay_uinfo(1:get_alipay_uinfo_req req),

    // 请求百度实体信息
    set_baidu_entity_res set_baidu_entity(1:set_baidu_entity_req req),
    baidu_entity_res baidu_entity(1:baidu_entity_req req),

    // 获取搜索记录
    search_history_res search_history(1:search_history_req req),

    // 搜索关键字补全
    search_item_res search_item(1:search_item_req req),

    // 更新修改试题的奖励积分
    refresh_modify_reward_res refresh_modify_reward(1:refresh_modify_reward_req req),

    // 更新用户的coin数值
    refresh_user_coin_res refresh_user_coin(1:refresh_user_coin_req req),
}
