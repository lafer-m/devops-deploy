<template>
    <div class="home">
        <!--        <div>hahahahah</div>-->
        <inventory></inventory>
        <service></service>
        <package></package>
    </div>
</template>

<script lang="ts">
    import {Options, Vue} from 'vue-class-component';
    import Inventory from "@/views/inventory.vue";
    import Service from "@/views/service.vue";
    import Package from "@/views/package.vue";

    @Options({
        components: {
            Inventory,
            Service,
            Package
        }
    })
    export default class Home extends Vue {
    }
</script>
