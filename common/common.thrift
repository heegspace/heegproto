namespace go common

struct authorize {
    1:string    key,
    2:string    value,
    3:map<string,string> extra,
}


enum Role {
    NORMAL          = 0x0000,        // 普通用户
    COOPERATOR      = 0x0001,        // 企业用户
    TEACHER         = 0x0002,        // 教师用户
    
    STAFFOR         = 0x0100,       // 员工同事
    SUPEROR         = 0x0200,       // 超级用户
}

enum vip_level {
    VIP_NOT         = 0,
    VIP_1LEVEL      = 1,        // 1 month
    VIP_2LEVEL      = 2,        // 6 month
    VIP_3LEVEL      = 3,        // 1 year

    VIP_16LEVEL     = 15,       // forever
}

enum user_status {
    FROM_PLATOM      =  0x00000,    // 平台帐号
    FROM_WECHAT      =  0x00001,    // wechat帐号
    FROM_ALIPAY      =  0x00002,    // alipay帐号

    FROM_INVALID     =  0x00801,    // 无效用户
}

enum from_platom {
    FROM_LOCAL      = 0x0000,     // 自己平台
    FROM_WECHAT     = 0x0001,    // 微信平台
    FROM_ALIPAY     = 0x0002,    // 支付宝平台
}

enum search_tyle {
    FROM_QUESTION   = 0x0000,   // 来自试题
    FROM_GRPAH      = 0x0001,   // 来自图谱
    FROM_CLOUD      = 0x0002,   // 来自云盘
    FROM_PRIVATE    = 0x0003,   // 来自自己
}

enum approve_dest {
    APPROVE_TIMU_MODIFY  = 0x0000,  // 审核修改试题
    APPROVE_TIMU_ADD     = 0x0001,  // 审核试题添加
}

// 教师用户的状态
enum teacher_status {
    CREATE          =   0x0000,     // 创建
    APPROVING       =   0x0001,     // 审核中
    APPOVED         =   0x0002,     // 已经审核
    PASSED          =   0x0004,     // 审核通过
    NOPASS          =   0x0008,     // 审核未通过
    LOGOUT          =   0x0010,     // 已注销
}

enum teacher2user_status {
    FOCUS           =   0x0000,     // 关注
    CANCEL          =   0x0001,     // 取消关注
}

enum note_status {
    CREATE          =   0x0000,     // Create
    DELETE          =   0x0001,     // Delete
}

enum cert_tyoe {
    ID              =   0x0000,     // 身份证实名
    FACE            =   0x0001,     // 刷脸实名
}

enum update_userinfo_field {
    USER_EMAIL      =   0x0001,     //  更新用户邮箱
    USER_PASS       =   0x0002,     //  更新用户密码
    USER_AVATAR     =   0x0003,     //  用户头像
    USER_NICKNAME   =   0x0004,     //  用户名
    USER_SEX        =   0x0005,     //  用户性别
}

struct update_userinfo_item {
    1:i64                               uid,
    2:update_userinfo_field             field,
    3:string                            new_value,
    4:string                            old_value,
    5:map<string,string>                extra,
}

struct question_query {
    1:string        roll_id (go.tag = 'form:"roll_id" json:"roll_id"'),
    2:string        roll_name   (go.tag = 'form:"roll_name" json:"roll_name"'),
    3:string        grade_id    (go.tag = 'form:"grade_id" json:"grade_id"'),
    4:string        grade_name  (go.tag = 'form:"grade_name" json:"grade_name"'),
    5:string        subject_id  (go.tag = 'form:"subject_id" json:"subject_id"'),
    6:string        subject_name    (go.tag = 'form:"subject_name" json:"subject_name"'),
    7:string        version_id  (go.tag = 'form:"version_id" json:"version_id"'),
    8:string        version_name    (go.tag = 'form:"version_name" json:"version_name"'),
    9:string        chapter_id  (go.tag = 'form:"chapter_id" json:"chapter_id"'),
    10:string       chapter_name    (go.tag = 'form:"chapter_name" json:"chapter_name"'),
    11:string       ti_xing_id  (go.tag = 'form:"ti_xing_id" json:"ti_xing_id"'),
    12:string       ti_xing_name    (go.tag = 'form:"ti_xing_name" json:"ti_xing_name"'),
    13:string       source_id   (go.tag = 'form:"source_id" json:"source_id"'),
    14:string       source_name (go.tag = 'form:"source_name" json:"source_name"'),
    15:i32          page    (go.tag = 'form:"page" json:"page"'),
    16:i32          size    (go.tag = 'form:"size" json:"size"'), 
    17:map<string,string> extra (go.tag = 'form:"extra" json:"extra"'),
}

struct version_item {
    1:string    uid,
    2:string    subject_id,
    3:string    name,
    4:string    sign,
    5:string    create_at,
}


