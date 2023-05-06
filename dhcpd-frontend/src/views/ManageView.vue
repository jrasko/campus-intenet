<script setup></script>
<template>
  <v-container fluid>
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
            </tr>
          </tbody>
        </v-table>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import { getConfigs } from '@/utils/axios'

export default {
  data() {
    return {
      people: []
    }
  },
  mounted() {
    getConfigs().then((data) => {
      this.people = data
    })
  }
}
</script>

<style scoped></style>
