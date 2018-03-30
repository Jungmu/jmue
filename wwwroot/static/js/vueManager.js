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

let indexPage = new Vue({
    el: '#index',
    data: {
        show: false,
        isInit : false,
        title: "Jongsun's Home Page"
    },
    methods: {
        init() {
            if(!this.isInit) {
                let url = new URL(window.location.href);
                let send = url.searchParams.get("send");
                if(send == "ok") {
                    openAlert("thank for send email");
                } else if (send == "fail") {
                    openAlert("mail send fail, email not vaild");
                }
                this.isInit = true;
            }
        }
    }
});

let contactPage = new Vue({
    el: '#contact',
    data: {
        show: false,
        isInit : false,
        title: 'Try Contact Me!'
    },
    methods: {
        init() {
            if(!this.isInit) {
                this.isInit = true;
            }
        }
    }
});