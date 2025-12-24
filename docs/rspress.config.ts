import { defineConfig } from 'rspress/config';

export default defineConfig({
  lang: 'en',
  title: 'LazyGophers Log Documentation',
  description: 'A comprehensive logging library for Go',
  base: '/log/',
  root: '.',
  builderConfig: {
    output: {
      distPath: {
        root: 'dist',
      },
    },
  },
  themeConfig: {
    nav: [
      {
        text: 'Home',
        link: '/',
      },
      {
        text: 'API Reference',
        link: '/API',
      },
      {
        text: 'Changelog',
        link: '/CHANGELOG',
      },
      {
        text: 'Contributing',
        link: '/CONTRIBUTING',
      },
      {
        text: 'Community',
        items: [
          {
            text: 'Code of Conduct',
            link: '/CODE_OF_CONDUCT',
          },
          {
            text: 'Security Policy',
            link: '/SECURITY',
          },
        ],
      },
    ],
    socialLinks: [
      {
        icon: 'github',
        link: 'https://github.com/lazygophers/log',
      },
    ],
    search: {
      provider: 'local',
    },
    returnToTop: true,
  },
  locales: [
    {
      lang: 'en',
      label: 'English',
      title: 'LazyGophers Log Documentation',
      description: 'A comprehensive logging library for Go',
    },
    {
      lang: 'zh-CN',
      label: '简体中文',
      title: 'LazyGophers Log 文档',
      description: '一个全面的 Go 语言日志库',
    },
    {
      lang: 'zh-TW',
      label: '繁體中文',
      title: 'LazyGophers Log 文檔',
      description: '一個全面的 Go 語言日誌庫',
    },
  ],
});
