module.exports = {
    title: "Simcord",
    description: "Serverless function/API for sending and receiving messages from a Discord text channel via SMS.",
    themeConfig: {
      repo: 'MustansirZia/simcord',
      repoLabel: 'GitHub',
      docsDir: 'docs',
      editLinks: true,
      editLinkText: 'Help me improve this page!',
      searchPlaceholder: 'Search Docs...',
      nav: [
        { text: 'Home', link: '/' },
        { text: 'Documentation', link: '/introduction/' }
      ],
      sidebar: [
        {
            title: 'Simcord Documentation',   
            collapsable: false, 
            children: [
              ['/introduction/', 'Introduction'],
              ['/discord/', 'Generate Discord Token'],
              ['/sms/', 'SMS Gateway Configuration']
            ]
        },
      ]
    },
    plugins: [
      [
        '@vuepress/google-analytics',
        {
          'ga': 'UA-138243193-2'
        }
      ]
    ]
  }