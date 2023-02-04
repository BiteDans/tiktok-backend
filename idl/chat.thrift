namespace go douyin.extra.message

struct douyinMessageChatRequest {
    1: string token
    2: i64 toUserId
}

struct douyinMessageChatResponse {
    1: i32 statusCode
    2: optional string statusMsg
    3: list<message> messageList
}

struct message {
    1: i64 id
    2: i64 toUserId
    3: i64 fromUserId
    4: string content
    5: optional string createTime
}

struct douyinRelationActionRequest {
    1: string token
    2: i64 toUserId
    3: i32 actionType
    4: string content
}

struct douyinRelationActionResponse {
    1: i32 statusCode
    2: optional string statusMsg
}

service douyinMessageService {
    douyinRelationActionResponse MessageSend(1: douyinRelationActionRequest req) (api.post="/douyin/message/action/")
    douyinMessageChatResponse MessageHistory(1: douyinMessageChatRequest req) (api.get="/douyin/message/chat/")
}
