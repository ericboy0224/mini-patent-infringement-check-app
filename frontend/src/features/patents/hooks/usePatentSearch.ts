import { useMutation } from '@tanstack/react-query'
import { searchInfringingProducts } from '@/apis/patent'
import type { PatentSearchParams } from '@/apis/patent'
import type { InfringementResult } from '@/apis/patent'

export function usePatentSearch() {
  return useMutation<InfringementResult[], Error, PatentSearchParams>({
    mutationFn: searchInfringingProducts,
    onError: (error: Error) => {
      console.error('Patent search failed:', error)
    },
    gcTime: 0
  })
} 