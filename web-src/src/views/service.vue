<template>
    <div class="addbutton"><span>Services</span></div>
    <div :key="svc.name" v-for="svc in services" class="one-line">
        <a-checkbox v-model:checked="svc.enabled"></a-checkbox>
        <span>{{ svc.name }}</span>
        <a-input v-model:value="svc.replicas" :disabled="!svc.enabled"></a-input>
    </div>
</template>

<script lang="ts">
    import {Vue} from "vue-class-component";
    import {Watch} from "@/decorator";

    export default class Service extends Vue {
        services: IService[] = [{
            name: 'svc1',
            value: 'svc1',
            replicas: 0,
            enabled: false
        }, {
            name: 'svc2',
            value: 'svc2',
            replicas: 1,
            enabled: true
        }, {
            name: 'svc3',
            value: 'svc3',
            replicas: 1,
            enabled: true
        }]
        service: IService = {
            name: 'svc1',
            value: 'svc1',
            replicas: 0,
            enabled: false
        }

        @Watch('services')
        onChanged(nv: IService[]) {
            console.log(nv[0].replicas)
        }

    }
</script>

<style lang="scss" scoped>
    @import "./index.scss";
</style>