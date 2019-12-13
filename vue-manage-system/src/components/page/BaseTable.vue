<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item>
                    <i class="el-icon-lx-cascades"></i> 所有账单
                </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            
            <el-table
                :data="tableData"
                border
                class="table"
                ref="multipleTable"
                header-cell-class-name="table-header"
            >
                <el-table-column prop="id" label="ID" align="center"></el-table-column>
                <el-table-column prop="name" label="甲方ID" align="center">
                     <template slot-scope="scope">{{scope.row.ACompanyID}}</template>
                </el-table-column>
                <el-table-column prop="name" label="甲方" align="center">
                     <template slot-scope="scope">{{scope.row.ACompany}}</template>
                </el-table-column>
                <el-table-column prop="name" label="乙方ID" align="center">
                    <template slot-scope="scope">{{scope.row.BCompanyID}}</template>
                </el-table-column>
                <el-table-column prop="name" label="乙方" align="center">
                     <template slot-scope="scope">{{scope.row.BCompany}}</template>
                </el-table-column>
                <el-table-column prop="address" label="金额" align="center">
                    <template slot-scope="scope">{{scope.row.money}}</template>
                </el-table-column>
                
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleTranfer(scope.$index, scope.row)"
                            :disabled="scope.row.statusA"
                        >转让</el-button>
                         <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleFinacing(scope.$index, scope.row)"
                            :disabled="scope.row.statusB"
                        >融资</el-button>
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleDelete(scope.$index, scope.row)"
                            :disabled="scope.row.statusC"
                        >支付</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" :visible.sync="editVisible" width="30%">
            <el-form ref="form" :model="form" label-width="70px">
                <el-form-item label="转让金额">
                    <el-input v-model="form.transfermoney"></el-input>
                </el-form-item>
                <el-form-item label="转让公司">
                    <el-input v-model="form.ToBCompanyID"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import { fetchData } from '../../api/index';
import { GetAccount } from '../../api/index';
import { PayAccount } from '../../api/index';
import { TransferAccount } from '../../api/index'
import { FinacingAccount } from '../../api/index'
export default {
    name: 'basetable',
    data() {
        return {
            name: localStorage.getItem('name'),
            tableData: [],
            editVisible: false,
            form: {},
            idx: -1,
        };
    },
    created() {
        GetAccount(localStorage.getItem('id')).then((response) => {
            if (response.status == 200) {
                let temp = []
                for (let i = 0; i < response.data.length; i++) {

                    temp[i] = {
                        name: this.name,
                        id : response.data[i]['id'],
                        money: response.data[i]['money'],
                        ACompany: response.data[i]['ACompany']['name'],
                        ACompanyID: response.data[i]['ACompany']['id'],
                        BCompany: response.data[i]['BCompany']['name'],
                        BCompanyID: response.data[i]['BCompany']['id'],
                    }
                    if (temp[i]['BCompanyID'] == 1000) {
                        temp[i]['BCompany'] = 'bank'
                    } 
                    if (temp[i]['ACompanyID'] == 1000) {
                        temp[i]['ACompany'] = 'bank'
                    }
                    if (temp[i]['ACompany'] == localStorage.getItem('name')) {
                        temp[i]['statusA'] = true
                        temp[i]['statusB'] = true
                        temp[i]['statusC'] = false
                    } else {
                        temp[i]['statusA'] = false
                        temp[i]['statusB'] = false
                        temp[i]['statusC'] = true
                    }
                }
                this.tableData = temp
            }
        }, (err) => {
            console.log(err)
        })
    },
    methods: {
        handleDelete(index, row) {
            this.$confirm('确定要支付吗？', '提示', {
                type: 'warning'
            })
                .then(() => {
                    PayAccount(row.id).then((response) => {
                        if (response.status == 200) {
                            this.$message.success('支付成功');
                            this.tableData.splice(index, 1);
                        } else {
                            this.$message.error('支付失败');
                        }
                    }, (err) => {
                        console.log(err)
                    })
                    
                })
                .catch(() => {});
        },
        // 编辑操作
        handleTranfer(index, row) {
            this.idx = index;
            this.form = row;
            this.editVisible = true;
        },
        // 保存编辑
        saveEdit() {
            this.editVisible = false;
            let query = {
                id: this.form.id,
                money: this.form.transfermoney,
                ACompany: this.form.ACompanyID,
                BCompany: this.form.ToBCompanyID
            }

            TransferAccount(query).then((response) => {
                if (response.status == 200) {
                    this.$message.success('转让成功');
                    if (this.form.money == this.form.transfermoney) {
                        this.tableData.splice(this.idx, 1);
                    } else {
                        this.form.money = this.form.money - this.form.transfermoney
                        this.$set(this.tableData, this.idx, this.form);
                    }
                   
                } else {
                    this.$message.error('转让失败');
                }
            }, (err) => {
                console.log(err)
            })
            
        },
        handleFinacing(index, row) {
            let query = {
                accountID: row.id,
                companyID: localStorage.getItem('id')
            }
            FinacingAccount(query).then((response) => {
                if (response.status == 200) {
                    this.$message.success('融资成功');
                    this.tableData.splice(index, 1);
                   
                } else {
                    this.$message.error('融资失败');
                }
            }, (err) => {
                console.log(err)
            })
        }
    }
};
</script>

<style scoped>
.handle-box {
    margin-bottom: 20px;
}

.handle-select {
    width: 120px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}
.table {
    width: 100%;
    font-size: 14px;
}
.red {
    color: #ff0000;
}
.mr10 {
    margin-right: 10px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}
</style>
