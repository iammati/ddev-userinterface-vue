import '../Scss/app.scss';

window.onload = function () {
    setInterval(async () => {
        const response = await fetch('/api/router_status');
        const json = await response.text();
        const status = JSON.parse(json).data;

        (document.getElementById('router-status') as HTMLDivElement).innerText = status;
    }, 1500)
};
