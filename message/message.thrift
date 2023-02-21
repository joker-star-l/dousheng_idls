namespace go api

include "../response.thrift"

struct LatestMessageResponse {
    1: required response.Response response
    2: string message
    3: i64 from
    4: i64 to
}

service Message {
    LatestMessageResponse latestMessage(1: i64 userId, 2: i64 friendId)
}