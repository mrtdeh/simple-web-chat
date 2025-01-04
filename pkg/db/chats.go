package database

import (
	"database/sql"
)

func (db *ChatDatabase) GetChats(userId uint32) (*sql.Rows, error) {

	rows, err := db.gormDB.Table("chats").
		// Get Chats of user with details from chat_member table
		InnerJoins("JOIN chat_members cm1 ON chats.id = cm1.chat_id AND cm1.user_id = ?", userId).
		// Find the left user of private chat if it exists
		Joins("LEFT JOIN chat_members cm2 ON chats.id = cm2.chat_id AND chats.type = 'personal' AND cm2.user_id <> ?", userId).
		// Find left user information from users table
		Joins("LEFT JOIN users leftUser ON cm2.user_id = leftUser.id").
		// Find the group chat if it exists
		Joins("LEFT JOIN groups g ON chats.id = g.chat_id").
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
		chats.type AS type,
		CASE
			WHEN chats.type <> 'personal' THEN g.name            
			WHEN leftUser.id > 0 THEN leftUser.username      
		END AS chat_title,
		cm1.mute AS mute,
		CASE 
			WHEN chats.type <> 'personal' THEN g.avatar_url  
			WHEN leftUser.id > 0 THEN leftUser.profile_picture        
		END AS avatar_url,
		latest_message.content AS last_message,
		unread_msg_count.count AS unreaded_messages_count,
		cm1.last_readed_message_id AS last_readed_message_id
		
	`).Rows()
	if err != nil {
		return nil, err
	}

	return rows, nil

}
