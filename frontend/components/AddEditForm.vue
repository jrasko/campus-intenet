<template>
  <v-form @submit.prevent="">
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field v-model="person.firstname" label="Vorname"/>
      </v-col>
      <v-col cols="12" sm="6">
        <v-text-field v-model="person.lastname" label="Nachname"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="person.dhcpConfig.mac"
          :maxlength="17"
          label="MAC-Adresse"
          @input="updateMac"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="person.dhcpConfig.ip"
          :readonly="true"
          clearable
          hint="autogeneriert wenn leer"
          label="IP Addresse"
          persistent-hint
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6">
        <v-select
          v-model="person.roomNr"
          :item-props="roomMapper"
          :items="availableRooms"
          label="Zimmernummer"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="4">
        <v-text-field v-model="person.email" label="Email"/>
      </v-col>
      <v-col cols="12" sm="4">
        <v-text-field v-model="person.phone" label="Telefonnummer"/>
      </v-col>
      <v-col cols="6" sm="2">
        <v-switch v-model="person.hasPaid" color="green" label="Bezahlt"/>
      </v-col>
      <v-col cols="6" sm="2">
        <v-switch v-model="person.dhcpConfig.disabled" color="green" label="Deaktiviert"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-text-field v-model="person.comment" label="Kommentar"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <NuxtLink to="/">
          <v-btn color="red"> Abbrechen</v-btn>
        </NuxtLink>
      </v-col>
      <v-col>
        <v-btn color="green" @click="$emit('submit',person)">Speichern</v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script lang="ts" setup>
  import {toInputMember} from "~/utils/utils";

  const props = defineProps<{ prefetchId?: string, room?: string }>()

  const person = ref<MemberInput>({
    id: 0,
    firstname: '',
    lastname: '',
    roomNr: '',
    dhcpConfig: {
      mac: '',
      ip: '',
      disabled: false
    },
    phone: '',
    email: '',
    hasPaid: false,
    comment: '',
  })

  const availableRooms = ref<Room[]>([])

  onMounted(() => nextTick(() => {
    if (props.prefetchId != undefined) {
      prefetchForID(<string>props.prefetchId)
    }
    if (props.room != undefined) {
      person.value.roomNr = props.room
    }
    fetchAvailableRooms()
  }))

  function updateMac() {
    person.value.dhcpConfig.mac = formatMac(person.value.dhcpConfig.mac)
  }

  async function prefetchForID(id: string) {
    const {data, error} = await getConfigFor(id)
    if (error.value == null) {
      person.value = toInputMember(data.value)
      availableRooms.value.unshift({roomNr: person.value.roomNr, block: '', wg: ''})
      return
    }
    console.log(error.value)
  }

  async function fetchAvailableRooms() {
    const {data, error} = await fetchRooms({occupied: false, block: []})
    if (error.value == null) {
      // append unoccupied rooms
      availableRooms.value = availableRooms.value.concat(data.value)
      return
    }
    console.log(error.value)
  }

  function roomMapper(item: Room | string) {
    if (typeof item == "string") {
      return {
        title: item,
      }
    }
    return {
      title: item.roomNr,
      subtitle: item.wg,
    }
  }
</script>

<style scoped></style>
