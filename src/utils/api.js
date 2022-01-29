import * as ddevUiConfig from "../../ddev-ui.config";

const config = ddevUiConfig.default;

export async function api(endpoint, args = {}, method = "POST") {
    const url = `${config.scheme}://${config.host}:${config.port}/api${endpoint}`;

    if (Object.keys(args).length > 0) {
        return await fetch(url, {
            method: method,
            mode: "cors",
            cache: "no-cache",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            redirect: "follow",
            referrerPolicy: "no-referrer",
            body: JSON.stringify(args),
        });
    }

    return await fetch(url);
}
