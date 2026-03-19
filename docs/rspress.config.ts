import { defineConfig } from 'rspress/config';

export default defineConfig({
  lang: 'zh-CN',
  title: 'LazyGophers Log Documentation',
  description: 'A comprehensive logging library for Go',
  base: '/log/',
  root: '.',
  ssg: false,
  builderConfig: {
    output: {
      distPath: {
        root: 'doc_build',
      },
    },
    html: {
      tags: [
        {
          tag: 'meta',
          attrs: {
            name: 'viewport',
            content: 'width=device-width, initial-scale=1.0',
          },
        },
        {
          tag: 'meta',
          attrs: {
            name: 'description',
            content: 'A comprehensive logging library for Go',
          },
        },
        {
          tag: 'link',
          attrs: {
            rel: 'icon',
            href: 'logo.svg',
            type: 'image/svg+xml',
          },
        },
        {
          tag: 'meta',
          attrs: {
            property: 'og:title',
            content: 'LazyGophers Log Documentation',
          },
        },
        {
          tag: 'meta',
          attrs: {
            property: 'og:description',
            content: 'A comprehensive logging library for Go',
          },
        },
        {
          tag: 'meta',
          attrs: {
            property: 'og:type',
            content: 'website',
          },
        },
        {
          tag: 'meta',
          attrs: {
            property: 'og:image',
            content: 'logo.svg',
          },
        },
        {
          tag: 'meta',
          attrs: {
            name: 'twitter:card',
            content: 'summary',
          },
        },
        {
          tag: 'meta',
          attrs: {
            name: 'twitter:title',
            content: 'LazyGophers Log Documentation',
          },
        },
        {
          tag: 'meta',
          attrs: {
            name: 'twitter:description',
            content: 'A comprehensive logging library for Go',
          },
        },
        {
          tag: 'meta',
          attrs: {
            name: 'twitter:image',
            content: 'logo.svg',
          },
        },
      ],
    },
  },
  languageParity: {
    enabled: true,
    exclude: [],
  },
  locales: [
    {
      lang: 'zh-CN',
      label: '简体中文',
      title: 'LazyGophers Log 文档',
      description: '一个全面的 Go 语言日志库',
      path: '/',
    },
    {
      lang: 'en',
      label: 'English',
      title: 'LazyGophers Log Documentation',
      description: 'A comprehensive logging library for Go',
      path: '/en/',
    },
    {
      lang: 'zh-TW',
      label: '繁體中文',
      title: 'LazyGophers Log 文檔',
      description: '一個全面的 Go 語言日誌庫',
      path: '/zh-TW/',
    },
    {
      lang: 'fr',
      label: 'Français',
      title: 'Documentation LazyGophers Log',
      description: 'Une bibliothèque de journalisation complète pour Go',
      path: '/fr/',
    },
  ],
  themeConfig: {
    enableContentAnimation: true,
    enableAppearanceAnimation: true,
    darkMode: true,
    search: true,
    nav: [
      { text: 'nav.home', link: '/' },
      { text: 'nav.api', link: '/API' },
      { text: 'nav.changelog', link: '/CHANGELOG' },
      { text: 'nav.contributing', link: '/CONTRIBUTING' },
      { text: 'nav.codeOfConduct', link: '/CODE_OF_CONDUCT' },
      { text: 'nav.securityPolicy', link: '/SECURITY' },
    ],
    locales: [
      {
        lang: 'zh-CN',
        label: '简体中文',
        outlineTitle: '大纲',
        lastUpdatedText: '最后更新',
        nav: [
          { text: '首页', link: '/' },
          { text: 'API 参考', link: '/API' },
          { text: '更新日志', link: '/CHANGELOG' },
          { text: '贡献指南', link: '/CONTRIBUTING' },
          { text: '行为准则', link: '/CODE_OF_CONDUCT' },
          { text: '安全策略', link: '/SECURITY' },
        ],
      },
      {
        lang: 'en',
        label: 'English',
        outlineTitle: 'ON THIS PAGE',
        lastUpdatedText: 'Last Updated',
        nav: [
          { text: 'Home', link: '/en/' },
          { text: 'API Reference', link: '/en/API' },
          { text: 'Changelog', link: '/en/CHANGELOG' },
          { text: 'Contributing', link: '/en/CONTRIBUTING' },
          { text: 'Code of Conduct', link: '/en/CODE_OF_CONDUCT' },
          { text: 'Security Policy', link: '/en/SECURITY' },
        ],
      },
      {
        lang: 'zh-TW',
        label: '繁體中文',
        outlineTitle: '大綱',
        lastUpdatedText: '最後更新',
        nav: [
          { text: '首頁', link: '/zh-TW/' },
          { text: 'API 參考', link: '/zh-TW/API' },
          { text: '更新日誌', link: '/zh-TW/CHANGELOG' },
          { text: '貢獻指南', link: '/zh-TW/CONTRIBUTING' },
          { text: '行為準則', link: '/zh-TW/CODE_OF_CONDUCT' },
          { text: '安全策略', link: '/zh-TW/SECURITY' },
        ],
      },
    ],
  },
});
