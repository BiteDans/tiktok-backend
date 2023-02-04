namespace go douyin.extra.message

struct douyinMessageChatRequest {
    1: string token
    2: i64 toUserId (api.query="to_user_id")
}

struct douyinMessageChatResponse {
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
    3: list<Message> messageList (api.body="message_list")
}

struct Message {
    1: i64 id
    2: i64 toUserId (api.body="to_user_id")
    3: i64 fromUserId (api.body="from_user_id")
    4: string content
    5: string createTime (api.body="create_time")
}

struct douyinRelationActionRequest {
    1: string token
    2: i64 toUserId (api.query="to_user_id")
    3: i32 actionType (api.query="action_type")
    4: string content
}

struct douyinRelationActionResponse {
    1: i32 statusCode (api.body="status_code")
    2: string statusMsg (api.body="status_msg")
}

service douyinMessageService {
    douyinRelationActionResponse MessageSend(1: douyinRelationActionRequest req) (api.post="/douyin/message/action/")
    douyinMessageChatResponse MessageHistory(1: douyinMessageChatRequest req) (api.get="/douyin/message/chat/")
}
