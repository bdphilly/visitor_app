import Vue from 'vue';
import VueRouter from 'vue-router';
import VisitorList from './views/VisitorList';
import Login from './views/Login';
import NewVisitorModal from './components/NewVisitorModal';
import { getToken } from './api/auth';

Vue.use(VueRouter);

const router = new VueRouter({
 mode: 'history',
 routes: [
  {
    path: '/',
    name: 'Home',
    component: VisitorList,
    meta: {
      requiresAuth: true
    },
    children: [{
      path: '/new',
      name: 'NewModal',
      component: NewVisitorModal,
      props: true,
      meta: {
        requiresAuth: true
      },
    }],
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '*',
    name: 'CatchAll',
    component: Login
  }
]});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const token = getToken();

    if (token) {
      next();
    } else {
      next({name: 'Login'});
    }
  }
  next();
});

export default router;