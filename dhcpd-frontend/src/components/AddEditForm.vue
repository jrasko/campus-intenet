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
        <v-text-field v-model="person.mac" label="MAC-Adresse" />
      </v-col>
      <v-col>
        <v-text-field
          v-model="person.ip"
          label="IP Addresse"
          :disabled="disableIp"
          persistent-hint
          hint="autogeneriert wenn leer"
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
