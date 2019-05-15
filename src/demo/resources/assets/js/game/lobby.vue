//大厅测试端

<template>
  <div :host="host" :port="port">
    <div>
      <p>大厅测试端</p>
    </div>
    <div class="row">
      <button @click="gameSelect('cheer')">选择干瞪眼游戏</button>
      <button @click="gameSelect('rand')">选择转盘游戏</button>
    </div>

    <div>
      <component v-bind:is="subGame"></component>
    </div>

    <div>
      <p>模拟断线重连(登陆事件触发)</p>
      <input v-model="uid">
      <button @click="TriggerPress">{{button_text}}</button>
    </div>

    <div>
      <p>广播</p>
    </div>
    <div>
      <input v-model="bmsg">
      <button @click="BoardCast">发送</button>
    </div>

    <div>
      <p>私聊</p>
    </div>
    <div>
      <input v-model="fid">
      <input v-model="fmsg">
      <button @click="SendMsg">给指定对象发送消息</button>
    </div>

    <p>消息窗口</p>
    <div>
      <textarea id="public-msg" v-model="public_msg"></textarea>
    </div>
  </div>
</template>

<style>
#public-msg {
  width: 600px;
  height: 400px;
}
</style>

<script>
import { Msg } from "@/game/msg";
import Cheer from "@/game/cheer/client";
import Rand from "@/game/rand/client";
import { WS } from "@/lib";

console.log(WS);

export default {
  components: { Cheer, Rand },
  props: {
    host: {
      required: true
    },
    port: {
      required: true
    }
  },

  data() {
    return {
      button_text: "连接", //为了保证按钮更新及时
      conn: null, //连接

      uid: 123,
      boardcast: null,
      fid: null,
      fmsg: null,
      bmsg: null,
      public_msg: "",

      subGame: "cheer" //动态组件
    };
  },

  mounted() {},

  methods: {
    TriggerPress: function() {
      if (this.isConnected()) {
        return this.close();
      } else {
        if (!this.uid) {
          alert("请输入您的身份ID!");
          return;
        }
        this.connect();
      }
    },
    connect() {
      if (!this.isConnected()) {
        let url = "ws://" + this.host + ":" + this.port + "/ws?uid=" + this.uid;
        this.conn = new WebSocket(url);
      }
      this.conn.onopen = this.onOpen;
      this.conn.onmessage = this.onMessage;
      this.conn.onerror = this.onError;
      this.conn.onclose = this.onClose;
    },

    gameSelect() {
      this.subGame = arguments[0];
    },

    close() {
      this.button_text = "连接";
      this.conn.close();
    },
    onOpen(e) {
      console.log(this.conn);
      this.button_text = "断开连接";
    },
    onMessage(e) {
      this.public_msg += e.data + "\n";
      console.log("msg from server:", e.data);
    },
    onError(e) {
      console.log("msg error");
    },
    onClose(e) {
      this.button_text = "重新连接";
      console.log("msg close");
    },
    BoardCast() {
      if (this.isConnected()) {
        let m = new Msg(this.conn);
        m.boardCast(this.bmsg);
      }
    },
    SendMsg() {
      if (this.isConnected()) {
        let m = new Msg(this.conn);
        m.sendMsg(this.fid, this.fmsg);
      }
    },
    isConnected() {
      if (this.conn && this.conn.readyState) {
        return this.conn.readyState === 1 ? true : false;
      }
      return false;
    },
    getNowTime() {
      return Math.floor(new Date().getTime() / 1000);
    }
  }
};
</script>