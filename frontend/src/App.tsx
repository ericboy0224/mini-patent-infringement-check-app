import { HashRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

import { RootLayout } from './layouts/RootLayout'
import Home from './pages/Home'

const queryClient = new QueryClient({
  defaultOptions: {
    mutations: {
      retry: false,
    },
  },
})

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <HashRouter>
        <RootLayout>
          <Routes>
            <Route path="/" element={<Home />} />
          </Routes>
        </RootLayout>
      </HashRouter>
    </QueryClientProvider>
  )
}

export default App
