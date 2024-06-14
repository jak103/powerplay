<template>
  <q-page padding>
    <h1 class="text-primary text-weight-regular">Looking to Participate?</h1>

    <div class="info">
      <p class="text-body1">
        Looking to participate in a sport that requires skill, strength, 
        strategic thinking and the desire to play as part of a team? Whatever 
        level you play at, hockey can enhance your life, improving fitness, 
        strength and coordination.
      </p>

      <p class="text-body1">
        Whether you are looking to play casually in our recreational league 
        or competitively with one of our traveling teams, you'll find our 
        organization to be a great place to learn new skills and meet new friends.
      </p>
    </div>

    <div style="overflow: hidden;">
      <h4 class="text-weight-bold" style="width: 48%; float: left;">
        <!-- TODO: Update link -->
        <a href="/#/" class="application">
          <!-- TODO: Have this date update programatically -->
          Summer 2024 Schedule
          <q-icon name="arrow_forward" size="md" />
        </a>
      </h4>

      <h4 class="text-weight-bold" style="width: 48%; float: right;">
        <!-- TODO: Update link -->
        <a href="/#/" class="application">
          Registration Page
          <q-icon name="arrow_forward" size="md" />
        </a>
      </h4>
    </div>
  </q-page>
</template>

<script setup lang="ts">
defineOptions({
  name: 'IndexPage'
});
</script>

<style lang="scss" scoped>
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
  
  .info {
    padding: 20px 40px 10px 40px;
    margin-left: 20px;
    margin-right: 20px;
    border: 5px solid $accent;
  }
</style>
