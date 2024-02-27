<template>
  <v-snackbar v-model="success" :timeout="2000" color="success"> Erfolg!</v-snackbar>
  <v-snackbar v-model="failure" :timeout="3000" color="error">{{ this.errorMessage }}</v-snackbar>
  <AddEditForm :disable-ip="true" :person="this.person" @macUpdate="updateMac" />
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

export default {
  components: { AddEditForm: AddEditForm },
  props: {},
  data: () => ({
    success: false,
    failure: false,
    errorMessage: '',
    person: {
      firstname: '',
      lastname: '',
      mac: '',
      wg: '',
      roomNr: '',
      phone: '',
      email: '',
      hasPaid: false,
      comment: ''
    }
  }),
  methods: {
    submit() {
      createConfig(this.person)
        .then(() => {
          this.$router.push('/')
        })
        .catch((e) => {
          if (e.response.status === 403) {
            this.errorMessage = 'no permissions for that'
            this.failure = true
          } else {
            this.errorMessage = e.response.data
            this.failure = true
          }
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
          if (e.response.status === 403) {
            this.errorMessage = 'no permissions for that'
            this.failure = true
          } else {
            this.errorMessage = e.response.data
            this.failure = true
          }
          console.log(e)
        })
    },
    updateMac() {
      this.person.mac = formatMac(this.person.mac)
    }
  }
}
</script>
<style scoped></style>
