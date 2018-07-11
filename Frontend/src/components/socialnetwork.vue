<template>
  <div class="socialnetwork">
    <h1>{{ title }}</h1>
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
        <button v-if="edit_index > -1" v-on:click="make_new_contact">Создать нового контакт</button>
     </form>
  </div>
</template>

<script>
export default {
  name: 'socialnetwork',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      contact_list: [
        {
          'name': 'Иван',
          'surname': 'Иванов',
          'email': 'Iivanov@domain.ru',
          'github': 'Iivanov'
        },
        {
          'name': 'Имя',
          'surname': 'Фамилия',
          'email': 'user@domain.ru',
          'github': 'user'
        }
      ],
      new_contact:
        {
          'name': 'Иван',
          'surname': 'Иванов',
          'email': 'Iivano@domain.ru',
          'github': 'Iivanov'
        }
    }
  },
  methods: {
    add_new_contact: function () {
      this.contact_list.push(this.new_contact)
    },
    remove_contact: function (item) {
      this.contact_list.splice(this.contact_list.indexOf(item), 1)
    },
    edit_contact: function (item) {
      this.edit_index = this.contact_list.indexOf(item)
      this.new_contact = this.contact_list[this.edit_index]
    },
    make_new_contact: function () {
      this.edit_index = -1
      this.new_contact = {
        'name': '',
        'surname': '',
        'email': '',
        'github': ''
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
