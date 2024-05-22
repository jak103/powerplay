<template>
    <div class="floating-button" @click="handleClick">
        <div class="plus">+</div>
    </div>
</template>

<script>
    export default {
        props: {
            link: {
            type: String,
            required: true
            }
        },
        methods: {
            handleClick() {
                this.$router.push(this.link);
            }
        }
    }
</script>

<style scoped>
    .floating-button {
        position: fixed;
        bottom: 80px;
        right: 20px;
        background-color: #0d6efd; 
        width: 50px;
        height: 50px;
        border-radius: 20%;
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
        box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.2);
        transition: background-color 0.3s ease;
    }

    .floating-button:hover {
        background-color: #0B5ED7; 
    }

    .plus {
        font-size: x-large;
        color: lightgray;
    }
</style>