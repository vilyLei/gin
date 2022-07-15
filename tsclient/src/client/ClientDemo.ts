
export class ClientDemo {


    constructor() { }

    initialize(): void {

        console.log("ClientDemo::initialize()......");        
		document.onmousedown = (evt: any): void => {
			this.mouseDown(evt);
		};
    }
    private mouseDown(evt: any): void {
        console.log("ClientDemo::mouseDown()");   
        let reqUrl = "http://localhost:9090/shopping/index";
        let request: XMLHttpRequest = new XMLHttpRequest();
        request.open('GET', reqUrl, true);
        console.log("reqUrl: ", reqUrl);
        request.onload = () => {
            if (request.status <= 206) {
                console.log("lrequest.responseText: ", request.responseText);
                // (request.responseText);
            }
            else {
                console.error("load error, url: ", reqUrl);
            }
        };
        request.onerror = e => {
            console.error("load error, url: ", reqUrl);
        };
        request.send(null);
    }
    run(): void {
    }
}
export default ClientDemo;