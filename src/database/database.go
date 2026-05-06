package database

import (
	"app/src/config"
	"app/src/model"
	"app/src/utils"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbHost, dbName string) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, config.DBUser, config.DBPassword, dbName, config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		utils.Log.Errorf("Failed to connect to database: %+v", err)
	}

	sqlDB, errDB := db.DB()
	if errDB != nil {
		utils.Log.Errorf("Failed to connect to database: %+v", errDB)
	}

	// Config connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)



	// Run Migrations
	if err := db.AutoMigrate(&model.User{}, &model.Profile{}, &model.Experience{}, &model.NavItem{}, &model.Achievement{}, &model.Photo{}, &model.Video{}, &model.Project{}, &model.CompareHistory{}); err != nil {
		utils.Log.Errorf("Failed to auto migrate: %+v", err)
	}

	// Seed Data if empty
	seedData(db)

	return db
}

func seedData(db *gorm.DB) {
	// Seed Profile
	var profileCount int64
	db.Model(&model.Profile{}).Count(&profileCount)
	if profileCount == 0 {
		profile := model.Profile{
			Name:      "Mabubul Alam Patwary Jaber",
			Headline:  "Software Engineer",
			Bio:       "I'm a software engineer based in Dhaka, Bangladesh specializing in building (and occasionally designing) exceptional digital experiences. Currently, I'm focused on building accessible, human-centered products.",
			About:     "Hello! My name is Jaber and I enjoy creating things that live on the internet. My interest in web development started back in 2020 when I decided to try editing custom Tumblr themes — turns out hacking together HTML & CSS is a lot of fun!\n\nFast-forward to today, and I've had the privilege of working at an advertising agency, a start-up, a huge corporation, and a student-led design studio. My main focus these days is building accessible, inclusive products and digital experiences at Upstatement for a variety of clients.",
			Email:     "jaber@example.com",
			Github:    "https://github.com/jaber",
			Linkedin:  "https://linkedin.com/in/jaber",
			Twitter:   "https://twitter.com/jaber",
			Instagram: "https://instagram.com/jaber",
			Skills:    "JavaScript (ES6+), Typscript, React, Vue, Node.js, Go, PostgreSQL, Docker",
			Image:     "https://i.ibb.co/IMG-URL-HERE/profile.jpg", // Placeholder
		}
		db.Create(&profile)
		utils.Log.Info("Default Profile created")
	}

	// Seed Experience
	var expCount int64
	db.Model(&model.Experience{}).Count(&expCount)
	if expCount == 0 {
		experiences := []model.Experience{
			{
				Company:      "MediaMind",
				Role:         "Software Engineer",
				StartDate:    "2024",
				EndDate:      "Present",
				Description:  "Deliver high-quality, robust production code for a diverse array of projects for clients using Go, Fiber, and React.",
				Technologies: "Go, Fiber, PostgreSQL, Docker, AWS",
			},
			{
				Company:      "TechSolutions",
				Role:         "Junior Web Developer",
				StartDate:    "2022",
				EndDate:      "2024",
				Description:  "Developed and maintained code for in-house and client websites primarily using HTML, CSS, Sass, JavaScript, and jQuery.",
				Technologies: "JavaScript, React, Node.js, MongoDB",
			},
			{
				Company:      "Creative Agency",
				Role:         "Intern",
				StartDate:    "2021",
				EndDate:      "2022",
				Description:  "Assisted in the design and development of creative web projects.",
				Technologies: "HTML, CSS, JavaScript",
			},
		}
		db.Create(&experiences)
		utils.Log.Info("Default Experiences created")
	}
	// Seed Default Admin User
	var adminCount int64
	db.Model(&model.User{}).Where("email = ?", "admin@admin.com").Count(&adminCount)
	if adminCount == 0 {
		hashedPassword, _ := utils.HashPassword("admin")
		adminUser := model.User{
			Name:         "Admin User",
			Email:        "admin@admin.com",
			Phone:        "01700000000",
			PasswordHash: hashedPassword,
			Status:       "active",
		}
		db.Create(&adminUser)
		utils.Log.Info("Default admin user created - Email: admin@admin.com, Password: admin")
	}

	// Seed Achievements
	var count int64
	db.Model(&model.Achievement{}).Count(&count)
	if count == 0 {
		achievements := []model.Achievement{
			{Title: "Best Developer Award 2024", Date: "Dec 2024", Description: "Recognized for outstanding contribution to open source community."},
			{Title: "Certified Kubernetes Administrator", Date: "Aug 2023", Description: "Achieved CKA certification with 95% score."},
			{Title: "Hackathon Winner", Date: "Mar 2023", Description: "First place in National Fintech Hackathon using Go and React."},
		}
		db.Create(&achievements)
	}

	// Seed Photos
	db.Model(&model.Photo{}).Count(&count)
	if count == 0 {
		photos := []model.Photo{
			{Caption: "Speaking at TechConf", Image: "https://images.unsplash.com/photo-1544531586-fde5298cdd40?q=80&w=1000&auto=format&fit=crop"},
			{Caption: "My Workspace Setup", Image: "https://images.unsplash.com/photo-1498050108023-c5249f4df085?q=80&w=1000&auto=format&fit=crop"},
			{Caption: "Team Brainstorming", Image: "https://images.unsplash.com/photo-1522071820081-009f0129c71c?q=80&w=1000&auto=format&fit=crop"},
			{Caption: "Deployment Night", Image: "https://images.unsplash.com/photo-1555099962-4199c345e5dd?q=80&w=1000&auto=format&fit=crop"},
		}
		db.Create(&photos)
	}

	// Seed Videos
	db.Model(&model.Video{}).Count(&count)
	if count == 0 {
		videos := []model.Video{
			{Title: "Building Scalable APIs with Go", YoutubeLink: "https://www.youtube.com/watch?v=2yZbeL26F7g", Description: "A deep dive into Fiber framework and GORM."},
			{Title: "System Design Interview Guide", YoutubeLink: "https://www.youtube.com/watch?v=bUHFg8CZFws", Description: "How to ace your system design interviews."},
			{Title: "My Coding Journey", YoutubeLink: "https://www.youtube.com/watch?v=k9WbpQCRRFs", Description: "From Hello World to Senior Engineer."},
		}
		db.Create(&videos)
	}

	// Seed Projects
	db.Model(&model.Project{}).Count(&count)
	if count == 0 {
		projects := []model.Project{
			{
				Title:      "E-Commerce Platform",
				Summary:    "A full-stack e-commerce solution with payment integration",
				Content:    "Built a comprehensive e-commerce platform using Go, Fiber, and PostgreSQL. Features include user authentication, product management, shopping cart, order processing, and payment gateway integration with bKash and Nagad.",
				Image:      "https://images.unsplash.com/photo-1557821552-17105176677c?q=80&w=1000&auto=format&fit=crop",
				Category:   "Go, Fiber, PostgreSQL, Payment Gateway",
				GithubLink: "https://github.com/yourusername/ecommerce",
				LiveLink:   "https://demo.example.com",
				Featured:   true,
				Order:      1,
			},
			{
				Title:      "Portfolio CMS",
				Summary:    "Content management system for portfolio websites",
				Content:    "Developed a headless CMS specifically designed for portfolio websites. Features include dynamic content management, media uploads, SEO optimization, and a beautiful admin panel.",
				Image:      "https://images.unsplash.com/photo-1460925895917-afdab827c52f?q=80&w=1000&auto=format&fit=crop",
				Category:   "Go, GORM, REST API, Admin Panel",
				GithubLink: "https://github.com/yourusername/portfolio-cms",
				LiveLink:   "",
				Featured:   true,
				Order:      2,
			},
			{
				Title:      "Real-time Chat Application",
				Summary:    "WebSocket-based chat with rooms and notifications",
				Content:    "Created a real-time chat application using WebSockets. Supports multiple chat rooms, private messaging, typing indicators, and push notifications.",
				Image:      "https://images.unsplash.com/photo-1611606063065-ee7946f0787a?q=80&w=1000&auto=format&fit=crop",
				Category:   "WebSocket, Go, Redis, React",
				GithubLink: "https://github.com/yourusername/chat-app",
				LiveLink:   "https://chat.example.com",
				Featured:   true,
				Order:      3,
			},
			{
				Title:      "Task Management API",
				Summary:    "RESTful API for task and project management",
				Content:    "Built a robust task management API with team collaboration features, task assignments, deadlines, and progress tracking.",
				Image:      "https://images.unsplash.com/photo-1484480974693-6ca0a78fb36b?q=80&w=1000&auto=format&fit=crop",
				Category:   "Go, REST API, JWT, PostgreSQL",
				GithubLink: "https://github.com/yourusername/task-api",
				LiveLink:   "",
				Featured:   false,
				Order:      4,
			},
			{
				Title:      "Weather Dashboard",
				Summary:    "Beautiful weather forecast dashboard",
				Content:    "Developed a weather dashboard that displays current weather and forecasts using external APIs. Features include location search, favorites, and weather alerts.",
				Image:      "https://images.unsplash.com/photo-1592210454359-9043f067919b?q=80&w=1000&auto=format&fit=crop",
				Category:   "JavaScript, API Integration, Charts",
				GithubLink: "https://github.com/yourusername/weather-dashboard",
				LiveLink:   "https://weather.example.com",
				Featured:   false,
				Order:      5,
			},
			{
				Title:      "Blog Platform",
				Summary:    "Modern blogging platform with markdown support",
				Content:    "Created a blogging platform with markdown editor, syntax highlighting, comments, and social sharing features.",
				Image:      "https://images.unsplash.com/photo-1499750310107-5fef28a66643?q=80&w=1000&auto=format&fit=crop",
				Category:   "Go, Markdown, SEO, Comments",
				GithubLink: "https://github.com/yourusername/blog-platform",
				LiveLink:   "",
				Featured:   false,
				Order:      6,
			},
		}
		db.Create(&projects)
	}
}

