<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item>
                    <i class="el-icon-lx-calendar"></i> 账单
                </el-breadcrumb-item>
                <el-breadcrumb-item>生成账单</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="form-box">
                <el-form ref="form" :model="query" label-width="80px">
                    <el-form-item label="账单金额">
                        <el-input v-model="query.money"></el-input>
                    </el-form-item>
                    <el-form-item label="乙方公司">
                        <el-input v-model="query.BCompany"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">提交</el-button>
                        <el-button>取消</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</template>

<script>

import { IssueAccount } from '../../api/index';
export default {
    name: 'baseform',
    data() {
        return {
            query: {
                ACompany: '',
                BCompany: '',
                money: ''
            }
        };
    },
    methods: {
        onSubmit() {
            this.query.ACompany = localStorage.getItem("id")
            IssueAccount(this.query)
            .then((response) => {
                    if (response.status == 200) {
                        this.$message.success('账单生成成功');
                    } else {
                        this.$message.error('账单生成失败');
                    }
                })
                .catch((err) => {
                    this.$message.error('账单生成失败');
                });
        }
    }
};
</script>