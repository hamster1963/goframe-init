package m_chat_content

import "github.com/gogf/gf/v2/os/gtime"

type AddChatContent struct {
	Content   *string `json:"content"     ` // 聊天内容
	SendId    *int    `json:"send_id"     ` // 发送者ID
	ReceiveId *int    `json:"receive_id"  ` // 接收者ID
}

type ReadChatContent struct {
	Id       *string     `json:"id" `        // 聊天记录ID
	Read     *int        `json:"read" `      // 聊天记录阅读状态
	ReadTime *gtime.Time `json:"read_time" ` // 聊天记录阅读时间
}
