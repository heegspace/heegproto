namespace go macipnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct ip_to_address_req {
    1:common.authorize      auth,
    2:string                ip,
    3:map<string,string>    extra,
}

struct address_item {
    1:string                ip,
    2:string                country,
    3:string                province,
    4:string                city,
    5:string                district,
    6:string                organization,
    7:string                isp,
    8:string                country_code,
}

struct ip_to_address_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:address_item          address,
    4:map<string,string>    extra,
}

service macipnode_service {
    // ip地址转换
    ip_to_address_res    ip_to_address(1:ip_to_address_req req),
}