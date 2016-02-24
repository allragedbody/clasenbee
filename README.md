"# clasenbee" 
curl -X POST http://127.0.0.1/api/createmodule -d '{"Name": "CDN_LEVE2CACHE_NGINX_MAIN", "ChName": "CDN二级缓存nginx主配置", "Path": "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", "Desc": "二级缓存nginx主配置文件"}'
 curl -X POST http://127.0.0.1/api/createmodule -d '{"name": "CDN_LEVE2CACHE_NGINX_MAIN", "chname": "CDN二级缓存nginx主配置", "path": "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", "desc": "二级缓存nginx主配置文件"}'

 func (this *TestController) Show() {
    o := orm.NewOrm()
    var maps []orm.Params
    o.Raw("select arc.id,arc.title,arc.typeid,art.typename from go_archives as arc left join go_arctype as art on art.id=arc.typeid where arc.typeid=?", 2).Values(&maps)
    fmt.Println(maps)
}