package main

import (
	"LollipopGo/LollipopGo/conf"
	"LollipopGo/LollipopGo/log"
	_ "LollipopGo/LollipopGo/match"
	"LollipopGo/LollipopGo/player"
	"LollipopGo/LollipopGo/util"
	"Proto"
	"Proto/Proto2"
	"fmt"
)

//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
// DSQ Server 子协议的处理
func (this *NetDataConn) HandleCltProtocol2DSQ(protocol2 interface{}, ProtocolData map[string]interface{}) {

	switch protocol2 {
	case float64(Proto2.DSQ2GW_ConnServerProto2):
		{
			// 网关主动链接进来，做数据链接的保存
			this.DSQConnServerFunc(ProtocolData)
		}
	case float64(Proto2.DSQ2GW_InitGameProto2):
		{
			// 网关初始化棋牌数据
			this.DSQGameInitFunc(ProtocolData)
		}
	case float64(Proto2.DSQ2GW_PlayerStirChessProto2):
		{
			// 翻牌的协议
			this.PlayerStirChessFunc(ProtocolData)
		}
	case float64(Proto2.DSQ2GW_PlayerMoveChessProto2):
		{
			// 移动的协议
			this.PlayerMoveChessFunc(ProtocolData)
		}
	case float64(Proto2.DSQ2GW_BroadCast_GameOverProto2):
		{
			// 结算的协议
			this.BroadCastGameOverFunc(ProtocolData)
		}
	case float64(Proto2.DSQ_GameHintProto2):
		{
			// DSQ提示
			this.BroadCastGameHintFunc(ProtocolData)
		}
	case float64(Proto2.DSQ2GW_PlayerRelinkGameProto2):
		{
			// DSQ 断线重新链接
			this.RelinkGameFunc(ProtocolData)
		}
	default:
		panic("子协议：不存在！！！")
	}

	return
}

func (this *NetDataConn) RelinkGameFunc(ProtocolData map[string]interface{}) {

	StrOpenIDA := ProtocolData["OpenIDA"].(string)
	iLeftTime := int(ProtocolData["LeftTime"].(float64))
	StrChessData := ProtocolData["ChessData"].([]interface{})

	data := &Proto2.S2GWS_PlayerRelinkGame{
		Protocol:   Proto.G_GateWay_Proto,
		Protocol2:  Proto2.S2GWS_PlayerRelinkGameProto2,
		LeftTime:   iLeftTime,
		ChessBoard: StrChessData,
	}

	this.SendClientDataFunc(StrOpenIDA, "connect", data)
	return
}

func (this *NetDataConn) BroadCastGameHintFunc(ProtocolData map[string]interface{}) {

	StrOpenIDA := ProtocolData["OpenIDA"].(string)
	StrOpenIDB := ProtocolData["OpenIDB"].(string)

	data := &Proto2.BroadCast_GameHint{
		Protocol:  Proto.G_GateWay_Proto,
		Protocol2: Proto2.BroadCast_GameHintProto2,
	}

	this.SendClientDataFunc(StrOpenIDA, "connect", data)
	this.SendClientDataFunc(StrOpenIDB, "connect", data)
	return
}

