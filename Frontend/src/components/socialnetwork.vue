<template>
  <div class="socialnetwork">
    <h1>{{ title }}</h1>
    <h3 v-if="error">Ошибка: {{error}}</h3>
    <table>
      <tr>
        <th>Имяя</th>
        <th>Фамилия</th>
        <th>E-mail</th>
        <th>github</th>
      </tr>
      <tr v-for="contact in contact_list" v-bind:key="contact.email">
        <td>{{contact.name}}</td>
        <td>{{contact.surname}}</td>
        <td>{{contact.email}}</td>
        <td>{{contact.github}}</td>
        <td><button v-on:click="edit_contact(contact)">Редактировать контакт</button></td>
        <td><button v-on:click="remove_contact(contact)">Удалить контакт</button></td>
      </tr>
    </table>

     <h3 v-if="edit_index == -1">Новый контакт</h3>
     <form>
        <p>Имя: <input type="text" v-model="new_contact.name"></p>
        <p>Фамилия: <input type="text" v-model="new_contact.surname"></p>
        <p>Email: <input type="text" v-model="new_contact.email"></p>
        <p>Github: <input type="text" v-model="new_contact.github"></p>
        <button v-if="edit_index == -1" v-on:click="add_new_contact">Добавить контакт</button>
        <button v-if="edit_index > -1" v-on:click="end_of_edition">Закончить редактирование</button>
     </form>
  </div>
</template>

<script>
const axios = require('axios')

export default {
  name: 'socialnetwork',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      error: '',
      contact_list: [],
      new_contact:
        {
          'name': 'Иван',
          'surname': 'Иванов',
          'email': 'Iivano@domain.ru',
          'github': 'Iivanov'
        }
    }
  },
  mounted: function () {
    this.get_contacts()
  },
  methods: {
    get_contacts: function () {
      this.error = ''
      const url = '/api/socialnetwork/users'
      axios.get(url).then(response => {
        this.contact_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_new_contact: function () {
      this.error = ''
      const url = '/api/socialnetwork/users'
      axios.post(url, this.new_contact).then(response => {
        this.contact_list.push(this.new_contact)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_contact: function (item) {
      this.error = ''
      const url = '/api/socialnetwork/users/' + this.contact_list.indexOf(item)
      axios.delete(url).then(response => {
        this.contact_list.splice(this.contact_list.indexOf(item), 1)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_contact: function (item) {
      this.edit_index = this.contact_list.indexOf(item)
      this.new_contact = this.contact_list[this.edit_index]
    },
    end_of_edition: function () {
      this.error = ''
      const url = '/api/socialnetwork/users/' + this.edit_index
      axios.put(url, this.new_contact).then(response => {
        this.edit_index = -1
        this.new_contact = {
          'name': '',
          'surname': '',
          'email': '',
          'github': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
