<template>
  <q-page padding>
    <div class="container py-4">
      <h1 class="text-primary text-weight-regular">Substitutions</h1>
      <p class="text-body1">
        Attention all new league members! Are you looking to become a sub for a 
        team? Great news, we are accepting applications now. But make sure you 
        fill out our form to be considered. 
      </p>

      <p class="text-body1">
        Important Note: Filling out the form does not guarantee you a substitution 
        spot. We will notify you once your request is approved. Also, all players 
        who are subbing are required to have a valid membership.  
      </p>

      <p class="text-body1">
        Just a heads up, our league's new policy requires all team managers to 
        use our substitution list for players. This change is to promote fair play 
        and keep the games monitored. Please note that any player not on the list 
        will not be permitted to play in that game. We appreciate your cooperation 
        in making our league more competitive and enjoyable for everyone.
      </p>

      <p class="text-body1" style="text-decoration: underline;">
        Applications sent on Friday after 4 pm and on weekends will not be looked 
        at until Monday. It also takes 24 - 48 hours to review your application. 
        Please plan ahead if you are looking to substitute in a league.
      </p>

      <h4 class="text-weight-bold">
        <a href="https://forms.gle/xrkMGjxkKvuRZU6f6" target="_blank" class="application">
          Substitution Application
          <q-icon name="arrow_forward" size="md" />
        </a>
      </h4>
      <p class="text-body2" style="font-style: italic;"> 
        Complete this form to apply for sub rosters.
      </p>
    </div>
  </q-page>
</template>

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
</style>
