import React, {useState} from 'react';
import MyInput from "./UI/input/MyInput";
import MyButton from "./UI/button/MyButton";

const PostForm = ({create}) => {
    const [post, setPost] = useState({title:'',body:''})

    //const bodyInputRef = useRef(); // Для неуправляемого

    const addNewPost = (e) => {
        e.preventDefault() // Не обнавляет страницу
        //console.log(newPost)
        //setPosts([...posts, {...post, id: Date.now()}]) // Не менять напрямую, только через функцию. Разворачиваем старый массив, прибавляем элемент
        const newPost={
            ...post, id: Date.now()
        }
        create(newPost)
        setPost({title:'',body:''})
        //console.log(bodyInputRef.current.value) // Для неуправляемого

    }

    return (
        <form>
            {/*Управляемый компонент*/}
            <MyInput
                value={post.title}
                onChange={e => setPost({...post, title: e.target.value})}
                type="text"
                placeholder="Название поста"
            />
            {/*/!*Неуправляемый компонент*!/*/}
            {/*<MyInput*/}
            {/*    ref={bodyInputRef}*/}
            {/*    type={"text"}*/}
            {/*    placeholder={"Описание поста"}*/}
            {/*/>*/}
            <MyInput
                value={post.body}
                onChange={e => setPost({...post, body: e.target.value})}
                type={"text"}
                placeholder={"Описание поста"}
            />
            <MyButton onClick={addNewPost}>Создать пост</MyButton>
        </form>
    );
};

export default PostForm;