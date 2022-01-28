<template>
    <div id="header">
        <div class="logo">
            <img alt="Vue logo" src="./assets/logo.svg" />
        </div>
    </div>

    <main>
        <section id="nav">
            <nav>
                <ul>
                    <li>
                        <router-link to="/">Home</router-link>
                    </li>

                    <li>
                        <router-link to="/projects">Projects</router-link>
                    </li>

                    <li id="ddev-router" class="online">
                        <a href="#">
                            ddev-router: 

                            <span>
                                <DdevRouterStatus :data="routerData" />
                            </span>
                        </a>
                    </li>
                </ul>
            </nav>
        </section>

        <section id="content">
            <span class="relative inline-flex">
      <button type="button" class="inline-flex items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-sky-500 bg-white dark:bg-slate-800 transition ease-in-out duration-150 cursor-not-allowed ring-1 ring-slate-900/10 dark:ring-slate-200/20" disabled="">
        Transactions
      </button>
      <span class="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
        <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-sky-400 opacity-75"></span>
        <span class="relative inline-flex rounded-full h-3 w-3 bg-sky-500"></span>
      </span>
    </span>

            <router-view />
        </section>
    </main>
</template>

<script>
import DdevRouterStatus from "@/components/DdevRouterStatus.vue";
import { api } from "./utils/api";

export default {
    name: "App",
    components: {
        DdevRouterStatus,
    },
    data() {
        return {
            routerData: {
                status: "unknown",
                warning: "none",
            },
        };
    },
    async mounted() {
        const routerData = await api("/router_status").then(response => response.json())
        this.routerData = routerData
    },
}
</script>

<style lang="scss">
$primary: #00325a;
$headerHeight: 85px;

body {
    margin: 0;
}

#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;

    #header {
        padding: 16px;
        display: flex;
        background-color: $primary;
        max-height: $headerHeight;

        .logo {
            width: 200px;
        }
    }
    
    main {
        display: grid;
        grid-template-columns: 240px 1fr;
        grid-template-rows: 50px 1fr 50px;
        height: calc(100vh - $headerHeight + 2px);

        section {
            height: calc(100vh - $headerHeight + 2px);

            &#nav {
                background-color: $primary;

                nav {
                    height: 100%;

                    ul {
                        position: relative;
                        margin: 0;
                        padding-left: 0;
                        list-style: none;
                        display: flex;
                        flex-direction: column;
                        align-items: center;
                        justify-content: center;
                        height: 100%;

                        li {
                            width: 100%;

                            &:hover a {
                                background-color: #002340;
                            }

                            &#ddev-router {
                                position: absolute;
                                bottom: 0;

                                a {
                                    cursor: default;

                                    &:hover {
                                        background-color: inherit;
                                    }
                                }
                            }

                            a {
                                color: #fff;
                                text-decoration: none;
                                display: block;
                                padding: 16px;

                                &.router-link-exact-active {
                                    background-color: #002340;
                                }
                            }
                        }
                    }
                }
            }

            &#content {
                height: calc(100vh - $headerHeight + 2px - 1rem);

                padding: {
                    top: 1rem;
                    left: 1rem;
                };
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
    transition: all .5s ease-in-out
}

#ngProgress-container {
    position: fixed;
    margin: 0;
    padding: 0;
    top: 0;
    left: 0;
    right: 0;
    z-index: 99999
}
</style>
