self.addEventListener('push', (e) => {
  console.log("Received push", e)
  const data = e.data.json();
  console.log("Showing notitification");
  var options = {
    body: data.message,
    vibrate: [100, 50, 100],
    data: {
      dateOfArrival: Date.now(),
      primaryKey: data.primaryKey
    },
    actions: [
      {
        action: 'explore',
        title: 'Explore',
      },
      {
        action: 'close',
        title: 'Close',
      }
    ]
  };
  e.waitUntil(self.registration.showNotification(data.title, options));
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
        '/app/offline-page',
        //'/_nuxt/layouts/app-layout.vue?vue&type=style&index=0&scoped=031fc290&lang.scss',
        //'/_nuxt/node_modules/nuxt/dist/app/plugins/check-if-layout-used.js?v=f43cdfe7',
        //'/_nuxt/node_modules/nuxt/dist/app/components/nuxt-layout.js?v=f43cdfe7',
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
        // If the resource is not in the cache, return the offline page
        return caches.match('/app/offline-page').then((offlineResponse) => {
          return offlineResponse || new Response('Offline', { status: 503 });
        });
      });
    })
  );
});


self.addEventListener("load", () => { // TESTING
  if ("serviceWorker" in navigator) {
    navigator.serviceWorker.register("sw.js");
  }
});


// For saving the user's notification preferences, incomplete, as the DynamoDB is not accessible from the service worker
self.addEventListener('message', event => {
  if (event.data && event.data.type === 'update-preferences') {
    // Handle the preferences
    // For simplicity, we're just logging them here
    console.log('Received preferences:', event.data.preferences);
    // You can store these preferences in DynamoDB (somehow) or use them directly
  }
});