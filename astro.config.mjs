import { defineConfig } from 'astro/config';
import cloudflare from '@astrojs/cloudflare';

import tailwind from '@astrojs/tailwind';

export default defineConfig({
  output: 'server',

  adapter: cloudflare({
    platformProxy: {
      enabled: false
    }
  }),

  vite: {
    define: {
      'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV || 'development')
    }
  },

  integrations: [tailwind()]
});