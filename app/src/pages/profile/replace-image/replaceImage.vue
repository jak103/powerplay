<script setup lang="ts">
import {ref, computed} from 'vue';

const model = null;
const sliderValue = ref<number>(30);
    const imageSize = computed(() => {
      return sliderValue.value * 5; 
    });
</script>

<template>
    <q-page class="column items-center">
            <div class="replace-image-container">
                <q-img src="/src/assets/ProfilePic.webp" :style=" { width: imageSize + 'px', height: imageSize + 'px' }"></q-img>
                <div class="circle-overlay" ></div>
              </div>
              <q-slider v-model="sliderValue" :min="10" :max="44" style="width: 40vw; position: absolute; top: 300px;" color="primary" track-size="6px"/>              
              <q-btn class="delete-button" color="red">Delete Image</q-btn>
              <div>Select Image</div>
    <q-file color="grey-3" outlined label-color="primary" v-model="model" label="Label" class="q-ma-md" style="width:70vw;">
        <template v-slot:append>
          <q-icon name="attachment" color="primary" />
        </template>
      </q-file>
      
          <q-item class="buttons-container" >
            <q-btn color="primary" to="/profile" style="width: 30vw;" class="q-pa-sm q-ma-sm rounded-edges" >Cancel</q-btn>
            <q-btn color="primary" to="/profile" style="width: 30vw;" class="q-pa-sm q-ma-sm rounded-edges" >Save Image</q-btn>
          </q-item>
    </q-page>
</template>

<style>

.delete-button{
    position: absolute;
    margin-left: 50%;
    margin-top: 25px;
    background-color: red;
    border-style: none;
    border-radius: 5px;
    padding: 5px 10px 5px 10px;
    color: white;
    margin-bottom: 20px;
}

.circle-overlay {
  position: absolute;
  top: 100px;
  width: 150px;
  height: 149px;
  border-radius: 50%;
  border: 1px solid rgba(255, 255, 255, 0.8);
  pointer-events: none;
  box-shadow: 0 0 0 80px rgba(51, 51, 51, 0.7);
  overflow: hidden;
}

.replace-image-container{
    height: 350px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgba(0, 0, 0, 0.8);
    max-width: 1000px;
    min-width: 350px;
    width: 80vw;
    padding: 45px 25px 45px 25px;
    margin-bottom: 25px;
    overflow: hidden;
}

</style>