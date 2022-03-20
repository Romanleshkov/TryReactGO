import React, {useMemo, useState} from "react";
import './styles/App.css'
import PostList from "./components/PostList";
import PostForm from "./components/PostForm";
import PostFilter from "./components/PostFilter";

const App = function () {
    const [posts, setPosts] = useState([
        {id: 1, title: 'ab', body: 'cd'},
        {id: 2, title: 'cd', body: 'ef'},
        {id: 3, title: 'ef', body: 'ab'}
    ])

    const [filter, setFilter] = useState({sort: '', query: ''})

    const sortedPosts = useMemo(() => {
        console.log("Sorted")
        if (filter.sort) {
            return [...posts].sort((a, b) => a[filter.sort].localeCompare(b[filter.sort]))
        }
        return posts
    }, [filter.sort, posts])


    const sortedAndSearchedPosts = useMemo(()=>{
        return sortedPosts.filter(post => post.title.toLowerCase().includes(filter.query))
    }, [filter.query,sortedPosts])

    const createPost = (newPost) => {
        setPosts([...posts, newPost])
    }

    const removePost = (post) => {
        setPosts(posts.filter(p => p.id !== post.id))
    }


    return (
        <div className="App">
            <PostForm create={createPost}/>
            <hr style={{margin: '15px'}}/>
            <div>
                <PostFilter
                filter={filter}
                setFilter={setFilter}/>
            </div>
            <PostList remove={removePost} posts={sortedAndSearchedPosts} title={"Посты про JS"}/>
        </div>
    );
};


export default App;
