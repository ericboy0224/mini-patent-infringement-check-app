import { getCompanies, getPatents } from '@/apis/patent'
import { useQuery } from '@tanstack/react-query'

export function usePatentsList() {
  return useQuery({
    queryKey: ['patents'],
    queryFn: getPatents
  })
}

export function useCompaniesList() {
  return useQuery({
    queryKey: ['companies'],
    queryFn: getCompanies
  })
}
