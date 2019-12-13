<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="8">
                <el-card shadow="hover" class="mgb20" style="height:252px;">
                    <div class="user-info">
                        <img src="../../assets/img/img.jpg" class="user-avator" alt />
                        <div class="user-info-cont">
                            <div class="user-info-name">{{username}}</div>
                            <div>{{role}}</div>
                        </div>
                    </div>
                </el-card>
                <el-card shadow="hover" style="height:252px;">
                    <div slot="header" class="clearfix">
                        <span>公司列表</span>
                    </div>
                    <el-table :show-header="false" :data="companyList" style="width:100%;">
                        <el-table-column width="100">
                            <template slot-scope="scope">
                                <div class="todo-item">{{scope.row.id}}</div>
                            </template>
                        </el-table-column>
                        <el-table-column>
                            <template slot-scope="scope">
                                <div class="todo-item">{{scope.row.name}}</div>
                            </template>
                        </el-table-column>
                        <el-table-column>
                            <template slot-scope="scope">
                                <div class="todo-item">{{scope.row.money}}万</div>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-card>
            </el-col>
            <el-col :span="16">
                <el-row :gutter="20" class="mgb20">
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{padding: '0px'}">
                            <div class="grid-content grid-con-1">
                                <i class="el-icon-lx-redpacket grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{funds}}万</div>
                                    <div>公司资金</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{padding: '0px'}">
                            <div class="grid-content grid-con-2">
                                <i class="el-icon-lx-servicefill grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{AAccount}}</div>
                                    <div>甲方账单数量</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{padding: '0px'}">
                            <div class="grid-content grid-con-3">
                                <i class="el-icon-lx-service grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{BAccount}}</div>
                                    <div>乙方账单数量</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                </el-row>
                <el-card shadow="hover" style="height:403px;">
                    <div slot="header" class="clearfix">
                        <span>账单列表</span>
                        <el-button style="float: right; padding: 3px 0" type="text">添加</el-button>
                    </div>
                    <el-table :show-header="false" :data="accountList" style="width:100%;">
                        <el-table-column>
                            <template slot-scope="scope">
                                <div class="todo-item">{{scope.row.id}}</div>
                            </template>
                        </el-table-column>
                        <el-table-column>
                            <template slot-scope="scope">
                                <div class="todo-item">{{scope.row.money}}</div>
                            </template>
                        </el-table-column>
                         <el-table-column>
                            <template slot-scope="scope">
                                <div class="todo-item">{{scope.row.ACompany}}</div>
                            </template>
                        </el-table-column>
                        <el-table-column>
                            <template slot-scope="scope">
                                <div class="todo-item">{{scope.row.BCompany}}</div>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import bus from '../common/bus';
import { GetCompany } from '../../api/index';
import { GetAccount } from '../../api/index';
export default {
    name: 'dashboard',
    
    data() {
        return {
            username : localStorage.getItem('name'),
            companyList : [],
            accountList : [],
            funds: 0,
            AAccount: 0,
            BAccount: 0
        }
    },
    
    created() {
        GetAccount(localStorage.getItem('id')).then((response) => {
            if (response.status == 200) {
                let temp = [{id:"id", money:"money", ACompany:"甲方", BCompany:"乙方"}]
                for (let i = 0; i < response.data.length; i++) {

                    temp[i + 1] = {
                        id : response.data[i]['id'],
                        money: response.data[i]['money'],
                        ACompany: response.data[i]['ACompany']['name'],
                        BCompany: response.data[i]['BCompany']['name']
                    }
                    if (response.data[i]['ACompany']['id'] == 1000) {
                        temp[i + 1]['ACompany'] = 'bank'
                    }
                    if (response.data[i]['BCompany']['id'] == 1000) {
                        temp[i + 1]['BCompany'] = 'bank'
                    }
                    if (temp[i+1]['ACompany'] == localStorage.getItem('name')) {
                        this.AAccount++
                    } else {
                        this.BAccount++
                    }
                }
                this.accountList = temp
            }
        }, (err) => {
            console.log(err)
        })
        GetCompany().then((response) => {
        if (response.status == 200) {
            this.companyList = response.data
            for (let i = 0; i < response.data.length; i++) {
                if (response.data[i]['id'] == localStorage.getItem('id')) {
                    this.funds = response.data[i].money
                }
            }
        }
        }, (err) => {
            console.log(err)
        })
        
        
    },
    
    computed: {
        role() {
            return this.name === 'admin' ? '超级管理员' : '普通用户';
        },
    },
    methods: {
    }
};
</script>


<style scoped>
.el-row {
    margin-bottom: 20px;
}

.grid-content {
    display: flex;
    align-items: center;
    height: 100px;
}

.grid-cont-right {
    flex: 1;
    text-align: center;
    font-size: 14px;
    color: #999;
}

.grid-num {
    font-size: 30px;
    font-weight: bold;
}

.grid-con-icon {
    font-size: 50px;
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    color: #fff;
}

.grid-con-1 .grid-con-icon {
    background: rgb(45, 140, 240);
}

.grid-con-1 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-2 .grid-con-icon {
    background: rgb(100, 213, 114);
}

.grid-con-2 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-3 .grid-con-icon {
    background: rgb(242, 94, 67);
}

.grid-con-3 .grid-num {
    color: rgb(242, 94, 67);
}

.user-info {
    display: flex;
    align-items: center;
    padding-bottom: 20px;
    border-bottom: 2px solid #ccc;
    margin-bottom: 20px;
}

.user-avator {
    width: 120px;
    height: 120px;
    border-radius: 50%;
}

.user-info-cont {
    padding-left: 50px;
    flex: 1;
    font-size: 14px;
    color: #999;
}

.user-info-cont div:first-child {
    font-size: 30px;
    color: #222;
}

.user-info-list {
    font-size: 14px;
    color: #999;
    line-height: 25px;
}

.user-info-list span {
    margin-left: 70px;
}

.mgb20 {
    margin-bottom: 20px;
}

.todo-item {
    font-size: 14px;
}

.todo-item-del {
    text-decoration: line-through;
    color: #999;
}

.schart {
    width: 100%;
    height: 300px;
}
</style>