struct question_option {
    1:string        name (go.tag = 'form:"name" json:"name"'),
    2:list<string>  images (go.tag = 'form:"images" json:"images"'),
}

struct question {
    1:string                timu (go.tag = 'form:"timu" json:"timu"'),
    2:list<string>          images (go.tag = 'form:"images" json:"images"'),
    3:list<question_option> options (go.tag = 'form:"options" json:"options"'),
}

struct timu_item {
    1:string    uid,
    2:string    roll_id,
    3:string    roll_name,
    4:string    grade_id,
    5:string    grade_name,
    6:string    subject_id,
    7:string    subject_name,
    8:string    version_id,
    9:string    version_name,
    10:string   chapter_id,
    11:string   chapter_name,
    12:string   chapter_gd,
    13:string   source_id,
    14:string   source_name,
    15:string   ti_xing_id,
    16:string   ti_xing_name,
    17:string   sign,
    18:question data,
    19:bool     is_collect,
    20:string   create_at,
}


struct tixing_item {
    1:string    uid,
    2:string    version_id,
    3:string    chapter_id,
    4:string    name,
    5:string    sign,
    6:string    create_at,
}

struct source_item {
    1:string    uid,
    2:string    version_id,
    3:string    chapter_id,
    4:string    name,
    5:string    sign,
    6:string    create_at,
}

struct chapter_item {
    1:string    uid,
    2:string    version_id,
    3:string    name,
    4:string    chapter_gd,
    5:string    sign,
    6:string    create_at,
}

struct modify_meta {
    1:string    uid,
    2:string    roll_id,
    3:string    roll_name,
    4:string    grade_id,
    5:string    grade_name,
    6:string    subject_id,
    7:string    subject_name,
    8:string    version_id,
    9:string    version_name,
    10:string   chapter_id,
    11:string   chapter_name,
    12:string   chapter_gd,
    13:string   source_id,
    14:string   source_name,
    15:string   ti_xing_id,
    16:string   ti_xing_name,
    17:string   sign,
}

struct modify_item {
    1:i64       id,
    2:i64       userid,
    3:string    ti_mu_id,
    4:question  source,
    5:question  data,
    6:string    sign,
    7:string    status,
    8:string    msg,
    9:modify_meta meta,
    10:string   confirmer,
    11:string   create_at,
    12:string   update_at,
    13:string   username,
    14:string   confirmname,
    15:double   reward,
}

struct search_hits_total {
    1:i64       value,
    2:string    relation,
}

struct search_hits_item {
   1:string         index,
   2:string         type,
   3:string         id,
   4:double         score,
   5:timu_item      source,
   6:map<string,list<string>> high_light,
}

struct search_hits {
    1:search_hits_total         total,
    2:double                    max_score,
    3:list<search_hits_item>    hits,
}

struct search_shards {
    1:i64       total,
    2:i64       successful,
    3:i64       skipped,
    4:i64       failed,
}

struct search_topic {
    1:i32           took,
    2:bool          time_out,
    3:search_shards shards,
    4:search_hits   hits,
}

struct grade_cate {
    1:string    uid,
    2:string    roll_id,
    3:i64       sort,
    4:string    name,
    5:string    create_at,
}

struct school_static {
    1:i64           type,
    2:string        title,
    3:list<string>  content,
    4:string        href,
    5:string        target,
}

struct grade_item {
    1:string       uid,
    2:string    name,
    3:i64       sort,
    4:string    href,
    5:string    heritid,
}

struct grade_subject {
    1:string            roll_uid,
    2:string            grade_uid,
    3:string            grade_name,
    4:i64               sort,
    5:list<grade_item>  subjects,
}

struct school_roll {
    1:string     uid,
    2:i64       index,
    3:string    name,
    4:string    create_at,
}

struct subject_cate {
    1:string       uid,
    2:string    roll_id,
    3:string    grade_id,
    4:string    subject,
    5:i64       sort,
    6:string    href,
    7:string    create_at,
}

struct search_history_item {
    1:string                keyword,
    2:double                weight,
    3:string                source,
    4:map<string,string>    extra,
}

struct user_info {
    1:string                avatar,
    2:string                nickname,
    3:i32                   sex,
    4:string                province,
    5:string                city,
    6:string                country,
    7:i64                   vip,
}

struct wechat_userinfo {
    1:string                openid,
    2:string                nickname,
    3:i32                   sex,
    4:string                province,
    5:string                city,
    6:string                country,
    7:string                headimgurl,
    8:list<string>          privilege,
    9:string                unionid,
}

struct alipay_userinfo {
    1:string    code,
    2:string    msg,
    3:string    user_id,
    4:string    avatar,
    5:string    province,
    6:string    city,
    7:string    nick_name,
    8:string    gender,
}

