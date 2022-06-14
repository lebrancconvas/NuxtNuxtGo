<template>
  <div id="display">
    <h1>Display</h1>
    <div v-for="user in userData" :key="user.id">
      <div>First Name: {{ user.firstname }}</div>
      <div>Last Name: {{ user.lastname }}</div>
      <div>Age: {{ user.age }}</div>
      <div>E-Mail: {{ user.email }}</div>
      <div>Tel: {{ user.tel }}</div>
      <div></div>
    </div>
  </div>
</template>

<script>
  import axios from '@nuxtjs/axios';
  export default {
    name: "Display",
    data() {
      return {
        userData: []
      }
    },
    // methods: {
    //   getData() {
    //     axios.get("http://localhost:8000/users")
    //       .then(res => this.userData === res.data)
    //   }
    // },
    async asyncData({ $axios }) {
      const userData = await $axios.$get('http://localhost:8000/users')
      return {userData}
    }
  }
</script>

<style>
  #display {
    text-align: center;
  }
</style>