func (this *NetDataConn) BroadCastGameOverFunc(ProtocolData map[string]interface{}) {
	fmt.Println("函数--BroadCastGameOverFunc")
	if ProtocolData["OpenIDB"] == nil {
		fmt.Println("玩家数据为空")
	}

	StrOpenIDA := ProtocolData["OpenIDA"].(string) // 失败者
	StrOpenIDB := ""
	if ProtocolData["OpenIDB"] != nil {
		StrOpenIDB = ProtocolData["OpenIDB"].(string) // 成功者
	}

	BIsDraw := ProtocolData["IsDraw"].(bool)

	data := &Proto2.BroadCast_GameOver{
		Protocol:  Proto.G_GateWay_Proto,
		Protocol2: Proto2.BroadCast_GameOverProto2,
		IsDraw:    BIsDraw,
		RoomID:    this.GateWayGetPlayerRoomID(StrOpenIDA),
	}

	fmt.Println("函数--BroadCastGameOverFunc StrOpenIDA", StrOpenIDA)
	fmt.Println("函数--BroadCastGameOverFunc StrOpenIDB", StrOpenIDB)

	fmt.Println("广播玩家认输数据", data)
	playerdataA := this.GateWayGetPalyerData(StrOpenIDA) //.(player.PlayerSt)
	playerdataB := this.GateWayGetPalyerData(StrOpenIDB)
	data.FailPlayer = playerdataA
	data.SuccPlayer = playerdataB
	data.FailGameLev_Exp = "0,0"
	data.SuccGameLev_Exp = "0,10"
	this.SendClientDataFunc(StrOpenIDA, "connect", data)
	this.SendClientDataFunc(StrOpenIDB, "connect", data)

	return

	// if len(StrOpenIDB) == 0 {
	// 	this.SendClientDataFunc(StrOpenIDA, "connect", data)
	// 	return
	// }

	if true {

		playerdataA := this.GateWayGetPalyerData(StrOpenIDA) //.(player.PlayerSt)
		playerdataB := this.GateWayGetPalyerData(StrOpenIDB)

		data.FailPlayer = playerdataA
		data.SuccPlayer = playerdataB

		if playerdataA == nil ||
			playerdataB == nil {
			fmt.Println("玩家数据为空")
			this.SendClientDataFunc(StrOpenIDA, "connect", data)
			return
		}

		gameexpB := 0
		if playerdataA["1"].(map[string]interface{})["GameData"] != nil {
			gameexpB = (playerdataA["1"].(map[string]interface{})["GameData"].(map[int]*player.PlayerGameLev))[DSQGameID].GameExp // .GameData[DSQGameID].GameExp
		}

		gamelevB := util.Sort_LollipopGo(conf.DSQGameExp, 10+gameexpB)
		fmt.Println("-----------------------------------", gamelevB)

		if BIsDraw { // 平局，都不加分
			data.FailGameLev_Exp = "0,0"
			if playerdataA["1"].(map[string]interface{})["GameData"] != nil {
				data.FailGameLev_Exp = util.Int2str_LollipopGo((playerdataA["1"].(map[string]interface{})["GameData"].(map[int]*player.PlayerGameLev))[DSQGameID].GameLev) + ",0"
			}
			data.SuccGameLev_Exp = "0,0"
			if playerdataB["1"].(map[string]interface{})["GameData"] != nil {
				data.SuccGameLev_Exp = util.Int2str_LollipopGo((playerdataB["1"].(map[string]interface{})["GameData"].(map[int]*player.PlayerGameLev))[DSQGameID].GameLev) + ",0"
			}
		} else {
			data.FailGameLev_Exp = "0,0"
			if playerdataA["1"].(map[string]interface{})["GameData"] != nil {
				data.FailGameLev_Exp = util.Int2str_LollipopGo((playerdataA["1"].(map[string]interface{})["GameData"].(map[int]*player.PlayerGameLev))[DSQGameID].GameLev) + ",0"
			}
			data.SuccGameLev_Exp = util.Int2str_LollipopGo(gamelevB) + ",10"
		}

		fmt.Println("广播玩家认输数据", data)
		this.SendClientDataFunc(StrOpenIDA, "connect", data)
		this.SendClientDataFunc(StrOpenIDB, "connect", data)
	}
	return
}

func (this *NetDataConn) PlayerMoveChessFunc(ProtocolData map[string]interface{}) {

	StrOpenIDA := ProtocolData["OpenIDA"].(string)
	StrOpenIDB := ProtocolData["OpenIDB"].(string)
	iRoomUID := int(ProtocolData["RoomUID"].(float64))
	StrOldPos := ProtocolData["OldPos"].(string)
	StrNewPos := ProtocolData["NewPos"].(string)
	iResultID := int(ProtocolData["ResultID"].(float64))

	data := &Proto2.S2GWS_PlayerMoveChess{
		Protocol:  Proto.G_GateWay_Proto,
		Protocol2: Proto2.S2GWS_PlayerMoveChessProto2,
		OpenID:    StrOpenIDA,
		RoomUID:   iRoomUID,
		OldPos:    StrOldPos,
		NewPos:    StrNewPos,
		ResultID:  iResultID,
	}
	this.SendClientDataFunc(StrOpenIDA, "connect", data)
	this.SendClientDataFunc(StrOpenIDB, "connect", data)
	return
}

func (this *NetDataConn) PlayerStirChessFunc(ProtocolData map[string]interface{}) {

	StrOpenID := ProtocolData["OpenID"].(string)
	StrOpenID_b := ProtocolData["OpenID_b"].(string)
	StrStirPos := ProtocolData["StirPos"].(string)
	iChessNum := int(ProtocolData["ChessNum"].(float64))
	iResultID := int(ProtocolData["ResultID"].(float64))
	// 组装数据
	data := &Proto2.S2GWS_PlayerStirChess{
		Protocol:  Proto.G_GateWay_Proto,
		Protocol2: Proto2.S2GWS_PlayerStirChessProto2,
		OpenID:    StrOpenID,
		StirPos:   StrStirPos,
		ChessNum:  iChessNum,
		ResultID:  iResultID,
	}
	this.SendClientDataFunc(data.OpenID, "connect", data)
	this.SendClientDataFunc(StrOpenID_b, "connect", data)
	return
}

// DSQ 返回给玩家
func (this *NetDataConn) DSQGameInitFunc(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil ||
		ProtocolData["RoomID"] == nil {
		panic("玩家数据错误!!!")
		return
	}
	StrOpenID := ProtocolData["OpenID"].(string)
	StrRoomID := ProtocolData["RoomID"].(string)
	iSeatNum := int(ProtocolData["SeatNum"].(float64))
	iiqipan := ProtocolData["InitData"].([]interface{})

	// 组装数据
	data := &Proto2.S2GWS_PlayerGameInit{
		Protocol:   Proto.G_GateWay_Proto,
		Protocol2:  Proto2.S2GWS_PlayerGameInitProto2,
		OpenID:     StrOpenID,
		RoomUID:    util.Str2int_LollipopGo(StrRoomID),
		SeatNum:    iSeatNum,
		ChessBoard: iiqipan,
	}

	this.SendClientDataFunc(data.OpenID, "connect", data)
	// 保存玩家的状态
	cacheGW.Add(data.OpenID, 0, player.DSQServer)
	return
}

