package robot_manager

import (
	"mahjong-table-go/msg"
	"testing"
)

func TestNewRobotManager(t *testing.T) {

	MsgIn := make(chan msg.Msg)
	MsgOut := make(chan msg.Msg)

	m := NewRobotManager(MsgIn, MsgOut)

	m.Run()

}
