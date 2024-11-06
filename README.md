# Patlytics Takehome
A mini patent infringement check application that analyzes potential patent infringements by comparing patent claims against company products.

## Features
- Patent infringement analysis using AI (Groq API)
- MongoDB integration for caching analysis results
- RESTful API endpoint for infringement checks
- Docker containerization for easy deployment
- Automated testing suite

## Tech Stack
- Go 1.22.3
- Gin Web Framework
- MongoDB
- Docker & Docker Compose
- Groq AI API (using gemma-7b-it model)

## Prerequisites
- Docker and Docker Compose installed
- Groq API key
- Go 1.22.3+ (for local development)

## Project Structure
```
.
├── data/                  # JSON data files
├── domains/              # Business logic
├── handlers/             # HTTP request handlers
├── models/               # Data models
├── prompts/              # AI prompt templates
├── services/             # External service integrations
├── .env                  # Environment variables
├── docker-compose.yml    # Docker compose configuration
├── Dockerfile           # Docker build instructions
└── Makefile            # Build and deployment commands
```

## Setup
1. Clone the repository
2. Create a `.env` file with the following variables:
```
GROQ_API_KEY=your_groq_api_key
MONGODB_URI=mongodb://mongodb:27017
```

3. Build and run the application:
```bash
make setup
```

## API Endpoints

### POST /infringement-check
Analyzes potential patent infringement for a given patent and company.

Request body:
```json
{
  "patent_id": "US-RE49889-E1",
  "company_name": "Company Name"
}
```

Response:
```json
{
  "status": "success",
  "data": {
    "analysis_id": "...",
    "analysis_date": "...",
    "patent_id": "...",
    "company_name": "...",
    "infringing_products": [...],
    "overall_risk_assessment": "..."
  },
  "message": "Infringement check completed successfully"
}
```

## Development Commands
- `make build` - Build Docker images
- `make run` - Run the application
- `make test` - Run tests
- `make stop` - Stop the application
- `make clean` - Clean up containers and volumes
- `make logs` - View application logs

## Testing
Run the test suite:
```bash
make test
```

## Architecture
The application follows a clean architecture pattern with the following components:

1. Handlers: HTTP request handling and response formatting
2. Domains: Core business logic and AI integration
3. Models: Data structures and database models
4. Services: External service connections (MongoDB, Groq)

## Code References
- Main application setup: 
main.go


- Infringement check handler:
handlers/infringement.go


- Analysis domain logic:
domains/analysis.go


## License
MIT

## Contributing
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request