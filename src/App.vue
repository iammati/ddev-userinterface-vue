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

                    <li id="ddev-router" data-status="unknown">
                        <a href="#" class="flex">
                            <span class="pulse">
                                <span class="animate-ping"></span>
                                <span class="static-bg"></span>
                            </span>

                            <span class="mr-2"> ddev-router: </span>

                            <DdevRouterStatus :data="routerData" />
                        </a>
                    </li>
                </ul>
            </nav>
        </section>

        <section id="content">
            <router-view />
        </section>
    </main>
</template>

<script>
import DdevRouterStatus from "@/components/DdevRouterStatus.vue";
import { api } from "./utils/api";
// import Swal from 'sweetalert2';
// import { Swal2Toastr } from "./utils/swal2toastr";

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
    mounted() {
        setInterval(async () => {
            const routerData = await api("/router_status").then((response) =>
                response.json()
            );
            this.routerData = routerData;

            if (routerData.status.length > 0) {
                document
                    .getElementById("ddev-router")
                    .setAttribute("data-status", routerData.status);
            }
        }, 1000);

        // Swal.fire({
        //     title: 'Error!',
        //     text: 'Do you want to continue',
        //     icon: 'error',
        //     confirmButtonText: 'Cool'
        // })

        // const a = (new Swal2Toastr()).createToastWrapper();
        // console.log(a);
    },
};
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

                                span {
                                    position: relative;
                                    display: block;
                                    width: 18px;
                                    height: 18px;

                                    .animate-ping {
                                        animation: ping 1s cubic-bezier(0, 0, .2, 1) infinite;
                                        opacity: .75;
                                        border-radius: 9999px;
                                        width: 100%;
                                        height: 100%;
                                        display: inline-flex;
                                        position: absolute;
                                    }
                                }

                                &[data-status="healthy"] {
                                    span {
                                        .animate-ping {
                                            background-color: rgba(
                                                52,
                                                211,
                                                153,
                                                0.75
                                            );

                                            & + span {
                                                background-color: rgba(
                                                    52,
                                                    211,
                                                    153,
                                                    0.75
                                                );
                                            }
                                        }
                                    }
                                }

                                &[data-status="unhealthy"],
                                &[data-status="stopped"] {
                                    span {
                                        .animate-ping {
                                            background-color: rgba(
                                                251,
                                                113,
                                                133,
                                                0.75
                                            );

                                            & + span {
                                                background-color: rgba(
                                                    244,
                                                    63,
                                                    94,
                                                    0.75
                                                );
                                            }
                                        }
                                    }
                                }

                                &[data-status="unknown"],
                                &[data-status="starting"] {
                                    span {
                                        .animate-ping {
                                            background-color: gold;

                                            & + span {
                                                background-color: gold;
                                            }
                                        }
                                    }
                                }

                                a {
                                    cursor: default;
                                    display: flex;

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
