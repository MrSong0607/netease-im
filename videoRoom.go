package netease

import (
	"errors"
	"net/http"
)

const (
	neteaseRoomBaseURL = "https://roomserver-dev.netease.im/v1/api"
	roomInfoPoint      = neteaseRoomBaseURL + "/rooms/{id}"
	roomMembersPoint   = neteaseRoomBaseURL + "/rooms/{id}/members"
)

const (
	//RoomModeDuet 双人模式
	RoomModeDuet = iota + 1
	//RoomModeMulti 多人模式
	RoomModeMulti
)

const (
	//RoomStatusInit 初始状态
	RoomStatusInit = iota + 1
	//RoomStatusRunning 进行中
	RoomStatusRunning
	//RoomStatusEnded 正常结束
	RoomStatusEnded
	//RoomStatusException 异常结束
	RoomStatusException
)

//GetRoomInfo .
func (c *ImClient) GetRoomInfo(roomID string) (*RoomInfo, error) {
	param := map[string]string{"id": roomID}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetPathParams(param)

	resp, err := client.Get(roomInfoPoint)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	ri := &RoomInfo{}
	if err := jsonTool.Unmarshal(resp.Body(), ri); err != nil {
		return nil, err
	}

	return ri, nil
}

//GetRoomMembers .
// func (c *ImClient) GetRoomMembers(roomID string) {
// 	c.setCommonHead()
// 	param := map[string]string{"id": roomID}

// 	client := c.client.R()
// 	client.SetPathParams(param)

// 	resp, err := client.Get(roomMembersPoint)
// }

//DeleteRoom 删除某个房间
func (c *ImClient) DeleteRoom(roomID string) error {
	param := map[string]string{"id": roomID}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetPathParams(param)

	if _, err := client.Delete(roomInfoPoint); err != nil {
		return err
	}

	return nil
}
