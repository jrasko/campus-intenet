<template>
  <v-snackbar v-model="failure" :timeout="2000" color="error"> Fehler!</v-snackbar>
  <AddEditForm :person="this.person" />
  <v-container>
    <v-row>
      <v-col>
        <RouterLink to="/">
          <v-btn color="red"> Abbrechen</v-btn>
        </RouterLink>
      </v-col>
      <v-col>
        <v-btn color="green" @click="this.submit">Speichern</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { getConfigFor, updateConfig } from '@/axios'
import AddEditForm from '@/components/AddEditForm.vue'

export default {
  components: { AddEditForm },
  props: {},
  data: () => ({
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
  mounted() {
    getConfigFor(this.$route.params.mac).then((resp) => {
      this.person = resp.data
    })
  },
  methods: {
    submit() {
      updateConfig(this.person)
        .then(() => {
          this.$router.push('/')
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
