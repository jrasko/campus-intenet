<template>
  <v-container fluid>
    <v-row>
      <v-spacer />
      <v-col cols="12" lg="10">
        <v-row justify="space-between">
          <v-col>
            <v-btn prepend-icon="mdi-theme-light-dark" @click="toggleTheme"> Lichtschalter</v-btn>
          </v-col>
          <v-spacer />
          <v-col cols="3" lg="1" md="2">
            <v-btn v-if="isLoggedIn()" color="red" @click="logout">Logout</v-btn>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" />
        </v-row>
        <RouterView :key="$route.path" />
      </v-col>
      <v-spacer />
    </v-row>
  </v-container>
</template>
<script>
import { useTheme } from 'vuetify'
import { isLoggedIn } from '@/utils'

export default {
  setup() {
    const theme = useTheme()
    return {
      theme,
      toggleTheme: () =>
        (theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark')
    }
  },
  methods: {
    isLoggedIn,
    logout() {
      localStorage.removeItem('jwt')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped></style>
