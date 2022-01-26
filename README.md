# ddev-ui

Experimental UI for ddev based on golang as backend and vue3 as frontend.

Note that minimum requirements to compile this are:
- go v1.17

ddev should take care for the leftover parts of the frontend via `ddev pnpm`.

## Project setup
```
git clone https://github.com/iammati/ddev-ui .
ddev pnpm install
```

### Compiles and hot-reloads for development
```
ddev pnpm run serve
```

Head over to `http://ddev-ui.ddev.site:2999/public/` in order to see hot-reloads being available.
Hot changes won't affect the golang runtime.

### Compiles and minifies for production
```
ddev pnpm run build
```

### Lints and fixes files
```
ddev pnpm run lint
```
