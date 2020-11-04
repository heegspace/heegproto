namespace go common

struct authorize {
    1:string    key,
    2:string    value,
    3:map<string,string> extra,
}

struct question_query {
    1:string        roll_id,
    2:string        roll_name,
    3:string        grade_id,
    4:string        grade_name,
    5:string        subject_id,
    6:string        subject_name,
    7:string        version_id,
    8:string        version_name,
    9:string        chapter_id,
    10:string       chapter_name,
    11:string       ti_xing_id,
    12:string       ti_xing_name,
    13:string       source_id,
    14:string       source_name,
    15:i32          page,
    16:i32          size, 
    17:map<string,string> extra,
}

struct version_item {
    1:string    uid,
    2:string    subject_id,
    3:string    name,
    4:string    sign,
    5:string    create_at,
}


struct question_option {
    1:string        name,
    2:list<string>  images,
}

struct question {
    1:string                timu,
    2:list<string>          images,
    3:list<question_option> options,
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
    2:string    userid,
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