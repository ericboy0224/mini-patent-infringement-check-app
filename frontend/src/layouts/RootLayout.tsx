import { ReactNode } from 'react'

interface RootLayoutProps {
  children: ReactNode
}

export function RootLayout({ children }: RootLayoutProps) {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-900 relative overflow-hidden">
      {/* Background blur effect */}
      <div 
        className="absolute inset-0 bg-[url('https://cdn.prod.website-files.com/661050b5482552f0f0296612/6611879a3cefc9f993d865b2_Patlytics-White.svg')] 
        bg-center bg-no-repeat opacity-5 blur-xl scale-150"
        aria-hidden="true"
      />
      
      {/* Large logo */}
      <img 
        src="https://cdn.prod.website-files.com/661050b5482552f0f0296612/6611879a3cefc9f993d865b2_Patlytics-White.svg"
        alt="Patlytics Logo"
        className="w-64 mb-12 relative z-10"
      />
      
      {children}
    </div>
  )
} 