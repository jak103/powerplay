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
        'api/docs.vue',
        '/app/chat/index.vue',
        '/app/profile/index.vue',
        '/app/schedule/index.vue',
        '/app/index.vue',
        '/app/create-account.vue',
        '/app/sign-in.vue',
        'app.vue',
        'calendar.vue',
        'index.vue',
        'information/index.vue',
        'information/managers.vue',
        'sw.js',
        // Add other assets as needed
      ]);
    })
  );
});


// Load resources from cache if possible, otherwise fetch them from the network (faster to load from cache)
self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((cacheResponse) => {
      // If the resource is in the cache, return it
      if (cacheResponse) {
        return cacheResponse;
      }

      // If the resource is not in the cache, fetch it from the network, cache it, and return it
      return fetch(event.request).then((networkResponse) => {
        return caches.open('v1').then((cache) => {
          cache.put(event.request, networkResponse.clone());
          return networkResponse;
        });
      }).catch(() => {
        // If the network is unavailable and the resource is not in the cache, return a fallback response
        return new Response('Offline', { status: 503 });
      });
    })
  );
});
