<template>
  <v-app-bar density="compact" justify="space-between">
    <v-icon/>
    <v-btn v-if="props.loggedIn" append-icon="mdi-bed" @click="()=> {navigateTo('/')}">Räume</v-btn>
    <v-btn v-if="props.loggedIn && props.role == roles.admin" append-icon="mdi-server" @click="navigateTo('/servers')">
      Server
    </v-btn>
    <v-btn v-if="props.loggedIn" append-icon="mdi-sheep" @click="navigateTo('/shame')">Shamelist</v-btn>
    <v-spacer/>
    <v-btn prepend-icon="mdi-theme-light-dark" @click="toggleTheme">Lichtschalter</v-btn>
    <v-btn v-if="props.loggedIn" append-icon="mdi-logout" color="red" @click="$emit('logout')">Logout</v-btn>
  </v-app-bar>
</template>
<script lang="ts" setup>
  import {useTheme} from "vuetify";
  
  const theme = useTheme()
  
  const props = defineProps<{ loggedIn: boolean, role: string }>()
  
  function toggleTheme() {
    theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  }
</script>

<style scoped>

</style>