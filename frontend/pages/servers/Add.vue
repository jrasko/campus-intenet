<template>
  <v-snackbar v-model="modal.failure" :timeout="3000" color="error">{{ modal.errorMessage }}</v-snackbar>
  <AddEditServerForm :room="<string>route.query.room" @submit="submit"/>
</template>

<script lang="ts" setup>
  import {createConfig} from "~/utils/fetch";

  const route = useRoute()

  const modal = ref({
    errorMessage: '',
    failure: false
  })


  async function submit(person: InputMember) {
    const {error} = await createConfig(person)
    if (error.value == null) {
      navigateTo('/')
      return
    }

    modal.value.failure = true
    modal.value.errorMessage = error.value.data
    console.log(error.value)
    return
  }
</script>


<style scoped></style>
