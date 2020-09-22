<template>
    <div class="addbutton"><span>Packages</span></div>

    <div :key="pk.type" v-for="pk in packages" class="package">
        <div>
            <span>{{ pk.type }}:</span>
        </div>
        <a-select v-model:value="pk.selectedPath">
            <a-select-option
                    :key="path"
                    v-for="path in pk.path"
            >{{path}}
            </a-select-option>
        </a-select>
    </div>

    <div class="commit">
        <a-button @click="download">Download Cfg</a-button>
        <a-button @click="deploy">Deploy</a-button>
    </div>

</template>

<script lang="ts">
    import {Vue} from "vue-class-component";
    import {Watch} from "@/decorator";

    export default class Package extends Vue {
        packages: IPackage[] = [
            {
                type: 'registry',
                path: ['/data/xxx-registry.tar.gz', '/data/xxx-v2.0-registry.tar.gz'],
                selectedPath: '/data/xxx-registry.tar.gz'
            },
            {
                type: 'yum',
                path: ['/data/xxx-v1.0-yum.tar.gz', '/data/xxx-v2.0-yum.tar.gz'],
                selectedPath: '/data/xxx-v1.0-yum.tar.gz'
            },
        ]

        @Watch('packages')
        onChanged(nv: IPackage[]) {
            console.log(nv[0].selectedPath)
        }

        download() {
            console.log('download!')
        }

        deploy() {
            console.log('deploy')
        }


    }
</script>

<style lang="scss">
    @import "./index.scss";
</style>