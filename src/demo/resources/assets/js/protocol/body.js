let Body = function () {
    let B = function () {
        this.CMD
        this.POS
        this.BODY
    }


    //静态
    B.prototype.stringify = function () {
        console.log(this)
        let body = {
            //前提 CMD POS BODY 不会存在0 null 特殊值
            CMD: this.CMD || null,
            POS: this.POS || null,
            Content: this.Content || null,
        }
        return JSON.stringify(body)
    }

    B.prototype.Set = function (ob) {
        this.CMD = ob.CMD
        return this
    }

    return new B
}

let ProtocalBody = new Body()

// ob = { CMD: 12 }
// console.log(ProtocalBody.Set(ob).JsonEncode())

export { ProtocalBody }