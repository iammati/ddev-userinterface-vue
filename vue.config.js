const production = process.env.NODE_ENV === "production";

/** @type {import('@vue/cli-service').ProjectOptions} */
module.exports = {
    runtimeCompiler: true,
    publicPath: production ? "/dist/" : "/public/",
    devServer: {
        port: 3000,
        public: "ddev-ui.ddev.site:2999/public/",
        overlay: true,
        disableHostCheck: true,
        hot: true,
    },
    configureWebpack: {
        watchOptions: {
            ignored: /node_modules/,
            aggregateTimeout: 300,

            // increase this in case of performance issues...
            // personally running on AMD ryzen 5 with 6-core processor
            // - worst case was an increasement of around
            // ~6% spikes on the CPU - also when spamming
            // ctrl+s only...
            poll: 5,
        },
    },
    css: {
        loaderOptions: {
            sass: {
                prependData: `@import "@/styles/app.scss";`,
            },
        },
    },
};
