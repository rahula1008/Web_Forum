import './App.css'
import { Route, Routes } from 'react-router'
import HomePage from './pages/HomePage/HomePage'

function App() {
  return (
    <Routes>
        <Route index element = {<HomePage />}></Route>
        <Route path='/topic' element={<p>hello</p>} />
    </Routes>
  )
}

export default App
