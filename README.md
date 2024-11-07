# Patlytics Takehome

![Result](/img/result.png)


A full-stack patent infringement check application that analyzes potential patent infringements by comparing patent claims against company products.

## Features
- Patent infringement analysis using AI (Groq API)
- MongoDB integration for caching analysis results
- RESTful API endpoint for infringement checks
- Modern React frontend with TypeScript
- Docker containerization for easy deployment
- Automated testing suite

## Tech Stack
### Backend
- Go 1.22.3
- Gin Web Framework
- MongoDB
- Groq AI API (using gemma-7b-it model)

### Frontend
- React 18.3
- TypeScript
- Vite
- TailwindCSS
- React Query
- React Router DOM
- Shadcn/ui Components

## Prerequisites
- Docker and Docker Compose installed
- Groq API key
- Go 1.22.3+ (for local development)
- Node.js 20+ (for frontend development)
- pnpm (for frontend package management)

## Project Structure
```
.
├── frontend/             # React frontend application
│   ├── src/             # Frontend source code
│   ├── public/          # Static assets
│   └── package.json     # Frontend dependencies
├── data/                # JSON data files
├── domains/             # Business logic
├── handlers/            # HTTP request handlers
├── models/              # Data models
├── prompts/             # AI prompt templates
├── services/            # External service integrations
├── .env                 # Environment variables
├── docker-compose.yml   # Docker compose configuration
├── Dockerfile          # Docker build instructions
└── Makefile           # Build and deployment commands
```

## Setup
Build and run the application(visit http://localhost:8080):
```bash
make setup
```

## Development
### Frontend Development
```bash
cd frontend
pnpm install
pnpm dev
```

### Backend Development
```bash
go run main.go
```

## API Endpoints

### POST /patlytics/v1/infringement-check
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
- `make frontend-install` - Install frontend dependencies

## Testing
Run the test suite:
```bash
make test
```

## Architecture
The application follows a clean architecture pattern with:

### Backend
1. Handlers: HTTP request handling and response formatting
2. Domains: Core business logic and AI integration
3. Models: Data structures and database models
4. Services: External service connections (MongoDB, Groq)

### Frontend
1. Components: Reusable UI components
2. Features: Feature-specific logic and types
3. Hooks: Custom React hooks
4. APIs: API integration layer
5. Layouts: Page layouts and structure

## License
MIT

## Contributing
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request