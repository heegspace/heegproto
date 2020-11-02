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