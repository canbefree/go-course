//大厅测试端

<template>
  <div :host="host" :port="port">
    <div>
      <p>大厅测试端</p>
    </div>
    <div class="row">
      <button @click="gameSelect('cheer')">选择干瞪眼</button>
      <button @click="gameSelect('rand')">选择转盘</button>
    </div>

    <div>
      <component v-bind:is="subGame"></component>
    </div>

    <div>
      <p>模拟断线重连(登陆事件触发)</p>
      <input v-model="uid">
      <button @click="TriggerPress">{{button_text}}{{Connected}}</button>
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
      <button @click="SendMsg">给指定对象发送消息1</button>
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
import { NormalMsg, BoardCast } from "@/protocol";

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
      Connected: false,

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

  mounted() {
    //注册状态修改
    WS.registerStatusFunction(v => {
      this.Connected = v;
    });
    //注册消息接受 不同的消息处理机制都要处理一遍
    WS.registerMsgHandleFunction(v => {
      this.public_msg += v;
    });
  },

  methods: {
    TriggerPress: function() {
      if (this.Connected) {
        return WS.Close();
      } else {
        if (!this.uid) {
          alert("请输入您的身份ID!");
          return;
        }
        this.connect();
      }
    },
    connect() {
      let url = "ws://" + this.host + ":" + this.port + "/ws?uid=" + this.uid;
      WS.Init(url);
    },

    gameSelect() {
      this.subGame = arguments[0];
    },

    BoardCast() {
      let msg = BoardCast(this.bmsg);
      WS.Send(msg);
    },

    SendMsg() {
      WS.Send(NormalMsg());
    }
  }
};
</script>