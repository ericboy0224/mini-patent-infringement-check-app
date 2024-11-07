import { PatentSearchParams } from '@/apis/patent'

interface SavedSearchItemProps {
  search: PatentSearchParams
  onSelect: (search: PatentSearchParams) => void
}

export function SavedSearchItem({ search, onSelect }: SavedSearchItemProps) {
  return (
    <button
      onClick={() => onSelect(search)}
      className="w-full p-2 text-left bg-gray-700 hover:bg-gray-600 rounded-md transition-colors"
    >
      <div className="text-sm">
        {search.patent_id && (
          <span className="text-blue-400">Patent: {search.patent_id}</span>
        )}
        {search.patent_id && search.company_name && (
          <span className="mx-2">|</span>
        )}
        {search.company_name && (
          <span className="text-green-400">Company: {search.company_name}</span>
        )}
      </div>
    </button>
  )
} 