const input: HTMLInputElement = document.getElementById("input") as HTMLInputElement;
const output = document.getElementById("output");
const ws = new WebSocket(
    `${window.location.href.includes("https")?"wss":"ws"}://${window.location.host}/infectClientWS${window.location.search}`
  );
  
output!.innerText = ""

input?.addEventListener("keyup", async (e) => {
    if (e.key == "Enter") {
        ws.send(input.value);
        output!.innerText = input.value + "\n"
        input!.value = "";         

    }
})
ws.onmessage = (e) => {
    console.log(e.data)
    
    output!.innerText += "\n"+JSON.parse( e.data)["output"] 
    if (output!.innerText.split("\n")!.length > 20) {
        output!.innerText = output!.innerText.split("\n").slice(-20).join("\n")
    }
}