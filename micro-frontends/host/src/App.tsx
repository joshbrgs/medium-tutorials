import React, { useEffect, useRef } from 'react'
import ReactDOM from 'react-dom/client'

import './index.scss'
import Counter from "remote/Counter"

const App = () => {
    const ref = useRef(null);

    useEffect(() => {
        Counter(ref.current)
    }, [])

    return(
      <div className="flex items-center justify-center min-h-screen bg-gray-200">
      <div className="relative p-8 bg-white border-4 border-black shadow-[8px_8px_0_0_rgba(0,0,0,1)]">
        <h1 className="text-2xl font-bold text-black">Neo Brutalist Card</h1>
        <p className="my-4 text-black">
          This is an example of a Neo Brutalist card styled with bold borders,
          stark contrasts, and a playful shadow effect.
        </p>
        <div className="active:translate-x-2 active:translate-y-2 active:shadow-[4px_4px_0_0_rgba(0,0,0,0)] transition-transform duration-150" ref={ref} /> 
      </div>
    </div>
)};

const rootElement = document.getElementById('app')
if (!rootElement) throw new Error('Failed to find the root element')

const root = ReactDOM.createRoot(rootElement as HTMLElement)

root.render(<App />)
