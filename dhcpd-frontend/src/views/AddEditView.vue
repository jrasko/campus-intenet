<template>
  <v-snackbar v-model="success" :timeout="2000" color="success"> Erfolg! </v-snackbar>
  <v-snackbar v-model="failure" :timeout="2000" color="error"> Fehler! </v-snackbar>
  <v-form @submit.prevent="">
    <v-container>
      <v-row>
        <v-col>
          <v-text-field v-model="config.firstname" label="Vorname" />
        </v-col>
        <v-col>
          <v-text-field v-model="config.lastname" label="Nachname" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field v-model="config.mac" label="MAC-Adresse" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field v-model="config.wg" label="WG" />
        </v-col>
        <v-col>
          <v-text-field v-model="config.roomNr" label="Zimmernummer" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field v-model="config.email" label="Email" />
        </v-col>
        <v-col>
          <v-text-field v-model="config.phone" label="Telefonnummer" />
        </v-col>
        <v-col>
          <v-switch v-model="config.hasPaid" color="green" label="Bezahlt" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <RouterLink to="/">
            <v-btn color="red"> Abbrechen </v-btn>
          </RouterLink>
        </v-col>
        <v-col>
          <v-btn color="blue" @click="this.goNext"> Speichern & Next </v-btn>
        </v-col>
        <v-col>
          <v-btn color="green" @click="this.submit">Speichern </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import { updateConfig } from '@/utils/axios'

export default {
  data: () => ({
    success: false,
    failure: false,
    config: {
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
      updateConfig(this.config)
        .then(() => {
          this.success = true
        })
        .catch((e) => {
          this.failure = true
          console.log(e)
        })
    },
    goNext() {
      updateConfig(this.config)
        .then(() => {
          this.success = true
          this.config = {}
        })
        .catch((e) => {
          this.failure = true
          console.log(e)
        })
    }
  }
}
</script>
<style scoped>
</style>
