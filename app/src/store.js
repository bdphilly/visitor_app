import Vue from 'vue';
import Vuex from 'vuex';
import { fetchAll } from './api/entries';
import { getToken } from './api/auth';
import moment from 'moment';

Vue.use(Vuex);

export const storeConfig = {
  state: {
    authenticated: false,
    visitors: [],
    filtered: []
  },
  getters: {
    isAuthenticated(state) {
      return state.authenticated;
    },
    visitors(state) {
      return state.visitors;
    },
    filtered(state) {
      return state.filtered;
    },
  },
  actions: {
    async fetchVisitors ({ commit }) {
      const resp = await fetchAll();

      if (resp.status === 200) {
        commit('setVisitors', resp.data);
      } else {
        console.error('error fetching visitors', resp);
      }
    },
    filterVisitors ({ commit, state }, query) {
      if (!query.length) {
        return commit('resetFiltered');
      }

      if (query.toLowerCase() === 'signed out') {
        const signOutFiltered = state.visitors.filter(el => !!el.SignOutTime);
        return commit('setFiltered', signOutFiltered);
      }

      const filtered = [];
      state.visitors.forEach((visitor) => {
        if (visitor.SignOutTime) {
          visitor.SignOutTime = moment(visitor.SignOutTime).format('MMM D YYYY @ h:mm a');
        }

        if (Object.values(visitor).join('').indexOf(query) > -1) {
          filtered.push(visitor);
        }
      });
      commit('setFiltered', filtered);
    },
  },
  mutations: {
    resetFiltered (state) {
      state.filtered = state.visitors;
    },
    setFiltered (state, filtered) {
      state.filtered = filtered;
    },
    setVisitors (state, visitors) {
      state.visitors = visitors;
      state.filtered = visitors;
    },
    initializeAuthInStore (state) {
      if (getToken()) {
        state.authenticated = true;
      }
    }
  }
};

export default new Vuex.Store(storeConfig);