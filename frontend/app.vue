<template>
  <v-app>
    <v-container fluid>
      <AppBar :logged-in="loggedIn" :role="role" @logout="logout"/>
      <v-main>
        <NuxtPage @login="setLogin" @logout="logout"/>
      </v-main>
    </v-container>
    <v-spacer />
    <footer>
      <v-footer height="30" class="justify-end">2024 @Jannik Raskob</v-footer>
    </footer>
  </v-app>
</template>

<script lang="ts" setup>
  const loggedIn = ref(false)
  const role = ref("")
  
  onMounted(() => setLogin())
  onUpdated(() => setLogin())
  
  function setLogin() {
    loggedIn.value = localStorage.getItem('jwt') != null
    let r = localStorage.getItem("role")
    if (r != null) {
      role.value = r
    }
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