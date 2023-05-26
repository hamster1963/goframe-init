// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo/internal/dao/internal"
)

// internalChatUserDao is internal type for wrapping internal DAO implements.
type internalChatUserDao = *internal.ChatUserDao

// chatUserDao is the data access object for table chat_user.
// You can define custom methods on it to extend its functionality as you wish.
type chatUserDao struct {
	internalChatUserDao
}

var (
	// ChatUser is globally public accessible object for table chat_user operations.
	ChatUser = chatUserDao{
		internal.NewChatUserDao(),
	}
)

// Fill with you ideas below.
