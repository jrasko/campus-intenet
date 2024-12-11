<template>
  <v-row align="end" justify="center">
    <v-col cols="6" md="2">
      <NuxtLink to="/servers/add">
        <v-btn prepend-icon="mdi-view-grid-plus">
          Server Hinzufügen
        </v-btn>
      </NuxtLink>
    </v-col>
    <v-col cols="6" md="2">
      <v-select
        v-model="filters.disabled"
        :items="serverFilter.disabled"
        append-inner-icon="mdi-filter"
        hide-details
        item-title="header"
        item-value="value"
        label="Deaktiviert"
        variant="underlined"
        @update:modelValue="refresh"
      />
    </v-col>
    <v-spacer/>
  </v-row>
  <v-row>
    <v-col cols="12">
      <v-table hover>
        <thead>
        <tr>
          <th>Aktiv</th>
          <th>Name</th>
          <th>MAC</th>
          <th>IP</th>
          <th/>
        </tr>
        </thead>
        <tbody>
        <tr v-for="s in servers">
          <td>
            <v-icon :color="s.disabled ? 'orange' : 'green'" icon="mdi-circle-medium"/>
          </td>
          <td>{{ s.name }}</td>
          <td>{{ s.mac }}</td>
          <td>{{ s.ip }}</td>
          <td>
            <v-row align="center" justify="center">
              <v-col cols="1">
                <NuxtLink :to="'/servers/edit/' + s.id">
                  <v-btn density="compact" icon="mdi-square-edit-outline"/>
                </NuxtLink>
              </v-col>
              <v-col cols="1">
                <v-btn density="compact" icon="mdi-delete" @click="removeServer(s.id)"/>
              </v-col>
            </v-row>
          </td>
        </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>
</template>
<script lang="ts" setup>
  import {listServers, deleteServer} from "~/utils/fetch_netconfig";
  import {serverFilter} from "~/utils/constants";

  const emit = defineEmits(['logout'])

  const servers = ref<Server[]>([])

  const filters = ref<ServerFilters>({
    disabled: null,
    server: true
  })

  onMounted(() => {
    nextTick(refresh)
  })

  async function refresh() {
    try {
      servers.value = await listServers(filters.value);
    } catch (e: any) {
      if (e.statusCode === 403) {
        emit('logout')
      }
      console.log(e)
    }
  }

  async function removeServer(id: number) {
    try {
      await deleteServer(id)
      refresh()
    } catch (e) {
      console.log(e)
    }
  }
</script>
