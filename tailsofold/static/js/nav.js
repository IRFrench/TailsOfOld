function toggleTheme() {
    var body = $("body");
    if (body.hasClass("dark")) {
        body.removeClass("dark");
        body.addClass("light");
        localStorage.theme = "light";
        return;
    }
    body.removeClass("light");
    body.addClass("dark");
    localStorage.theme = "dark";
}

function toggleThemeMenu() {
    var menu = $("#themeMenu")
    if (menu.hasClass("hidden")) {
        menu.removeClass("hidden")
        return
    }
    menu.addClass("hidden")
}

$("#hueRange").on("input", function () {
    var hue = $("#hueRange").val();
    console.log(hue);
    document.documentElement.style.setProperty("--hue", hue)
    document.documentElement.style.setProperty("hue", hue)

})

if (localStorage.theme) {
    $("body").addClass(localStorage.theme);
} else {
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        $("body").addClass("dark");
    }
}

function mobileNav() {
    var mn = $("#mobile_nav");
    mn.toggle();
}