let Body = function () {

    let B = function () {
        this.CMD = 1
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

    B.prototype.AddCMD = function (value) {
        this.CMD = value
        return this
    }

    B.prototype.AddContent = function (value) {
        this.Content = value
        return this
    }

    let New = function () {
        return new B()
    }

    return {
        New: New,
    }
}()


// s = Body.New().AddCMD(3)
// console.log(s.CMD)
let ProtocalBody = BODY

export { ProtocalBody }

