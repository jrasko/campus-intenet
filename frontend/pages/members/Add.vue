<template>
  <v-snackbar v-model="modal.failure" :timeout="3000" color="error">{{ modal.errorMessage }}</v-snackbar>
  <AddEditMemberForm :room="<string>route.query.room" @submit="submit"/>
</template>

<script lang="ts" setup>
  import {createMemberConfig} from "~/utils/fetch_members";

  const route = useRoute()

  const modal = ref({
    errorMessage: '',
    failure: false
  })

  async function submit(person: InputMember) {
    try {
      await createMemberConfig(person)
      navigateTo('/')
    } catch (error: any) {
      modal.value.failure = true
      modal.value.errorMessage = error.data
      console.log(error)
    }
  }
</script>


<style scoped></style>
