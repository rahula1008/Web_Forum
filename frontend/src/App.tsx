import './App.css'
import { Route, Routes } from 'react-router'
import HomePage from './pages/HomePage/HomePage'

function App() {
  return (
    <Routes>
        <Route index element = {<HomePage />}></Route>
    </Routes>
  )
}

export default App
