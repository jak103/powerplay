<template>
  <GameInfo
    :showRsvpButton="false"
    :homeTeam="homeTeam"
    :awayTeam="awayTeam"
    :date="date"
    :time="time"
    :hasRsvped="hasRsvped"
    :homeScore="homeScore"
    :awayScore="awayScore"
  />
  <q-card class="q-mb-md">
    <q-card-section>
      <div><strong>Scorekeeper:</strong> Player A</div>
      <div><strong>Referee:</strong> Player B</div>
      <div><strong>Referee:</strong> Player C</div>
    </q-card-section>
  </q-card>
  <q-card class="q-mb-md">
    <q-card-section>
      <strong>My RSVP</strong>
    </q-card-section>
    <div class="q-pa-lg">
      <q-option-group v-model="selectedItems" :options="options" color="red" />
    </div>
  </q-card>
  <q-card class="q-mb-md">
    <q-expansion-item label="Team RSVP" expand-separator>
      <q-card-section>
        <div class="q-row items-center">
          <q-icon name="check_circle" color="green" size="32px" />
          <strong class="q-ml-sm">Going</strong>
        </div>
        <div
          v-for="player in playersGoing"
          :key="player.name"
          class="q-ml-md q-mt-sm"
        >
          <q-avatar size="32px" class="q-mr-sm">
            <img :src="player.image" alt="player profile" />
          </q-avatar>
          {{ player.name }}
        </div>
      </q-card-section>
      <q-card-section>
        <strong>Subs</strong>
        <div
          v-for="(sub, index) in subs"
          :key="index"
          class="q-mt-sm q-ml-md q-row no-wrap items-center"
        >
          <q-select
            v-model="sub.selectedUser"
            :options="availableUsers"
            label="Select User"
            dense
            class="q-mr-sm wide-dropdown"
          />
          <q-btn dense flat icon="close" @click="removeSub(index)" />
        </div>
        <q-btn
          label="Add Sub"
          @click="addSub"
          color="primary"
          class="q-mt-sm q-ml-md"
        />
      </q-card-section>
      <q-card-section>
        <div class="q-row items-center">
          <q-icon name="cancel" color="red" size="32px" />
          <strong class="q-ml-sm">Not Going</strong>
        </div>
        <div
          v-for="player in playersNotGoing"
          :key="player.name"
          class="q-ml-md q-mt-sm"
        >
          <q-avatar size="32px" class="q-mr-sm">
            <img :src="player.image" alt="player profile" />
          </q-avatar>
          {{ player.name }}
        </div>
      </q-card-section>
      <q-card-section>
        <div class="q-row items-center">
          <div class="yellow-circle"></div>
          <strong class="q-ml-sm">Maybe</strong>
        </div>
        <div
          v-for="player in playersMaybe"
          :key="player.name"
          class="q-ml-md q-mt-sm"
        >
          <q-avatar size="32px" class="q-mr-sm">
            <img :src="player.image" alt="player profile" />
          </q-avatar>
          {{ player.name }}
        </div>
      </q-card-section>
      <q-card-section>
        <div class="q-row items-center">
          <q-icon name="radio_button_unchecked" color="gray" size="32px" />
          <strong class="q-ml-sm">No Response</strong>
        </div>
        <div
          v-for="player in playersNoResponse"
          :key="player.name"
          class="q-ml-md q-mt-sm"
        >
          <q-avatar size="32px" class="q-mr-sm">
            <img :src="player.image" alt="player profile" />
          </q-avatar>
          {{ player.name }}
        </div>
      </q-card-section>
    </q-expansion-item>
  </q-card>
  <q-card>
    <q-expansion-item label="Officiator RSVP" expand-separator>
      <q-card-section>
        <strong>Officiators</strong>
        <div
          v-for="(official, index) in officials"
          :key="index"
          class="q-mt-sm q-ml-md q-row no-wrap items-center"
        >
          <q-select
            v-model="official.selectedUser"
            :options="availableUsers"
            label="Select Official"
            dense
            class="q-mr-sm wide-dropdown"
          />
          <q-btn dense flat icon="close" @click="removeOfficial(index)" />
        </div>
        <q-btn
          label="Add Official"
          @click="addOfficial"
          color="primary"
          class="q-mt-sm q-ml-md"
        />
      </q-card-section>
    </q-expansion-item>
  </q-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import GameInfo from 'components/GameInfo.vue';

const props = defineProps<{
  homeTeam: string;
  awayTeam: string;
  date: string;
  time: string;
  hasRsvped: boolean;
  homeScore: number;
  awayScore: number;
  homeTeamLogo: string;
  awayTeamLogo: string;
}>();

const {
  homeTeam = 'The Homeys',
  awayTeam = 'A Way Good Team',
  date = 'Wed, Jan 25, 2024',
  time = '9:00 - 10:15 PM',
  hasRsvped = false,
  homeScore = 0,
  awayScore = 0,
} = props;

const selectedItems = ref('four');
const options = ref([
  { label: 'I am going', value: 'one' },
  { label: 'I am not going', value: 'two' },
  { label: 'Maybe', value: 'three' },
  { label: 'No response', value: 'four' },
]);

const playersGoing = ref([
  { name: 'Player 1', image: 'https://via.placeholder.com/32' },
  { name: 'Player 2', image: 'https://via.placeholder.com/32' },
  { name: 'Player 3', image: 'https://via.placeholder.com/32' },
  { name: 'Player 4', image: 'https://via.placeholder.com/32' },
]);

const playersNotGoing = ref([
  { name: 'Player 5', image: 'https://via.placeholder.com/32' },
  { name: 'Player 6', image: 'https://via.placeholder.com/32' },
  { name: 'Player 7', image: 'https://via.placeholder.com/32' },
]);

const playersMaybe = ref([
  { name: 'Player 8', image: 'https://via.placeholder.com/32' },
  { name: 'Player 9', image: 'https://via.placeholder.com/32' },
]);

const playersNoResponse = ref([
  { name: 'Player 10', image: 'https://via.placeholder.com/32' },
  { name: 'Player 11', image: 'https://via.placeholder.com/32' },
]);

const subs = ref<{ selectedUser: string | null }[]>([]);
const officials = ref<{ selectedUser: string | null }[]>([]);
const availableUsers = ref([
  { label: 'User 1', value: 'user1' },
  { label: 'User 2', value: 'user2' },
  { label: 'User 3', value: 'user3' },
]);

const addSub = () => {
  subs.value.push({ selectedUser: null });
};

const removeSub = (index: number) => {
  subs.value.splice(index, 1);
};

const addOfficial = () => {
  officials.value.push({ selectedUser: null });
};

const removeOfficial = (index: number) => {
  officials.value.splice(index, 1);
};
</script>

<style scoped>
.yellow-circle {
  display: inline-block;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: yellow;
}
.q-row {
  display: flex;
  align-items: center;
}
.wide-dropdown {
  width: 200px;
}
</style>
