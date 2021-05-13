var navItems = document.querySelectorAll(".nav__item");

for (var i = 0; i < navItems.length; i++) {
  var item = navItems[i];
  if (item.firstElementChild.getAttribute("href") == window.location.pathname) {
    item.classList.add("active");
    break;
  }
}
