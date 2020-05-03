<template>
  <div class="container">
    <div class="title-wrapper">
      <span>Convert Qiita articles to pdf</span>
    </div>
    <div class="form-wrapper">
      <input type="text" v-model="url">
    </div>
    <div class="alert-wrapper">
      {{ this.alert }}
    </div>
    <div class="form-wrapper">
      <button @click="receiveData">ダウンロード</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'
import qs from 'qs'

export default Vue.extend({
  data() {
    return {
      url: "",
      alert: ""  
    }
  },
  methods: {
    receiveData: function() {
      let reg = new RegExp(/https\:\/\/qiita.com\/\.*/),
          checkReg = reg.test(this.url)
      if(checkReg != true) {
        this.alert = "不正なURLです。"
        return
      } else {
        this.alert = ""
      }
      let formData = new FormData()
      formData.append("url", this.url)
      axios.post("/api", formData,
        { responseType: "blob" })
        .then(res => {
          let blob = new Blob([res.data], { type: "application/pdf" }),
              path = (window.URL || window.webkitURL).createObjectURL(blob),
              a = document.createElement("a")

          a.href = path
          a.download = "test.pdf"

          document.body.appendChild(a)
          a.click()

          document.body.removeChild(a)
        })
    }
  }
})
</script>

<style>

.container {
  min-height: 100vh;
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  padding-top: 100px;
}

.title-wrapper {
  width: 100%;
  text-align: center;
}

.title-wrapper span {
  display: inline-block;
  font-size: calc(20px + 1vw);
}

.form-wrapper {
  margin: 30px 0px;
  width: 100%;
  text-align: center;
}

.alert-wrapper {
  text-align: center;
  color: #f30b0b;
}

input {
  width: 100%;
  margin-top: 100px;
  height: 2.5rem;
  font-size: calc(15px + 0.6vw);
  border: 2px solid #7c7c7c;
  border-radius: 5px;
}

button {
  font-family: 'Roboto', 'Noto Sans JP';
  color: #FFFFFF;
  background: #55C500;
  width: 50%;
  height: 2.5rem;
  border: 0px solid;
  border-radius: 3px;
  font-size: calc(15px + 0.6vw);
}

button:hover {
  cursor: pointer;
} 

</style>
