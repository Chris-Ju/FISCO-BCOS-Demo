module.exports = {
    baseUrl: './',
    assetsDir: 'static',
    productionSourceMap: false,
    devServer: {
        proxy: {
            '/':{
                target:'http://localhost:8000/',
                changeOrigin:true,
                pathRewrite:{
                    '/':''
                }
            }
        }
    }
}