package service

import "websocketDemo2/server/component"

type Service struct {
}

func (s *Service) Run() {
	cm := component.GetChannelManager()
	cm.Start()
}
