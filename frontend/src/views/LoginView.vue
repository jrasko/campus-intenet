<template>
  <v-form>
    <v-row justify="center">
      <h1>Login</h1>
    </v-row>
    <v-row justify="center">
      <v-col lg="6" sm="8">
        <v-text-field v-model="this.credentials.username" label="Nutzername" />
        <v-text-field
          v-model="this.credentials.password"
          label="Passwort"
          type="password"
          @keyup.enter="login"
        />
        <v-btn @click="login">Login</v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
import { login } from '@/axios'

export default {
  data() {
    return {
      credentials: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    login() {
      login(this.credentials)
        .then((r) => {
          localStorage.setItem('jwt', r.data.token)
          localStorage.setItem('role', r.data.role)
          this.$router.push('/')
        })
        .catch((e) => {
          if (e.response.status === 403) {
            window.location.href = 'https://youtu.be/dQw4w9WgXcQ'
          } else {
            console.log(e)
          }
        })
    }
  }
}
</script>
<style scoped></style>
