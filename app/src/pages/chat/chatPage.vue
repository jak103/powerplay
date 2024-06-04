<template>
    <q-page class="row items-center justify-evenly">
        <CardComponent :key="key" :type="type" :name="test" :unread="missed" :image="image" :link="link"/>
    </q-page>
</template>

<script setup lang="ts">
    import { defineComponent } from 'vue';
    import { QPage } from 'quasar';
    import CardComponent from '../chat/CardComponent.vue';

    defineComponent({
        name: 'ChatPage',
        components: {
            CardComponent,
        },
    });

    const test = 'atest';
    const missed = 20;
    const image = null
    const key = 'exampleKey';
    const type = 'team';
    const link = '/'
</script>