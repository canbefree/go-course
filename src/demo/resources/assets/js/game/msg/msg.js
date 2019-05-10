const CMD_BOARDCAST = 1;
const CMD_NORMAL = 0;

//发送数据Msg
let Msg = (function () {
    let Msg = function (conn) {
        this.conn = conn;
    };

    Msg.prototype.sendMsg = function (fid, msg) {
        let body = {CMD: CMD_NORMAL, Content: msg, FID: parseInt(fid)}
        this.conn.send(JSON.stringify(body))
    }
    Msg.prototype.boardCast = function (msg) {
        let body = {CMD: CMD_BOARDCAST,Content: msg}
        this.conn.send(JSON.stringify(body))
    }

    return Msg
}());


export {Msg}




