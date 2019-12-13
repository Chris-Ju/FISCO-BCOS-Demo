import axios from "axios"


export const QueryCompany = query => {
    return axios({
        url: '/api/user/?username='+query.username+'&password='+query.password,
        method: 'get'
    })
}

export const GetCompany = query => {
    return axios({
        url: '/admin/api/company/',
        method: 'get'
    })
}

export const RegisterCompany = query => {
    return axios({
        url: '/api/user/',
        method: 'post',
        data: query,
        transformRequest: [function (data) {
            let ret = ''
            for (let it in data) {
              ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
            }
            return ret
        }],
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        }
    })
}

export const IssueAccount = query => {
    return axios({
        url: '/api/account/',
        method: 'post',
        data: query,
        transformRequest: [function (data) {
            let ret = ''
            for (let it in data) {
              ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
            }
            return ret
        }],
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        }
    })
}

export const GetAccount = query => {
    return axios({
        url: '/api/account/?id='+query,
        method: 'get'
    })
}

export const PayAccount = query => {
    return axios({
        url: '/api/account/'+query,
        method: 'delete'
    })
}

export const TransferAccount = query => {
    return axios({
        url: '/api/account/'+query.id,
        method: 'put',
        data: query,
        transformRequest: [function (data) {
            let ret = ''
            for (let it in data) {
              ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
            }
            return ret
        }],
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        }
    })
}

export const FinacingAccount = query => {
    return axios({
        url: '/api/account/bank/'+query.accountID+'?companyID='+query.companyID,
        method: 'post'
    })
}