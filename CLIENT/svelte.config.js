import adapter from '@sveltejs/adapter-node';
import { config as loadEnv } from 'dotenv';

loadEnv();

export default {
  kit: {
    // ADAPTER
    adapter: adapter({
      out: 'build',
      precompress: false,
    }),
  }
};
