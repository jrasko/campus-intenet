<template>
  <v-snackbar v-model="modal.failure" :timeout="3000" color="error"> {{ modal.errorMessage }}</v-snackbar>
  <AddEditForm @submit="submit" :prefetch-id="<string>route.params.id"/>

</template>

<script lang="ts" setup>
  const route = useRoute()
  const modal = ref({
    failure: false,
    errorMessage: ''
  })

  async function submit(person: MemberInput) {
    const {error} = await updateConfig(person)
    if (error.value == null){
      navigateTo('/')
      return
    }

    modal.value.failure = true
    modal.value.errorMessage = error.value.data
    console.log(error)
  }
</script>
