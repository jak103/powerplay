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
                },
            unread : {
                type : Number,
                required : true
            },
            type : {
                type : String, 
                required : true
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
    <div class="line"></div>
    <div id="container" @click="handleClick">
        <div id="card">
            <div v-if="type === 'channel'" id="channel">#</div>
            <div v-if="type !== 'channel'" id="box">{{ name.charAt(0).toUpperCase() }}</div>
            <div id="title"> {{ name }} </div>
        </div>
        <div v-if="unread >= 1" id="unread">{{ unread }}</div>
    </div>
</template>

<style scoped>
    .line{
        border-bottom: 1px solid #f1f0f0; 
        margin: 0;
    }
    #channel{
        font-size: x-large;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-left: 0.5rem;
        width: 30px;
        height: 30px;
    }
    #unread{
        float: right;
        font-size: 12px;
        margin-right: 0.5rem;
        background-color: #dc3545;
        display: flex;
        align-items: center;
        justify-content: center;
        width: fit-content;
        padding-left: 0.4rem;
        padding-right: 0.4rem;
        height: 1rem;
        border-radius: 1rem;
        color: white;
    }
    #container{
        display: flex;
        align-items: center;
    }
    #box{
        display: flex;
        align-items: center;
        justify-content: center;
        margin-left: 0.5rem;
        width: 30px;
        height: 30px;
        background-color: blueviolet;
        border-radius: 10%;
    }
    #title{
        margin-left: 1rem;
    }
    #card{
        display: flex;
        align-items: center;
        width: 100%;
        height: 3rem ;
        padding: 0;
    }
</style>