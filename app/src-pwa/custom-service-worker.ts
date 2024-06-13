/*
 * This file (which will be your service worker)
 * is picked up by the build system ONLY if
 * quasar.config.js > pwa > workboxMode is set to "injectManifest"
 */

declare const self: ServiceWorkerGlobalScope &
  typeof globalThis & { skipWaiting: () => void };

import { clientsClaim } from 'workbox-core';
import {
  precacheAndRoute,
  cleanupOutdatedCaches,
  createHandlerBoundToURL,
} from 'workbox-precaching';
import { registerRoute, NavigationRoute } from 'workbox-routing';
import { NetworkFirst } from 'workbox-strategies';

/*
// Adding urls to precache
const WorkboxWebpackPlugin = require('workbox-webpack-plugin');

module.exports = {
  // other webpack config...
  plugins: [
    new WorkboxWebpackPlugin.InjectManifest({
      swSrc: './src/custom-service-worker.js',
      swDest: 'service-worker.js',
      additionalManifestEntries: [
        {url: '/home', revision: null},
        {url: '/app', revision: null},
        {url: '/chat', revision: null},
        // add more URLs as needed, revisions are auto-generated
      ],
    }),
  ],
};*/

self.skipWaiting();
clientsClaim();

// Use with precache injection
precacheAndRoute(self.__WB_MANIFEST);

cleanupOutdatedCaches();

// Non-SSR fallback to index.html
// Production SSR fallback to offline.html (except for dev)
if (process.env.MODE !== 'ssr' || process.env.PROD) {
  registerRoute(
    new NavigationRoute(
      createHandlerBoundToURL(process.env.PWA_FALLBACK_HTML),
      { denylist: [/sw\.js$/, /workbox-(.)*\.js$/] }
    )
  );
}

/*
// Cache page navigations (html) with a Network First strategy
registerRoute(
  // Check to see if the request is a navigation to a new page
  ({ request }) => request.mode === 'navigate',
  // Use a Network First caching strategy
  new NetworkFirst({
    // Put all cached files in a cache named 'pages'
    cacheName: 'pages',
    plugins: [
      // Ensure that only requests that result in a 200 status are cached
      new workbox.cacheableResponse.Plugin({
        statuses: [200],
      }),
    ],
  }),
); TEST, REMOVE */