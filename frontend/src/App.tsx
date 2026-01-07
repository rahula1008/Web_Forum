import './App.css'
import { Route, Routes } from 'react-router'
import HomePage from './pages/HomePage/HomePage'
import TopicPage from './pages/TopicPage/TopicPage'
import PostPage from './pages/PostPage/PostPage'
import LoginPage from './pages/LoginPage/LoginPage'
import RegisterPage from './pages/RegisterPage/RegisterPage'

function App() {
  return (
    <Routes>
        <Route index element = {<HomePage />}></Route>
        <Route path='/topics/:id/posts' element={<TopicPage />} />
        <Route path='/topics/:id/posts/:postId' element={<PostPage />} />
        <Route path='/login' element={<LoginPage />} />
        <Route path='/register' element={<RegisterPage />} />
    </Routes>
  )
}

export default App
