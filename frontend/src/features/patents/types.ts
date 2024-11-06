export interface PatentSearchParams {
  patentId?: string
  companyName?: string
}

export interface InfringingProduct {
  product_name: string
  infringement_likelihood: string
  relevant_claims: number[]
  explanation: string
  specific_features: string[]
}

export interface InfringementResponse {
  status: string
  message: string
  data: {
    analysis_date: string
    analysis_id: string
    company_name: string
    infringing_products: InfringingProduct[]
    overall_risk_assessment: string
    patent_id: string
  }
} 