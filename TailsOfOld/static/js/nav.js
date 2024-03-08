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

if (localStorage.theme) {
    $("body").addClass(localStorage.theme)
} else {
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        $("body").addClass("dark")
    }
}