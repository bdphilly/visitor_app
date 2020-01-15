<template>
  <transition name="bounce">
    <div class="modal">
      <h1 class="title">New Visitor</h1>
        <router-link to="/" class="link-back">
          <font-awesome-icon class="cancel" icon="times" />
        </router-link>
      <hr>
      <form class="form" @submit.prevent="newVisitor">
        <div class="input-wrapper">
          <label >
            <div class="label">Name: *</div>
            <input
              type="text"
              id="inputName"
              v-bind:class="{ 'has-error': !!error }"
              placeholder="ie. John Smith"
              v-model="form.name"
              class="new-input">
          </label>
        </div>
        <div class="input-wrapper">
          <label>
            <div class="label">Notes:</div>
            <input
              id="inputNotes"
              type="text"
              placeholder="ie. Visitor is vegan"
              v-model="form.notes"
              class="new-input">
          </label>
        </div>
        <div class="req-text">( * denotes required )</div>
        <input class="submit-btn" type="submit" value="Submit" />
      </form>
      <div class="error-message" v-if="error">{{error}}</div>
    </div>
  </transition>
</template>

<script>

import { create } from '../api/entries';

// For now, the modal is also a route.
export default {
  name: 'NewVisitorModal',
  props: {
    msg: String
  },
  data () {
    return {
      form: {},
      error: null
    };
  },
  created () {
    window.scrollTo(0, 0);
  },
  methods: {
    newVisitor() {
      const { name, notes } = this.form;
      const visitor = {
        name,
        notes
      };

      create(visitor)
        .then(() => {
          window.location.assign('/');
          // this.$router.replace({name: 'Home', force: true}).catch(() => {}); // force reload not working
        })
        .catch(err => {
          const { data, status } = err.response;

          if (status == 400) {
            this.error = data;
          } else {
            console.error('Unhandled error.', err);
          }
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.modal {
  background: #ccc;
  bottom: 10%;
  box-shadow: 5px 5px 40px rgba(0,0,0,0.4);
  color: #333;
  left: 20%;
  position: absolute;
  right: 20%;
  top: 10%;
}
.link-back {
  position: absolute;
  right: 25px;
  top: 25px;
}
.link-back:visited {
  color: #333;
}
.cancel:hover {
  opacity: 0.7;
}
.form {
  align-items: center;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.input-wrapper {
  margin: 15px;
}
.new-input {
  border-radius: 0;
  border: 0;
  font-size: 14px;
  height: 30px;
  line-height: 24px;
  padding: 5px 10px;
  width: 300px;
}
.submit-btn {
  border-radius: 5px;
  border: 0;
  padding: 10px 15px;
  font-size: 16px;
}
.submit-btn:hover {
  cursor: pointer;
  opacity: 0.7;
}
.label {
  margin-bottom: 5px;
  text-align: left;
}
.title {
  text-align: center;
}
.error-message {
  color: #e15656;
  margin-top: 30px;
  text-align: center;
}
.has-error {
  border: 2px solid #e15656;
}

.bounce-enter-active {
  animation: bounce-in .5s;
}
.bounce-leave-active {
  animation: bounce-in .5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
.req-text {
  font-size: 10px;
  margin: 15px;
}
</style>
