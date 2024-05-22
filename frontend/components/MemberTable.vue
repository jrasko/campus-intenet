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
    <tr v-for="p in props.people">
      <td>
        <v-icon
          :color="p.dhcpConfig.disabled ? 'orange' : 'green'"
          icon="mdi-circle-medium"
          @click="swapNetworkActivation(p)"
        />
      </td>
      <td>
        <v-icon
          v-if="p.hasPaid"
          color="green"
          icon="mdi-checkbox-marked-circle"
          @click="swapPayment(p)"
        />
        <v-icon v-else color="red" icon="mdi-close-circle" @click="swapPayment(p)"/>
      </td>
      <td v-for="c in columns">
        {{ getValue(p, c) }}
      </td>
      <td>
        <v-row align="center" justify="center">
          <v-col cols="1">
            <NuxtLink :to="'/members/edit/' + p.id">
              <v-btn density="compact" icon="mdi-square-edit-outline"/>
            </NuxtLink>
          </v-col>
          <v-col cols="1">
            <v-btn density="compact" icon="mdi-delete" @click="deleteUser(p)"/>
          </v-col>
        </v-row>
      </td>
    </tr>
    </tbody>
  </v-table>
</template>
<script lang="ts" setup>
  import {tableData} from "~/utils/constants";
  import {toggleNetworkActivation} from "~/utils/fetch_members";

  const props = defineProps<{
    people: MemberConfig[]
    columns: Column[]
  }>()

  const emit = defineEmits(['refresh'])

  const modals = ref({
    success: false,
    failure: false,
    errorMessage: '',
  })

  async function swapPayment(p: MemberConfig) {
    try {
      await togglePayment(p.id)
      modals.value.success = true
      emit('refresh')
    } catch (e) {
      handleError(e)
    }
  }
  
  async function swapNetworkActivation(p: MemberConfig) {
    try {
      await toggleNetworkActivation(p.dhcpConfig.id)
      modals.value.success = true
      emit('refresh')
    } catch (e) {
      handleError(e)
    }
  }
  
  async function deleteUser(u: MemberConfig) {
    if (confirm('Wirklich löschen?')) {
      try {
        await deleteMemberConfigFor(u.id)
        modals.value.success = true
        emit('refresh')
      } catch (e) {
        handleError(e)
      }
    }
  }

  function getValue(p: MemberConfig, c: Column): string | undefined {
    const selector = tableData[c].field.split('.')
    let value: any = p
    for (let s of selector) {
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