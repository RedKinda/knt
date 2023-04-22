<template>
  <!-- d-flex align-items-center flex-column -->
  <div class="product-page">

    <!-- NAVIGATION TEST -->
    <router-link to="/products">Products</router-link>

    <!--Search Bar-->
    <input type="text" class="input user-search-bar" v-model="search" @input="onInputChange"
      placeholder="Search your name..." />

    <!-- align-self-start -->
    <div class="top-section">

      <!--Displaying users-->
      <div v-for="user in searchedUsers()" :key="user.id"
        class="d-flex user-card align-items-center border-bottom border-2">
        <div class="p-3">
          <h5>{{ user.firstName + " " + user.lastName }}</h5>
        </div>
      </div>
    </div>

    <div class="bottom-section d-flex align-items-center flex-column">

      <div class="keyboard d-flex align-items-center">
        <SimpleKeyboard @onChange="onChange" @onKeyPress="onKeyPress" :input="input" />
      </div>

    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import axios from 'axios'
import SimpleKeyboard from "../components/SimpleKeyboardUserPage.vue";

interface User {
  id: number;
  firstName: string;
  lastName: string;
  money: number;
}

const listOfUsers: User[] = [];

export default defineComponent({

  components: {
    SimpleKeyboard
  },

  data() {
    return {
      search: "",
      input: "",
      users: listOfUsers,
    };
  },

  methods: {
    onChange(input: string) {
      this.search = input;
    },
    onKeyPress(button: string) {
      console.log("button", button);
    },
    onInputChange(input: any) {
      this.search = input.target.value;
    },
    searchedUsers() {
      return this.users.filter(user => (user.firstName + " " + user.lastName).toLowerCase().includes(this.search.toLowerCase()))
    }
  },
  mounted(){
    //change the URL to http://127.0.0.1:5000/users later. Also idk how to do SSR yet
    this.users = [];
    axios
    .get('http://127.0.0.1:5000/admin/users')
      .then(response => {
        for(let i = 0; i < response.data.length; i++){
          const userToAdd: User = {id: response.data[i].id, firstName: response.data[i].firstName, lastName: response.data[i].lastName, money: response.data[i].balance};
          console.log(userToAdd)
          this.users.push(userToAdd);
        }
      })
  }
});
</script>

<style>
.top-section,
.bottom-section {
  position: fixed;
  left: 0;
  right: 0;
}

.top-section {
  overflow: scroll;
  top: 55px;
  height: 45%;
}

.bottom-section {
  bottom: 0;
  height: 45%;
  display: flex;
  justify-content: center;
}

.header {
  display: flex;
  justify-content: end;
  align-items: center;
}

.navigation-button {
  display: inline-block;
  margin: 0.5rem 0.5rem 0;
}

.user-card:hover {
  background-color: #e5e5e5;
  transition-timing-function: ease;
  transition: 0.5s;
}

.user-search-bar {
  display: block;
  margin: 0.5rem 0.5rem 0.5rem auto;
  padding: 5px;
  width: 20%;
}

.keyboard{
  width: 100%;
  height: 100%;
}

</style>