namespace go douyin.extra.follow

struct User {
    1: i64 id
    2: string name
    3: i64 followCount (api.body="follow_count")
    4: i64 followerCount (api.body="follower_count")
    5: bool isFollow (api.body="is_follow")
}

struct douyinRelationActionRequest {
    1: string token (api.query="token")
    2: i64 toUserId (api.query="to_user_id")
    3: i32 actionType (api.query="action_type")
}
struct douyinRelationActionResponse {
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
}
struct douyinRelationFollowListRequest {
    1: i64 userId (api.query="user_id")
    2: string token (api.query="token")
}
struct douyinRelationFollowListResponse {
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: list<User> userList (api.body="user_list")
}
struct douyinRelationFollowerListRequest {
    1: i64 userId (api.query="user_id")
    2: string token (api.query="token")
}
struct douyinRelationFollowerListResponse {
    1: i32 statusCode (api.body="staus_code")
    2: string statusMsg (api.body="status_msg")
    3: list<User> userList (api.body="user_list")
}

service douyinFollowService {
    douyinRelationActionResponse FollowAction(1: douyinRelationActionRequest req) (api.post="/douyin/relation/action")
    douyinRelationFollowListResponse FollowList(1: douyinRelationFollowListRequest req) (api.get="/douyin/relation/follow/list/")
    douyinRelationFollowerListResponse FollowerList(1: douyinRelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/")
}