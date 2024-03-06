<template>
  <v-snackbar v-model="modal.failure" :timeout="3000" color="error"> {{ modal.errorMessage }}</v-snackbar>
  <AddEditMemberForm @submit="submit" :prefetch-id="<string>route.params.id"/>

</template>

<script lang="ts" setup>
  const route = useRoute()
  const modal = ref({
    failure: false,
    errorMessage: ''
  })

  async function submit(person: InputMember) {
    try {
      await updateMemberConfig(person)
      navigateTo('/manage')
    } catch (error: any) {
      modal.value.failure = true
      modal.value.errorMessage = error.data
      console.log(error)
    }
  }
</script>
