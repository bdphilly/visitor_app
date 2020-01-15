<template>
  <div class="btn-wrapper">
    <div v-if="!loading && time">{{ time | formatDate }}</div>
    <div
      v-else
      v-on:click="signOut"
      class="sign-out-btn">
      <div v-if="loading">Signing Out<font-awesome-icon class="loading-icon" icon="spinner" /></div>
      <div v-else class="sign-out-text">Sign out</div>
    </div>
  </div>
</template>

<script>

import moment from 'moment';

export default {
  name: 'SignOutButton',
  components: {
  },
  props: {
    time: {
      type: String,
    },
  }, 
  data() {
    return {
      loading: false
    };
  },
  created() {
  },
  methods: {
    signOut() {
      this.loading = true;
      this.$emit('signOutClick');

      // timeout hack for animation
      setTimeout(() => {
        this.loading = false;
      }, 1000);
    },
  },
  filters: {
    formatDate: function (date) {
      return moment(date).format('MMM D YYYY @ h:mm a');
    }
  },
};
</script>

<style scoped lang="scss">
.sign-out-btn {
  cursor: pointer;
  color: #6c6c6c;

  &:hover {
    opacity: 0.7;
  }
}
.loading-icon {
  padding-left: 10px;
}
.sign-out-text {
  padding: 5px 10px;
  border: 1px solid rgba(#de5625, 0.5);
  border-radius: 5px;
}
</style>