<template>
  <v-row align="baseline" justify="center">
    <v-col cols="6" md="2">
      <v-select
        v-model="filters.block"
        :items="[1,2,3,4,5]"
        append-inner-icon="mdi-filter"
        hide-details
        item-title="header"
        item-value="value"
        label="Block"
        multiple
        variant="underlined"
        @update:modelValue="refresh"
      />
    </v-col>
    <v-col cols="6" md="2">
      <v-select
        v-model="filters.occupied"
        :items="roomFiler.occupied"
        append-inner-icon="mdi-filter"
        hide-details
        item-title="header"
        item-value="value"
        label="Belegt"
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
          <th>Block</th>
          <th>Zimmer Nr</th>
          <th>WG</th>
          <th>Belegt</th>
          <th>Vorname</th>
          <th>Nachname</th>
          <th/>
        </tr>
        </thead>
        <tbody>
        <tr v-for="r in rooms">
          <td>{{ r.block }}</td>
          <td>{{ r.roomNr }}</td>
          <td>{{ r.wg }}</td>
          <td>
            <v-icon v-if="r.member == undefined" color="red" icon="mdi-close-circle"/>
            <v-icon v-else color="green" icon="mdi-checkbox-marked-circle"/>
          </td>
          <td>{{ r.member?.firstname }}</td>
          <td>{{ r.member?.lastname }}</td>
          <td>
            <NuxtLink v-if="r.member != undefined" :to="'/edit/'+r.member?.id">
              <v-icon density="compact" icon="mdi-account-edit"/>
            </NuxtLink>
            <NuxtLink v-else :to="'/add?room='+r.roomNr">
              <v-icon density="compact" icon="mdi-account-plus"/>
            </NuxtLink>
          </td>
        </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>
</template>
<script lang="ts" setup>
  import {fetchRooms} from "~/composables/fetch";
  import {roomFiler} from "~/utils/constants";

  const emits = defineEmits(['logout'])

  const rooms = ref<Room[]>([])
  
  const filters = ref<RoomFilters>({
    occupied: null,
    block: [],
  })


  onMounted(() => {
    nextTick(refresh)
  })

  async function refresh() {
    const {data, error} = await fetchRooms(filters.value);
    if (error.value == null) {
      rooms.value = data.value
      return
    }
    if (error.value.statusCode === 403) {
      emits('logout')
      navigateTo('/login')
      return
    }
  }

</script>