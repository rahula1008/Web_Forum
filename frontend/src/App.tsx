import './App.css'
import { Route, Routes } from 'react-router'
import HomePage from './pages/HomePage/HomePage'
import TopicPage from './pages/TopicPage/TopicPage'

function App() {
  return (
    <Routes>
        <Route index element = {<HomePage />}></Route>
        <Route path='/topics/:id/posts' element={<TopicPage />} />
    </Routes>
  )
}

export default App
