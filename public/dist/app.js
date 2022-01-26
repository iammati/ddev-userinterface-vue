window.onload=function(){setInterval(async()=>{const t=await(await fetch("/api/router_status")).text(),e=JSON.parse(t).data;document.getElementById("router-status").innerText=e},1500)};
//# sourceMappingURL=app.js.map
