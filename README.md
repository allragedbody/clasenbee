"# clasenbee" 
curl -X POST http://127.0.0.1/api/createmodule -d '{"Name": "CDN_LEVE2CACHE_NGINX_MAIN", "ChName": "CDN二级缓存nginx主配置", "Path": "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", "Desc": "二级缓存nginx主配置文件"}'
 curl -X POST http://127.0.0.1/api/createmodule -d '{"name": "CDN_LEVE2CACHE_NGINX_MAIN", "chname": "CDN二级缓存nginx主配置", "path": "/export/server/UCCMS/mod/nginx/cdn/nginx.mod", "desc": "二级缓存nginx主配置文件"}'

 func (this *TestController) Show() {
    o := orm.NewOrm()
    var maps []orm.Params
    o.Raw("select arc.id,arc.title,arc.typeid,art.typename from go_archives as arc left join go_arctype as art on art.id=arc.typeid where arc.typeid=?", 2).Values(&maps)
    fmt.Println(maps)
}


git使用方法

本地操作。
初始化一个Git仓库，使用git init命令。
第一步，使用命令git add <file>，注意，可反复多次使用，添加多个文件；
第二步，使用命令git commit，完成。


查找
 git log --grep readme_3

查看版本差异
git diff HEAD^^ -- readme.txt


查看修改情况
git status
提交修改
git add --all && git commit -m "nginx_version_2"
撤销修改
git checkout -- LICENSE 


回滚版本
git reset --hard 803461d
git reset --hard HEAD@{6}