namespace go douyin.core.user

struct User {
    1: i64 id
    2: string name
    3: i64 followCount (api.body="follow_count")
    4: i64 followerCount (api.body="follower_count")
    5: bool isFollow (api.body="is_follow")
    6: string avatar 
    7: string backgroundImage (api.body="background_image")
    8: string signature
    9: i64 totalFavorited (api.body="total_favorited")
    10: i64 workCount (api.body="work_count")
    11: i64 favoriteCount (api.body="favorite_count")
}

struct douyinUserRegisterRequest {
    1: string username (api.query="username")
    2: string password (api.query="password")
}

struct douyinUserRegisterResponse {
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: i64 userId (api.body="user_id")
    4: string token
}

struct douyinUserRequest {
    1: i64 userId (api.query="user_id")
    2: string token (api.query="token")
}

struct douyinUserResponse {
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: User user
}

struct douyinUserLoginRequest {
    1: string username (api.query="username")
    2: string password (api.query="password")
}

struct douyinUserLoginResponse {
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: i64 userId (api.body="user_id")
    4: string token (api.body="token")
}

service douyinUserService {
    douyinUserResponse UserInfo(1: douyinUserRequest req) (api.get="/douyin/user/")
    douyinUserRegisterResponse UserRegister(1: douyinUserRegisterRequest req) (api.post="/douyin/user/register/")
    douyinUserLoginResponse UserLogin(1: douyinUserLoginRequest req) (api.post="/douyin/user/login/")
}
