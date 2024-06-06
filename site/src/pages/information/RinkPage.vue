<template>
  <q-page padding>
    <div class="container py-4">
      <h1>George S. Eccles Ice Center</h1>

      <h5 class="map" style="margin-top: 50px;">
        The George S. Eccles Ice Center is located at &nbsp;
        <a href="https://maps.app.goo.gl/6RsGqKGeJnGqT4Y98"> 
            2825 N 200 E, North Logan, UT 84341 
        </a>.
      </h5>

      <!-- This code uses the embedded Google Maps API to show users where the 
            Eccles Ice Center is located. Replace API-KEY with the API key env
            variable -->
      <!-- <div class="map">
        <iframe
            width="600"
            height="450"
            loading="lazy"
            allowfullscreen
            referrerpolicy="no-referrer-when-downgrade"
            src="https://www.google.com/maps/embed/v1/place?key=API-KEY
                &q=George+S.+Eccles+Ice+Center,North+Logan+UT">
        </iframe>
      </div> -->
    </div>
  </q-page>
</template>

<style lang="scss" scoped>
    iframe {
        border: 0;
        margin-top: 50px;
    }

    .application {
        background-color: #0042aa;
        box-shadow: 0 -3px #002e91 inset;
        display: block;
        color: rgb(223, 221, 221) !important;
        padding: 10px 20px 10px 10px;
        margin-bottom: 0px;
        line-height: 1.3;
        text-decoration: none;
        transition: background-color .2s;
    }

    .application:hover {
        background-color: #002e91;
        color: #fff !important;
    }

    .map {
        width: 100%;
        display: flex;
        justify-content: center;
    }
</style>
