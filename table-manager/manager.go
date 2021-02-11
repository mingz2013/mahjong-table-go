package table_manager

import (
	"log"
	"mahjong-table-go/msg"
	"mahjong-table-go/table"
	"sync"
	"time"
)

type Manager struct {
	Id              string
	tableSessionMap map[int]TableSession

	userSessionMap map[int]UserSession

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg
}

func (mgr *Manager) Init() {
	mgr.tableSessionMap = map[int]TableSession{}
	mgr.userSessionMap = map[int]UserSession{}
}

func NewTableManager(id string, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Manager {
	m := Manager{Id: id}
	m.MsgIn = msgIn
	m.MsgOut = msgOut
	m.Init()
	return m
}

func (mgr Manager) Run() {

	for {
		select {
		case m, ok := <-mgr.MsgIn:
			{
				if !ok {
					continue
				}
				mgr.onMsg(m)
			}
		case <-time.After(1 * time.Second):
			continue

		}

	}

	wg.Wait()

}

func (mgr *Manager) onMsg(m msg.Msg) {

	userId := m["userId"].(int)

	switch m.GetCmd() {
	case "create":
		mgr.onCmdCreate(m, userId)
	case "delete":
		break
	case "join":
		mgr.onCmdJoin(m, userId)

	case "table":
		// 通过session找到对应的table，然后调用table的onMsg
		mgr.onCmdTable(m, userId)

	case "play":
		// 自动转到桌子里
		mgr.onCmdTable(m, userId)

	default:
		log.Println("unknown cmd", m)
	}

}

func (mgr *Manager) onCmdTable(m msg.Msg, userId int) {

	session, ok := mgr.userSessionMap[userId]
	if !ok {
		// 没有session
		session := NewUserSession(userId)
		mgr.userSessionMap[userId] = session
	}

	tableId := session.TableId
	if tableId == 0 {
		// 没在桌子里，忽略，或者返回错误信息
		return
	}

	tableSession, ok := mgr.tableSessionMap[tableId]
	if !ok {
		// 桌子已经不存在
		return
	}

	tableSession.MsgIn <- m

}

func (mgr *Manager) onCmdCreate(m msg.Msg, userId int) {
	session, ok := mgr.userSessionMap[userId]
	if !ok {
		// 没有session
		session := NewUserSession(userId)
		mgr.userSessionMap[userId] = session
	}

	if session.TableId != 0 {
		// 在桌子里，应当直接发送进桌命令
		return
	}

	tableSession, ok := mgr.createOneTable()

	if !ok {
		// ccreate false
		return
	}

	// usersession 关联tableid
	session.TableId = tableSession.TableId

	mgr.userSessionMap[userId] = session

	// 自动给table发送一个sit消息
	m.SetResults(map[string]interface{}{
		"msg":     "ok",
		"retcode": 0,
		"tableId": session.TableId,
	})

	mgr.SendRes(userId, m)
}

var wg sync.WaitGroup

func (mgr *Manager) createOneTable() (TableSession, bool) {
	tableMsgIn := make(chan msg.Msg)
	tableMsgOut := make(chan msg.Msg)
	tableId, ok := mgr.findOneTableId()
	if !ok {
		log.Println("not found tableid...")
		return TableSession{}, false
	}
	t := table.NewTable(tableId, tableMsgIn, tableMsgOut)

	tableSession := NewTableSession(t, tableMsgIn, tableMsgOut, tableId)

	mgr.tableSessionMap[tableId] = tableSession

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case m, ok := <-tableSession.MsgOut:
				{
					if !ok {
						continue
					}
					mgr.MsgOut <- m
				}
			case <-time.After(1 * time.Second):
				{
					continue
				}

			}
		}

	}()

	return tableSession, true

}

func (mgr *Manager) findOneTableId() (int, bool) {
	return 1, true
}

func (mgr *Manager) SendRes(userId int, m msg.Msg) {
	m.SetKey("userId", userId)
	mgr.MsgOut <- m
}

func (r *Manager) onCmdDelete(m msg.Msg) {
	userId := m["userId"].(int)
	tableId := m["tableId"].(int)

	session, ok := r.findTableSessionByRoomId(tableId)

	if !ok {
		log.Println("cant found room", m)
		return
	}

	if session.Table.Creator().UserId() != userId {
		log.Println("not the creator")
		return
	}

	// delete one room
	ok = r.deleteOneTable(tableId)
	if !ok {
		log.Println("delete not ok..")
		return
	}

	m.SetResults(map[string]interface{}{
		"retcode": 0,
		"msg":     "ok",
		"tableId": tableId,
	})

	r.SendRes(userId, m)

}

func (r *Manager) deleteOneTable(roomId int) (ok bool) {
	session, ok := r.findTableSessionByRoomId(roomId)
	session.Close()
	delete(r.tableSessionMap, roomId)
	return
}

func (r *Manager) findTableSessionByRoomId(roomId int) (session TableSession, ok bool) {
	session, ok = r.tableSessionMap[roomId]
	return
}

func (mgr *Manager) onCmdJoin(m msg.Msg, userId int) {
	params := m.GetParams()
	tableId := params["tableId"].(int)

	session, ok := mgr.userSessionMap[userId]
	if !ok {
		// 没有session
		session := NewUserSession(userId)
		mgr.userSessionMap[userId] = session
	}

	if session.TableId != 0 {
		// 在桌子里，应当直接发送进桌命令
		return
	}

	tableSession, ok := mgr.tableSessionMap[tableId]
	if !ok {
		// 要加入的桌子不存在
		return
	}

	session.TableId = tableSession.TableId
	mgr.userSessionMap[userId] = session

	// 发送一个sit命令

}
