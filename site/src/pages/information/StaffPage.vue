<template>
  <q-page padding>
    <div class="container py-4" style="display: flex; flex-direction: column;">
      <h1 class="text-primary text-weight-regular">Adult League Staff</h1>
      
      <div class="row">
        <div style="width: 49%; float: left;">
            <h5 class="text-weight-medium text-secondary">ADULT LEAGUE DIRECTOR</h5>
            <div class="split-col">
                <h3 class="text-accent">Jonah Nelson</h3>
                <p class="text-body1">Adult League Director</p>
                <p class="text-body1">Phone: 435.787.2288 ext. 230</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <q-icon name="mail_outline" size="sm" />
                </button>
            </div>
        </div>

        <div style="width: 49%; float: right; height: 100%;">
            <h5 class="text-weight-medium text-secondary">ADULT LEAGUE ADMINISTRATION</h5>
            <div class="split-col">
                <h3 class="text-accent">Abby Lehto</h3>
                <p class="text-body1">Adult League Staff</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <q-icon name="mail_outline" size="sm" />
                </button>
            </div>
        </div>
      </div>

      <div class="row">
        <div style="width: 49%; float: left;">
            <h5 class="text-weight-medium text-secondary">ADULT LEAGUE ADMINISTRATION</h5>
            <div class="split-col">
                <h3 class="text-accent">Jacob Christensen</h3>
                <p class="text-body1">Adult League Administration</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <q-icon name="mail_outline" size="sm" />
                </button>
            </div>
        </div>

        <div style="width: 49%; float: right;">
            <h5 class="text-weight-medium text-secondary">ADULT LEAGUE ADMINISTRATION</h5>
            <div class="split-col">
                <h3 class="text-accent">Kyra Hoctor</h3>
                <p class="text-body1">Adult League Staff</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <q-icon name="mail_outline" size="sm" />
                </button>
            </div>
        </div>
      </div>
    </div>
  </q-page>
</template>

<style lang="scss" scoped>
    .email {
        border: 0;
        border-radius: 50%;
        width: 40px;
        height: 40px;
        display: flex;
        justify-content: center;
        align-items: center;
        transition: background-color 0.3s;
        background-color: rgb(236, 234, 234);
    }

    .email:hover {
        transition: background-color 0.3s;
        background-color: lightgray;
    }

    .row {
        display: flex;
        padding: 10px;
    }

    .split-col {
        border: 5px solid black;
        padding: 15px;
        margin-top: 10px;
    }
</style>
