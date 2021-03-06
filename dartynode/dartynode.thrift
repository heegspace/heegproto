namespace go dartynode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct login_wechat_req {
    1:string                appid,
    2:string                Code,
    3:string                State,
    4:string                style, // normal: 普通用户   cooper: 合作者   school: 学校
    5:map<string,string>   extra,
}

struct login_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct refresh_wechat_req {
    1:i64                   uid,
    2:string                appid,
    3:string                refresh_token,
    4:string                access_token,
    5:map<string,string>    extra,
}

struct refresh_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct logout_wechat_req {
    1:i64                   uid,
    2:string                appid,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct logout_wechat_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct userinfo_wechat_req {
    1:i64                   uid,
    2:string                openid,
    3:string                appid,
    4:string                access_token,
    5:string                refresh_token,  // 用于刷新access_token
    6:map<string,string>    extra,
}

struct userinfo_wechat_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:common.wechat_userinfo    userinfo,
    4:string                    cookie, // 是否更新cookie
    5:map<string,string>        extra,
}

struct login_alipay_req {
    1:string                app_id,
    2:string                source,
    3:string                scope,
    4:string                auth_code,
    5:string                style,  // normal: 普通用户   cooper: 合作者   school: 学校
    6:map<string,string>   extra,
}

struct login_alipay_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct userinfo_alipay_req {
    1:i64                   uid,
    2:string                openid,
    3:string                appid,
    4:string                access_token,
    5:string                refresh_token,  // 用于刷新access_token
    6:map<string,string>    extra,
}

struct userinfo_alipay_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:common.alipay_userinfo    userinfo,
    4:string                    cookie, // 是否更新cookie
    5:map<string,string>        extra,
}

struct refresh_alipay_req {
    1:i64                   uid,
    2:string                appid,
    3:string                refresh_token,
    4:string                access_token,
    5:map<string,string>    extra,
}

struct refresh_alipay_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct logout_alipay_req {
    1:i64                   uid,
    2:string                appid,
    3:string                cookie,
    4:map<string,string>    extra,
}

struct logout_alipay_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct baidu_entity_req {
    1:i32                   page,
    2:i32                   size,
    3:string                statement,
    4:map<string,string>    extra,
}

struct baidu_entity_res {
    1:rescode.code              rescode,
    2:string                    resmsg,
    3:i32                       page,
    4:i32                       size,
    5:list<common.baidu_entity> entitys,
    6:map<string,string>        extra,
}

struct baidu_doc_analysis_req {
    1:i64                   uid,
    2:string                image_base64,
    3:map<string,string>    extra,
}

struct baidu_doc_analysis_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:common.baidu_doc_analysis         analysis,
    4:map<string,string>                extra,
}

struct baidu_idcard_ident_req {
    1:i64                   uid,
    2:string                idcard_side,
    3:string                image_base64,
    4:map<string,string>    extra,
}

struct baidu_idcard_ident_res {
    1:rescode.code                      rescode,
    2:string                            resmsg,
    3:common.baidu_idcard_ident         ident,
    4:map<string,string>                extra,
}

service dartynode_service {
    // 登录微信
    login_wechat_res login_wechat(1:login_wechat_req req),

    // 刷新微信的token
    refresh_wechat_res refresh_wechat(1:refresh_wechat_req req),

    // 退出微信登录
    logout_wechat_res logout_wechat(1:logout_wechat_req req),

    // 获取用户信息
    userinfo_wechat_res userinfo_wechat(1:userinfo_wechat_req req),

    // 支付宝登陆
    login_alipay_res login_alipay(1:login_alipay_req req),

    // 刷新alipay的token
    refresh_alipay_res refresh_alipay(1:refresh_alipay_req req),

    // 退出alipay登录
    logout_alipay_res logout_alipay(1:logout_alipay_req req),

    // 获取用户信息
    userinfo_alipay_res userinfo_alipay(1:userinfo_alipay_req req),

    // 百度实体
    baidu_entity_res baidu_entity(1:baidu_entity_req req),

    // 试卷识别
    baidu_doc_analysis_res baidu_doc_analysis(1:baidu_doc_analysis_req req),

    // 身份证识别
    baidu_idcard_ident_res baidu_idcard_ident(1:baidu_idcard_ident_req req),
}