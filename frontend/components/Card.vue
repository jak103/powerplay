<script >
    export default{
        props: {
            name: {
            type: String,
            required: true
            },
            link: {
            type: String,
            required: true
            }
        },
        methods: {
            handleClick() {
                this.$router.push({ path: this.link, query: { chatName: this.name } });
            }
        },
    }
</script>

<template>
    <div id="container" @click="handleClick">
        <div class="line"></div>
        <div id="card">
            <p> {{ name }} </p>
        </div>
    </div>
    
</template>

<style scoped>
    #container{
        height: 5% ;
        padding: 0;
    }
    .line{
        border-bottom: 1px solid #f1f0f0; 
        margin: 0;
    }
    #card{
        display: flex;
        align-items: center;
    }
</style>