namespace go douyin.extra.interaction

struct douyinFavoriteActionRequest {
    1: string token
    2: i64 videoId (api.query="video_id")
    3: i32 actionType (api.query="action_type")
}

struct douyinFavoriteActionResponse{
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
}

struct douyinFavoriteListRequest{
    1: i64 userId (api.query="user_id")
    2: string token
}

struct douyinFavoriteListResponse{
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: list<Video> videoList (api.body="video_list")
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

struct douyinCommentActionRequest{
    1: string token
    2: i64 videoId (api.query="video_id")
    3: i32 actionType (api.query="action_type")
    4: string commentText (api.query="comment_text")
    5: i64 commentId (api.query="comment_id")
}

struct douyinCommentActionResponse{
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: Comment comment
}

struct Comment {
    1: i64  id
    2: User user
    3: string content
    4: string createDate (api.body="create_date")
}

struct douyinCommentListRequest{
    1: string token
    2: i64 videoId (api.query="video_id")
}

struct douyinCommentListResponse{
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: list<Comment> commentList (api.body="comment_list")
}

service douyinInteractionService {
    douyinFavoriteActionResponse FavoriteInteraction(1: douyinFavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    douyinFavoriteListResponse FavoriteList(1: douyinFavoriteListRequest req) (api.get="/douyin/favorite/list/")
    douyinCommentActionResponse CommentInteraction(1: douyinCommentActionRequest req) (api.post="/douyin/comment/action/")
    douyinCommentListResponse CommentList(1: douyinCommentListRequest req) (api.get="/douyin/comment/list/")
}
