import { API_URL } from './config'

export interface PatentSearchParams {
  patent_id: string
  company_name: string
}

interface InfringingProduct {
  product_name: string
  infringement_likelihood: string
  relevant_claims: number[]
  explanation: string
  specific_features: string[]
}

interface InfringementResponse {
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

export interface InfringementResult {
  productId: string
  productName: string
  companyName: string
  confidenceScore: number
  matchedFeatures: string[]
  analysisDate: string
  overallRiskAssessment: string
  patentId: string
}

export async function searchInfringingProducts(params: PatentSearchParams): Promise<InfringementResult[]> {
  try {
    console.log(API_URL)
    const response = await fetch(`${API_URL}/infringement-check`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(params),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const result: InfringementResponse = await response.json()

    // Transform the API response to match our internal format
    return result.data.infringing_products.map(product => ({
      productId: result.data.analysis_id,
      productName: product.product_name,
      companyName: result.data.company_name,
      confidenceScore: convertLikelihoodToScore(product.infringement_likelihood),
      matchedFeatures: product.specific_features,
      analysisDate: result.data.analysis_date,
      overallRiskAssessment: result.data.overall_risk_assessment,
      patentId: result.data.patent_id
    }))
  } catch (error) {
    console.error('Error searching for infringing products:', error)
    throw error
  }
}

function convertLikelihoodToScore(likelihood: string): number {
  switch (likelihood.toLowerCase()) {
    case 'high':
      return 0.9
    case 'moderate':
      return 0.6
    case 'low':
      return 0.3
    default:
      return 0.5
  }
} 