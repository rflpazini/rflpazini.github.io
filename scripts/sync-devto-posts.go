// sync-devto-posts - Dev.to posts synchronizer
//
// This Go program replaces the Python scripts for synchronizing Dev.to posts.
// It supports both RSS feed and API methods for fetching posts.
//
// Usage:
//   ./sync-devto-posts           # Use RSS method (default, recommended)
//   ./sync-devto-posts rss       # Use RSS method explicitly
//   ./sync-devto-posts api       # Use API method
//
// Build:
//   go build -o sync-devto-posts sync-devto-posts.go
//
// Features:
//   - Fetches latest 10 posts from Dev.to
//   - Supports both RSS feed and API methods
//   - Automatic fallback from RSS to API if RSS fails
//   - Writes data to data/devto_posts.toml in TOML format
//   - Cross-platform single binary (no Python dependencies needed)

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mmcdole/gofeed"
)

// DevtoPost represents a Dev.to post
type DevtoPost struct {
	Title   string `json:"title" toml:"title"`
	URL     string `json:"url" toml:"url"`
	Date    string `json:"date" toml:"date"`
	Excerpt string `json:"excerpt" toml:"excerpt"`
}

// DataFile represents the TOML structure
type DataFile struct {
	Posts []DevtoPost `toml:"posts"`
}

// Config holds the application configuration
type Config struct {
	DevtoUsername string
	DataFile      string
	UseRSS        bool
}

// NewConfig creates a new configuration
func NewConfig() *Config {
	return &Config{
		DevtoUsername: "rflpazini",
		DataFile:      "data/devto_posts.toml",
		UseRSS:        true, // Default to RSS as it's more reliable
	}
}

