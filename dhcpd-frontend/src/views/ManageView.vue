<script setup>
</script>
<template>
    <v-table>
        <thead>
        <tr>
            <th>Bezahlt?</th>
            <th>Name</th>
            <th>MAC</th>
            <th>WG</th>
            <th>Zimmer-Nr</th>
            <th>Telefonnummer</th>
            <th>E-Mail</th>
        </tr>
        </thead>
        <tbody>
        <tr
                v-for="person in this.people"
        >
            <td v-if="person.hasPaid">
                <v-icon icon="mdi-checkbox-marked-circle"/>
            </td>
            <td v-else>
                <v-icon icon="mdi-close-circle"/>
            </td>
            <td>{{ person.name }}</td>
            <td>{{ person.mac }}</td>
            <td>{{ person.wg }}</td>
            <td>{{ person.phone }}</td>
            <td>{{ person.email }}</td>
        </tr>
        </tbody>
    </v-table>
</template>
<script>
export default {
    data: () => {
        return {
            people: [
                {
                    name: "Jannik",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "31-13",
                    hasPaid: true,
                    wg: "18a",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                },
                {
                    name: "other dude",
                    mac: "ab:cd:12:34:56:78",
                    roomNr: "12-34",
                    hasPaid: false,
                    wg: "29",
                    email: "test@test.de",
                    phone: "01234 5678910"
                }
            ]
        }
    }
}
</script>


<style scoped>

</style>