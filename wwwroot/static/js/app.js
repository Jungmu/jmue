var IE_warning = new Vue({
    el: '#IE',
    data: {
        show: false,
        message: 'Internet Explorer is not support! please use other browser.'
    }
})

new Vue({
    el: '#index',
    data: {
        message: 'Hello Vue.js!'
    }
})

function detectIE(){
    var agent = navigator.userAgent.toLowerCase();

    if( agent.indexOf('msie') != -1 || agent.indexOf('trident') != -1 ) {
        IE_warning.show = true;
    }
}

function onload() {
    detectIE();
}