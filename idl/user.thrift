namespace go douyin.core

struct douyinUserRegisterRequest {
    1: string username (api.query="username");
    2: string password (api.query="password");
}
struct douyinUserRegisterResponse {
    1: i32 statusCode
    2: string statusMsg
    3: i64 userId
    4: string token
}

service douyinUserService {
    douyinUserRegisterResponse RegisterUser(1: douyinUserRegisterRequest req) (api.post="/douyin/user/register/");
}