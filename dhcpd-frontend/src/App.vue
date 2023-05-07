<template>
  <v-container fluid>
    <v-row justify="center">
      <v-col cols="10">
        <v-btn prepend-icon="mdi-theme-light-dark" @click="toggleTheme"> Lichtschalter </v-btn>
      </v-col>
    </v-row>
  </v-container>
  <RouterView />
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
  }
}
</script>

<style scoped></style>
