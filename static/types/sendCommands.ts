const input:HTMLInputElement=document.getElementById("input") as HTMLInputElement;
input?.addEventListener("keyup",(e)=>{
    if (e.key=="Enter"){
        console.log(input?.value);

    }
})