struct baidu_concept {
    1:string    level1, // 一级概
    2:string    level2, // 二级概念
}

struct baidu_entity {
    1:string            status,             // 用于对关联结果进行标识
    2:string            confidence,         // 实体关联至该@id的置信度
    3:string            annoType,           // 标注类型：Instance | Category | Property
    4:baidu_concept    concept,            // 概念标注结果
    5:string            bdbkKgId,           // 百科newid
    6:string            mainReqRankList,    // 主需求实体，列出所有候选实体信息，根据热度从高到低排列
    7:string            bdbkUrl,            // 百科url
    8:string            offset,
    9:string            desc,               // 实体的简介
    10:string           mention,            // 实体名
}

struct add_timu_item {
    1:string    roll_id,
    2:string    roll_name,
    3:string    grade_id,
    4:string    grade_name,
    5:string    subject_id,
    6:string    subject_name,
    7:string    version_id,
    8:string    version_name,
    9:string    chapter_id,
    10:string   chapter_name,
    11:string   chapter_gd,
    12:string   source_id,
    13:string   source_name,
    14:string   ti_xing_id,
    15:string   ti_xing_name,
    16:string   sign,
    17:question data,
    18:i64      user_id,
    19:string   user_name,
    20:i64      approvor,
    21:string   approvor_name,
    22:string   create_at,
    23:string   msg,
    24:string   status,
}

struct baidu_word_location {
    1:i32   top,
    2:i32   left,
    3:i32   width,
    4:i32   height,
}

struct baidu_line_probability {

}

struct baidu_words_item {
    1:baidu_word_location     location,
    2:baidu_line_probability  probability,
    3:string            word,
}

struct baidu_result_item {
    1:string                words_type,
    2:baidu_words_item      word,
}

struct baidu_doc_analysis {
    1:i64                       log_id,
    2:i32                       img_direction,
    3:i32                       results_num,
    4:i32                       layouts_num,
    5:list<baidu_result_item>   results,
}

struct baidu_ident_item {
    1:i64                           uid,
    2:string                        iid,
    3:string                        sign,
    4:string                        status,
    5:baidu_doc_analysis            result,
    6:string                        create_at,
    7:string                        update_at,
    8:double                        reward,
    9:string                        image_url,
}

struct note_meta {
    1:i64           uid,
    2:i64           user_id,
    3:string        title,
    4:string        thumb,
    5:string        desc,
    6:string        sign,
    7:string        create_at,
    8:string        update_at,
    9:list<string>  tag,
    10:string       data1,
    11:string       data2,
    12:i64          status,
    13:string       bg_color,
    14:i64          note_type,
}

struct note_tag {
    1:i64       id,
    2:i64       user_id,
    3:string    tag,
    4:string    lang,
    5:i64       status,
    6:i64       create_at,
}

struct person_cert {
    1:i64                   uid,
    2:string                idname,
    3:string                idnum,
    4:string                image_up,
    5:string                image_down,
    6:string                image_face,
    7:i32                   cert_type,
    8:string                status,
    9:string                info,
    10:string               mark,
    11:i64                  approv_id,
    12:string               approv_name,
    13:i64                  submit_at,
    14:i64                  create_at,
    15:i64                  update_at,
    16:string               data1,
    17:string               data2,
    18:map<string,string>   extra,
}

struct baidu_idcard_ident {
    1:i32                   image_status, // 0:成功  1:失败
    2:string                idname,
    3:string                idnum,
    4:string                address,
    5:string                brithday,
    6:string                sex,
    7:string                ethnic, // 民族
    8:map<string,string>    extra,
}

// 日志类型,日志类型并不重复，所有采用对应的整型
enum log_type {
    CREATE_USER         =   0x0001, // 用户创建
    UPDATE_USER         =   0x0002, // 用户更新
    LOGIN_USER          =   0x0003, // 用户登录
    LOGOUT_USER         =   0x0004, // 用户退出
    SUBMIT_CERT         =   0x0005, // 提交实名
    CANCEL_CERT         =   0x0006, // 取消实名
    CERT_FAIL           =   0x0007, // 实名失败
    CERT_SUCCESS        =   0x0008, // 实名成功
    VIP_LEVELUP         =   0x0009, // 升级vip
    COIN_APPEND         =   0x000A, // 增加coin
    COIN_USED           =   0x000B, // 减少coin
}

// 用户日志
struct user_log {
    1:i64       id,
    2:i64       uid,
    3:i32       log_type,
    4:string    log_info,
    5:string    create_at,
}

struct cron_item {
    1:string                interval,
    2:string                start_at,
    3:bool                  repeadted,
    4:string                func,
    5:bool                  mutitask,
    6:string                info,
    7:string                version,
    8:map<string,string>    extra,    
}
