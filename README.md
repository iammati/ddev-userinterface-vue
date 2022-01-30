# ddev-ui

Experimental UI for ddev based on golang as backend and vue3 as frontend.

Note that minimum requirements to compile this are:
- go v1.17

ddev should take care for the leftover parts of the frontend via `ddev pnpm` commands as you can see below.

## Project setup
```
git clone https://github.com/iammati/ddev-ui .
ddev pnpm install
```

### Building the ui

In case you're running on wsl2 – as I do – you may want to [checkout this video here](https://www.youtube.com/watch?v=o6H6q3xg8wg)
as it installs chrome inside your wsl2 distribution to be able running the built golang binary.

```
ddev pnpm run build
go run app.go
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

### Preview

![ddev-ui-projects-view](https://user-images.githubusercontent.com/41418763/151707931-335a6ca4-f9d5-4ea2-bfa0-aa405af4c5e1.png)
