<template>
        <div class="product-page">
        <b-navbar class="header">
            <b-navbar-nav class="navigation-bar">
                <b-button v-b-modal.modal-cart class="navigation-button" href="#" variant="primary">
                    Cart
                </b-button>
                <b-button class="navigation-button">
                    Buy
                </b-button>
            </b-navbar-nav>
        </b-navbar>
</template>

<script lang="ts">

</script>

<style>
.header {
    display: flex;
    justify-content: end;
    align-items: center;
}
.navigation-button {
    display: inline-block;
    margin: 0.5rem 0.5rem 0;
}
</style>