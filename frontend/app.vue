<template>
  <v-app>
    <v-container fluid>
      <AppBar :logged-in="loggedIn" @logout="logout"/>
      <v-main>
        <NuxtPage @login="setLogin" @logout="logout"/>
      </v-main>
    </v-container>
  </v-app>
</template>

<script lang="ts" setup>
  const loggedIn = ref(false)

  onMounted(() => setLogin())
  onUpdated(() => setLogin())

  function setLogin() {
    loggedIn.value = localStorage.getItem('jwt') != null
  }

  function logout() {
    localStorage.removeItem('jwt')
    loggedIn.value = false
    navigateTo('/login')
  }
</script>
<style>
  a {
    text-decoration: none;
    color: inherit;
  }
</style>