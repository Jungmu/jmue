new Vue({
    el: '#app',
    data: {
        message: 'Hello Vue.js!'
    }
})

var navVue = new Vue({
    el: '#nav',
    data: {
        show: true
    }
})

function navVueShow() {
    if(window.innerHeight < 1080) {
        navVue.show = false
    }
}
document.getElementById("nav").onload = function() {navVueShow()};
document.getElementById("nav").onresize = function(){navVueShow()};