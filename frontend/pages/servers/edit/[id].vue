<template>
  <v-snackbar v-model="modal.failure" :timeout="3000" color="error"> {{ modal.errorMessage }}</v-snackbar>
  <AddEditServerForm @submit="submit" :prefetch-id="<string>route.params.id"/>

</template>

<script lang="ts" setup>

  import {updateConfig} from "~/utils/fetch";

  const route = useRoute()
  const modal = ref({
    failure: false,
    errorMessage: ''
  })

  async function submit(server: InputMember) {
    const {error} = await updateConfig(server)
    if (error.value == null){
      navigateTo('/')
      return
    }

    modal.value.failure = true
    modal.value.errorMessage = error.value.data
    console.log(error)
  }
</script>
