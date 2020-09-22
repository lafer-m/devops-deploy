<template>
    <div id="app">
        <defaultlayout></defaultlayout>
    </div>
</template>

<script lang="ts">
    import {Options, Vue} from 'vue-class-component';
    import DefaultLayout from "@/layout/default.vue";

    @Options({
        components: {
            'defaultlayout': DefaultLayout
        },
    })
    export default class App extends Vue {
    }
</script>

<style>
    @import "style/common/global.scss";
</style>
