package db

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/rohinish404/social-go/internal/store"
)

var usernames = []string{
	"alex", "jordan", "taylor", "morgan", "casey",
	"riley", "avery", "quinn", "charlie", "dakota",
	"skyler", "phoenix", "river", "sage", "rowan",
	"parker", "blake", "jamie", "drew", "sam",
	"austin", "devon", "hayden", "jessie", "bailey",
	"cameron", "elliott", "finley", "gray", "harley",
	"kendall", "peyton", "reese", "sawyer", "emerson",
	"logan", "hunter", "shane", "tyler", "spencer",
	"mason", "payton", "rory", "robin", "remy",
	"adrian", "kai", "angel", "corey", "dallas",
}

var titles = []string{
	"Getting Started with Go",
	"Understanding Concurrency",
	"Web Development Best Practices",
	"Building RESTful APIs",
	"Database Design Patterns",
	"Microservices Architecture",
	"Clean Code Principles",
	"Docker for Beginners",
	"Testing Strategies",
	"Git Workflow Tips",
	"Performance Optimization",
	"Security Best Practices",
	"Debugging Techniques",
	"CI/CD Pipeline Setup",
	"Cloud Native Applications",
	"API Design Guidelines",
	"Code Review Checklist",
	"Monitoring and Logging",
	"Scalability Patterns",
	"Modern DevOps Practices",
}

var contents = []string{
	"This is a comprehensive guide to help you understand the basics and get started quickly.",
	"In this post, we'll explore key concepts and practical examples that you can apply immediately.",
	"Learn the essential techniques that will improve your development workflow significantly.",
	"Discover best practices and common pitfalls to avoid in your projects.",
	"A deep dive into the fundamental principles that every developer should know.",
	"Step-by-step tutorial with real-world examples and hands-on exercises.",
	"Exploring advanced topics and how to implement them in production environments.",
	"Tips and tricks from experienced developers to level up your skills.",
	"Understanding the core concepts with practical demonstrations and code samples.",
	"A complete walkthrough of building scalable and maintainable solutions.",
	"Common challenges and how to overcome them with proven strategies.",
	"Best practices for writing clean, efficient, and testable code.",
	"An introduction to modern tools and frameworks that boost productivity.",
	"How to architect robust systems that can handle growing demands.",
	"Practical insights into optimizing performance and reducing bottlenecks.",
	"Essential patterns and approaches for building reliable applications.",
	"A beginner-friendly guide with clear explanations and examples.",
	"Advanced techniques for experienced developers looking to refine their craft.",
	"Real-world case studies and lessons learned from production systems.",
	"The complete guide to mastering this important aspect of development.",
}

var tags = []string{
	"golang", "programming", "web-development", "backend",
	"tutorial", "best-practices", "architecture", "database",
	"api", "microservices", "devops", "docker",
	"kubernetes", "testing", "security", "performance",
	"cloud", "git", "ci-cd", "beginners",
}

var comments = []string{
	"Great post! This really helped me understand the concept better.",
	"Thanks for sharing this. Very informative and well-written.",
	"I have a question about the implementation details. Could you elaborate?",
	"This is exactly what I was looking for. Appreciate the detailed explanation.",
	"Interesting approach! I'll definitely try this in my next project.",
	"Well explained! The examples made it very easy to follow.",
	"Thanks for the tutorial. Looking forward to more content like this.",
	"This solved my problem perfectly. Thank you!",
	"Excellent article. The step-by-step guide was really helpful.",
	"Nice work! Would love to see a follow-up on advanced topics.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(100, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i:=0; i<num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email: usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i:=0; i< num; i++ {
		user := users[rand.IntN(len(users))]


		posts[i] = &store.Post{
			UserID: user.ID,
			Title: titles[rand.IntN(len(titles))],
			Content: contents[rand.IntN(len(contents))],
			Tags: []string{
				tags[rand.IntN(len(tags))],
				tags[rand.IntN(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i:=0; i<num; i++ {
		cms[i] = &store.Comment{
			PostID: posts[rand.IntN(len(posts))].ID,
			UserID: users[rand.IntN(len(users))].ID,
			Content: comments[rand.IntN(len(comments))],
		}
	}

	return cms
}