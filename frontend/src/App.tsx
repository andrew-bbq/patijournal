import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import { Home } from './pages/home/home'
import { Entries } from './pages/entries/entries'
import { CreateEntry } from './pages/create-entry/create-entry'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/entries" element={<Entries />} />
        <Route path="/entries/create" element={<CreateEntry />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
