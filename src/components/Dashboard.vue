<template>
    <section id="nav">
        <nav>
            <ul>
                <li class="waves-effect waves-light">
                    <router-link to="/">
                        <i class="fas fa-home me-2"></i>
                        <span>Home</span>
                    </router-link>
                </li>

                <li class="waves-effect waves-light">
                    <router-link to="/projects">
                        <i class="fas fa-list-ul me-2"></i>
                        <span>Projects</span>
                    </router-link>
                </li>

                <li id="ddev-router" data-status="unknown">
                    <a href="#" class="flex">
                        <span class="pulse me-2">
                            <span class="animate-ping"></span>
                            <span class="static-bg"></span>
                        </span>

                        <span class="me-2"> ddev-router: </span>

                        <DdevRouterStatus :data="routerData" />
                    </a>
                </li>
            </ul>
        </nav>
    </section>
</template>

<script>
import DdevRouterStatus from "@/components/DdevRouterStatus.vue";
import { api } from "../utils/api";
import Swal from "sweetalert2";

export default {
    name: "Dashboard",
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
        let retrys = 0;

        const Toast = Swal.mixin({
            toast: true,
            position: 'top-end',
            showConfirmButton: false,
            timer: 10,
        });

        const that = this;

        window.dashboardInterval = setInterval(function() {
            if (retrys == 5) {
                clearInterval(window.dashboardInterval);

                Toast.fire({
                    icon: 'error',
                    title: 'No connection established to backend.',
                    timerProgressBar: true,
                });
            } else {
                retrys = retrys + 1;

                (async () => {
                    const response = await api("/router_status");
                    const routerData = await response.json();
                    that.routerData = routerData;

                    if (routerData.status.length > 0) {
                        document
                            .getElementById("ddev-router")
                            .setAttribute("data-status", routerData.status);
                    }
                })();
            }
        }, 1000);
    },
}
</script>

<style lang="scss">
main section#nav {
    position: relative;
    background-color: $primary;
    box-shadow: 0px 0px 20px rgb(0 0 0 / 70%), 0 0px 10px 0 #000;
    padding-top: 3rem;

    nav {
        height: 100%;

        ul {
            position: relative;
            margin: 0;
            padding-left: 0;
            list-style: none;
            height: 100%;

            li {
                width: 100%;

                &:hover a,
                a.router-link-exact-active {
                    background-color: darken($primary, 5%);
                }

                &#ddev-router {
                    background-color: darken($primary, 3.33%);
                    position: absolute;
                    bottom: 0;

                    span.pulse {
                        position: relative;
                        display: block;
                        width: 12px;
                        height: 12px;

                        span {
                            position: relative;
                            display: inline-flex;
                            opacity: .75;
                            border-radius: 9999px;
                            width: 100%;
                            height: 100%;

                            &.animate-ping {
                                position: absolute;
                                top: 6px;
                                animation: ping 1s cubic-bezier(0, 0, .2, 1) infinite;

                                @keyframes ping {
                                    75%, 100% {
                                        transform: scale(2);
                                        opacity: 0;
                                    }
                                }
                            }
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
                        // color: #00325a;
                    }
                }
            }
        }
    }
}
</style>
