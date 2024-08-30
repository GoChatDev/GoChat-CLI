package client

import (
    "log"
    "golang.org/x/net/websocket"
)

type WebSocketClient struct {
    UserID string
    Conn   *websocket.Conn
}

func NewWebSocketClient(userID string) *WebSocketClient {
    return &WebSocketClient{
        UserID: userID,
    }
}

func (c *WebSocketClient) Connect() error {
    wsURL := "ws://localhost:8080/ws" // Update with your backend WebSocket URL
    var err error
    c.Conn, err = websocket.Dial(wsURL, "", "http://localhost/")
    if err != nil {
        return err
    }
    log.Println("Connected to WebSocket server")
    return nil
}

func (c *WebSocketClient) SendMessage(message string) error {
    if c.Conn == nil {
        return nil
    }
    return websocket.Message.Send(c.Conn, message)
}

func (c *WebSocketClient) ReceiveMessage() (string, error) {
    var msg string
    if err := websocket.Message.Receive(c.Conn, &msg); err != nil {
        return "", err
    }
    return msg, nil
}
