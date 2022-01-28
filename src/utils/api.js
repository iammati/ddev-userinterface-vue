import * as ddevUiConfig from "../../ddev-ui.config";

const config = ddevUiConfig.default;

export async function api(endpoint) {
    return await fetch(
        `${config.scheme}://${config.host}:${config.port}/api${endpoint}`
    );
}
