<template>
  <v-snackbar v-model="success" :timeout="2000" color="success"> Erfolg!</v-snackbar>
  <v-snackbar v-model="failure" :timeout="3000" color="error"> {{ errorMessage }}</v-snackbar>
  <v-row>
    <v-alert v-model="warning" closable type="warning" variant="tonal">
      <v-alert-title>
        inconsistent user-list.json file
        <v-spacer />
        <v-btn
          append-icon="mdi-reload-alert"
          density="compact"
          variant="text"
          @click="writeDhcp"
        >
          Regenerate File
        </v-btn>
      </v-alert-title>
    </v-alert>
  </v-row>
  <v-row>
    <v-col>
      <RouterLink to="/add">
        <v-btn prepend-icon="mdi-account-plus"> Person hinzufügen</v-btn>
      </RouterLink>
    </v-col>
    <v-col>
      <v-btn prepend-icon="mdi-credit-card-refresh" @click="resetPaymentsForAll">
        Zahlungen zurücksetzen
      </v-btn>
    </v-col>
    <v-col>
      <a :href="'mailto:' + copyEmails()"><v-btn prepend-icon="mdi-content-copy" @click="copyEmails">Emails kopieren</v-btn></a>
    </v-col>
  </v-row>
  <v-row justify="center" align="baseline">
    <v-col cols="0" md="4"></v-col>
    <v-col cols="12" md="4" @input="refresh">
      <v-text-field
        v-model="search"
        append-inner-icon="mdi-magnify"
        clearable
        hide-details
        label="Suche"
        variant="underlined"
      />
    </v-col>
    <v-col cols="12" md="4">
      <v-select
        v-model="columns"
        :items="Object.values(tableData)"
        item-title="header"
        item-value="field"
        multiple
        label="Spalten"
        variant="underlined"
      />
    </v-col>
  </v-row>
  <v-row>
    <v-col cols="12">
      <v-table hover>
        <thead>
          <tr>
            <th v-for="c in columns">
              <div v-if="tableData[c].field === sortKey">
                <b>{{ tableData[c].header }}</b>
              </div>              
              <div v-else @click="sort(tableData[c].field)">
                {{ tableData[c].header }}
              </div>
            </th>
            <th></th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in people">
            <td v-for="c in columns">
              <div v-if="tableData[c].kind === 'text'">
                {{ p[c] }}
              </div>
              <div v-else-if="tableData[c].kind === 'icon'">
                <v-icon v-if="p[c]" :color="tableData[c].trueColor" :icon="tableData[c].trueIcon" />
                <v-icon v-else :color="tableData[c].falseColor" :icon="tableData[c].falseIcon" />
              </div>
            </td>
            <td>
              <v-row align="center" justify="center">
                <v-col cols="1">
                  <RouterLink :to="'/edit/' + p.id">
                    <v-btn density="compact" icon="mdi-square-edit-outline" />
                  </RouterLink>
                </v-col>
                <v-col cols="1">
                  <v-btn density="compact" icon="mdi-delete" @click="deleteUser(p)" />
                </v-col>
              </v-row>
            </td>
            <td></td>
          </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>
</template>

<script>
import { deleteConfigFor, getConfigs, resetPayments, updateDhcp } from '@/axios'

export default {
  data() {
    return {
      people: [],
      success: false,
      failure: false,
      warning: false,
      errorMessage: '',
      search: '',
      columns: ['disabled', 'hasPaid', 'firstname', 'lastname', 'wg', 'roomNr', 'comment'],
      sortKey: 'roomNr'
    }
  },
  mounted() {
    this.refresh()
  },
  methods: {
    refresh() {
      getConfigs(this.search)
        .then((resp) => {
          this.people = resp.data
          this.sort(this.sortKey)
          if (resp.status === 210) {
            this.warning = true
          }
        })
        .catch((e) => {
          if (e.response.status === 403) {
            this.$router.push('/login')
          }
          console.log(e)
        })
    },
    copyEmails() {
      let mails = ''
      for (const p of this.people) {
        mails += p.email + ','
      }
      return mails
    },
    deleteUser(p) {
      if (confirm('Wirklich löschen?')) {
        deleteConfigFor(p.id)
          .then(() => {
            this.success = true
            this.refresh()
          })
          .catch((e) => {
            this.errorMessage = e.response.data
            this.failure = true
          })
      }
    },
    sort(field){
      let sortKey = field
      this.people.sort(function (a, b) {
        let aVal = a[sortKey]
        let bVal = b[sortKey]
        if (typeof aVal === 'string'){
          return aVal.localeCompare(bVal)
        }
        if (typeof aVal === 'boolean'){
          return aVal === bVal?0:aVal?-1:1
        }
        return 0
      })
      this.sortKey = sortKey
    },
    resetPaymentsForAll() {
      if (confirm('Zahlungen zurücksetzen?')) {
        resetPayments()
          .then(() => {
            this.success = true
            this.refresh()
          })
          .catch((e) => {
            this.failure = true
            this.errorMessage = e.response.data
          })
      }
    },
    writeDhcp() {
      updateDhcp()
        .then(() => {
          this.success = true
          this.refresh()
        })
        .catch((e) => {
          this.failure = true
          this.errorMessage = e.response.data
        })
    }
  }
}
</script>
<script setup>
const tableData = {
  firstname: {
    header: 'Vorname',
    field: 'firstname',
    kind: 'text'
  },
  lastname: {
    header: 'Nachname',
    field: 'lastname',
    kind: 'text'
  },
  mac: {
    header: 'MAC',
    field: 'mac',
    kind: 'text'
  },
  roomNr: {
    header: 'Zimmer-Nr.',
    field: 'roomNr',
    kind: 'text'
  },
  wg: {
    header: 'WG',
    field: 'wg',
    kind: 'text'
  },
  email: {
    header: 'E-Mail',
    field: 'email',
    kind: 'text'
  },
  phone: {
    header: 'Telefonnr.',
    field: 'phone',
    kind: 'text'
  },
  ip: {
    header: 'IP',
    field: 'ip',
    kind: 'text'
  },
  manufacturer: {
    header: 'Hersteller',
    field: 'manufacturer',
    kind: 'text'
  },
  comment: {
    header: 'Kommentar',
    field: 'comment',
    kind: 'text'
  },
  hasPaid: {
    header: 'Zahlung',
    kind: 'icon',
    field: 'hasPaid',
    trueIcon: 'mdi-checkbox-marked-circle',
    trueColor: 'green',
    falseIcon: 'mdi-close-circle',
    falseColor: 'red'
  },
  disabled: {
    header: 'Aktiv',
    kind: 'icon',
    field: 'disabled',
    falseIcon: 'mdi-circle-medium',
    falseColor: 'green',
    trueIcon: 'mdi-circle-medium',
    trueColor: 'orange'
  }
}
</script>

<style scoped></style>
