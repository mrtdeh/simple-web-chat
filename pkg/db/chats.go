package database

import (
	"api-channel/proto"
)

type Chat struct {
	ChatID             uint
	ChatName           string
	Type               string
	Mute               bool
	AvatarUrl          string
	LastMessage        string
	LastMessageRead    string
	UnreadMessageCount int
}

func GetChats(userId uint) (*proto.ChatsResponse, error) {

	var chats []struct {
		ChatID             uint
		ChatName           string
		Type               string
		Mute               bool
		AvatarUrl          string
		LastMessage        string
		UnreadMessageCount int
	}

	db.Table("chats").
		// فیلتر کردن چت‌ها برای کاربر خاص
		InnerJoins("JOIN chat_members cm1 ON chats.id = cm1.chat_id AND cm1.user_id = ?", userId).
		// پیدا کردن طرف دیگر چت خصوصی
		Joins("LEFT JOIN private_chats pc ON chats.id = pc.chat_id").
		Joins("LEFT JOIN users u1 ON pc.user1_id = u1.id").
		Joins("LEFT JOIN users u2 ON pc.user2_id = u2.id").
		// اتصال به جدول گروه‌ها برای چت‌های گروهی
		Joins("LEFT JOIN groups g ON chats.id = g.chat_id").

		// پیوستن به آخرین پیامی که کاربر خوانده است
		Joins("LEFT JOIN last_message_reads lmr ON lmr.chat_id = chats.id AND lmr.user_id = ?", userId).
		// Joins("LEFT JOIN LATERAL (SELECT m.content, m.created_at FROM messages m WHERE m.chat_id = chats.id ORDER BY m.created_at DESC LIMIT 1) AS last_msg ON chats.id = last_msg.chat_id").
		Joins("LEFT JOIN LATERAL (SELECT m.content, m.created_at FROM messages m WHERE m.chat_id = chats.id AND m.id = lmr.last_message_id) AS last_msg ON chats.id = last_msg.chat_id").
		// پیوستن به پیام خوانده‌شده توسط کاربر
		Joins("LEFT JOIN messages last_read_msg ON last_read_msg.id = lmr.last_readed_message_id").
		// محاسبه تعداد پیام‌های خوانده‌نشده
		Joins("LEFT JOIN LATERAL (SELECT COUNT(*) FROM messages m WHERE m.chat_id = chats.id AND m.created_at > COALESCE(last_read_msg.created_at, '1970-01-01')) AS unread_msg_count ON chats.id = unread_msg_count.chat_id").

		// فیلتر کردن چت‌ها برای کاربر خاص
		// Where("cm1.user_id = ?", userId).
		Select(`
		chats.id AS chat_id,
		CASE
			WHEN chats.is_group THEN g.name            
			WHEN pc.user1_id = ? THEN u2.username      
			WHEN pc.user2_id = ? THEN u1.username      
		END AS chat_name,
		CASE 
			WHEN chats.is_group THEN 'group'          
			ELSE 'private'                            
		END AS type,
		cm1.mute AS mute,
		CASE 
			WHEN chats.is_group THEN g.avatar          
			WHEN pc.user1_id = ? THEN u2.profile_picture 
			WHEN pc.user2_id = ? THEN u1.profile_picture 
		END AS avatar_url,
		last_msg.content AS last_message,              
		unread_msg_count.count AS unread_message_count 
		/* last_read_msg.content AS last_message_read*/   

	`, userId, userId).
		Scan(&chats)

	return &proto.ChatsResponse{}, nil

}