// DSQ server 保存
func (this *NetDataConn) DSQConnServerFunc(ProtocolData map[string]interface{}) {
	if ProtocolData["ServerID"] == nil {
		panic("ServerID 数据为空!")
		return
	}

	fmt.Println("DSQ server conn entry gateway!!!")
	StrServerID := ProtocolData["ServerID"].(string)
	strDSQServer = StrServerID
	// 1 发送数据
	data := &Proto2.GW2DSQ_ConnServer{
		Protocol:  Proto.G_GameDSQ_Proto, // 游戏主要协议
		Protocol2: Proto2.GW2DSQ_ConnServerProto2,
		ServerID:  StrServerID,
	}
	// 发送数据
	this.PlayerSendMessage(data)

	// 2 保存DSQ的链接信息
	//================================推送消息处理===================================
	// 保存在线的玩家的数据信息
	onlineServer := &NetDataConn{
		Connection:    this.Connection, // 链接的数据信息
		MapSafeServer: this.MapSafeServer,
	}
	// 保存玩家数据到内存
	this.MapSafeServer.Put(StrServerID+"|DSQ_Server", onlineServer)
	//==============================================================================
	return
}

//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
// Global Server 子协议的处理
func (this *NetDataConn) HandleCltProtocol2GL(protocol2 interface{}, ProtocolData map[string]interface{}) {

	switch protocol2 {
	case float64(Proto2.G2GW_ConnServerProto2):
		{
			// 网关主动链接进来，做数据链接的保存
			this.GLConnServerFunc(ProtocolData)
		}
	case float64(Proto2.GW2G_PlayerEntryHallProto2):
		{
			// Global server 返回给服务器
			this.GWPlayerLoginGL(ProtocolData)
		}
	case float64(Proto2.GW2G_PlayerMatchGameProto2):
		{
			// Global server 玩家匹配的协议
			this.GWPlayerMatchGameGL(ProtocolData)
		}
	case float64(Proto2.G2GW_PlayerQuitMatchGameProto2):
		{
			// Global server 玩家退出匹配的协议
			this.GWPlayerQuitMatchGameGL(ProtocolData)
		}
	case float64(Proto2.G2GW_GetPlayerEmailListProto2):
		{
			// Global server 玩家获取邮件列表
			this.GWPlayerGetPlayerEmailListGL(ProtocolData)
		}
	case float64(Proto2.G2GW_ReadOrDelPlayerEmailProto2):
		{
			// Global server 邮件列表读取、删除
			this.GWPlayerReadOrDelPlayerEmailGL(ProtocolData)
		}
	case float64(Proto2.G_Broadcast_MsgNoticePlayerProto2):
		{
			// Global server 消息通知
			this.GWPlayerBroadcast_MsgNoticeGL(ProtocolData)
		}
	case float64(Proto2.G_Broadcast_NoticePlayerEmailProto2):
		{
			// Global server 邮件通知
			this.GWPlayerBroadcast_NoticePlayerEmailGL(ProtocolData)
		}

	default:
		panic("子协议：不存在！！！")
	}

	return
}

// 邮件通知
func (this *NetDataConn) GWPlayerBroadcast_NoticePlayerEmailGL(ProtocolData map[string]interface{}) {

	//	StrOpenID := ProtocolData["OpenID"].(string)
	EmailDataSt := ProtocolData["EmailData"].(map[string]interface{})

	data := &Proto2.Broadcast_NoticePlayerEmail{
		Protocol:  Proto.G_GateWay_Proto, // 游戏主要协议
		Protocol2: Proto2.Broadcast_NoticePlayerEmailProto2,
		EmailData: EmailDataSt,
	}
	fmt.Println("全服通知", data)
	//this.SendClientDataFunc(StrOpenID, "connect", data)
	this.XC_Data_Send_AllPlayer_State("", data)
	return
}

// 消息通知
func (this *NetDataConn) GWPlayerBroadcast_MsgNoticeGL(ProtocolData map[string]interface{}) {

	//StrOpenID := ProtocolData["OpenID"].(string)
	MsgDataSt := ProtocolData["MsgData"].(map[string]interface{})

	data := &Proto2.Broadcast_MsgNoticePlayer{
		Protocol:  Proto.G_GateWay_Proto, // 游戏主要协议
		Protocol2: Proto2.Broadcast_MsgNoticePlayerProto2,
		MsgData:   MsgDataSt,
	}
	fmt.Println("全服通知", data)
	// 全服通知
	this.XC_Data_Send_AllPlayer_State("", data)
	// this.SendClientDataFunc(StrOpenID, "connect", data)

	return
}

