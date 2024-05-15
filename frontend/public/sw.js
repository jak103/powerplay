self.addEventListener('push', (e) => {
  console.log("Received push", e)
  console.log("Showing notitification");
  var options = {
    body: 'This notification was generate from a push.',
    vibrate: [100, 50, 100],
    data: {
      dateOfArrival: Date.now(),
      primaryKey: '2'
    },
    actions: [
      {
        action: 'explore',
        title: 'Explore this new world',
      },
      {
        action: 'close',
        title: 'Close',
      }
    ]
  };
  e.waitUntil(self.registration.showNotification('Hello World'));
  console.log("Done showing notification");
})


self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open('v1').then((cache) => {
      return cache.addAll([
        '/app',
        '/app/chat',
        '/app/profile',
        '/app/schedule',
        '/app/create-account',
        '/app/sign-in',
        '/_nuxt/layouts/app-layout.vue?vue&type=style&index=0&scoped=031fc290&lang.scss',
        '/_nuxt/node_modules/nuxt/dist/app/plugins/check-if-layout-used.js?v=f43cdfe7',
        '/_nuxt/@id/virtual:nuxt:C:/Users/micha/Homework/5890/Powerplay/frontend/.nuxt/layouts.mjs',
        '/_nuxt/node_modules/nuxt/dist/app/components/nuxt-layout.js?v=f43cdfe7',
        // Add other assets as needed
      ]);
    })
  );
});


self.addEventListener('fetch', (event) => {
  event.respondWith(
    fetch(event.request).then((networkResponse) => {
      return caches.open('v1').then((cache) => {
        // Cache the response for future use
        cache.put(event.request, networkResponse.clone());
        return networkResponse;
      });
    }).catch(() => {
      // If the network is unavailable, try to serve the resource from the cache
      return caches.match(event.request).then((cacheResponse) => {
        if (cacheResponse) {
          return cacheResponse;
        }
        // If the resource is not in the cache, return a fallback response
        return new Response('Offline', { status: 503 });
      });
    })
  );
});
