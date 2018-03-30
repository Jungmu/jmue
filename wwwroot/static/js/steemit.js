let converter = new showdown.Converter({
    'tables':true,
    'strikethrough':true,
    'simpleLineBreaks':true,
    "simplifiedAutoLink":true
});

let steemit = new Vue({
    el: '#steemit',
    data: {
        show: false,
        isInit : false,
        title : "My Steemit Blog",
        posts: []
    },
    methods: {
        addToPost(item) {
            this.posts.push({
                title: item.title,
                content: item.content
            });
        },
        init() {
            if(!this.isInit) {
                showBlog();
                this.isInit = true;
            }
        }
    }
});

let lastAuthor;
let lastPermlink;

function showSteemit(){
    steemit.show = true;
}

function hideSteemit(){
    steemit.show = false;
}

function getOnlyMyDevPostOne(element){
    // category == dev but, no post for dev yet
    if(element.author == "jungmu" && element.permlink != lastPermlink && element.category == "kr")
    {                
        let item = {title:"", content:""};
        item.title = element.title;
        item.content = converter.makeHtml(element.body);
        steemit.addToPost(item);
        lastAuthor=element.author;
        lastPermlink = element.permlink;
        document.getElementById("postLoader").style.display = "none";
        return false;
    }
    return true;
}

function showBlog(){    
    document.getElementById("postLoader").style.display = "block";
    steem.api.getDiscussionsByBlog({"tag": "jungmu", "limit": 5}, (err, result) => {
        // if this true, no more post
        if(result.every(getOnlyMyDevPostOne)) {
            document.getElementById("postLoader").style.display = "none";
        }
    });
}

function getPostMore(){
    document.getElementById("postLoader").style.display = "block";
    steem.api.getDiscussionsByBlog({"tag": "jungmu", "limit": 5, "start_author": lastAuthor, "start_permlink": lastPermlink}, function(err, result) {
        // if this true, no more post
        if(result.every(getOnlyMyDevPostOne)) {
            document.getElementById("postLoader").style.display = "none";
            openAlert("no more post");    
        }
    });
}

window.onscroll = () => {
    if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight) {
        if(steemit.show == true){
            getPostMore();            
        }
    }
};