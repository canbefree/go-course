import { ProtocalBody } from "../protocol";

let Version = 1;
let CMD_NORMAL = 1;
let CMD_BOARDCAST = 2;
let CMD_GAME = 3;


let WS = function () {
    let conn = null
    let connected = false
    let _statusFunc = null
    let _msgHandleFunc = null

    let wsUrl = null

    let Init = (url) => {
        if (!connected) {
            wsUrl = url
            conn = new WebSocket(wsUrl);
        }

        conn.onopen = OnOpen
        conn.onclose = OnClose
        conn.onmessage = OnMessage
        return true
    }

    let Send = (protocolBody) => {
        if (!check()) return false;
        console.log("send:", protocolBody.stringify())
        conn.send(protocolBody.stringify())
    }

    let check = () => {
        if (!wsUrl) {
            console.log('请先初始化连接')
            return false;
        }
        if (!connected && wsUrl) {
            if (!Init(wsUrl)) {
                console.log('websocket error ')
                return false;
            }
        }
        return true
    }



    //利用回调修改 vue前端的连接状态
    let registerStatusFunction = (f) => {
        _statusFunc = f
    }


    let registerMsgHandleFunction = (f) => {
        _msgHandleFunc = f
    }

    let OnOpen = (e) => {
        console.log('onopen')
        connected = true
        _statusFunc(true)
    }

    let OnClose = (e) => {
        console.log('onclose')
        connected = false
        _statusFunc(false)
    }
    let OnMessage = (e) => {
        return _msgHandleFunc(e.data)
    }

    //关闭连接
    let Close = () => {
        conn.close()
        _statusFunc(false)
    }

    let getNowTime = () => {
        return Math.floor(new Date().getTime() / 1000);
    }

    let NormalMsg = () => {
        ProtocalBody.New().AddCMD
        conn.send();
    }

    return {
        Connected: connected,
        Init: Init,
        Send: Send,
        Close: Close,
        registerStatusFunction: registerStatusFunction,
        registerMsgHandleFunction: registerMsgHandleFunction,


        //发送普通消息
        NormalMsg: NomalMsg,
        //广播
        BoardCast: BoardCast,
        //游戏逻辑协议
        GameMsg: GameMsg,
    }


}()




export { WS }
