package database

import (
	"api-channel/proto"
)

func GetChats(userId uint) (*proto.ChatsResponse, error) {

	var Data []*proto.ChatsResponse_ChatData

	db.Table("chats").
		// Get Chats of user with details from chat_member table
		InnerJoins("JOIN chat_members cm1 ON chats.id = cm1.chat_id AND cm1.user_id = ?", userId).
		// Find the private chat if it exists
		Joins("LEFT JOIN chat_members cm2 ON chats.id = cm2.chat_id AND chats.is_group = false AND cm2.user_id <> ?", userId).
		// Find users of private chat if it exists
		Joins("LEFT JOIN users u ON cm2.user_id = u.id").
		// Find the group chat if it exists
		Joins("LEFT JOIN groups g ON chats.id = g.chat_id").
		// Find last readed message id and last message id of chat
		// Joins("LEFT JOIN last_message_reads lmr ON lmr.chat_id = chats.id AND lmr.user_id = ?", userId).
		// Find last readed message of chats
		Joins("LEFT JOIN messages last_read_msg ON last_read_msg.id = cm1.last_readed_message_id").
		// Find last message of chats
		Joins(`
			LEFT JOIN LATERAL (
				SELECT * FROM messages 
				WHERE id = chats.last_message_id 
				LIMIT 1
			) AS latest_message ON latest_message.chat_id = chats.id
		`).
		// Find unreaded messages count of chats
		Joins(`
		    LEFT JOIN LATERAL (
		        SELECT COUNT(*) as count
		        FROM messages m
		        WHERE m.chat_id = chats.id
		         AND m.created_at > COALESCE(last_read_msg.created_at, '1970-01-01')
		    ) AS unread_msg_count ON true
		`).
		Select(`
		
		chats.id AS chat_id,
		CASE
			WHEN chats.is_group THEN g.name            
			WHEN u.id > 0 THEN u.username      
		END AS chat_title,
		CASE 
			WHEN chats.is_group THEN 'group'          
			ELSE 'private'                            
		END AS type,
		cm1.mute AS mute,
		CASE 
			WHEN chats.is_group THEN g.avatar_url  
			WHEN u.id > 0 THEN u.profile_picture        
		END AS avatar_url,
		latest_message.content AS last_message,
		unread_msg_count.count AS unreaded_messages_count
		
	`).
		Scan(&Data)

	return &proto.ChatsResponse{Data: Data}, nil

}
