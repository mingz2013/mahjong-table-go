package main

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/base"
	"github.com/mingz2013/lib-go/msg"
	"github.com/mingz2013/mahjong-table-go/robot"
	"github.com/mingz2013/mahjong-table-go/table"
	"github.com/mingz2013/table-manager-server-go/server"
	"log"
	"sync"
)

//
//type RobotContext struct {
//	Robot  robot2.Robot
//	MsgIn  chan msg.Msg
//	MsgOut chan msg.Msg
//}
//
//func makeRobots() []RobotContext {
//	var robots []RobotContext
//	for i := 0; i < 4; i++ {
//		robotMsgIn := make(chan msg.Msg)
//		robotMsgOut := make(chan msg.Msg)
//
//		robot := robot2.NewRobot(i+1000, "", robotMsgIn, robotMsgOut)
//
//		robots = append(robots, RobotContext{robot, robotMsgIn, robotMsgOut})
//
//	}
//	return robots
//}

func StartLocalTest() {
	//sdk := sdk2.MakerSdk("1")

	//tableManager := table.NewTableManager("1")

	var wg sync.WaitGroup // 工作goroutine个数

	//go sdk.Run()
	//go tableManager.Run()
	//RunProcessor(wg, sdk)
	//RunProcessor(wg, tableManager)

	tableMsgIn := make(chan msg.Msg)
	tableMsgOut := make(chan msg.Msg)

	table_new := table.NewTable(1, tableMsgIn, tableMsgOut)

	robotManager := robot.NewRobotManager(tableMsgOut, tableMsgIn)
	//var robots []RobotContext

	//robots := makeRobots()
	//
	//log.Println("make obj down")
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	for {
	//
	//		select {
	//		case m, ok := <-tableMsgOut:
	//			{
	//				//m, ok := <-tableMsgOut
	//				log.Println("on msg table", m)
	//				if !ok {
	//					continue
	//				}
	//				id := m["id"].(int)
	//
	//				for i := 0; i < len(robots); i++ {
	//					if robots[i].Robot.Id == id {
	//						robots[i].MsgIn <- m
	//					}
	//				}
	//			}
	//
	//		case <-time.After(1 * time.Second):
	//			continue
	//
	//		}
	//
	//	}
	//
	//}()
	//
	//for i := 0; i < len(robots); i++ {
	//	wg.Add(1)
	//	go func(index int) {
	//		defer wg.Done()
	//		for {
	//
	//			select {
	//			case m, ok := <-robots[index].MsgOut:
	//				{
	//					log.Println("on msg robot", index, m)
	//					if !ok {
	//						continue
	//					}
	//					tableMsgIn <- m
	//				}
	//			case <-time.After(1 * time.Second):
	//				continue
	//
	//			}
	//
	//		}
	//
	//	}(i)
	//}

	log.Println("bound ch down")

	base.RunProcessor(&wg, table_new)
	base.RunProcessor(&wg, robotManager)
	//for i := 0; i < len(robots); i++ {
	//
	//	base.RunProcessor(wg, robots[i].Robot)
	//}

	log.Println("run down")

	wg.Wait()
}

func StartApp() {
	confMap := map[string]interface{}{

		"host":    "redis-mq",
		"port":    "6379",
		"db":      1,
		"channel": "connector-server",
	}
	data, _ := json.Marshal(confMap)
	a := server.NewApp(data)
	a.Start()
}

func main() {
	//StartLocalTest()
	StartApp()
}
