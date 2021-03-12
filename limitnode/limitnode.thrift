namespace go limitnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct modify_timu_limit_req {
    1:common.authorize     auth,
    2:i64                  uid,
    3:map<string,string>   extra,
}

struct modify_timu_limit_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

struct add_timu_limit_req {
    1:common.authorize     auth,
    2:i64                  uid,
    3:map<string,string>   extra,
}

struct add_timu_limit_res {
    1:rescode.code      rescode,
    2:string            resmsg,
    3:map<string,string> extra,
}

service limitnode_service {
    // 检查修改试题是否受限
    modify_timu_limit_res modify_timu_limit(1:modify_timu_limit_req req),
    // 检查添加试题是否受限
    add_timu_limit_res add_timu_limit(1:add_timu_limit_req req),
}