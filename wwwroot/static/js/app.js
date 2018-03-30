let IE_warning = new Vue({
    el: '#IE',
    data: {
        show: false,
        title: 'Internet Explorer is not support! please use other browser.'
    }
});

function detectIE(){
    let agent = navigator.userAgent.toLowerCase();

    if( agent.indexOf('msie') != -1 || agent.indexOf('trident') != -1 ) {
        IE_warning.show = true;
        return ture;
    }
}



function onload() {
    if(detectIE()) {
        // only show warning text,
        return;
    }
    route(indexPage);
}