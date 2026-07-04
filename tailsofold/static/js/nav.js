function toggleTheme() {
    var body = $("body");
    if (body.hasClass("dark")) {
        body.removeClass("dark");
        body.addClass("light");
        localStorage.theme = "light";
        document.documentElement.style.setProperty("--hue", 300);
        return;
    }
    body.removeClass("light");
    body.addClass("dark");
    localStorage.theme = "dark";
    document.documentElement.style.setProperty("--hue", 90);
}

if (localStorage.theme) {
    $("body").addClass(localStorage.theme);
    document.documentElement.style.setProperty("--hue", 90);
} else {
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        $("body").addClass("dark");
        document.documentElement.style.setProperty("--hue", 90);
    }
}

function mobileNav() {
    var mn = $("#mobile_nav");
    mn.toggle();
}