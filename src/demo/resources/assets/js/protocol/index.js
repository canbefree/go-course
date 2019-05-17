import { ProtocalBody } from './body';
const CMD_BOARDCAST = 1;
const CMD_NORMAL = 0;




let BoardCast = function (cmd) {
    return ProtocalBody.New().AddCMD(cmd).AddContent(content);
}


export * from './pos'
export * from './body'
export {
    BoardCast
}
