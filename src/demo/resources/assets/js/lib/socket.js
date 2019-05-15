import { ProtocalBody } from "../protocol";

let WS = function () {
    let conn = null
    let connected = false
    let func = null

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
        conn.send(ProtocalBody.stringify())
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
        func = f
    }

    let OnOpen = (e) => {
        console.log('onopen')
        connected = true
        func(true)
    }

    let OnClose = (e) => {
        console.log('onclose')
        connected = false
        func(false)
    }
    let OnMessage = (e) => {

    }
    let OnError = (e) => {
        connected = false
    }

    let Close = () => {
        conn.close()
        func(false)
    }
    return {
        Connected: connected,
        Init: Init,
        Send: Send,
        Close: Close,
        registerStatusFunction: registerStatusFunction
    }

}()




export { WS }
