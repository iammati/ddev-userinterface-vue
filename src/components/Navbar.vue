<template>
    <div id="header">
        <div id="ngProgress-container">
            <div id="ngProgress"></div>
        </div>

        <div class="logo">
            <Logo subline="Userinterface v0.1.0" />
        </div>
    </div>
</template>

<script>
import Logo from "./Logo.vue";

export default {
    name: "Navbar",
    components: {
        Logo,
    },
};
</script>

<style lang="scss">
#header {
    padding: 10px;
    display: flex;
    height: 65px;
    background-color: $primary;
    position: relative;
    z-index: 1;
    box-shadow: 0px 0px 12px rgba(0, 0, 0, .7), 230px -20px 30px 0 #000;

    .logo {
        max-width: 120px;
        width: 100%;
    }
}
</style>
