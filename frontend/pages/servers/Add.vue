<template>
  <v-snackbar v-model="modal.failure" :timeout="3000" color="error">{{ modal.errorMessage }}</v-snackbar>
  <AddEditServerForm :room="<string>route.query.room" @submit="submit"/>
</template>

<script lang="ts" setup>

  const route = useRoute()

  const modal = ref({
    errorMessage: '',
    failure: false
  })


  async function submit(server: Server) {
    try {
      await createServer(server)
      navigateTo('/servers')
    } catch (e: any) {
      modal.value.failure = true
      modal.value.errorMessage = e.value.data
      console.log(e.value)
    }
  }
</script>

<style scoped></style>
