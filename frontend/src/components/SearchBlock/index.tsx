import { InfringementResult } from '@/apis/patent'
import { Button } from '@/components/ui/button'
import { useCompaniesList, usePatentsList } from '@/features/patents/hooks/usePatentsList'

import { SearchableDropdown } from './SearchableDropdown'

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
  const { data: patents, isLoading: isLoadingPatents } = usePatentsList()
  const { data: companies, isLoading: isLoadingCompanies } = useCompaniesList()
  
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
        <SearchableDropdown
          items={patents || []}
          value={patentId}
          onChange={setPatentId}
          placeholder="Select patent ID"
          label="Patent ID"
          disabled={isSearching}
          isLoading={isLoadingPatents}
        />

        <SearchableDropdown
          items={companies || []}
          value={companyName}
          onChange={setCompanyName}
          placeholder="Select company"
          label="Company Name"
          disabled={isSearching}
          isLoading={isLoadingCompanies}
        />

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