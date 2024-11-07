import { useState, useEffect } from 'react'
import { usePatentSearch } from '@/features/patents/hooks/usePatentSearch'
import { InfringementResult, PatentSearchParams } from '@/apis/patent'

const STORAGE_KEY = 'savedPatentSearches'

interface SearchBlockStates {
  patentId: string
  companyName: string
  isSearching: boolean
  error: Error | null
  searchResults?: InfringementResult[]
  savedSearches: PatentSearchParams[]
  previousPatentId?: string
  previousCompanyName?: string;
  showSaveButton: boolean;
  isCurrentSearchSaved: boolean;
}

interface SearchBlockOperations {
  setPatentId: (value: string) => void
  setCompanyName: (value: string) => void
  handleSearch: () => Promise<void>
  handleSave: () => void
  handleSelect: (search: PatentSearchParams) => void
}

export function useSearchBlock(): [SearchBlockStates, SearchBlockOperations] {
  const [patentId, setPatentId] = useState('')
  const [companyName, setCompanyName] = useState('')
  const [previousPatentId, setPreviousPatentId] = useState<string>()
  const [previousCompanyName, setPreviousCompanyName] = useState<string>()
  const [savedSearches, setSavedSearches] = useState<PatentSearchParams[]>([])
  const { mutate, isPending, error, data, reset } = usePatentSearch()
  const showSaveButton = !!(data?.length && previousPatentId && previousCompanyName);

  useEffect(() => {
    const saved = localStorage.getItem(STORAGE_KEY)
    if (saved) {
      setSavedSearches(JSON.parse(saved))
    }
  }, [])

  const handleSearch = async (search?: PatentSearchParams) => {
    if (!patentId && !companyName && !search) {
      console.warn('Please enter at least one search term')
      return
    }

    reset();

    let patent_id = patentId;
    let company_name = companyName

    if (search && search.company_name && search.patent_id) {
      patent_id = search.patent_id;
      company_name = search.company_name;
    }

    mutate(
      {
        patent_id,
        company_name
      },
      {
        onSuccess: () => {
          setPreviousPatentId(patentId)
          setPreviousCompanyName(companyName)
        }
      },
    )
  }

  const handleSave = () => {
    if (!previousPatentId && !previousCompanyName) {
      console.warn('No previous search to save')
      return
    }

    const newSearch: PatentSearchParams = {
      patent_id: previousPatentId as string,
      company_name: previousCompanyName as string
    }

    const updatedSearches = [
      newSearch,
      ...savedSearches.filter(search =>
        !(search.patent_id === previousPatentId && search.company_name === previousCompanyName)
      )
    ]

    setSavedSearches(updatedSearches)
    localStorage.setItem(STORAGE_KEY, JSON.stringify(updatedSearches))
  }

  const handleSelect = (search: PatentSearchParams) => {
    setPatentId(search.patent_id)
    setCompanyName(search.company_name)
    handleSearch(search);
  }

  const isCurrentSearchSaved = savedSearches.some(
    search =>
      search.patent_id === previousPatentId &&
      search.company_name === previousCompanyName
  );

  const states: SearchBlockStates = {
    patentId,
    companyName,
    isSearching: isPending,
    error,
    searchResults: data,
    savedSearches,
    previousPatentId,
    previousCompanyName,
    showSaveButton,
    isCurrentSearchSaved,
  }

  const operations: SearchBlockOperations = {
    setPatentId,
    setCompanyName,
    handleSearch,
    handleSave,
    handleSelect
  }

  return [states, operations]
} 