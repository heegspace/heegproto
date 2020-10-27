namespace go cloudnode

include "../rescode/rescode.thrift"
include "../common/common.thrift"

struct dir_req {
    1:common.authorize      auth,
    2:string                uid,
    3:string                path,
    4:map<string,string>    extra,
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

struct dir_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:list<FileItem>        dirs,
    4:map<string,string>    extra,
}

struct add_dir_req {
    1:common.authorize      auth,
    2:string                uid,
    3:string                parent,
    4:string                name,
    5:map<string,string>    extra, 
}

struct add_dir_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct file_item {
    1:i64    size,
    2:string path,
    3:string name,
    4:string url,
    5:string thumb,
    6:string content_type,
}

struct add_file_req {
    1:common.authorize      auth,
    2:string                uid,
    3:file_item             files,
    4:map<string,string>    extra,
}

struct add_file_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:map<string,string>    extra,
}

struct attr_req {
    1:common.authorize      auth,
    2:string                uid,
    3:map<string,string>    extra,
}

struct attr_item {
    1:i64     capacity,
    2:i32     number,
}

struct attr_res {
    1:rescode.code          rescode,
    2:string                resmsg,
    3:attr_item             attr,
    4:map<string,string>    extra,
}

service cloudnode_service {
    // 获取目录
    dir_res   dir(1:dir_req req), 

    // 添加目录
    add_dir_res add_dir(1:add_dir_req req),

    // 添加文件
    add_file_res add_file(1:add_file_req req),

    // 云存储属性
    attr_res attr(1:attr_req req),
}