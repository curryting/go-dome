// Code generated by goctl. DO NOT EDIT.
package types

type ChatList struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

type ChatListReq struct {
	Uid string `json:"uid"`
}

type ChatListRes struct {
	List []ChatList `json:"list"`
}