package main

// 阶段
const (
	GAME_STATUS_DOWNBET = 10 + iota // 下注
	GAME_STATUS_LOTTERY             // 开奖
	GAME_STATUS_BALANCE             // 结算
	GAME_STATUS_READY               //准备
)

// 通信协议
const (
	MSG_GAME_INFO_AUTO              = 410000 //410000匹配
	MSG_GAME_INFO_QDESKINFO         = 410001 //410001,请求桌子消息
	MSG_GAME_INFO_RDESKINFO         = 410002 //410002,请求桌子消息返回
	MSG_GAME_INFO_NSTATUS_CHANGE    = 410003 //410003,阶段改变通知
	MSG_GAME_INFO_QDOWNBET          = 410004 //410004,玩家请求下注
	MSG_GAME_INFO_RDOWNBET          = 410005 //410005,玩家请求下注返回
	MSG_GAME_INFO_NDOWNBET          = 410006 //410006,玩家请求下注通知
	MSG_GAME_INFO_NLOTTERY          = 410007 //410007,开奖通知
	MSG_GAME_INFO_NBALANCE          = 410008 //410008,结算通知
	MSG_GAME_INFO_NTIPS             = 410009 //410009,x局未下注提示通知
	MSG_GAME_INFO_RRECONNECT_REPLAY = 410010 //410010,断线重连返回
	MSG_GAME_INFO_NDESKCHANGE       = 410011 //410011,局号改变通知
	MSG_GAME_INFO_NRECORD           = 410012 //410012,开奖记录通知
	MSG_GAME_INFO_MOREPLAYER        = 410013 //请求更多玩家
	MSG_GAME_INFO_MOREPLAYER_REPLY  = 410014 //请求更多玩家响应
	MSG_GAME_INFO_BETAGAIN          = 410015 //重复下注请求
	MSG_GAME_INFO_BETAGAIN_REPLAY   = 410016 //重复下注请求响应
	MSG_GAME_INFO_GET_RECORD        = 410017 //410017 获取游戏记录
	MSG_GAME_INFO_GET_RECORD_REPLY  = 410018 //410018 获取游戏记录应答
	MSG_GAME_INFO_CHAIR_UPDATE      = 410019 //座位变更通知
)
