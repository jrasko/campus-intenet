<script setup></script>
<template>
  <v-snackbar v-model="success" :timeout="2000" color="success"> Erfolg!</v-snackbar>
  <v-snackbar v-model="failure" :timeout="2000" color="error"> Fehler!</v-snackbar>
  <v-container fluid>
    <v-row justify="center">
      <v-col v-if="!(this.$route.name === 'add')" cols="10">
        <RouterLink to="/update">
          <v-btn prepend-icon="mdi-account-plus"> Person hinzufügen</v-btn>
        </RouterLink>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="10">
        <v-table hover>
          <thead>
            <tr>
              <th>Zahlung</th>
              <th>Vorname</th>
              <th>Nachname</th>
              <th>MAC</th>
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
              <td>{{ p.wg }}</td>
              <td>{{ p.roomNr }}</td>
              <td>{{ p.phone }}</td>
              <td>{{ p.email }}</td>
              <td>
                <v-container>
                  <v-row align="center" justify="center">
                    <v-col cols="1">
                      <RouterLink :to="'/update/' + p.mac">
                        <v-btn density="compact" icon="mdi-square-edit-outline" />
                      </RouterLink>
                    </v-col>
                    <v-col cols="1">
                      <v-btn density="compact" icon="mdi-delete" @click="this.delete(p)" />
                    </v-col>
                  </v-row>
                </v-container>
              </td>
              <td></td>
            </tr>
          </tbody>
        </v-table>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import { deleteConfigFor, getConfigs } from '@/axios'

export default {
  data() {
    return {
      people: [],
      success: false,
      failure: false
    }
  },
  mounted() {
    this.refresh()
  },
  methods: {
    refresh() {
      getConfigs().then((data) => {
        this.people = data
      })
    },
    delete(p) {
      deleteConfigFor(p.mac)
        .then(() => {
          this.success = true
          this.refresh()
        })
        .catch(() => {
          this.failure = true
        })
    }
  }
}
</script>

<style scoped></style>
