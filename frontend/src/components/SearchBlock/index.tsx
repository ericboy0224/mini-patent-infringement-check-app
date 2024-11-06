import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { type InfringementResult } from '@/apis/patent'

interface SearchBlockProps {
  states: {
    patentId: string
    companyName: string
    isSearching: boolean
    error: Error | null
    searchResults?: InfringementResult[]
    previousPatentId?: string
    previousCompanyName?: string
  }
  operations: {
    setPatentId: (value: string) => void
    setCompanyName: (value: string) => void
    handleSearch: () => Promise<void>
    handleSave: () => void
  }
}

export function SearchBlock({ states, operations }: SearchBlockProps) {
  const { 
    patentId, 
    companyName, 
    isSearching, 
    error
  } = states
  const { setPatentId, setCompanyName, handleSearch } = operations

  return (
    <div className="space-y-6">
      <div className="flex flex-col md:flex-row gap-4 items-end">
        <div className="flex-1 space-y-2">
          <label htmlFor="patentId" className="text-sm font-medium text-gray-300">
            Patent ID
          </label>
          <Input
            id="patentId"
            type="text"
            value={patentId}
            onChange={(e) => setPatentId(e.target.value)}
            placeholder="Enter patent ID"
            className="bg-gray-700 border-gray-600 text-white placeholder-gray-400 focus:ring-blue-500 focus:border-blue-500"
            disabled={isSearching}
          />
        </div>

        <div className="flex-1 space-y-2">
          <label htmlFor="companyName" className="text-sm font-medium text-gray-300">
            Company Name
          </label>
          <Input
            id="companyName"
            type="text"
            value={companyName}
            onChange={(e) => setCompanyName(e.target.value)}
            placeholder="Enter company name"
            className="bg-gray-700 border-gray-600 text-white placeholder-gray-400 focus:ring-blue-500 focus:border-blue-500"
            disabled={isSearching}
          />
        </div>

        <div className="flex gap-2">
          <Button
            onClick={handleSearch}
            size="lg"
            disabled={isSearching || (!patentId && !companyName)}
          >
            {isSearching ? 'Searching...' : 'Search'}
          </Button>
        </div>
      </div>

      {error && (
        <div className="text-red-400 text-sm">
          {error.message}
        </div>
      )}
    </div>
  )
} 