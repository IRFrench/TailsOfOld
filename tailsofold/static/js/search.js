function searchRedirect(id) {
    var search = $('#' + id).val();
    if (search == "") {
        return;
    }
    window.location = '/search?q=' + search;
}

$("#nav-search").on("keypress", function (event) {
    if (event.key === "Enter") {
        searchRedirect("nav-search")
    }
});
