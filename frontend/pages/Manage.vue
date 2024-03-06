<template>
  <v-snackbar v-model="modals.success" :timeout="2000" color="success"> Erfolg!</v-snackbar>
  <v-snackbar v-model="modals.failure" :timeout="3000" color="error"> {{ modals.errorMessage }}</v-snackbar>
  <!--  are retries the way to handle this problem?
  <v-row>
      <v-alert v-model="modals.warning" closable type="warning" variant="tonal">
        <v-alert-title>
          inconsistent user-list.json file
          <v-spacer/>
          <v-btn append-icon="mdi-reload-alert" density="compact" variant="text" @click="writeDhcp">
            Regenerate File
          </v-btn>
        </v-alert-title>
      </v-alert>
    </v-row>
    -->
  <v-row>
    <v-col>
      <v-btn prepend-icon="mdi-credit-card-refresh" @click="resetPaymentsForAll">
        Zahlungen zurücksetzen
      </v-btn>
    </v-col>
    <v-col>
      <a :href="'mailto:wohnheimsprecher@scj.fh-aachen.de?bcc=' + copyEmails()">
        <v-btn prepend-icon="mdi-content-copy" @click="copyEmails">Emails kopieren</v-btn>
      </a>
    </v-col>
    <v-spacer />
    <v-spacer />
    <v-spacer />
    <v-spacer />
  </v-row>
  <v-row align="baseline" justify="center">
    <v-col cols="6" md="2">
      <v-select
        v-model="filters.payment"
        :items="manageFilter.payment"
        append-inner-icon="mdi-filter"
        hide-details
        item-title="header"
        item-value="value"
        label="Bezahlung"
        variant="underlined"
        @update:modelValue="refresh()"
      />
    </v-col>
    <v-col cols="6" md="2">
      <v-select
        v-model="filters.disabled"
        :items="manageFilter.disabled"
        append-inner-icon="mdi-filter"
        hide-details
        item-title="header"
        item-value="value"
        label="Status"
        variant="underlined"
        @update:modelValue="refresh()"
      />
    </v-col>
    <v-col cols="12" md="4" @input="refresh">
      <v-text-field
        v-model="filters.search"
        append-inner-icon="mdi-magnify"
        clearable
        hide-details
        label="Suche"
        variant="underlined"
        @click:clear="refresh"
      />
    </v-col>
    <v-col cols="12" md="4">
      <v-select
        v-model="columns"
        :items="Object.values(tableData)"
        item-title="header"
        item-value="key"
        label="Spalten"
        multiple
        variant="underlined"
        @update:modelValue="changeColumns"
      />
    </v-col>
  </v-row>
  <v-row>
    <v-col cols="12">
      <MemberTable
        :columns="columns"
        :people="members"
        @refresh="refresh"
      />
    </v-col>
  </v-row>
</template>

<script lang="ts" setup>
  import {manageFilter, tableData} from "~/utils/constants";

  const members = ref<MemberConfig[]>([])
  const modals = ref({
    success: false,
    failure: false,
    warning: false,
    errorMessage: '',
  })

  const filters = ref<ManageFilters>({
    search: '',
    payment: null,
    disabled: null,
  })

  const columns = ref<Column[]>(['firstname', 'lastname', 'wg', 'roomNr', 'comment'])

  const emit = defineEmits(['logout'])

  onMounted(() => {
    let storedColumns = localStorage.getItem('columns');
    if (storedColumns != null) {
      columns.value = <Column[]>storedColumns.split(',')
    }
    refresh()
  })

  async function refresh() {
    try {
      members.value = await getMemberConfigs(filters.value)
    } catch (error: any) {
      if (error.statusCode === 403) {
        emit('logout')
        return
      }
      handleError(error)
    }
  }

  function copyEmails() {
    let mails = ''
    for (const p of members.value) {
      mails += p.email + ','
    }
    return mails
  }

  async function resetPaymentsForAll() {
    let answer = prompt('Zahlungen zurücksetzen?\n\nSchreibe "reset", wenn du dir sicher bist!',);
    if (answer == 'reset') {
      try {
        await resetPayments()
        modals.value.success = true
        await refresh()
      } catch (e) {
        handleError(e)
      }
    } else {
      modals.value.errorMessage = 'Abbruch!'
      modals.value.failure = true
    }
  }

  function handleError(error: any) {
    console.log(error)
    if (error.status === 403) {
      modals.value.errorMessage = 'no permissions for that'
      modals.value.failure = true
    } else {
      modals.value.errorMessage = error.data
      modals.value.failure = true
    }
  }

  function changeColumns() {
    localStorage.setItem('columns', columns.value.toString())
  }
</script>
