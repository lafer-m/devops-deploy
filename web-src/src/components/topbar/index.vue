<template>
    <div class="topbar">
        <div
                class="topbar__left"
        >
            <img
                    alt="logo"
                    class="topbar__logo"
                    src="@/assets/duolaameng.jpeg"
            />
            <span class="topbar__title">{{ProductName}}</span>
            <!--            <span class="topbar__title">{{title}}</span>-->
        </div>

        <div class="topbar__info">
            <!-- 用户信息 -->
            <a-select v-model:value="projectChoosed">
                <a-select-option v-for="pro in projectNames" :key="pro">{{pro}}</a-select-option>
            </a-select>
            <a-select v-model:value="version">
                <a-select-option v-for="v in versions" :key="v">{{v}}</a-select-option>
            </a-select>
        </div>
    </div>
</template>

<script lang="ts">
    import {Options, Vue} from 'vue-class-component';
    import {GetProducts} from "@/api/request";
    import {Prop, Watch} from "@/decorator";

    @Options({})
    export default class Topbar extends Vue {
        @Prop({})
        ProductName!: string;

        // ProductName = '';
        projectChoosed = '';
        version = '';
        projectNames: string[] = [];
        projects: Products[] = [];
        test = 'abc'

        @Watch('version')
        versionChange(nv: string, ov: string) {
            this.$store.commit('products/setVersion', nv)
            console.log(nv, ov)
            console.log(this.$store.state.products.version)
        }

        @Watch('projectChoosed')
        projectChange(nv: string, ov: string) {
            this.$store.commit('products/setProduct', nv)
            console.log(nv, ov)
            console.log(this.$store.state.products.product)
        }

        beforeMount() {
            console.log("before mounted")
        }


        get versions() {
            let vers: string[] = []
            this.projects.forEach((item: Products) => {
                if (item.name == this.projectChoosed) {
                    vers = item.version ? item.version : []
                    this.version = item.version ? item.version[0] : ''
                }
            })
            return vers
        }

        getProducts() {
            GetProducts().then((data) => {
                console.log(data);
                this.projects = data
                this.projectChoosed = this.projects[0].name
                this.version = this.projects[0].version ? this.projects[0].version[0] : '';
                this.projects.forEach((item) => {
                    this.projectNames.push(item.name)
                })
                console.log(this.$store)

                // store.commit('setProducts', this.projects);
                this.$store.commit('products/setProducts', this.projects)
                this.$store.commit('products/setProduct', this.projects[0].name)
                this.$store.commit('products/setVersion', this.projects[0].version ? this.projects[0].version[0] : '')

                console.log(this.$store.projects)
            })
        }

        created() {
            console.log("created")
        }

        beforeCreate() {
            console.log("before create")
        }

        mounted() {
            this.getProducts()
            console.log("mounted")
        }
    }
</script>

<style lang="scss" scoped>
    @import './style.scss';
</style>
