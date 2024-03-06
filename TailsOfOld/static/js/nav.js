$("body").addClass(localStorage.theme)
console.log(localStorage.theme)

function toggleTheme() {
    var body = $("body")
    if (body.hasClass("dark")) {
        body.removeClass("dark")
        body.addClass("light")
        localStorage.theme = "light"
        return
    }
    body.removeClass("light")
    body.addClass("dark")
    localStorage.theme = "dark"
}