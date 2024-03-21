function calcuateReadTime() {
    blog = document.getElementById("article");
    readTime = document.getElementById("read-time")

    entire_blog = blog.textContent.split(" ")
    readTime.textContent = "Read Time: " + Math.round(entire_blog.length / 238 + 1) + " Mins"
}

calcuateReadTime()