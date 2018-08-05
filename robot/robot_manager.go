package robot

import (
	"github.com/mingz2013/mahjong-table-go/base"
	"github.com/mingz2013/mahjong-table-go/msg"
	"sync"
	"time"
)

var wg sync.WaitGroup

type RobotSession struct {
	Robot  Robot
	MsgIn  chan msg.Msg
	MsgOut chan msg.Msg
}

type RobotManager struct {
	Sessions []RobotSession

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg
}

func (r *RobotManager) Init() {

}

func NewRobotManager(msgIn chan msg.Msg, msgOut chan msg.Msg) RobotManager {
	r := RobotManager{MsgIn: msgIn, MsgOut: msgOut}
	r.Init()
	return r
}

func (r RobotManager) Run() {
	for {
		select {
		case m, ok := <-r.MsgIn:
			{
				if !ok {
					continue
				}
				r.onMsg(m)
			}
		case <-time.After(time.Second * 1):
			continue

		}
	}

	wg.Wait()
}

func (r *RobotManager) onMsg(m msg.Msg) {

}

func (r *RobotManager) NewRobot() {
	userId, ok := r.findOneUserId()
	if !ok {
		return
	}

	robot := NewRobot(userId)
	session := RobotSession{}
	session.Robot = robot
	r.Sessions = append(r.Sessions, session)

	base.RunProcessor(wg, robot)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case m, ok := <-session.MsgOut:
				{
					if !ok {
						continue
					}
					r.MsgOut <- m
				}
			case <-time.After(time.Second * 1):
				continue

			}
		}
	}()

}

func (r *RobotManager) findOneUserId() (int, bool) {
	return 1, true
}
