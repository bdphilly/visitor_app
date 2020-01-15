<template>
  <div class="btn-wrapper">
    Filter by:
    <input
      type="text"
      class="searchbar"
      v-on:keyup="filter"
      v-model="query" />
  </div>
</template>

<script>

import { login } from '../api/auth';

export default {
  name: 'Login',
  data () {
    return {
      query: ''
    };
  },
  methods: {
    lookup () {
      login();
    },
    filter () {
      this.$store.dispatch('filterVisitors', (this.query));
    },
  }
};
</script>

<style scoped lang="scss">
.btn-wrapper {
  padding: 10px;
  text-align: center;

  &:hover {
    cursor: pointer;
    opacity: 0.7;
  }
}
.searchbar {
  margin: 5px;
  padding: 10px;
  font-size: 16px;
  width: 200px;

  &:focus {
    outline:0;
  }
}
</style>