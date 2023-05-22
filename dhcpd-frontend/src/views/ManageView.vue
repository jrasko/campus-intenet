<template>
  <v-snackbar v-model="success" :timeout="2000" color="success"> Erfolg!</v-snackbar>
  <v-snackbar v-model="failure" :timeout="3000" color="error"> {{ this.errorMessage }}</v-snackbar>
  <v-row>
    <v-col v-if="!(this.$route.name === 'add')">
      <RouterLink to="/add">
        <v-btn prepend-icon="mdi-account-plus"> Person hinzufügen</v-btn>
      </RouterLink>
    </v-col>
    <v-col>
      <v-btn prepend-icon="mdi-credit-card-refresh" @click="this.resetPayments">
        Zahlungen zurücksetzen
      </v-btn>
    </v-col>
    <v-col>
      <v-btn prepend-icon="mdi-content-copy" @click="this.copyEmails"> Emails kopieren </v-btn>
    </v-col>
  </v-row>
  <v-row>
    <v-col cols="12">
      <v-table hover>
        <thead>
          <tr>
            <th>Zahlung</th>
            <th>Vorname</th>
            <th>Nachname</th>
            <th>MAC</th>
            <th>IP</th>
            <th>WG</th>
            <th>Zimmer-Nr.</th>
            <th>Telefonr.</th>
            <th>E-Mail</th>
            <th></th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in this.people">
            <td v-if="p.hasPaid">
              <v-icon color="green" icon="mdi-checkbox-marked-circle" />
            </td>
            <td v-else>
              <v-icon color="red" icon="mdi-close-circle" />
            </td>
            <td>{{ p.firstname }}</td>
            <td>{{ p.lastname }}</td>
            <td>{{ p.mac }}</td>
            <td>{{ p.ip }}</td>
            <td>{{ p.wg }}</td>
            <td>{{ p.roomNr }}</td>
            <td>{{ p.phone }}</td>
            <td>{{ p.email }}</td>
            <td>
              <v-row align="center" justify="center">
                <v-col cols="1">
                  <RouterLink :to="'/edit/' + p.id">
                    <v-btn density="compact" icon="mdi-square-edit-outline" />
                  </RouterLink>
                </v-col>
                <v-col cols="1">
                  <v-btn density="compact" icon="mdi-delete" @click="this.delete(p)" />
                </v-col>
              </v-row>
            </td>
            <td></td>
          </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>
</template>
<script>
import { deleteConfigFor, getConfigs, resetPayments } from '@/axios'

export default {
  data() {
    return {
      people: [],
      success: false,
      failure: false,
      errorMessage: ''
    }
  },
  mounted() {
    this.refresh()
  },
  methods: {
    refresh() {
      getConfigs()
        .then((resp) => {
          this.people = resp.data
        })
        .catch((e) => {
          if (e.response.status === 403) {
            this.$router.push('/login')
          }
          console.log(e)
        })
    },
    async copyEmails() {
      let mails = ''
      for (const p of this.people) {
        mails += p.email + ';'
      }
      await navigator.clipboard.writeText(mails)
    },
    delete(p) {
      if (confirm('Wirklich löschen?')) {
        deleteConfigFor(p.id)
          .then(() => {
            this.success = true
            this.refresh()
          })
          .catch((e) => {
            this.errorMessage = e.response.data
            this.failure = true
          })
      }
    },
    resetPayments() {
      if (confirm('Zahlungen zurücksetzen?')) {
        resetPayments().then(() => {
          this.success = true
          this.refresh()
        })
      }
    }
  }
}
</script>

<style scoped></style>
