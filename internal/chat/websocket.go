package chat

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"github.com/thanh2k4/Chat-app/internal/chat/infras/postgres"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Manager struct {
	clients     map[string]*websocket.Conn
	lock        sync.Mutex
	redisClient *redis.Client
	queries     *postgres.Queries
}

func NewManager(redisClient *redis.Client, queries *postgres.Queries) *Manager {
	return &Manager{
		clients:     make(map[string]*websocket.Conn),
		redisClient: redisClient,
	}
}

func (m *Manager) HandleConnection(c *gin.Context) {
	userID := c.Query("user_id")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	m.lock.Lock()
	m.clients[userID] = conn
	m.lock.Unlock()

	m.redisClient.Set(context.Background(), "user_status"+userID, "online", time.Minute*5)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	m.lock.Lock()
	delete(m.clients, userID)
	m.lock.Unlock()
	m.redisClient.Del(context.Background(), "user_status"+userID)
	conn.Close()
}

func (m *Manager) SendToUser(ctx context.Context, chatID pgtype.UUID, message string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	members, err := m.queries.GetMemberByChatID(ctx, chatID)
	if err != nil {
		return
	}
	for _, member := range members {
		if conn, ok := m.clients[member.String()]; ok {
			err = conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				conn.Close()
				delete(m.clients, member.String())
			}
		}
	}
}
