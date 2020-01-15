import { shallowMount } from '@vue/test-utils';
import SignOutButton from '@/components/SignOutButton.vue';
import moment from 'moment';

describe('SignOutButton.vue', () => {
  it('Only renders the right button when on the first page', () => {
    const component = shallowMount(SignOutButton, {
      propsData: {
        time: null,
      }
    });
    
    expect(component.find('.sign-out-text').exists()).toBe(true);
  });

  it('Renders both buttons when it is a middle page', () => {
    const component = shallowMount(SignOutButton, {
      propsData: {
        time: new moment().format(),
      }
    });
    
    expect(component.find('.sign-out-text').exists()).toBe(false);
  });
});