<template>
  <v-form>
    <v-container>
      <v-row justify="center">
          <h1>Login</h1>
      </v-row>
      <v-row justify="center">
        <v-col lg="6" sm="8">
          <v-text-field v-model="this.credentials.username" label="Nutzername"/>
          <v-text-field v-model="this.credentials.password" type="password" label="Passwort"/>
          <v-btn @click="login">Login</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
    export default {
        data() {
            return {
                credentials: {
                    username: '',
                    password: '',
                }
            }
        },
        methods: {
            login() {
                let resp = login(this.credentials)
            }
        }
    }
</script>
<style scoped>

</style>