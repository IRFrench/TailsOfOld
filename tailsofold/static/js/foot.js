$("#newsletter").submit(function (e) {
    e.preventDefault();
    $("#news_submit").attr('disabled', 'disabled')
    var fullName = $("#news_full_name").val();
    var email = $("#news_email").val();
    $.post("/news/subscribe", `{ "full_name": "${fullName}", "email": "${email}" }`)
        .done(subscribed)
        .fail(failedToSubscribe)
});

function subscribed() {
    $("#news_full_name").hide();
    $("#news_email").hide();
    $("#form_note").show();
    console.log("subscribed");
}

function failedToSubscribe() {
    $("#news_full_name").hide();
    $("#news_email").hide();
    $("#form_note").text("Sorry, something went wrong");
    $("#form_note").show();
    console.log("failed to subscribe");
}

function Subscribe() {
    console.log("here");
}