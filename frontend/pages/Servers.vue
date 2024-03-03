<template>
  <v-row>
    <v-col cols="12">
      <v-table hover>
        <thead>
        <tr>
          <th>Name</th>
          <th>MAC</th>
          <th>IP</th>
          <th>Deaktiviert</th>
          <th>Kommentar</th>
          <th/>
          <th/>
        </tr>
        </thead>
        <tbody>
        <tr v-for="r in servers">
          <td>{{ r.name }}</td>
          <td>{{ r.mac }}</td>
          <td>{{ r.ip }}</td>
          <td>
            <v-icon v-if="r.disabled" color="red" icon="mdi-close-circle"/>
            <v-icon v-else color="green" icon="mdi-checkbox-marked-circle"/>
          </td>
          <td>{{ r.disabled }}</td>
          <td>{{ r.comment }}</td>
          <td>
            <NuxtLink :to="'/edit/server/'+r.id">
              <v-icon density="compact" icon="mdi-square-edit-outline"/>
            </NuxtLink>
          </td>
        </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>
</template>
<script lang="ts" setup>
  import {fetchRooms} from "~/utils/fetch";

  const emits = defineEmits(['logout'])

  const servers = ref<Server[]>([])

  const filters = ref<RoomFilters>({
    occupied: null,
    block: [],
  })

  onMounted(() => {
    nextTick(refresh)
  })

  async function refresh() {
    try {
      const resp = await fetchRooms(filters.value);
    } catch (e: any) {
      if (e.statusCode === 403) {
        emits('logout')
      }
      console.log(e)
    }
  }
</script>
