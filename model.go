package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ConnMap map[string]*websocket.Conn

var WSConnections = &WSConnectionsStruct{
	connections: make(ConnMap),
}

type WSConnectionsStruct struct {
	sync.RWMutex
	connections ConnMap
	conn        *websocket.Conn
}

func (c *WSConnectionsStruct) Add(addr string, conn *websocket.Conn) {
	c.Lock()
	defer c.Unlock()
	c.connections[addr] = conn
}

func (c *WSConnectionsStruct) Remove(addr string) {
	c.Lock()
	defer c.Unlock()
	delete(c.connections, addr)
}
func (c *WSConnectionsStruct) Get(addr string) *websocket.Conn {
	c.RLock()
	defer c.RUnlock()
	return c.connections[addr]
}

func (c *WSConnectionsStruct) GetAll() ConnMap {
	c.RLock()
	defer c.RUnlock()
	return c.connections
}
