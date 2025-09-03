# Dev.to Automatic Integration

This document explains how the automatic integration between your Hugo site and Dev.to works.

## How It Works

### 1. Automatic Synchronization
- **Daily**: GitHub Actions workflow runs automatically every day at 3:00 BRT
- **On every push**: Every time you push to the `main` branch
- **Manual**: You can run manually through GitHub Actions

### 2. Data Flow
```
Dev.to RSS Feed → Go Binary → TOML File → Hugo → Site
```

### 3. Files Involved

#### Scripts
- `scripts/sync-devto-posts.go` - Go source code for synchronization
- `scripts/sync-devto-posts` - Compiled Go binary (generated)

#### Data
- `data/devto_posts.toml` - Data file with synchronized posts

#### Templates
- `layouts/shortcodes/devto-posts.html` - Shortcode that displays posts

#### Workflow
- `.github/workflows/gh-pages.yml` - Automated workflow

## 🔧 Configuration

### Dependencies
```bash
# Go dependencies are managed automatically via go.mod
go mod download
```

### Manual Execution
```bash
# Build the binary (if not already built)
go build -o scripts/sync-devto-posts scripts/sync-devto-posts.go

# Run synchronization (uses RSS by default)
./scripts/sync-devto-posts

# Or specify method explicitly
./scripts/sync-devto-posts rss  # Use RSS feed
./scripts/sync-devto-posts api  # Use API
```

## Data Structure

The `data/devto_posts.toml` file contains:

```toml
# Last update: 2024-01-15 09:30:00

[[posts]]
title = "Post Title"
url = "https://dev.to/username/post-slug"
date = "2024-01-15"
excerpt = "Post summary..."
```

## Customization

### Change Dev.to Username
Edit the `DEVTO_USERNAME` constant in the scripts:
```python
DEVTO_USERNAME = "your_username_here"
```

### Change Number of Posts
In the `layouts/shortcodes/devto-posts.html` file, change the number:
```html
{{ range first 5 $devtoData.posts }}  <!-- Shows 5 posts -->
```

### Change Synchronization Frequency
In the `.github/workflows/gh-pages.yml` file:
```yaml
schedule:
  - cron: '0 6 * * *'  # Every day at 6:00 UTC (3:00 BRT)
```

## 🔍 Monitoring

### Workflow Logs
Access **Actions** in your GitHub repository to view execution logs.

### Synchronization Status
- ✅ **Success**: Posts were synchronized
- ⚠️ **Warning**: RSS may have temporary issues
- ❌ **Error**: Synchronization failed (doesn't affect build)

## 🛠️ Troubleshooting

### Problem: Posts don't appear on site
1. Check if the `data/devto_posts.toml` file was updated
2. Confirm the `{{< devto-posts >}}` shortcode is on the page
3. Run the workflow manually in GitHub Actions

### Problem: RSS Error
1. Dev.to RSS may be temporarily unavailable
2. The script has fallbacks to continue working even with errors

### Problem: Build fails
1. The workflow uses `continue-on-error: true` to not break the build
2. Even if synchronization fails, the site will be built with old data

## 📈 Future Improvements

- [ ] Support for multiple RSS feeds
- [ ] Smart caching to avoid unnecessary requests
- [ ] Notifications when new posts are detected
- [ ] Support for Dev.to webhooks
- [ ] Synchronization metrics
