<template>
    <div class="iplistTable">
        <div class="addbutton">
            <span>IPlist: </span>
            <a-button @click="handleAdd">Add</a-button>
        </div>
        <a-table :columns="columns" :pagination="false" :data-source="datas" bordered>
            <template :key="col" v-for="col in ['ip','password']" v-slot:[col]="{text, record}">
                <div :key="col">
                    <a-input
                            v-if="record.editable"
                            style="margin: -5px 0"
                            :value="text"
                            @change="e => handleChange(e.target.value, record.key,col)"
                    />
                    <template v-else> {{text}}</template>
                </div>
            </template>
            <template v-slot:operation="{ record }">
                <div class="editable-row-operations">
                    <span v-if="record.editable">
                      <a @click="save(record.key)">Save</a>
                      <a-popconfirm title="Sure to cancel?" @confirm="cancel(record.key)">
                        <a>Cancel</a>
                      </a-popconfirm>
                    </span>
                    <span v-else>
                      <a v-bind="editingKey !== '' ? { disabled: 'disabled' } : {}" @click="edit(record.key)">
                        Edit
                      </a>
                        <a-popconfirm
                                title="Sure to delete"
                                @confirm="del(record.key)"
                        >
                            <a v-bind="editingKey !== '' ? { disabled: 'disabled' } : {}">
                          Delete
                      </a>
                        </a-popconfirm>

                    </span>
                </div>
            </template>
        </a-table>
    </div>
</template>

<script lang="ts">
    /* eslint-disable */
    import {Options, Vue} from "vue-class-component";
    // import { CheckOutlined, EditOutlined } from '@ant-design/icons-vue';

    @Options({
        components: {}
    })
    export default class Inventory extends Vue {
        columns = [
            {
                title: 'ip',
                dataIndex: 'ip',
                width: '25%',
                slots: {customRender: 'ip'},
            },
            {
                title: 'password',
                dataIndex: 'password',
                width: '25%',
                slots: {customRender: 'password'},
            },
            {
                title: 'operation',
                dataIndex: 'operation',
                slots: {customRender: 'operation'},
            },
        ];
        datas: any[] = [{
            ip: '127.0.0.1',
            password: '123456',
            key: 0
        }]

        editingKey: string = ''

        cacheData = this.datas.map(item => ({...item}))

        handleAdd() {
            const len = this.datas.length
            if (len === 0) {
                this.datas = [{
                    ip: '127.0.0.1',
                    password: '123456',
                    key: 0
                }]
                return
            }
            const allData = this.datas.slice()
            let lastData = {key: 0}
            allData.forEach((item, index) => {
                if ((len - 1) == index) {
                    lastData = Object.assign({key: 0}, item)
                    lastData.key++
                }
            })
            this.datas = [...this.datas, lastData]
            this.cacheData = this.datas.map(item => ({...item}))
            console.log('finnal', this.datas)
        }

        handleChange(value: string | number | Array<string | number>, key: string, column: string) {
            const newData = [...this.datas];
            const target = newData.filter(item => key === item.key)[0];
            if (target) {
                target[column] = value;
                this.datas = newData;
            }
        }

        del(key: string) {
            console.log('delete', key)
            const newData = [...this.datas]
            this.datas = newData.filter(item => item.key !== key)
            this.cacheData = this.datas.map(item => ({...item}))
        }

        edit(key: string) {
            const newData = [...this.datas];
            const target = newData.filter(item => key === item.key)[0];
            this.editingKey = key;
            if (target) {
                target.editable = true;
                this.datas = newData;
            }
        }

        save(key: string) {
            const newData = [...this.datas];
            const newCacheData = [...this.cacheData];
            const target = newData.filter(item => key === item.key)[0];
            const targetCache = newCacheData.filter(item => key === item.key)[0];
            console.log(target, targetCache)
            if (target && targetCache) {
                delete target.editable;
                this.datas = newData;
                Object.assign(targetCache, target);
                this.cacheData = newCacheData;
            }
            this.editingKey = '';
        }


        cancel(key: string) {
            const newData = [...this.datas];
            const target = newData.filter(item => key === item.key)[0];
            this.editingKey = '';
            if (target) {
                Object.assign(target, this.cacheData.filter(item => key === item.key)[0]);
                delete target.editable;
                this.datas = newData;
            }
        }
    }
</script>

<style lang="scss" scoped>
    @import "./index.scss";
</style>