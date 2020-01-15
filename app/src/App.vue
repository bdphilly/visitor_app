<template>
  <div id="app">
    <div v-if="isAuthenticated" v-on:click="signOut" class="logout-btn">Logout</div>
    <div class="main-container">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import { logout } from '../src/api/auth';
import { mapGetters } from 'vuex';

export default {
  name: 'app',
  methods: {
    signOut () {
      logout();
    },
  },
  computed: {
    ...mapGetters({
      isAuthenticated: 'isAuthenticated'
    })
  }
};
</script>

<style lang="scss">
#app {
  color: #2c3e50;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
}
.main-container {
  border: 1px solid #dedede;
  margin: 5% 20% 10%;
  padding: 50px;
}
.logout-btn {
  cursor: pointer;
  position: absolute;
  right: 3%;
  top: 3%;

  &:hover {
    opacity: 0.7;
  }
}
</style>
