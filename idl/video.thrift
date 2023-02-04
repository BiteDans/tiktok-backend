namespace go douyin.core.video

struct douyinVideoFeedRequest {
    1: i64  latestTime (api.body="latest_time")
    2: string  token
}

struct douyinVideoFeedResponse {
    1: i32  statusCode (api.body="status_code")
    2: string  statusMsg (api.body="status_msg")
    3: list<Video>  videoList (api.body="video_list")
    4: i64 nextTime (api.body="next_time")
}

struct douyinVideoPublishRequest {
    1: string  token
    2: string  title
}

struct douyinVideoPublishResponse {
    1: i32  statusCode (api.body="status_code")
    2: string  statusMsg (api.body="status_msg")
}

struct douyinVideoPublishListRequest {
    1: i64  userId (api.body="user_id")
    2: string  token
}

struct douyinVideoPublishListResponse {
    1: i32  statusCode (api.body="status_code")
    2: string  statusMsg (api.body="status_msg")
    3: list<Video>  videoList (api.body="video_list")
}

struct Video {
    1: i64  id
    2: User  author
    3: string  playUrl (api.body="play_url")
    4: string  coverUrl (api.body="cover_url")
    5: i64  favoriteCount (api.body="favorite_count")
    6: i64  commentCount (api.body="comment_count")
    7: bool  isFavorite (api.body="is_favorite")
    8: string  title
}

struct User {
    1: i64  id
    2: string  name
    3: i64  followCount (api.body="follow_count")
    4: i64  followerCount (api.body="follower_count")
    5: bool  isFollow (api.body="is_follow")
}

service douyinVideoService {
    douyinVideoFeedResponse VideoFeed(1: douyinVideoFeedRequest req) (api.get="/douyin/feed")
    douyinVideoPublishResponse VideoPublish(1: douyinVideoPublishRequest req) (api.post="/douyin/publish/action/")
    douyinVideoPublishListResponse VideoPublishList(1: douyinVideoPublishListRequest req) (api.get="/douyin/publish/list/")
}
