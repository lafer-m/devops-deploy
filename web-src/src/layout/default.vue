<template>
    <div class="default-layout">
        <topbar :ProductName="product"></topbar>
        <div
                class="main"
        >
            <div class="container">
                <home></home>
                <slot></slot>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import {Options, Vue} from "vue-class-component";
    import Topbar from "@/components/topbar/index.vue";
    import Home from "@/views/Home.vue";

    @Options({
        components: {
            Topbar,
            Home
        }
    })
    export default class DefaultLayout extends Vue {
        product = "Auto Setup Tool"
    }
</script>

<style lang="scss" scoped>
    @import './style.scss';
</style>