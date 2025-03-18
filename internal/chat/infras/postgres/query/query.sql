-- name: CreateChat :one
INSERT INTO chat.chats ( type, name, creator_id)
VALUES ($1, $2, $3, )
    RETURNING *;

-- name: GetChatByID :one
SELECT * FROM chat.chats WHERE id = $1;

-- name: AddChatMember :one
INSERT INTO chat.chat_members (id, chat_id, user_id, role, status)
VALUES ($1, $2, $3, $4, 'active')
    ON CONFLICT (chat_id, user_id) DO UPDATE SET status = 'active'
                                          RETURNING *;

-- name: RemoveChatMember :exec
UPDATE chat.chat_members SET status = 'left' WHERE chat_id = $1 AND user_id = $2;

-- name: GetChatMembers :many
SELECT * FROM chat.chat_members WHERE chat_id = $1;

-- name: SendMessage :one
INSERT INTO chat.messages (id, chat_id, sender_id, content, type, media_url, status)
VALUES ($1, $2, $3, $4, $5, $6, 'sent')
    RETURNING *;

-- name: GetMessagesByChatID :many
SELECT * FROM chat.messages WHERE chat_id = $1 ORDER BY created_at ASC;

-- name: UpdateMessageStatus :exec
UPDATE chat.messages SET status = $2 WHERE id = $1;

-- name: MarkMessageAsRead :exec
INSERT INTO chat.message_status (message_id, user_id, status, read_at)
VALUES ($1, $2, 'read', NOW())
    ON CONFLICT (message_id, user_id) DO UPDATE SET status = 'read', read_at = NOW();

-- name: GetUnreadMessages :many
SELECT * FROM chat.messages WHERE chat_id = $1 AND id NOT IN (
    SELECT message_id FROM chat.message_status WHERE user_id = $2 AND status = 'read'
) ORDER BY created_at ASC;

-- name: GetMemberByChatID :many
SELECT user_id FROM chat.chat_members WHERE chat_id = $1 ;