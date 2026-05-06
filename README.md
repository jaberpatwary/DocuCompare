# Portfolio Website

A modern, clean portfolio website built with Go (Fiber framework) and vanilla JavaScript.

## Tech Stack

**Backend:**
- Go 1.21+
- Fiber (Web Framework)
- GORM (ORM)
- PostgreSQL
- JWT Authentication

**Frontend:**
- HTML5
- CSS3 (Vanilla)
- JavaScript (Vanilla)

## Features

- ✅ Dynamic profile management
- ✅ Work experience timeline
- ✅ Project showcase
- ✅ Customizable navigation menu
- ✅ Social media links
- ✅ Skills section
- ✅ Admin panel for content management
- ✅ Responsive design

## Setup

### Prerequisites
- Go 1.21 or higher
- PostgreSQL
- Docker (optional)

### Installation

1. Clone the repository
```bash
git clone <repository-url>
cd potp
```

2. Copy `.env` file and configure
```bash
cp .env.example .env
# Edit .env with your database credentials
```

3. Install dependencies
```bash
go mod download
```

4. Run the application
```bash
make start
```

The application will be available at `http://localhost:7777`

### Using Docker

```bash
docker-compose up -d
```

## Project Structure

```
potp/
├── frontend/           # Frontend assets
│   ├── admin.html     # Admin panel
│   ├── index.html     # Main website
│   ├── css/           # Stylesheets
│   └── js/            # JavaScript files
├── src/
│   ├── config/        # Configuration
│   ├── controller/    # Request handlers
│   ├── database/      # Database connection
│   ├── middleware/    # Middleware functions
│   ├── model/         # Data models
│   ├── router/        # Route definitions
│   ├── service/       # Business logic
│   ├── utils/         # Utility functions
│   └── validation/    # Input validation
├── .env               # Environment variables
├── Makefile           # Build commands
└── go.mod             # Go dependencies
```

## Admin Panel

Access the admin panel at `/admin.html`

Default credentials:
- Username: `admin`
- Password: `admin`

**Features:**
- Profile management
- Experience CRUD
- Project CRUD
- Navigation menu customization
- Social media links

## API Endpoints

### Public
- `GET /` - Main website
- `GET /api/profile` - Get profile data
- `GET /api/experience` - Get work experience
- `GET /api/projects` - Get projects
- `GET /api/nav-items` - Get navigation menu

### Admin (Requires Authentication)
- `POST /api/profile` - Update profile
- `POST /api/experience` - Create experience
- `PUT /api/experience/:id` - Update experience
- `DELETE /api/experience/:id` - Delete experience
- `POST /api/projects` - Create project
- `PUT /api/projects/:id` - Update project
- `DELETE /api/projects/:id` - Delete project
- `POST /api/nav-items` - Create nav item
- `PUT /api/nav-items/:id` - Update nav item
- `DELETE /api/nav-items/:id` - Delete nav item

## Development

### Available Commands

```bash
make start      # Start the application
make build      # Build the application
make test       # Run tests
make clean      # Clean build artifacts
```

## License

MIT License

## Credits

Design inspired by [Brittany Chiang](https://brittanychiang.com)
