import { mount } from '@vue/test-utils';
import Pagination from '@/components/Pagination.vue';

describe('Pagination.vue', () => {
  it('Only renders the right button when on the first page', () => {
    const component = mount(Pagination, {
      propsData: {
        numPerPage: 10,
        total: 100,
        page: 0
      }
    });
    
    expect(component.find("#pagination-back").attributes().style).toBe('visibility: hidden;');
  });

  it('Renders both buttons when it is a middle page', () => {
    const component = mount(Pagination, {
      propsData: {
        numPerPage: 10,
        total: 100,
        page: 2
      }
    });
    
    expect(component.find("#pagination-back").attributes().style).toBe('visibility: visible;');
    expect(component.find("#pagination-next").attributes().style).not.toBe('visibility: hidden;');
  });
});