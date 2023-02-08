namespace go api

enum StatusCode {
    Success=0,
    Error=1
}

struct Response {
    1: required i32 statusCode
    2: string statusMsg
}