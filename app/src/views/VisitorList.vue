<template>
  <div class="">
    <div class="title-wrapper">
      <h1 title="title">Visitor List</h1>
      <Searchbar />
      <router-link to="/new" class="add-user-btn">
        <font-awesome-icon class="" icon="plus-circle" />
        Add
      </router-link>
    </div>
    <table class="visitor-table">
      <thead class="visitor-head">
        <tr>
          <th width="20%" class="th-name">Name</th>
          <th width="50%" class="th-notes">Notes</th>
          <th width="30%" class="th-time">Signed out</th>
        </tr>
      </thead>
      <tbody class="visitor-body">
        <tr v-for="(visitor) in currentPage" :key="visitor.ID">
          <td class="td-name">{{visitor.Name}}</td>
          <td class="td-notes">{{visitor.Notes}}</td>
          <td class="td-time">
            <SignOutButton
              :time="visitor.SignOutTime"
              v-on:signOutClick="signOutVisitor(visitor.ID)" />
          </td>
        </tr>
      </tbody>
    </table>
    <Pagination
      :page="page"
      :total="visitors.length"
      :numPerPage="numPerPage"
      v-on:changePage="pageChange" />
    <router-view></router-view>
  </div>
</template>

<script>

import { mapActions, mapGetters } from 'vuex';
import { signOut } from '../api/entries';
import Pagination from '../components/Pagination';
import SignOutButton from '../components/SignOutButton';
import Searchbar from '../components/Searchbar';

export default {
  name: 'VisitorList',
  components: {
    Pagination,
    SignOutButton,
    Searchbar
  },
  data() {
    return {
      loading: false,
      error: null,
      page: 0,
      numPerPage: 10
    };
  },
  created() {
    this.fetchVisitors();
  },
  methods: {
    signOutVisitor(id) {
      this.loading = true;

      signOut(id)
        .then(resp => {
          const visitor = this.visitors.find(v => v.ID === resp.data.ID);
          visitor.SignOutTime = resp.data.SignOutTime;
        })
        .catch(err => {
          this.loading = false;
          this.error = err.toString();
        });
    },
    pageChange(increment) {
      // TODO: move the page logic to the vuex store
      if (increment) {
        this.page++;
      } else {
        this.page--;
      }
    },
    ...mapActions(['fetchVisitors'])
  },
  computed: {
    ...mapGetters({
      visitors: 'filtered'
    }),
    currentPage() {
      const { page, numPerPage } = this;
      return (this.visitors || []).slice(page * numPerPage, page * numPerPage + numPerPage);
    },
  }
};
</script>

<style scoped lang="scss">
.visitor-table {
  border-collapse: collapse; 
  width: 100%;

  tr { 
    border-bottom: 1px solid #ccc;
    height: 50px;
  }
}
.td-notes {
  text-align: center;
}
.td-time {
  display: flex;
  justify-content: center;
  align-items: center; 
  height: 50px;
}
.visitor-body {
  tr {
    &:nth-child(2n) {
      color: #969696;
    }
  }
  text-align: left;
}
.title {
  text-align: center;
}
.add-user-btn {
  color: #343434;
  text-decoration: none;

  &:hover {
    opacity: 0.7;
  }
}
.title-wrapper {
  align-items: center;
  display: flex;
  justify-content: space-between;
}
</style>