// 玩家获取邮件读取
func (this *NetDataConn) GWPlayerReadOrDelPlayerEmailGL(ProtocolData map[string]interface{}) {

	StrOpenID := ProtocolData["OpenID"].(string)
	iItype := int(ProtocolData["Itype"].(float64))
	iEmailID := int(ProtocolData["EmailID"].(float64))

	data := &Proto2.S2GWS_ReadOrDelPlayerEmail{
		Protocol:  Proto.G_GateWay_Proto, // 游戏主要协议
		Protocol2: Proto2.S2GWS_ReadOrDelPlayerEmailProto2,
		Itype:     iItype,
		EmailID:   iEmailID,
	}
	this.SendClientDataFunc(StrOpenID, "connect", data)
	return
}

// 玩家获取邮件列表
func (this *NetDataConn) GWPlayerGetPlayerEmailListGL(ProtocolData map[string]interface{}) {

	StrOpenID := ProtocolData["OpenID"].(string)
	EmailDataST := make(map[string]interface{})

	if ProtocolData["EmailData"] != nil {
		EmailDataST = ProtocolData["EmailData"].(map[string]interface{})
	}

	data := &Proto2.S2GWS_GetPlayerEmailList{
		Protocol:  Proto.G_GateWay_Proto, // 游戏主要协议
		Protocol2: Proto2.S2GWS_GetPlayerEmailListProto2,
		EmailData: EmailDataST,
	}
	this.SendClientDataFunc(StrOpenID, "connect", data)
	return
}

// 玩家退出匹配协议
func (this *NetDataConn) GWPlayerQuitMatchGameGL(ProtocolData map[string]interface{}) {
	if ProtocolData["ResultID"] == nil {
		log.Debug("Global server data is wrong:ResultID is nil!")
		return
	}
	StrOpenID := ProtocolData["OpenID"].(string)
	iResultID := ProtocolData["ResultID"].(float64)

	data := &Proto2.S2GWS_QuitMatch{
		Protocol:  Proto.G_GateWay_Proto, // 游戏主要协议
		Protocol2: Proto2.S2GWS_QuitMatchProto2,
		OpenID:    StrOpenID,
		ResultID:  int(iResultID),
	}
	// 发送数据  --
	this.SendClientDataFunc(data.OpenID, "connect", data)

	return
}

// Global server 返回给gateway server
func (this *NetDataConn) GWPlayerMatchGameGL(ProtocolData map[string]interface{}) {

	if ProtocolData["RoomUID"] == nil {
		log.Debug("Global server data is wrong:RoomUID is nil!")
		return
	}
	// 获取数据
	StrOpenID := ProtocolData["OpenID"].(string)
	//StrRoomUID := int(ProtocolData["RoomUID"].(float64))
	MatchPlayerST := make(map[string]interface{})
	if ProtocolData["MatchPlayer"] != nil {
		MatchPlayerST = ProtocolData["MatchPlayer"].(map[string]interface{})
	}
	var ChessBoard []interface{}
	if ProtocolData["ChessBoard"] != nil {
		ChessBoard = ProtocolData["ChessBoard"].([]interface{})
	}
	iResultID := int(ProtocolData["ResultID"].(float64))

	// 数据
	data_send := &Proto2.S2GWS_PlayerChooseGameMode{
		Protocol:  Proto.G_GateWay_Proto, // 游戏主要协议
		Protocol2: Proto2.S2GWS_PlayerChooseGameModeProto2,
		OpenID:    StrOpenID, // 玩家唯一标识
		// RoomUID:     StrRoomUID,
		MatchPlayer: MatchPlayerST,
		ChessBoard:  ChessBoard,
		ResultID:    iResultID,
	}

	fmt.Println("data_senddata_senddata_senddata_send", data_send)

	// 发送给匹配的人的
	iStrRoomUID := 0
	stropenidA := ""
	stropenidB := ""

	for k, v := range data_send.MatchPlayer {
		fmt.Println("-----------57575757", k)
		fmt.Println(v)
		fmt.Println("stropenidA8888888888888888888--------", v.(map[string]interface{})["RoomUID"])
		fmt.Println("stropenidA8888888888888888888--------", (v.(map[string]interface{})["RoomPlayerMap"]).(map[string]interface{})["OpenID"])
		icounts := 0
		iStrRoomUID = util.Str2int_LollipopGo(k)
		for kk, _ := range (v.(map[string]interface{})["RoomPlayerMap"]).(map[string]interface{}) {
			fmt.Println("stropenidA", kk)
			if icounts == 0 {
				stropenidA = kk
				icounts++
			} else {
				stropenidB = kk
			}
		}
	}
	//_ = iStrRoomUID
	data_send.RoomUID = iStrRoomUID
	//data_send.RoomUID = util.Str2int_LollipopGo(this.GateWayGetPlayerRoomID(StrOpenID))
	// fmt.Println("stropenidA", data_send)
	this.SendClientDataFunc(stropenidA, "connect", data_send)
	this.SendClientDataFunc(stropenidB, "connect", data_send)
	return
}

