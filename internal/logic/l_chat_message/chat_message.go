package l_chat_message

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/model/m_chat_content"
	"demo/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

// ==========================================================================
// logic 初始化
// ==========================================================================

type sChatMessage struct {
}

func init() {
	service.RegisterChatMessage(New())
}

func New() service.IChatMessage {
	return &sChatMessage{}
}

// GetMessageById
//
//	@dc:主键查询表信息
//	@params:表主键id-in
//	@response:表原结构信息entity
//	@author:auto @date:2023/5/21 16:57:37
func (s *sChatMessage) GetMessageById(ctx context.Context, in *string) (out *entity.ChatContent, err error) {
	db := dao.ChatContent.Ctx(ctx)
	err = db.Where("id", in).Scan(&out)
	if err != nil {
		return nil, fmt.Errorf("sChatMessage GetMessageById 数据库查询错误 %v", err)
	}
	return
}

// AddMessage
//
//	@dc: 添加消息
//	@params: 聊天内容-in
//	@response: 聊天内容ID-out, 错误信息-err
//	@author:laixin @date:2023/5/21 17:00:12
func (s *sChatMessage) AddMessage(ctx context.Context, in *m_chat_content.AddChatContent) (out *string, err error) {
	db := dao.ChatContent.Ctx(ctx)
	data, err := db.OmitEmpty().InsertAndGetId(in)
	if err != nil {
		return nil, fmt.Errorf("sChatMessage AddMessage 数据库新增错误 %v", err)
	} else {
		out = gconv.PtrString(data)
	}
	return
}

// ReadChatContent
//
//	@dc: 消息已读方法
//	@params:
//	@response:
//	@author:laixin @date:2023/5/24 09:06:16
func (s *sChatMessage) ReadChatContent(ctx context.Context, in *m_chat_content.ReadChatContent) (out *string, err error) {
	db := dao.ChatContent.Ctx(ctx)
	_, err = db.OmitEmpty().Where("id", in.Id).Update(in)
	if err != nil {
		return nil, fmt.Errorf("sChatMessage ReadChatContent 数据库更新错误 %v", err)
	} else {
		out = gconv.PtrString(in.Id)
	}
	return
}
