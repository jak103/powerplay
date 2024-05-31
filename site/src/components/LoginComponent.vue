<template>
    <q-card class="my-card">
      <q-card-section class="bg-primary text-white">
        <div class="text-h6">Login</div>
      </q-card-section>

      <q-separator />

      <q-card-section>

        <q-input class="q-mb-xl" filled v-model="email" label="Email" />
        <q-input filled v-model="password" label="Password" type="password" />
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat @click="handleLogin">Login</q-btn>
        <q-btn flat @click="handleCancel">Cancel</q-btn>
      </q-card-actions>
    </q-card>
</template>

<script setup lang="ts">
import { defineComponent, ref } from 'vue'
import { api } from 'boot/axios'

const email = ref('')
const password = ref('')

async function handleLogin() {
  const res = await api.post('http://localhost:9002/api/v1/auth', {email: email.value, password: password.value})

  console.log(res);
  // Test
}

function handleCancel() {
  alert('Handling cancel')
}


defineComponent({
  name: 'LoginComponent'
})

</script>