// Global server 返回给gateway server
func (this *NetDataConn) GWPlayerLoginGL(ProtocolData map[string]interface{}) {

	if ProtocolData["OpenID"] == nil {
		log.Debug("Global server data is wrong:OpenID is nil!")
		return
	}

	StrOpenID := ProtocolData["OpenID"].(string)
	StrPlayerName := ProtocolData["PlayerName"].(string)
	StrHeadUrl := ProtocolData["HeadUrl"].(string)
	StrConstellation := ProtocolData["Constellation"].(string)
	StrSex := ProtocolData["Sex"].(string)
	StGamePlayerNum := ProtocolData["GamePlayerNum"].(map[string]interface{})

	StRacePlayerNum := make(map[string]interface{})
	if ProtocolData["RacePlayerNum"] != nil {
		StRacePlayerNum = ProtocolData["RacePlayerNum"].(map[string]interface{})
	}
	StPersonal := ProtocolData["Personal"].(map[string]interface{})

	StAllPlayer := make(map[string]interface{})
	if ProtocolData["AllPlayer"] != nil {
		StAllPlayer = ProtocolData["AllPlayer"].(map[string]interface{})
	}

	StDefaultMsg := ProtocolData["DefaultMsg"].(map[string]interface{})
	// StDefaultAward := ProtocolData["DefaultAward"].(map[string]interface{})

	bIsNewEmail := ProtocolData["IsNewEmail"].(bool)
	//DefaultMsgST := ProtocolData["DefaultMsg"].(map[string]interface{})

	// 发给客户端模拟
	data := &Proto2.S2GWS_PlayerLogin{
		Protocol:      6,
		Protocol2:     2,
		PlayerName:    StrPlayerName,
		HeadUrl:       StrHeadUrl,
		Constellation: StrConstellation,
		Sex:           StrSex,
		OpenID:        StrOpenID,
		GamePlayerNum: StGamePlayerNum,
		RacePlayerNum: StRacePlayerNum,
		Personal:      StPersonal,
		DefaultMsg:    StDefaultMsg,
		// DefaultAward:  StDefaultAward,
		AllPlayer:  StAllPlayer, // 玩家的所有的结构数据
		IsNewEmail: bIsNewEmail, // 是否有新邮件
	}
	// 发送数据  --
	this.SendClientDataFunc(data.OpenID, "connect", data)
	//--------------------------------------------------------------------------
	//	playerdata := &player.PlayerSt{
	//		UID:           util.Str2int_LollipopGo(StrPlayerUID),
	//		Name:          StrPlayerName,
	//		HeadURL:       StrHeadUrl,
	//		Sex:           StrSex,
	//		PlayerSchool:  StrPlayerSchool,
	//		Constellation: StrConstellation,
	//		OpenID:        data.OpenID,
	//	}
	//	_ = playerdata
	// this.GateWaySavePalyerData(data.OpenID, data.AllPlayer)
	this.GateWaySavePalyerData(data.OpenID, data.Personal)
	// 保存状态
	cacheGW.Add(data.OpenID, 0, player.GateWayState)
	//--------------------------------------------------------------------------
	return
}

// Global server 保存
func (this *NetDataConn) GLConnServerFunc(ProtocolData map[string]interface{}) {
	if ProtocolData["ServerID"] == nil {
		panic("ServerID 数据为空!")
		return
	}

	fmt.Println("Global server conn entry gateway!!!")

	// Globla server 发过来的可以加密的数据
	StrServerID := ProtocolData["ServerID"].(string)
	strGlobalServer = StrServerID
	// 1 发送数据
	data := &Proto2.GW2G_ConnServer{
		Protocol:  9,
		Protocol2: 2,
		ServerID:  StrServerID,
	}
	// 发送数据
	this.PlayerSendMessage(data)

	// 2 保存Global的链接信息
	//================================推送消息处理===================================
	// 保存在线的玩家的数据信息
	onlineServer := &NetDataConn{
		Connection:    this.Connection, // 链接的数据信息
		MapSafeServer: this.MapSafeServer,
	}
	// 保存玩家数据到内存
	this.MapSafeServer.Put(StrServerID+"|Global_Server", onlineServer)
	//==============================================================================
	return
}

//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
//------------------------------------------------------------------------------

