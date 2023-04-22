<template>
<h1>Welcome to Product Page</h1>
<!-- Test for going back -->
<router-link to="/">Back</router-link>
<p>User id: {{ $route.params.id }}</p>

</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useRoute } from 'vue-router';

export default defineComponent({
    mounted(){
        console.log(useRoute().params.id);
    }
})
</script>

<style>
</style>