<template>
  <v-snackbar v-model="success" :timeout="2000" color="success"> Erfolg!</v-snackbar>
  <v-snackbar v-model="failure" :timeout="2000" color="error"> Fehler!</v-snackbar>
  <AddEditForm :person="this.person" disable-ip="true" />
  <v-row>
    <v-col>
      <RouterLink to="/">
        <v-btn color="red"> Abbrechen</v-btn>
      </RouterLink>
    </v-col>
    <v-col>
      <v-btn color="blue" @click="this.goNext"> Speichern & Next</v-btn>
    </v-col>
    <v-col>
      <v-btn color="green" @click="this.submit">Speichern</v-btn>
    </v-col>
  </v-row>
</template>

<script>
import { createConfig } from '@/axios'
import AddEditForm from '@/components/AddEditForm.vue'
import { th } from 'vuetify/locale'

export default {
  computed: {
    th() {
      return th
    }
  },
  components: { AddEditForm: AddEditForm },
  props: {},
  data: () => ({
    success: false,
    failure: false,
    person: {
      firstname: '',
      lastname: '',
      mac: '',
      wg: '',
      roomNr: '',
      phone: '',
      email: '',
      hasPaid: false
    }
  }),
  methods: {
    submit() {
      createConfig(this.person)
        .then(() => {
          this.$router.push('/')
        })
        .catch((e) => {
          this.failure = true
          console.log(e)
        })
    },
    goNext() {
      createConfig(this.person)
        .then(() => {
          this.success = true
          this.person = {}
        })
        .catch((e) => {
          this.failure = true
          console.log(e)
        })
    }
  }
}
</script>
<style scoped></style>