// client 子协议的处理
func (this *NetDataConn) HandleCltProtocol2GW(protocol2 interface{}, ProtocolData map[string]interface{}) {

	switch protocol2 {
	case float64(Proto2.C2GWS_PlayerLoginProto2):
		{
			// 功能函数处理 --  用户登陆协议
			this.GWPlayerLogin(ProtocolData)
		}
	case float64(Proto2.GateWay_HeartBeatProto2):
		{
			// 功能函数处理 --  心跳函数处理
			this.GWHeartBeat(ProtocolData)
		}
	case float64(Proto2.C2GWS_PlayerChooseGameProto2):
		{
			// 功能函数处理 --  选择游戏列表的数据
			this.PlayerEntryGame(ProtocolData)
		}
	case float64(Proto2.C2GWS_PlayerChooseGameModeProto2):
		{
			// 功能函数处理 --  选择游戏对战类型
			this.PlayerChooseGameModeGame(ProtocolData)
		}
	case float64(Proto2.GateWay_RelinkProto2):
		{
			// 断线重新链接
			this.PlayerRelinkGateWay(ProtocolData)
		}
	case float64(Proto2.C2GWS_QuitMatchProto2):
		{
			// 玩家取消匹配
			this.PlayerQuitMacthGame(ProtocolData)
		}
		/*
			--------------------------------------------------------------------
			                    Game server 斗兽棋
			--------------------------------------------------------------------
		*/
	case float64(Proto2.C2GWS_PlayerGameInitProto2):
		{
			// 功能函数处理 --  选择游戏对战类型
			this.PlayerEntryGameModeDSQGame(ProtocolData)
		}
	case float64(Proto2.C2GWS_PlayerStirChessProto2):
		{
			// 玩家翻棋
			this.PlayerStirChessDSQGame(ProtocolData)
		}
	case float64(Proto2.C2GWS_PlayerMoveChessProto2):
		{
			// 玩家移动棋子
			this.PlayerMoveChessDSQGame(ProtocolData)
		}
	case float64(Proto2.C2GWS_PlayerGiveUpProto2):
		{
			// 玩家认输--放弃
			this.PlayerGiveUpDSQGame(ProtocolData)
		}
	case float64(Proto2.C2GWS_PlayerRelinkGameProto2):
		{
			// 玩家 断线重新链接
			this.PlayerRelinkGameGame(ProtocolData)
		}
		/*
			--------------------------------------------------------------------
			      邮件系统
			--------------------------------------------------------------------
		*/
	case float64(Proto2.C2GWS_GetPlayerEmailListProto2):
		{
			// 玩家 获取邮件列表
			this.PlayerEmailListFunc(ProtocolData)
		}
	case float64(Proto2.C2GWS_ReadOrDelPlayerEmailProto2):
		{
			// 玩家 邮件列表add、del
			this.PlayerReadOrDelPlayerEmailFunc(ProtocolData)
		}
	default:
		panic("子协议：不存在！！！")
	}

	return
}

func (this *NetDataConn) PlayerReadOrDelPlayerEmailFunc(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil {
		panic("邮件列表 openid 错误！")
	}

	StrOpenID := ProtocolData["OpenID"].(string)
	iItype := int(ProtocolData["Itype"].(float64)) // 1:读取打开，2：删除，3：领取附件
	iEmailID := int(ProtocolData["EmailID"].(float64))

	data := &Proto2.GW2G_ReadOrDelPlayerEmail{
		Protocol:  Proto.G_GameGlobal_Proto,
		Protocol2: Proto2.GW2G_ReadOrDelPlayerEmailProto2,
		OpenID:    StrOpenID,
		Itype:     iItype,
		EmailID:   iEmailID,
	}

	this.SendServerDataFunc(strGlobalServer, "Global_Server", data)
	return
}

func (this *NetDataConn) PlayerEmailListFunc(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil {
		panic("获取邮件列表 openid 错误！")
	}

	StrOpenID := ProtocolData["OpenID"].(string)
	fmt.Println(StrOpenID)

	data := &Proto2.GW2G_GetPlayerEmailList{
		Protocol:  Proto.G_GameGlobal_Proto,
		Protocol2: Proto2.GW2G_GetPlayerEmailListProto2,
		OpenID:    StrOpenID,
	}

	fmt.Println(data)

	this.SendServerDataFunc(strGlobalServer, "Global_Server", data)
	return
}

// 玩家断线重新链接
func (this *NetDataConn) PlayerRelinkGameGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil {
		panic("断线重新链接 openid 错误！")
		return
	}
	StrOpenID := ProtocolData["OpenID"].(string)
	iRoomID := ProtocolData["RoomUID"].(float64)

	data := &Proto2.GW2DSQ_PlayerRelinkGame{
		Protocol:  Proto.G_GameDSQ_Proto,
		Protocol2: Proto2.GW2DSQ_PlayerRelinkGameProto2,
		OpenID:    StrOpenID,
		RoomUID:   int(iRoomID),
	}

	this.SendServerDataFunc(strDSQServer, "DSQ_Server", data)
	return
}

//------------------------------------------------------------------------------
// 缓存玩家数据
func (this *NetDataConn) GateWaySavePalyerData(stropenid string, data map[string]interface{}) {

	fmt.Println("缓存数据stropenid:", stropenid)
	fmt.Println("缓存数据前:", data)
	cacheGW.Add(stropenid+"huancun", 0, data)
	fmt.Println("缓存数据后:", this.GateWayGetPalyerData(stropenid+"huancun"))
	return
}

func (this *NetDataConn) GateWayGetPalyerData(stropenid string) map[string]interface{} {
	res, err1 := cacheGW.Value(stropenid + "huancun")
	if err1 != nil {
		// panic("没有对应数据")
		return nil
	}
	// return res.Data().(*player.PlayerSt)
	return res.Data().(map[string]interface{})
}

