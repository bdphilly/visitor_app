<template>
  <div class="pagination">
    <button
      class="pagination-button"
      id="pagination-back"
      :style="{visibility: page > 0 ? 'visible' : 'hidden'}"
      v-on:click="left()">
      <font-awesome-icon class="pagination-icon" icon="chevron-left" />
    </button>
    <span class="page">{{page}}</span>
    <button
      class="pagination-button"
      id="pagination-next"
      :style="{visibility: page * numPerPage < total - numPerPage ? 'visible' : 'hidden'}"
      v-on:click="right()">
      <font-awesome-icon class="pagination-icon" icon="chevron-right" />
    </button>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'Pagination',
  props: {
    numPerPage: {
      type: Number,
      required: true
    },
    page: {
      type: Number,
      required: true
    },
    total: {
      type: Number,
      required: true
    },
  },
  methods: {
    left () {
      this.$emit('changePage', false);
    },
    right () {
      this.$emit('changePage', true);
    }
  }
};
</script>

<style lang="scss" scoped>
.pagination {
  display: flex;
  justify-content: center;
}
.page {
  padding: 10px;
}
.pagination-button {
  border: none;
  border-radius: 5px;

  &:focus {
    outline:0;
  }
}
.pagination-icon {
  cursor: pointer;

  &:hover {
    opacity: 0.7;
  }
}
</style>
