import { shallowMount, createLocalVue} from '@vue/test-utils';
import VisitorList from '@/views/VisitorList.vue';
import Vuex from 'vuex';

const localVue = createLocalVue();
localVue.use(Vuex);

describe("VisitorList", () => {
  let store;
  let actions;

  beforeEach(() => {
    actions = {
      fetchVisitors: jest.fn()
    };
    store = new Vuex.Store({
      state: {
        visitors: [],
        filtered: [],
      },
      getters: {
        filtered: () => [
          { ID: "1", Name: "test1", Notes: "asdf1" },
          { ID: "2", Name: "test2", Notes: "asdf2" },
          { ID: "3", Name: "test3", Notes: "asdf3" }
        ],
        visitors: () => [],
      },
      actions
    });
  });
  afterEach(() => {
  });

  it('Shows 3 visitors when the getters filter returns 3', async () => {
    const wrapper = shallowMount(VisitorList, {
      localVue,
      store,
      stubs: ['font-awesome-icon', 'router-link', 'router-view']
    });

    expect(actions.fetchVisitors).toHaveBeenCalled();

    expect(wrapper.findAll('.td-name').length).toEqual(3);
  });
});
