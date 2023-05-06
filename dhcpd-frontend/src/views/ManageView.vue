<script setup>
</script>
<template>
    <v-table>
        <thead>
        <tr>
            <th>Zahlung</th>
            <th>Vorname</th>
            <th>Nachname</th>
            <th>MAC</th>
            <th>WG</th>
            <th>Zimmer-Nr</th>
            <th>Telefonnummer</th>
            <th>E-Mail</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="person in this.people">
            <td v-if="person.hasPaid">
                <v-icon icon="mdi-checkbox-marked-circle" color="green"/>
            </td>
            <td v-else>
                <v-icon icon="mdi-close-circle" color="red"/>
            </td>
            <td>{{ person.firstname }}</td>
            <td>{{ person.lastname }}</td>
            <td>{{ person.mac }}</td>
            <td>{{ person.wg }}</td>
            <td>{{ person.roomNr }}</td>
            <td>{{ person.phone }}</td>
            <td>{{ person.email }}</td>
        </tr>
        </tbody>
    </v-table>
</template>
<script>
import {getConfigs} from "@/utils/axios";

export default {
    data() {
        return {
            people: []
        }
    },
    mounted() {
        getConfigs()
            .then(data => {
                this.people = data
            })
    }
}
</script>


<style scoped>

</style>