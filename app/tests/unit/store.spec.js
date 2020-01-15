import { createLocalVue } from '@vue/test-utils';
import Vuex from 'vuex';
import { storeConfig } from '../../src/store';
import cloneDeep from 'lodash.clonedeep';
import moment from 'moment';

const localVue = createLocalVue();
localVue.use(Vuex);

describe('store', () => {
  let store;

  const testVisitors = [
    { ID: "1", Name: "test1", Notes: "asdf1", SignOutTime: moment().format() },
    { ID: "2", Name: "test2", Notes: "asdf2", SignOutTime: moment().format() },
    { ID: "3", Name: "test3", Notes: "asdf3" }
  ];

  beforeEach(() => {
    store = new Vuex.Store(cloneDeep(storeConfig));
  });

  it('sets visitors correctly', () => {
    expect(store.state.visitors.length).toEqual(0);
    store.commit('setVisitors', testVisitors);
    expect(store.state.visitors.length).toEqual(3);
  });

  it('filters visitors correctly', () => {
    store.commit('setVisitors', testVisitors);
    expect(store.state.filtered.length).toEqual(3);
    store.dispatch('filterVisitors', 'test3');
    expect(store.state.filtered.length).toEqual(1);
  });

  it('filters signed out visitors correctly', () => {
    store.commit('setVisitors', testVisitors);
    expect(store.state.filtered.length).toEqual(3);
    store.dispatch('filterVisitors', 'Signed out');
    expect(store.state.filtered.length).toEqual(2);
  });
});