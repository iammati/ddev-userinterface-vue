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
                            ddev-router: <span></span>
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
</style>
