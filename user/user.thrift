namespace go api

include "../response.thrift"

struct UserInfoResponse {
    1: required response.Response response
    2: i64 id
    3: string name
}

service User {
    UserInfoResponse userInfo(1: i64 userId)
}