//保存玩家RoomID
func (this *NetDataConn) GateWaySavePlayerRoomID(stropenid, data string) {
	cacheGW.Add(stropenid+"RoomID", 0, data)
}

func (this *NetDataConn) GateWayGetPlayerRoomID(stropenid string) string {
	res, _ := cacheGW.Value(stropenid + "RoomID")
	return res.Data().(string)
}

//------------------------------------------------------------------------------

// 玩家退出匹配
func (this *NetDataConn) PlayerQuitMacthGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil {
		panic("玩家退出匹配的协议 openid 错误！")
		return
	}
	strOPenID := ProtocolData["OpenID"].(string)
	data := &Proto2.GW2G_PlayerQuitMatchGame{
		Protocol:  Proto.G_GameGlobal_Proto,
		Protocol2: Proto2.GW2G_PlayerQuitMatchGameProto2,
		OpenID:    strOPenID,
	}
	// 发送给 global server
	this.SendServerDataFunc(strGlobalServer, "Global_Server", data)
	return
}

// 断线重新链接
func (this *NetDataConn) PlayerRelinkGateWay(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil {
		panic("断线重新链接 openid 错误！")
		return
	}
	//	strOPenID := ProtocolData["OpenID"].(string)
	return
}

// 玩家认输--放弃
func (this *NetDataConn) PlayerGiveUpDSQGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil ||
		ProtocolData["RoomUID"] == nil {
		panic(ProtocolData)
	}

	StrOpenID := ProtocolData["OpenID"].(string)
	iRoomID := ProtocolData["RoomUID"].(float64)

	data := &Proto2.GW2DSQ_PlayerGiveUp{
		Protocol:  Proto.G_GameDSQ_Proto,
		Protocol2: Proto2.GW2DSQ_PlayerGiveUpProto2,
		OpenID:    StrOpenID,
		RoomUID:   int(iRoomID),
	}

	this.SendServerDataFunc(strDSQServer, "DSQ_Server", data)
	return
}

// 玩家移动棋子
func (this *NetDataConn) PlayerMoveChessDSQGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil ||
		ProtocolData["RoomUID"] == nil ||
		ProtocolData["OldPos"] == nil ||
		ProtocolData["MoveDir"] == nil {
		panic("玩家移动棋子协议错误!")
	}

	StrOpenID := ProtocolData["OpenID"].(string)
	iRoomID := ProtocolData["RoomUID"].(float64)
	StrOldPos := ProtocolData["OldPos"].(string)
	iMoveDir := int(ProtocolData["MoveDir"].(float64))

	data := &Proto2.GW2DSQ_PlayerMoveChess{
		Protocol:  Proto.G_GameDSQ_Proto,
		Protocol2: Proto2.GW2DSQ_PlayerMoveChessProto2,
		OpenID:    StrOpenID,
		RoomUID:   int(iRoomID),
		OldPos:    StrOldPos,
		MoveDir:   iMoveDir,
	}

	this.SendServerDataFunc(strDSQServer, "DSQ_Server", data)
	return
}

// 玩家翻棋
func (this *NetDataConn) PlayerStirChessDSQGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil ||
		ProtocolData["RoomUID"] == nil ||
		ProtocolData["StirPos"] == nil {
		panic("玩家翻棋子协议错误!")
	}

	StrOpenID := ProtocolData["OpenID"].(string)
	iRoomID := ProtocolData["RoomUID"].(float64)
	StrStirPos := ProtocolData["StirPos"].(string)
	// 发送给DSQ服务器
	data := &Proto2.GW2DSQ_PlayerStirChess{
		Protocol:  Proto.G_GameDSQ_Proto,
		Protocol2: Proto2.GW2DSQ_PlayerStirChessProto2,
		OpenID:    StrOpenID,
		RoomUID:   int(iRoomID),
		StirPos:   StrStirPos,
	}

	// 发送给 DSQ server
	this.SendServerDataFunc(strDSQServer, "DSQ_Server", data)
	return
}

//------------------------------------------------------------------------------
func (this *NetDataConn) PlayerEntryGameModeDSQGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil ||
		ProtocolData["RoomUID"] == nil {
		panic("初始化游戏错误！")
		return
	}

	// 获取数据
	StrOpenID := ProtocolData["OpenID"].(string)
	iRoomID := int(ProtocolData["RoomUID"].(float64))

	data := &Proto2.GW2DSQ_InitGame{
		Protocol:  Proto.G_GameDSQ_Proto,
		Protocol2: Proto2.GW2DSQ_InitGameProto2,
		OpenID:    StrOpenID,                        // 玩家唯一标识
		RoomID:    util.Int2str_LollipopGo(iRoomID), // 房间ID
	}

	// 发送给 DSQ server
	this.SendServerDataFunc(strDSQServer, "DSQ_Server", data)
	return
}

