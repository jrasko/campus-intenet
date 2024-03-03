<template>
  <v-form @submit.prevent="">
    <v-row>
      <v-col>
        <v-text-field
          v-model="server.name"
          label="Name"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="server.mac"
          :maxlength="17"
          label="MAC-Adresse"
          @input="updateMac"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="server.ip"
          :readonly="true"
          clearable
          hint="autogeneriert wenn leer"
          label="IP Addresse"
          persistent-hint
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="6" sm="2">
        <v-switch v-model="server.disabled" color="green" label="Deaktiviert"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-text-field v-model="server.comment" label="Kommentar"/>
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

  const props = defineProps<{ prefetchId?: string, room?: string }>()

  const server = ref<Server>({
    name: '',
    id: 0,
    mac: '',
    ip: '',
    disabled: false,
    comment: ''
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
      const resp = await getConfigFor(id)
    } catch (e) {
      console.log(e)

    }
  }
</script>
