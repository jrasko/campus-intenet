<template>
  <v-snackbar v-model="modals.success" :timeout="2000" color="success"> Erfolg!</v-snackbar>
  <v-snackbar v-model="modals.failure" :timeout="3000" color="error">{{ modals.errorMessage }}</v-snackbar>
  <v-table hover>
    <thead>
    <tr>
      <th>Aktiv</th>
      <th>Zahlung</th>
      <th v-for="c in columns">
        {{ tableData[c].header }}
      </th>
      <th/>
    </tr>
    </thead>
    <tbody>
    <tr v-for="r in props.rooms">
      <td v-if="r.member">
        <v-icon
          :color="r.member?.dhcpConfig.disabled ? 'orange' : 'green'"
          icon="mdi-circle-medium"
          @click="swapNetworkActivation(r.member?.dhcpConfig.id)"
        />
      </td>
      <td v-else/>
      <td v-if="r.member">
        <v-icon
          v-if="r.member?.hasPaid"
          color="green"
          icon="mdi-checkbox-marked-circle"
          @click="swapPayment(r.member?.id)"
        />
        <v-icon v-else color="red" icon="mdi-close-circle" @click="swapPayment(r.member?.id)"/>
      </td>
      <td v-else/>
      <td v-for="c in columns">
        <div v-if="tableData[c].kind === 'bool' && tableData[c].banNull && getValue(r,c) == null" />
        <div v-else-if="tableData[c].kind === 'bool'">
          <v-icon
            v-if="getValue(r, c)"
            color="green"
            icon="mdi-checkbox-marked-circle"
          />
          <v-icon v-else color="red" icon="mdi-close-circle"/>
        </div>
        <div v-else>
          {{ getValue(r, c) }}
        </div>
      </td>
      <td>
        <v-row v-if="r.member" justify="center">
          <v-col cols="1">
            <NuxtLink :to="'/members/edit/' + r.member.id">
              <v-btn density="compact" icon="mdi-square-edit-outline"/>
            </NuxtLink>
          </v-col>
          <v-col cols="1">
            <v-btn density="compact" icon="mdi-delete" @click="deleteUser(r.member.id)"/>
          </v-col>
        </v-row>
        <v-row v-else justify="center">
          <v-col cols="1">
            <NuxtLink :to="'/members/add?room='+r.roomNr">
              <v-icon density="compact" icon="mdi-account-plus"/>
            </NuxtLink>
          </v-col>
        </v-row>
      </td>
    </tr>
    </tbody>
  </v-table>
</template>
<script lang="ts" setup>
  import {tableData} from "~/utils/constants";
  import {toggleNetworkActivation, togglePayment, deleteMemberConfigFor} from "~/utils/fetch_members";

  const props = defineProps<{
    rooms: Room[]
    columns: Column[]
  }>()

  const emit = defineEmits(['refresh'])

  const modals = ref({
    success: false,
    failure: false,
    errorMessage: '',
  })

  async function swapPayment(id: number) {
    try {
      await togglePayment(id)
      modals.value.success = true
      emit('refresh')
    } catch (e) {
      handleError(e)
    }
  }
  
  async function swapNetworkActivation(id: number) {
    try {
      await toggleNetworkActivation(id)
      modals.value.success = true
      emit('refresh')
    } catch (e) {
      handleError(e)
    }
  }
  
  async function deleteUser(id: number) {
    if (confirm('Wirklich löschen?')) {
      try {
        await deleteMemberConfigFor(id)
        modals.value.success = true
        emit('refresh')
      } catch (e) {
        handleError(e)
      }
    }
  }

  function getValue(r: Room, c: Column): string | boolean | undefined {
    const selector = tableData[c].field.split('.')
    let value: any = r
    
    for (let s of selector) {
      if (value[s] == null) {
        return undefined
      }
      value = value[s]
    }

    if (tableData[c].kind === "date") {
      let dateTime = new Date(value)
      return dateTime.toLocaleString('de-DE', {dateStyle: 'short', timeStyle: 'short'})
    }

    return value
  }

  function handleError(error: any) {
    console.log(error)
    if (error.status === 403) {
      modals.value.errorMessage = 'no permissions for that'
    } else {
      modals.value.errorMessage = 'something went wrong'
    }
    modals.value.failure = true
  }
</script>