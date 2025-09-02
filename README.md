# @rflpazini - Personal Blog

[![github pages](https://github.com/rflpazini/rflpazini.github.io/actions/workflows/gh-pages.yml/badge.svg)](https://github.com/rflpazini/rflpazini.github.io/actions/workflows/gh-pages.yml)
[![Hugo](https://img.shields.io/badge/Hugo-0.149.0-blue.svg)](https://gohugo.io)
[![License](https://img.shields.io/badge/License-CC%20BY--NC%204.0-green.svg)](https://creativecommons.org/licenses/by-nc/4.0/)

> Rafael Pazini, Software engineer, caipira do interior paulista 🇧🇷, Open Source enthusiast and I think I am an interesting person.

## 🌐 Live Site

Visit my blog at: **[https://rflpazini.com](https://rflpazini.com)**

## 🚀 About

This is my personal blog built with [Hugo](https://gohugo.io) - a fast and flexible static site generator. The site features a clean, minimal design using the [hello-friend-ng](https://github.com/rhazdon/hugo-theme-hello-friend-ng) theme.

### ✨ Features

- **Multilingual Support**: Available in English and Portuguese (pt-BR)
- **Responsive Design**: Optimized for all devices
- **Dark/Light Theme**: Automatic theme switching based on system preferences
- **Social Integration**: Links to YouTube, Twitter, Dev.to, Medium, GitHub, and LinkedIn
- **RSS Feed**: Subscribe to stay updated with new posts
- **Fast Loading**: Static site generation for optimal performance

## 🛠️ Tech Stack

- **[Hugo](https://gohugo.io)** - Static site generator
- **[hello-friend-ng](https://github.com/rhazdon/hugo-theme-hello-friend-ng)** - Theme
- **[GitHub Pages](https://pages.github.com)** - Hosting
- **[GitHub Actions](https://github.com/features/actions)** - CI/CD

## 📁 Project Structure

```
rflpazini.github.io/
├── content/           # Blog posts and pages
│   ├── about.md      # About page (EN)
│   ├── about.pt-br.md # About page (PT-BR)
│   └── slides.md     # Slides page
├── layouts/          # Custom layouts and shortcodes
├── static/           # Static assets (images, docs, etc.)
├── themes/           # Hugo theme
├── i18n/            # Internationalization files
└── config.toml      # Site configuration
```

## 🚀 Getting Started

### Prerequisites

- [Hugo Extended](https://gohugo.io/installation/) (v0.149.0 or later)
- Git

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/rflpazini/rflpazini.github.io.git
   cd rflpazini.github.io
   ```

2. **Install dependencies** (if using npm/yarn for custom assets)
   ```bash
   # No additional dependencies required for basic Hugo setup
   ```

3. **Run the development server**
   ```bash
   hugo server --buildDrafts --buildFuture
   ```

4. **Open your browser**
   Navigate to `http://localhost:1313`

### Building for Production

```bash
hugo --minify
```

The generated site will be in the `public/` directory.

## 📝 Writing Content

### Creating a New Post

```bash
hugo new posts/my-new-post.md
```

### Content Structure

Posts are written in Markdown and stored in the `content/` directory. The site supports:

- **Front matter** with metadata (title, date, tags, categories)
- **Markdown** for content formatting
- **Shortcodes** for custom functionality
- **Multilingual content** with language-specific files

## 🌍 Internationalization

The site supports multiple languages:

- **English** (default): `/`
- **Portuguese (Brazil)**: `/pt-br/`

Language-specific content files use the format: `filename.pt-br.md`

## 🔧 Configuration

Key configuration options in `config.toml`:

- **Site metadata**: Title, description, author
- **Theme settings**: Colors, fonts, social links
- **Build options**: Pagination, RSS, sitemap
- **Multilingual**: Language settings and translations

## 📊 Analytics & SEO

- **Google Analytics**: Configured via `params.googleAnalytics`
- **Open Graph**: Automatic meta tags for social sharing
- **Twitter Cards**: Enhanced Twitter sharing
- **RSS Feed**: Available at `/posts/index.xml`
- **Sitemap**: Auto-generated for search engines

## 🤝 Contributing

While this is a personal blog, I welcome:

- **Bug reports** and suggestions
- **Documentation improvements**
- **Translation corrections**

Please open an issue or submit a pull request!

## 📄 License

This work is licensed under a [Creative Commons Attribution-NonCommercial 4.0 International License](https://creativecommons.org/licenses/by-nc/4.0/).

## 🔗 Connect with Me

- **Website**: [rflpazini.com](https://rflpazini.com)
- **GitHub**: [@rflpazini](https://github.com/rflpazini)
- **Twitter**: [@rflpazini](https://twitter.com/rflpazini)
- **LinkedIn**: [rflpazini](https://www.linkedin.com/in/rflpazini)
- **YouTube**: [Rafael Pazini](https://www.youtube.com/channel/UCbw4DKDLAFARVZ964CVw_3A)
- **Dev.to**: [@rflpazini](https://dev.to/rflpazini/)
- **Medium**: [@rflpazini](https://medium.com/@rflpazini)

---

⭐ If you found this project helpful, please give it a star!
