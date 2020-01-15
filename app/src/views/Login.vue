<template>
  <div class="btn-wrapper" v-on:click="signIn">
    Login
  </div>
</template>

<script>

import { login } from '../api/auth';

export default {
  name: 'Login',
  methods: {
    signIn () {
      login();
    }
  },

};
</script>

<style scoped lang="scss">
.btn-wrapper {
  padding: 10px;
  border: 1px solid #969696;
  text-align: center;

  &:hover {
    cursor: pointer;
    opacity: 0.7;
  }
}
</style>