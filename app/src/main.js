import Vue from 'vue';
import App from './App.vue';
import router from './routes';
import { library } from '@fortawesome/fontawesome-svg-core';
import {
  faTimes,
  faSpinner,
  faChevronLeft,
  faChevronRight,
  faPlusCircle
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import axios from 'axios';
import VueAxios from 'vue-axios';
import store from './store';

library.add(
  faTimes,
  faSpinner,
  faChevronLeft,
  faChevronRight,
  faPlusCircle
);
Vue.component('font-awesome-icon', FontAwesomeIcon);
Vue.use(VueAxios, axios);

Vue.config.productionTip = false;

new Vue({
  router,
  render: createEle => createEle(App),
  store,
  beforeCreate() {
    this.$store.commit('initializeAuthInStore');
  }
}).$mount('#app');