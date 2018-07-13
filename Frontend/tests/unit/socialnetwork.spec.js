import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import socialnetwork from '@/components/socialnetwork.vue'

describe('socialnetwork.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(socialnetwork, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders titles', () => {
    const wrapper = shallowMount(socialnetwork, {})
    expect(wrapper.text()).to.include('Имя')
    expect(wrapper.text()).to.include('Фамилия')
    expect(wrapper.text()).to.include('Email')
    expect(wrapper.text()).to.include('github')
  })
})