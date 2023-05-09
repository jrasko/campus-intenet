<template>
  <v-container fluid>
    <v-row>
      <v-col cols="1" />
      <v-col cols="2">
        <v-col>
          <v-btn prepend-icon="mdi-theme-light-dark" @click="toggleTheme"> Lichtschalter</v-btn>
        </v-col>
      </v-col>
      <v-col cols="7" />
      <v-col cols="1">
        <v-col>
          <v-btn @click="logout" color="red" v-if="this.$route.name !== 'login'">Logout</v-btn>
        </v-col>
      </v-col>
      <v-col cols="1" />
    </v-row>
    <v-row>
      <v-col cols="1" />
      <v-col cols="10">
        <RouterView />
      </v-col>
      <v-col cols="1" />
    </v-row>
  </v-container>
</template>
<script>
import { useTheme } from 'vuetify'
import { th } from 'vuetify/locale'

export default {
  computed: {
    th() {
      return th
    }
  },
  setup() {
    const theme = useTheme()
    return {
      theme,
      toggleTheme: () =>
        (theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark')
    }
  },
  methods: {
    logout() {
      localStorage.removeItem('jwt')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped></style>
