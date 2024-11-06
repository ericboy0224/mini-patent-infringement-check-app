import { type InfringementResult } from '@/apis/patent'
import { Progress } from '@/components/ui/progress'

interface ListBlockProps {
    searchResults: InfringementResult[]
    analysisDate: string
    overallRiskAssessment: string
    patentId: string
}

export function ListBlock({ 
    searchResults, 
    analysisDate, 
    overallRiskAssessment,
    patentId 
}: ListBlockProps) {
    return (
        <div className="mt-6 space-y-4 text-white">
            {/* Analysis Summary */}
            <div className="p-4 bg-gray-800 rounded-lg">
                <div className="grid grid-cols-2 gap-4 text-sm">
                    <div>
                        <p className="text-gray-400">Analysis Date:</p>
                        <p>{new Date(analysisDate).toLocaleDateString()}</p>
                    </div>
                    <div>
                        <p className="text-gray-400">Patent ID:</p>
                        <p>{patentId}</p>
                    </div>
                    <div className="col-span-2">
                        <p className="text-gray-400">Overall Risk Assessment:</p>
                        <p className="text-yellow-400 font-medium">{overallRiskAssessment}</p>
                    </div>
                </div>
            </div>

            {/* Product Results */}
            {searchResults.map((product) => (
                <div
                    key={product.productId}
                    className="p-4 bg-gray-700 rounded-lg"
                >
                    <h4 className="font-medium">{product.productName}</h4>
                    <p className="text-sm text-gray-300 mt-1">
                        Company: {product.companyName}
                    </p>
                    <div className="space-y-1.5">
                        <div className="flex items-center justify-between text-sm text-gray-300">
                            <span>Infringing lookalike</span>
                            <span className="font-medium">
                                {(product.confidenceScore * 100).toFixed(0)}%
                            </span>
                        </div>
                        <Progress 
                            value={product.confidenceScore * 100} 
                            className="h-1.5"
                        />
                    </div>
                    <div className="text-sm mt-2">
                        <p>Matched Features:</p>
                        <ul className="list-disc list-inside mt-1">
                            {product.matchedFeatures.map((feature, index) => (
                                <li key={index}>{feature}</li>
                            ))}
                        </ul>
                    </div>
                </div>
            ))}
        </div>
    )
} 