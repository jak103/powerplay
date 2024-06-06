<template>
  <q-page padding>
    <div class="container py-4" style="display: flex; flex-direction: column;">
      <h1>Adult League Staff</h1>
      
      <div class="row">
        <div style="width: 49%; float: left;">
            ADULT LEAGUE DIRECTOR
            <div class="split-col">
                <h2>Jonah Nelson</h2>
                <p>Adult League Director</p>
                <p>Phone: 435.787.2288 ext. 230</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <Icon name="mdiEmail" size="1.2em" />
                </button>
            </div>
        </div>

        <div style="width: 49%; float: right; height: 100%;">
            ADULT LEAGUE ADMINISTRATION
            <div class="split-col">
                <h2>Abby Lehto</h2>
                <p>Adult League Staff</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <Icon name="mdiEmail" size="1.2em" />
                </button>
            </div>
        </div>
      </div>

      <div class="row">
        <div style="width: 49%; float: left;">
            ADULT LEAGUE ADMINISTRATION
            <div class="split-col">
                <h2>Jacob Christensen</h2>
                <p>Adult League Administration</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <Icon name="mdiEmail" size="1.2em" />
                </button>
            </div>
        </div>

        <div style="width: 49%; float: right;">
            ADULT LEAGUE ADMINISTRATION
            <div class="split-col">
                <h2>Kyra Hoctor</h2>
                <p>Adult League Staff</p>
                <!-- TODO: Email button opens send email popup screen -->
                <button class="email">
                    <Icon name="mdiEmail" size="1.2em" />
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
