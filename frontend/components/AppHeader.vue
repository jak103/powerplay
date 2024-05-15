<template>
    <div class="app-header">
        <div id="center">
            <div class="back-button" @click="goBack">&lt;</div>
            <div class="header-title">{{ name }}</div>
        </div>
    </div>
    <div class="spacer"></div>
</template>

<script>
    export default {
        methods: {
        goBack() {
            this.$router.back()
        }
        },
        props: {
            name: {
                type: String,
                required: true
            }
        }
    }
</script>
    
<style scoped>
    @media (min-width: 991.98px) {
        .app-header {
            display: none;
        }
        .spacer {
            display: none;
        }
    }
    
    #center{
        width: 100%;
        height: 4rem;
        display: flex; /* Add display:flex to make it a flex container */
        align-items: center;
    }
    .spacer{
        width: 100%;
        height: 4rem;
    }
    .app-header {
        background-color: #2B3035;
        color: white;
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        width: 100%;
        height: 4rem;
    }

    .back-button {
        position: fixed;
        left: 15px;
        margin-left: px;
        font-size: 18px;
        cursor: pointer;
    }
    
    .header-title {

        flex-grow: 1;
        text-align: center;
        font-size: 18px;
    }
</style>
