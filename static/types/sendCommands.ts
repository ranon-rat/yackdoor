const input: HTMLInputElement = document.getElementById("input") as HTMLInputElement;
const output=document.getElementById("output") ;
output!.innerText=""
input?.addEventListener("keyup", async (e) => {
    
    if (e.key == "Enter") {

        const response = await fetch("", {
            method: "POST",
            body: input?.value
        }).catch(e=>null)
        output!.innerText=input.value+"\n"
        input!.value = "";
        
        const reader = response!.body!.getReader();
        var gayInterval = setInterval(async function () {
            const { value, done } = await reader.read();
            if (done) {clearInterval(gayInterval)
            };
            const string=new TextDecoder().decode(value)
            console.log(string)
            try{
           output!.innerText+=JSON.parse(string)["output"]+"\n"
          if (output!.innerText.split("\n")!.length>20){
           output!.innerText=output!.innerText.split("\n").slice(-19).join("\n")}}
            catch{}
        }, 100)

    }
})