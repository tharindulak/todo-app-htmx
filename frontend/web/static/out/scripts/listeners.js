document.addEventListener("htmx:afterRequest", function(evt) {
    switch (evt.target.id) {
        case "create-form":
            document.getElementById("create-form").reset();
            break;
    }
})
