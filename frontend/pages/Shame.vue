<template>
  <v-row>
    <v-col>
      <v-img aspect-ratio="16/9" height="350" src="/shame.gif"/>
    </v-col>
    <v-col>
      <v-img height="350" src="/shamebox.gif"></v-img>
    </v-col>
  </v-row>
  <v-row>
    <v-col>
      <v-table hover>
        <thead>
        <tr>
          <th>WG</th>
          <th>Vorname</th>
          <th>Nachname</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="p in moochers">
          <td>{{ p.wg }}</td>
          <td>{{ p.member?.firstname }}</td>
          <td>{{ p.member?.lastname }}</td>
        </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>
</template>
<script lang="ts" setup>
  import {fetchRooms} from "~/utils/fetch_rooms";

  const moochers = ref<Room[]>([])
  onMounted(() => nextTick(() => refresh()))

  async function refresh() {
    try {
      moochers.value = await fetchRooms({payment: false, occupied: true})
    } catch (e) {
      console.log(e)
    }
  }
</script>
<style scoped></style>
