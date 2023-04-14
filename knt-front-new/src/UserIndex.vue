<template>
  <!-- d-flex align-items-center flex-column -->
  <div class="product-page">

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

      <div class="keyboard">
        <SimpleKeyboard @onChange="onChange" @onKeyPress="onKeyPress" :input="input" />
      </div>

    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import SimpleKeyboard from "./SimpleKeyboard.vue";

interface User {
  id: number;
  firstName: string;
  lastName: string;
  money: number;
}

//Static Data
const listOfUsers: User[] = [];
const userOne: User = { id: 1, firstName: "Name1", lastName: "Surname1", money: 100 };
const userTwo: User = { id: 2, firstName: "Name2", lastName: "Surname3", money: 200 };
const userFour: User = { id: 3, firstName: "Name3", lastName: "Surname3", money: 300 };
const userThree: User = { id: 3, firstName: "Name3", lastName: "Surname3", money: 300 };
const userFive: User = { id: 3, firstName: "Name3", lastName: "Surname3", money: 300 };
const userSix: User = { id: 3, firstName: "Name3", lastName: "Surname3", money: 300 };
const userSeven: User = { id: 3, firstName: "Name3", lastName: "Surname3", money: 300 };
const userEight: User = { id: 3, firstName: "Name3", lastName: "Surname3", money: 300 };

listOfUsers.push(userOne);
listOfUsers.push(userTwo);
listOfUsers.push(userThree);
listOfUsers.push(userFour);
listOfUsers.push(userFive);
listOfUsers.push(userSix);
listOfUsers.push(userSeven);
listOfUsers.push(userEight);

//http://127.0.0.1:5000/admin/users

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
  top: 5%;
  height: 45%;
}

.bottom-section {
  bottom: 0;
  height: 45%;
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

.keyboard {
  width: 80%;
}
</style>