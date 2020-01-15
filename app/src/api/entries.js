import Vue from 'vue';
import axios from 'axios';
import { Token } from './auth';

const apiBase = 'http://localhost:9090';
const auth = {
  headers: { 'X-Secret': Token },
};

export function fetchAll() {
  return Vue.axios.get(`${apiBase}/entries`, auth);
}

export function search(query) {
  return Vue.axios.get(`${apiBase}/entries?${query}`, auth);
}

export function create(data) {
  return axios.post(`${apiBase}/entries`, JSON.stringify(data), auth);
}

export function signOut(id) {
  return Vue.axios.patch(`${apiBase}/entries/${id}`, null, auth);
}
