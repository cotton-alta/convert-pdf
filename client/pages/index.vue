<template>
  <div class="container">
    <div>
      <button @click="receiveData">取得</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'

export default Vue.extend({
  methods: {
    receiveData: function() {
      axios.get("/api",{
        responseType: "blob"
      })
      .then(res => {
        let blob = new Blob([res.data], { type: "application/pdf" }),
            url = (window.URL || window.webkitURL).createObjectURL(blob),
            a = document.createElement("a")

        a.href = url
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
}

</style>