// cleanHTML removes HTML tags and cleans the text
func cleanHTML(text string) string {
	// Remove HTML tags
	htmlRegex := regexp.MustCompile(`<[^>]*>`)
	text = htmlRegex.ReplaceAllString(text, "")

	// Remove extra line breaks and spaces
	text = regexp.MustCompile(`\n+`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

// fetchPostsViaAPI fetches posts using the Dev.to API
func fetchPostsViaAPI(username string) ([]DevtoPost, error) {
	apiURL := fmt.Sprintf("https://dev.to/api/articles?username=%s&per_page=10", username)

	fmt.Printf("📡 Fetching posts from Dev.to API: %s\n", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var posts []map[string]interface{}
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	var processedPosts []DevtoPost
	for _, post := range posts {
		excerpt := ""
		if desc, ok := post["description"].(string); ok {
			excerpt = desc
			if len(excerpt) > 200 {
				excerpt = excerpt[:200] + "..."
			}
		}

		publishedAt := ""
		if date, ok := post["published_at"].(string); ok && len(date) >= 10 {
			publishedAt = date[:10] // YYYY-MM-DD format
		}

		processedPost := DevtoPost{
			Title:   post["title"].(string),
			URL:     post["url"].(string),
			Date:    publishedAt,
			Excerpt: excerpt,
		}
		processedPosts = append(processedPosts, processedPost)
	}

	return processedPosts, nil
}

// fetchPostsViaRSS fetches posts using the Dev.to RSS feed
func fetchPostsViaRSS(username string) ([]DevtoPost, error) {
	rssURL := fmt.Sprintf("https://dev.to/feed/%s", username)

	fmt.Printf("📡 Fetching RSS feed: %s\n", rssURL)

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	if len(feed.Items) == 0 {
		return nil, fmt.Errorf("no posts found in RSS feed")
	}

	fmt.Printf("📄 Found %d posts in RSS feed\n", len(feed.Items))

	var processedPosts []DevtoPost
	limit := 10
	if len(feed.Items) < limit {
		limit = len(feed.Items)
	}

	for i := 0; i < limit; i++ {
		item := feed.Items[i]

		// Extract excerpt from content or summary
		excerpt := ""
		if item.Description != "" {
			excerpt = cleanHTML(item.Description)
		} else if len(item.Content) > 0 {
			excerpt = cleanHTML(item.Content)
		}

		// Limit excerpt size
		if len(excerpt) > 200 {
			excerpt = excerpt[:200] + "..."
		}

		// Parse published date
		publishedDate := ""
		if item.PublishedParsed != nil {
			publishedDate = item.PublishedParsed.Format("2006-01-02")
		} else if item.UpdatedParsed != nil {
			publishedDate = item.UpdatedParsed.Format("2006-01-02")
		}

		processedPost := DevtoPost{
			Title:   item.Title,
			URL:     item.Link,
			Date:    publishedDate,
			Excerpt: excerpt,
		}
		processedPosts = append(processedPosts, processedPost)
	}

	return processedPosts, nil
}

// updateDataFile writes the posts to the TOML data file
func updateDataFile(posts []DevtoPost, dataFile string) error {
	fmt.Printf("📝 Processing %d posts...\n", len(posts))

	// Validate posts
	var validPosts []DevtoPost
	for _, post := range posts {
		if post.Title == "" {
			post.Title = "Title not available"
		}
		validPosts = append(validPosts, post)
	}

	if len(validPosts) == 0 {
		return fmt.Errorf("no valid posts found")
	}

	// Create TOML structure
	data := DataFile{
		Posts: validPosts,
	}

	// Ensure directory exists
	dir := filepath.Dir(dataFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write TOML file
	file, err := os.Create(dataFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(data); err != nil {
		return fmt.Errorf("failed to write TOML file: %w", err)
	}

	fmt.Printf("✅ File %s updated successfully!\n", dataFile)
	fmt.Printf("📝 %d posts synchronized\n", len(validPosts))

	// Show titles of first 3 posts
	for i, post := range validPosts {
		if i >= 3 {
			break
		}
		fmt.Printf("  %d. %s\n", i+1, post.Title)
	}

	return nil
}

// runAPISync runs the API-based synchronization
func runAPISync(config *Config) error {
	fmt.Println("🚀 Starting synchronization with Dev.to via API...")
	fmt.Printf("👤 User: @%s\n", config.DevtoUsername)

	posts, err := fetchPostsViaAPI(config.DevtoUsername)
	if err != nil {
		return fmt.Errorf("failed to fetch posts via API: %w", err)
	}

	if err := updateDataFile(posts, config.DataFile); err != nil {
		return fmt.Errorf("failed to update data file: %w", err)
	}

	fmt.Println("✅ API synchronization completed!")
	return nil
}

// runRSSSync runs the RSS-based synchronization
func runRSSSync(config *Config) error {
	fmt.Println("🚀 Starting synchronization with Dev.to via RSS...")
	fmt.Printf("👤 User: @%s\n", config.DevtoUsername)

	posts, err := fetchPostsViaRSS(config.DevtoUsername)
	if err != nil {
		fmt.Printf("❌ RSS sync failed: %v\n", err)
		fmt.Println("🔄 Trying API method as fallback...")

		// Fallback to API method
		return runAPISync(config)
	}

	if err := updateDataFile(posts, config.DataFile); err != nil {
		return fmt.Errorf("failed to update data file: %w", err)
	}

	fmt.Println("✅ RSS synchronization completed successfully!")
	return nil
}

func main() {
	config := NewConfig()

	// Check command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "api":
			config.UseRSS = false
		case "rss":
			config.UseRSS = true
		default:
			fmt.Printf("Usage: %s [api|rss]\n", os.Args[0])
			fmt.Println("  api: Use Dev.to API method")
			fmt.Println("  rss: Use RSS feed method (default)")
			os.Exit(1)
		}
	}

	var err error
	if config.UseRSS {
		err = runRSSSync(config)
	} else {
		err = runAPISync(config)
	}

	if err != nil {
		log.Fatalf("❌ Synchronization failed: %v", err)
	}
}
