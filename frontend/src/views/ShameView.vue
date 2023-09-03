<template>
  <v-row>
    <v-col>
      <v-img aspect-ratio="16/9" height="350" src="/shame.gif" />
    </v-col>
    <v-col>
      <v-img src="/shamebox.gif" height="350"></v-img>
    </v-col>
  </v-row>
  <v-row>
    <v-col>
      <v-table hover>
        <thead>
          <tr>
            <th>Vorname</th>
            <th>Nachname</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in moochers">
            <td>{{ p.firstname }}</td>
            <td>{{ p.lastname }}</td>
          </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>
</template>
<script>
import { getShameList } from '@/axios'

export default {
  data() {
    return {
      moochers: []
    }
  },
  mounted() {
    getShameList().then((r) => {
      this.moochers = r.data
    })
  }
}
</script>
<style scoped></style>
