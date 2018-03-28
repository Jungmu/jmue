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

steem.api.getDiscussionsByBlog({"tag": "jungmu", "limit": 20}, function(err, result) {
    result.forEach(element => {
        if(element.author == "jungmu")
        {
            let item = {title:"", content:""};
            item.title = element.title;
            item.content = converter.makeHtml(element.body);
            steemit.addToPost(item);
        }
    });
    console.log(steemit.posts);    
  });