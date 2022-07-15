


import { ClientDemo as Demo } from "./client/ClientDemo";

document.title = "goTsClient";
let demoIns: Demo = new Demo();
let ins: any = demoIns;

function main(): void {
    console.log("------ demo --- init ------");
    demoIns.initialize();
    function mainLoop(now: any): void {
        demoIns.run();
        window.requestAnimationFrame(mainLoop);
    }
    window.requestAnimationFrame(mainLoop);
    console.log("------ demo --- running ------");
}
//
main();