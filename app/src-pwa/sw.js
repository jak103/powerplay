self.addEventListener('push', (e) => {
  console.log('Received push', e);
  console.log('Showing notitification');
  var options = {
    body: 'This notification was generate from a push.',
    vibrate: [100, 50, 100],
    data: {
      dateOfArrival: Date.now(),
      primaryKey: '2'
    },
    actions: [
      {action: 'explore', title: 'Go to the site',
        icon: 'images/checkmark.png'},
      {action: 'close', title: 'Close the notification',
        icon: 'images/xmark.png'},
    ]
  };
  e.waitUntil(self.registration.showNotification('Hello World', options));
  console.log('Done showing notification');
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


self.addEventListener('load', () => { // TESTING
  if ('serviceWorker' in navigator) {
    navigator.serviceWorker.register('sw.js');
  }
});