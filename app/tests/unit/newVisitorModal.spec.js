import { shallowMount, createLocalVue} from '@vue/test-utils';
import NewVisitorModal from '@/components/NewVisitorModal.vue';
import axios from 'axios';
import flushPromises from "flush-promises";
import VueRouter from 'vue-router';
import MockAdapter from 'axios-mock-adapter';
import { routes } from '../../src/routes';

const localVue = createLocalVue();
localVue.use(VueRouter);

describe("NewVisitorModal", () => {
  const mock = new MockAdapter(axios);

  beforeEach(() => {
    mock.reset();
    window.scrollTo = jest.fn();
    delete window.location;
    window.location = { assign: jest.fn() };
  });
  afterEach(() => {
    mock.restore();
  });

  it('Shows the error from the server when the request is bad.', async () => {
    const errorMessage = 'Bad Payload';
    mock
      .onPost()
      .reply(400, errorMessage);

    const wrapper = shallowMount(NewVisitorModal, {
      stubs: ['font-awesome-icon']
    });

    wrapper.find('#inputName').setValue('alice jones');
    wrapper.find('#inputNotes').setValue('this is a note');
    wrapper.find('form').trigger('submit.prevent');

    await flushPromises();

    const msg = wrapper.find('.error-message');
    expect(msg.text()).toEqual(errorMessage);
  });

  xit('Returns to main page when successfully adds new visitor.', async () => {
    const router = new VueRouter({
      routes
    });

    mock
      .onPost()
      .reply(202, {});

    jest.spyOn(window.location, 'assign').mockImplementation(url => url);

    const wrapper = shallowMount(NewVisitorModal, {
      localVue,
      router,
      stubs: ['font-awesome-icon']
    });

    wrapper.find('#inputName').setValue('alice jones');
    wrapper.find('#inputNotes').setValue('this is a note');
    wrapper.find('form').trigger('submit.prevent');

    await flushPromises();
    // still not quite working
    expect(window.location.assign).toHaveBeenCalledWith('/');
  });
});
