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
<script lang="ts" setup>

  const moochers = ref<ReducedPerson[]>([])
  onMounted(() => nextTick(() => refresh()))

  async function refresh() {
    try {
      moochers.value = await getShameList()
    } catch (e) {
      console.log(e)
    }
  }
</script>
<style scoped></style>
