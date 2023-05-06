<template>
    <v-snackbar
            color="success"
            :timeout="2000"
            v-model="success"
    >
        Erfolg!
    </v-snackbar>
    <v-snackbar
            color="error"
            :timeout="2000"
            v-model="failure"
    >
        Fehler!
    </v-snackbar>
    <v-form
            @submit.prevent=""
    >
        <v-container>
            <v-row>
                <v-col>
                    <v-text-field
                            label="Vorname"
                            v-model="config.firstname"
                    />
                </v-col>
                <v-col>
                    <v-text-field
                            label="Nachname"
                            v-model="config.lastname"
                    />
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-text-field
                            label="MAC-Adresse"
                            v-model="config.mac"
                    />
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-text-field
                            label="WG"
                            v-model="config.wg"
                    />
                </v-col>
                <v-col>
                    <v-text-field
                            label="Zimmernummer"
                            v-model="config.roomNr"
                    />
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-text-field
                            label="Email"
                            v-model="config.email"
                    />
                </v-col>
                <v-col>
                    <v-text-field
                            label="Telefonnummer"
                            v-model="config.phone"
                    />
                </v-col>
                <v-col>
                    <v-switch
                            color="green"
                            label="Bezahlt"
                            v-model="config.hasPaid"
                    />
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-btn
                            color="red"
                    >Abbrechen
                    </v-btn>
                </v-col>
                <v-col>
                    <v-btn
                            color="blue"
                            type="submit"
                    >Speichern & Next
                    </v-btn>
                </v-col>
                <v-col>
                    <v-btn
                            color="green"
                            @click="this.submit"
                    >Speichern
                    </v-btn>
                </v-col>
            </v-row>
        </v-container>
    </v-form>
</template>

<script>

import {updateConfig} from "@/utils/axios";

export default {
    data: () => ({
        success: false,
        failure: false,
        config: {
            firstname: '',
            lastname: '',
            mac: '',
            wg: '',
            roomNr: '',
            phone: '',
            email: '',
            hasPaid: false
        }
    }),
    methods: {
        submit() {
            updateConfig(this.config)
                .then(() => {
                    this.success = true
                })
                .catch(e => {
                    this.failure = true
                    console.log(e)
                })
        }
    }
}
</script>
<style scoped>

</style>