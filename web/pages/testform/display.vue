<template>
  <div>
    <h1>Display</h1>
    <div>
      <div>{{ userData.firstname }}</div>
      <div></div>
      <div></div>
      <div></div>
      <div></div>
    </div>
    <button @click="getData">Show Data</button>
  </div>
</template>

<script>
  import axios from 'axios';
  export default {
    name: "Display",
    data() {
      return {
        userData: []
      }
    },
    methods: {
      getData() {
        axios.get("http://localhost:8000/users")
          .then(res => this.userData === res.data)
      }
    },
    async asyncData({ $axios }) {
      const userData = await $axios.$get('http://localhost:8000/users')
      return {userData}
    }
  }
</script>
