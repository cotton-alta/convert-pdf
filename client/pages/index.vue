<template>
  <div class="container">
    <div class="form-wrapper">
      <input type="text" v-model="url">
    </div>
    <div class="form-wrapper">
      <button @click="receiveData">取得</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'
import qs from 'qs'

export default Vue.extend({
  data() {
    return { url: "" }
  },
  methods: {
    receiveData: function() {
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
  max-width: 1000px;
  margin: 100px auto;
}

.form-wrapper {
  margin: 30px 0px;
  width: 100%;
  text-align: center;
}

input {
  width: 60%;
}

</style>
