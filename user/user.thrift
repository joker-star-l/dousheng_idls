namespace go api

include "../response.thrift"

struct UserInfoResponse {
    1: required response.Response response
    2: i64 id
    3: string name
    4: i64 followCount
    5: i64 followerCount
    6: bool isFollow
    7: string avatar
    8: string backgroundImage
    9: string signature
    10: i64 totalFavorited
    11: i64 workCount
    12: i64 favoriteCount
}

service User {
    UserInfoResponse userInfo(1: i64 userId, 2: i64 queryId)
}