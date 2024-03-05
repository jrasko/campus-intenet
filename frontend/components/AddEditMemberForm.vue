<template>
  <v-form @submit.prevent="">
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field v-model="member.firstname" label="Vorname"/>
      </v-col>
      <v-col cols="12" sm="6">
        <v-text-field v-model="member.lastname" label="Nachname"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="member.dhcpConfig.mac"
          :maxlength="17"
          label="MAC-Adresse"
          @input="updateMac"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="member.dhcpConfig.ip"
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
          v-model="member.roomNr"
          :item-props="roomMapper"
          :items="availableRooms"
          label="Zimmernummer"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="4">
        <v-text-field v-model="member.email" label="Email"/>
      </v-col>
      <v-col cols="12" sm="4">
        <v-text-field v-model="member.phone" label="Telefonnummer"/>
      </v-col>
      <v-col cols="6" sm="2">
        <v-switch v-model="member.hasPaid" color="green" label="Bezahlt"/>
      </v-col>
      <v-col cols="6" sm="2">
        <v-switch v-model="member.dhcpConfig.disabled" color="green" label="Deaktiviert"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-text-field v-model="member.comment" label="Kommentar"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <NuxtLink to="/">
          <v-btn color="red"> Abbrechen</v-btn>
        </NuxtLink>
      </v-col>
      <v-col>
        <v-btn color="green" @click="$emit('submit',member)">Speichern</v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script lang="ts" setup>
  import {toInputMember} from "~/utils/utils";

  const props = defineProps<{ prefetchId?: string, room?: string }>()

  const member = ref<InputMember>({
    id: 0,
    firstname: '',
    lastname: '',
    roomNr: '',
    dhcpConfig: {
      id: 0,
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
      member.value.roomNr = props.room
    }
    fetchAvailableRooms()
  }))

  function updateMac() {
    member.value.dhcpConfig.mac = formatMac(member.value.dhcpConfig.mac)
  }

  async function prefetchForID(id: string) {
    try {
      const data: MemberConfig = await getMemberConfigFor(id)
      member.value = toInputMember(data)
      availableRooms.value.unshift({
        roomNr: data.room.roomNr,
        wg: data.room.wg,
        block: ''
      })
    } catch (error) {
      console.log(error)
    }
  }

  async function fetchAvailableRooms() {
    try {
      const data : Room[] = await fetchRooms({occupied: false, block: []})
      availableRooms.value = availableRooms.value.concat(data)
    } catch (error) {
      console.error(error)
    }
  }

  function roomMapper(item: Room) {
    return {
      title: item.roomNr,
      value: item.roomNr,
      subtitle: item.wg,
    }
  }
</script>
