<template>
  <v-snackbar v-model="failure" :timeout="3000" color="error">something went wrong</v-snackbar>

  <v-form>
    <v-row justify="center">
      <h1>Login</h1>
    </v-row>
    <v-row justify="center">
      <v-col lg="6" sm="8">
        <v-text-field v-model="credentials.username" label="Nutzername"/>
        <v-text-field
          v-model="credentials.password"
          label="Passwort"
          type="password"
          @keyup.enter="login"
        />
        <v-btn @click="login">Login</v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script lang="ts" setup>
  import {loginUser} from "~/composables/fetch";
  const emit = defineEmits(['login'])

  const failure = ref(false)
  const credentials = ref<Credentials>({
    username: '',
    password: ''
  })

  async function login() {
    let {data, error} = await loginUser(credentials.value)
    if (data.value != null) {
      localStorage.setItem('jwt', data.value.token)
      localStorage.setItem('role', data.value.role)
      emit('login')
      navigateTo('/')
      return
    }
    if (error.value.statusCode === 403) {
      window.location.href = 'https://youtu.be/dQw4w9WgXcQ'
      return
    }
    failure.value = true
    console.log(error.value)
  }
</script>
<style scoped></style>
