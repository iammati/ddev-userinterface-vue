<template>
    <Navbar />

    <main>
        <Dashboard />

        <section id="content">
            <router-view />
        </section>
    </main>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import Dashboard from "@/components/Dashboard.vue";

export default {
    name: "App",
    components: {
        Navbar,
        Dashboard,
    },
};
</script>

<style lang="scss">
$primary: #00325a;
$headerHeight: 85px;

#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    color: rgb(27, 28, 38);
    overflow: hidden;

    main {
        display: grid;
        grid-template-columns: 240px 1fr;
        grid-template-rows: 50px 1fr 50px;
        height: calc(100vh - 65px);

        section {
            height: calc(100vh - 65px);

            &#content {
                height: calc(100vh - $headerHeight + 2px - 1rem);
                background-color: $lightgrey;

                padding: {
                    top: 3rem;
                    left: 3rem;
                    right: 3rem;
                }
            }
        }
    }
}

#ngProgress {
    margin: 0;
    padding: 0;
    z-index: 99998;
    background-color: gold !important;
    color: gold !important;
    box-shadow: 0 0 10px 0;
    height: 2px;
    width: 0%;
    opacity: 1;
    transition: all 0.5s ease-in-out;
}

#ngProgress-container {
    position: fixed;
    margin: 0;
    padding: 0;
    top: 0;
    left: 0;
    right: 0;
    z-index: 99999;
}
</style>
