<template>
  <v-form @submit.prevent="">
    <v-row>
      <v-col cols="8">
        <v-text-field
          v-model="server.name"
          label="Name"
        />
      </v-col>
      <v-col cols="4">
        <v-switch v-model="server.disabled" color="green" label="Deaktiviert"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="6">
        <v-text-field
          v-model="server.mac"
          :maxlength="17"
          label="MAC-Adresse"
          @input="updateMac"
        />
      </v-col>
      <v-col cols="6">
        <v-text-field
          v-model="server.ip"
          label="IP Addresse"
          persistent-hint
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <NuxtLink to="/">
          <v-btn color="red">Abbrechen</v-btn>
        </NuxtLink>
      </v-col>
      <v-col>
        <v-btn color="green" @click="$emit('submit',server)">Speichern</v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script lang="ts" setup>
  import {fetchServer} from "~/utils/fetch_netconfig";

  const props = defineProps<{ prefetchId?: string, room?: string }>()

  const server = ref<Server>({
    name: '',
    id: 0,
    mac: '',
    ip: '',
    disabled: false,
  })

  onMounted(() => nextTick(() => {
    if (props.prefetchId != undefined) {
      prefetchForID(<string>props.prefetchId)
    }
  }))

  function updateMac() {
    server.value.mac = formatMac(server.value.mac)
  }

  async function prefetchForID(id: string) {
    try {
      server.value = await fetchServer(id)
    } catch (e) {
      console.log(e)
    }
  }
</script>
