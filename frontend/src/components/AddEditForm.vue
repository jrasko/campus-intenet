<template>
  <v-form @submit.prevent="">
    <v-row>
      <v-col>
        <v-text-field v-model="person.firstname" label="Vorname" />
      </v-col>
      <v-col>
        <v-text-field v-model="person.lastname" label="Nachname" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field
          v-model="person.mac"
          :maxlength="17"
          label="MAC-Adresse"
          @input="this.$emit('macUpdate')"
        />
      </v-col>
      <v-col>
        <v-text-field
          v-model="person.ip"
          :disabled="disableIp"
          hint="autogeneriert wenn leer"
          label="IP Addresse"
          persistent-hint
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field v-model="person.wg" label="WG" />
      </v-col>
      <v-col>
        <v-text-field v-model="person.roomNr" label="Zimmernummer" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field v-model="person.email" label="Email" />
      </v-col>
      <v-col>
        <v-text-field v-model="person.phone" label="Telefonnummer" />
      </v-col>
      <v-col>
        <v-switch v-model="person.hasPaid" color="green" label="Bezahlt" />
      </v-col>
      <v-col>
        <v-switch v-model="person.disabled" color="green" label="Deaktiviert" />
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
export default {
  props: {
    disableIp: false,
    person: {}
  }
}
</script>
<style scoped></style>
