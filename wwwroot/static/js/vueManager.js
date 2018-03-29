function route(id) {
    document.getElementById("menuCheckBox").checked = false;
    hideAllVueObject();
    id.show = true;
    id.init();
}

function hideAllVueObject() {
    indexPage.show = false;
    contactPage.show = false;
    steemit.show =false;
}

var indexPage = new Vue({
    el: '#index',
    data: {
        show: false,
        isInit : false,
        title: 'Hello Vue.js!'
    },
    methods: {
        init() {
            if(!this.isInit) {
                var url = new URL(window.location.href);
                var send = url.searchParams.get("send");
                if(send == "ok") {
                    openAlert("thank for send email");
                }
                this.isInit = true;
            }
        }
    }
});

var contactPage = new Vue({
    el: '#contect',
    data: {
        show: false,
        isInit : false,
        title: 'Send Email!'
    },
    methods: {
        init() {
            if(!this.isInit) {
                this.isInit = true;
            }
        }
    }
});