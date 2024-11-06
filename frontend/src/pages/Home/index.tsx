import { ListBlock } from '@/components/ListBlock'
import { SearchBlock } from '@/components/SearchBlock'
import { useSearchBlock } from '@/components/SearchBlock/hooks/useSearchBlock'
import { Button } from '@/components/ui/button'
import { Star, StarIcon } from 'lucide-react'

function Home() {
  const [states, operations] = useSearchBlock()

  return (
    <div className="w-full max-w-lg space-y-4 p-8 bg-gray-800/90 backdrop-blur-sm rounded-xl shadow-lg relative z-10 text-white">
      <h2 className="text-2xl font-bold text-center mb-6">
        Patent Infringement Search
      </h2>
      <SearchBlock states={states} operations={operations} />

      {/* Saved Searches Section */}
      {states.savedSearches.length > 0 && (
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Saved Searches</h3>
          <div className="space-y-2">
            {states.savedSearches.map((search, index) => (
              <button
                key={index}
                onClick={() => {
                  operations.setPatentId(search.patent_id)
                  operations.setCompanyName(search.company_name)
                  operations.handleSearch()
                }}
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
            ))}
          </div>
        </div>
      )}

      <div className='gap-3 flex justify-between'>
        <h3 className="text-xl font-semibold">Search Results</h3>

        {states.showSaveButton && (
          <Button
            onClick={operations.handleSave}
            size="icon"
            variant="outline"
            className="text-yellow-400 hover:text-yellow-500"
          >
            {states.isCurrentSearchSaved ? (
              <StarIcon className="h-4 w-4 fill-current" />
            ) : (
              <Star className="h-4 w-4" />
            )}
          </Button>
        )}
      </div>
      {states.searchResults &&
        <ListBlock 
            searchResults={states.searchResults || []}
            analysisDate={states.searchResults?.[0]?.analysisDate || ''}
            overallRiskAssessment={states.searchResults?.[0]?.overallRiskAssessment || ''}
            patentId={states.searchResults?.[0]?.patentId || ''}
        />
      }
    </div>
  )
}

export default Home