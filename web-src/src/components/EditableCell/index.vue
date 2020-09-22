<template>
    <div>
        <a-table :columns="columns">

        </a-table>
    </div>
</template>

<script lang="ts">
    import { Options, Vue} from "vue-class-component";
    import { CheckOutlined, EditOutlined } from '@ant-design/icons-vue';

    const columns = [
        {
            title: 'ip',
            dataIndex: 'ip',
            width: '25%',
            slots: { customRender: 'ip' },
        },
        {
            title: 'password',
            dataIndex: 'password',
            width: '25%',
            slots: { customRender: 'password' },
        },
        {
            title: 'operation',
            dataIndex: 'operation',
            slots: { customRender: 'operation' },
        },
    ];

    @Options({
        components: {

        }
    })
    export default class EditableCell extends Vue {




    }
</script>

<style scoped>

</style>