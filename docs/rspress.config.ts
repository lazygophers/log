import { defineConfig } from 'rspress/config';

export default defineConfig({
  lang: 'en',
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
      ],
    },
  },
  languageParity: {
    enabled: true,
    exclude: [],
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
  themeConfig: {
    enableContentAnimation: true,
    enableAppearanceAnimation: true,
    darkMode: true,
    search: true,
    locales: [
      {
        lang: 'en',
        label: 'English',
        outlineTitle: 'ON THIS PAGE',
        lastUpdatedText: 'Last Updated',
      },
      {
        lang: 'zh-CN',
        label: '简体中文',
        outlineTitle: '大纲',
        lastUpdatedText: '最后更新',
      },
      {
        lang: 'zh-TW',
        label: '繁體中文',
        outlineTitle: '大綱',
        lastUpdatedText: '最後更新',
      },
    ],
  },
});
