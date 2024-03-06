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
  const emit = defineEmits(['login'])

  const failure = ref(false)
  const credentials = ref<Credentials>({
    username: '',
    password: ''
  })


  async function login() {
    try {
      const login = await <any>loginUser(credentials.value)
      localStorage.setItem('jwt', login.token)
      localStorage.setItem('role', login.role)
      emit('login')
      navigateTo('/')
    } catch (error: any) {
      console.log(error)
      if (error.status === 403) {
        window.location.href = 'https://youtu.be/dQw4w9WgXcQ'
        return
      }
      failure.value = true
    }
  }
</script>
