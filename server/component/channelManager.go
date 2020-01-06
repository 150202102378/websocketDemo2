package component

import "sync"

var channelManager *ChannelManager

//ChannelManager 频道管理中心
type ChannelManager struct {
	channels []Channel
}

func init() {
	channelManager = &ChannelManager{}
	for i := 0; i < 2; i++ {
		channelManager.channels = append(channelManager.channels, Channel{
			clients:      make(map[*Client]bool),
			broadcastMsg: make(chan []byte),
			register:     make(chan *Client),
			unregister:   make(chan *Client),
			lock:         &sync.RWMutex{},
		})
	}
}

func GetChannelManager() *ChannelManager {
	return channelManager
}

func (cm *ChannelManager) GetChannel(channelNum int) *Channel {
	return &cm.channels[channelNum]
}

func (cm *ChannelManager) Start() {
	for i := 0; i < 2; i++ {
		go cm.channels[i].Start()
	}
}
