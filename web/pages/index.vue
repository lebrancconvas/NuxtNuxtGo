<template>
  <div>
    <header>
      <h1>Nuxt Nuxt Go</h1>
      <h2>FORM</h2>
    </header>
    <form action="POST">
      <div>
        <label>First Name *</label>
        <input type="text" placeholder="First Name" v-model="firstname" required>
      </div>
      <div>
        <label>Last Name *</label>
        <input type="text" placeholder="Last Name" v-model="lastname" required>
      </div>
      <div>
        <label>Age *</label>
        <input type="number" placeholder="Age" v-model="age" required>
      </div>
      <div>
        <label>E-Mail *</label>
        <input type="email" placeholder="E-Mail" v-model="email" required>
      </div>
      <div>
        <label>Tel. *</label>
        <input type="tel" placeholder="Telephone" v-model="tel" required>
      </div>
      <div>
        <label>Image Upload (.jpg, .jpeg, .png)</label>
        <input type="file" name="image" id="image">
      </div>
      <div>
        <input class="submit" type="submit" value="Submit">
      </div>
    </form>
    <div id="display">
      <h3>Test Display</h3>
      <p>First Name: {{ firstname }}</p>
      <p>Last Name: {{ lastname }}</p>
      <p>Age: {{ age }}</p>
      <p>E-Mail: {{ email }}</p>
      <p>Telephone: {{ tel }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IndexPage',
  data() {
    return {
      firstname: '',
      lastname: '',
      age: '',
      email: '',
      tel: ''
    }
  }
}
</script>

<style>
  header {
    text-align: center;
  }
  form {
    max-width: 450px;
    margin: 30px auto;
    background: rgb(255, 236, 151);
    text-align: left;
    padding: 35px;
    border-radius: 10px;
  }
  label {
    color: rgb(37, 0, 141);
    display: inline-block;
    margin: 25px 0 15px;
    font-size: 0.7em;
    text-transform: uppercase;
    letter-spacing: 1px;
    font-weight: bold;
  }
  input {
    display: block;
    padding: 10px 5px;
    width: 100%;
    box-sizing: border-box;
    border: none;
    border-bottom: 1px solid #ddd;
    color: #001100;
  }
  .submit {
    cursor: pointer;
  }
  .submit:active {
    transform: scale(0.98);
  }
  .submit:hover {
    background: #9f9f9f
  }
</style>
