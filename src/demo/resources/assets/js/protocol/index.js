import { ProtocalBody } from './body';
const CMD_BOARDCAST = 1;
const CMD_NORMAL = 0;



const VERSION = 1.0

BoardCast = function (content) {
    return ProtocalBody.New().AddCMD(CMD_BOARDCAST).AddContent(content);
}

let NormalMsg = function (from, to, content) {
    return ProtocalBody.New().AddCMD(CMD_NORMAL).AddContent(content);
}

export * from './pos'
export * from './body'
export {
    BoardCast,
    NormalMsg
}
