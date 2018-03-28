var converter = new showdown.Converter({
    'tables':true,
    'strikethrough':true,
    'simpleLineBreaks':true,
    "simplifiedAutoLink":true
});

var steemit = new Vue({
    el: '#steemit',
    data: {
        posts: []
    },
    methods: {
        addToPost(item) {
            this.posts.push({
                title: item.title,
                content: item.content
            });
        }
    }
})

var lastAuthor;
var lastPermlink;

steem.api.getDiscussionsByBlog({"tag": "jungmu", "limit": 5}, function(err, result) {
    document.getElementById("postLoader").style.display = "block";
    result.forEach(element => {
        if(element.author == "jungmu")
        {
            let item = {title:"", content:""};
            item.title = element.title;
            item.content = converter.makeHtml(element.body);
            steemit.addToPost(item);
            lastAuthor=element.author;
            lastPermlink = element.permlink;
            document.getElementById("postLoader").style.display = "none";
            return;
        }
    });
    document.getElementById("postLoader").style.display = "none";
});

function getPostMore(){
    document.getElementById("postLoader").style.display = "block";
    steem.api.getDiscussionsByBlog({"tag": "jungmu", "limit": 5, "start_author": lastAuthor, "start_permlink": lastPermlink}, function(err, result) {
        result.forEach(element => {
            if(element.author == "jungmu" && element.permlink != lastPermlink && element.category == "dev")
            {                
                let item = {title:"", content:""};
                item.title = element.title;
                item.content = converter.makeHtml(element.body);
                steemit.addToPost(item);
                lastAuthor=element.author;
                lastPermlink = element.permlink;
                document.getElementById("postLoader").style.display = "none";
                return;
            }
        });
        document.getElementById("postLoader").style.display = "none";
    });
}

window.onscroll = function(ev) {
    if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight) {
        getPostMore();
    }
};