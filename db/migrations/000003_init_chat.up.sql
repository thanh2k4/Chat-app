START TRANSACTION;

CREATE SCHEMA IF NOT EXISTS chat;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE chat.chats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    type TEXT NOT NULL CHECK (type IN ('private', 'group')),
    name TEXT NULL,
    creator_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE chat.chat_members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    chat_id UUID NOT NULL REFERENCES chat.chats(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    role TEXT NOT NULL CHECK (role IN ('member', 'admin', 'owner')),
    status TEXT NOT NULL CHECK (status IN ('active', 'left', 'banned')),
    joined_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_chat_member UNIQUE (chat_id, user_id)
);

CREATE TABLE chat.messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    chat_id UUID NOT NULL REFERENCES chat.chats(id) ON DELETE CASCADE,
    sender_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    content TEXT NULL,
    type TEXT NOT NULL CHECK (type IN ('text', 'image', 'video', 'file', 'sticker')),
    media_url TEXT NULL,
    status TEXT NOT NULL CHECK (status IN ('sent', 'delivered', 'read')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE chat.message_status (
    message_id UUID NOT NULL REFERENCES chat.messages(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    status TEXT NOT NULL CHECK (status IN ('delivered', 'read')),
    read_at TIMESTAMP NULL,
    PRIMARY KEY (message_id, user_id)
);

COMMIT;
