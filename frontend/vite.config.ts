import { resolve } from 'path'
import { UserConfig, defineConfig } from 'vite'
import timeReporter from 'vite-plugin-time-reporter'
import { viteCommonjs } from '@originjs/vite-plugin-commonjs'

const production = process.env.NODE_ENV === 'production'

const configuration: UserConfig = {
    base: '/dist/',

    plugins: [
        viteCommonjs(),
        timeReporter(),
    ],

    build: {
        minify: production,
        outDir: resolve(__dirname, '../public/dist'),
        target: 'es2021',
        emptyOutDir: true,
        manifest: true,
        sourcemap: true,
        watch: {
            include: './src/**'
        },
        rollupOptions: {
            input: ['./src/TypeScript/app.ts'],
            output: [
                {
                    assetFileNames: '[name].[ext]',
                    entryFileNames: '[name].js',
                },
            ],
        },
    },

    // HMR server-port which is exposed by DDEV-Local in .ddev/docker-compose.hmr.yaml
    server: {
        port: 3000,

        // WSL2 support
        watch: {
            usePolling: true,
        },
    },
}

export default defineConfig(configuration)