func (this *NetDataConn) PlayerChooseGameModeGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil ||
		ProtocolData["RoomID"] == nil ||
		ProtocolData["Itype"] == nil {
		panic("选择游戏对战类型协议参数错误！")
		return
	}
	fmt.Println("iRoomID:", typeof(ProtocolData["RoomID"]))
	fmt.Println("Itype:", typeof(ProtocolData["Itype"]))

	// 获取数据
	StrOpenID := ProtocolData["OpenID"].(string)
	iRoomID := ProtocolData["RoomID"].(string)
	Itype := ProtocolData["Itype"].(string)

	data := &Proto2.G2GW_PlayerMatchGame{
		Protocol:  Proto.G_GameGlobal_Proto,
		Protocol2: Proto2.G2GW_PlayerMatchGameProto2,
		OpenID:    StrOpenID, // 玩家唯一标识
		Itype:     Itype,     // Itype == 1：表示主动选择房间；Itype == 2：表示快速开始
		RoomID:    iRoomID,   // 房间ID
	}
	// 保存roomiD
	this.GateWaySavePlayerRoomID(StrOpenID, data.RoomID)
	if Itype == "1" {
		if (util.Str2int_LollipopGo(iRoomID)) < 1000000 {
			data.RoomID = this.GateWayGetPlayerRoomID(StrOpenID)
		}
	}
	// 发送给 global server
	this.SendServerDataFunc(strGlobalServer, "Global_Server", data)

	return
}

func (this *NetDataConn) PlayerEntryGame(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil ||
		ProtocolData["GameID"] == nil {
		panic("进入游戏协议参数错误！")
		return
	}
	StrOpenID := ProtocolData["OpenID"].(string)
	StrGameID := ProtocolData["GameID"].(string)
	StrTimestamp := ProtocolData["Timestamp"].(string)
	_ = StrOpenID
	_ = StrTimestamp
	data := &Proto2.S2GWS_PlayerChooseGame{
		Protocol:  Proto.G_GateWay_Proto,
		Protocol2: Proto2.S2GWS_PlayerChooseGameProto2,
		RoomList:  conf.G_RoomList[StrGameID],
	}
	// 发送数据
	fmt.Println("StrGameID:", StrGameID)
	fmt.Println("房间列表:", data)
	this.PlayerSendMessage(data)
	return
}

func (this *NetDataConn) GWHeartBeat(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil {
		panic("心跳协议参数错误！")
		return
	}

	StrOpenID := ProtocolData["OpenID"].(string)
	// 将我们解析的数据 --> token --->  redis 验证等等
	// 主要看TTL的时间是否正确
	data := &Proto2.GateWay_HeartBeat{
		Protocol:  6,
		Protocol2: 3,
		OpenID:    StrOpenID,
	}
	// 发送数据
	this.PlayerSendMessage(data)
	return
}

func (this *NetDataConn) GWPlayerLogin(ProtocolData map[string]interface{}) {
	if ProtocolData["Token"] == nil ||
		ProtocolData["PlayerUID"] == nil {
		panic("网关登陆协议错误！！！")
		return
	}

	StrPlayerUID := ProtocolData["PlayerUID"].(string)
	StrPlayerName := ProtocolData["PlayerName"].(string)
	StrHeadUrl := ProtocolData["HeadUrl"].(string)
	StrConstellation := ProtocolData["Constellation"].(string)
	StrPlayerSchool := ProtocolData["PlayerSchool"].(string)
	StrSex := ProtocolData["Sex"].(string)
	StrToken := ProtocolData["Token"].(string)

	// 1 将我们解析的数据 --> token --->  redis 验证等等
	// 主要看TTL的时间是否正确
	// 2 发送给Global server 获取数据  在线人数等
	data := &Proto2.G2GW_PlayerEntryHall{
		Protocol:      Proto.G_GameGlobal_Proto,
		Protocol2:     Proto2.G2GW_PlayerEntryHallProto2,
		UID:           StrPlayerUID,
		OpenID:        util.MD5_LollipopGO(StrPlayerUID),
		PlayerName:    StrPlayerName,
		HeadUrl:       StrHeadUrl,
		Constellation: StrConstellation,
		PlayerSchool:  StrPlayerSchool,
		Sex:           StrSex,
		Token:         StrToken,
	}
	this.SendServerDataFunc(strGlobalServer, "Global_Server", data)
	//============================================================================
	// 缓存玩家数据--但是数据不全  ---  最好在global server返回数据的时候保存
	//	playerdata := &player.PlayerSt{
	//		UID:           util.Str2int_LollipopGo(StrPlayerUID),
	//		Name:          StrPlayerName,
	//		HeadURL:       StrHeadUrl,
	//		Sex:           StrSex,
	//		PlayerSchool:  StrPlayerSchool,
	//		Constellation: StrConstellation,
	//		OpenID:        data.OpenID,
	//	}
	//	this.GateWaySavePalyerData(data.OpenID, playerdata)
	//================================推送消息处理===================================
	// 保存在线的玩家的数据信息
	onlineUser := &NetDataConn{
		Connection: this.Connection, // 链接的数据信息
		MapSafe:    this.MapSafe,
	}
	// 保存玩家数据到内存
	this.MapSafe.Put(data.OpenID+"|connect", onlineUser)
	//==============================================================================

	return
}
