import{_ as b,o,c as l,a as e,F as y,m as C,H as p,b as i,O as v,t as g,w as c,h,M as x}from"./index-Drs-in-b.js";/* empty css                  */import{p as f,d as k}from"./server_group-zlpj1QQW.js";const z={name:"index",data(){return{list:[],page:{current:1,size:10,total:0}}},mounted(){this.pageChange()},methods:{pageChange(){f({page:this.page.current,size:this.page.size}).then(s=>{this.list=s.list,this.page.total=s.total,console.log("pageServerGroupList",s)}).catch(()=>{this.$message({type:"error",message:"获取数据失败!"})})},handleEdit(s){this.$router.push(`/server_group/edit/${s}`)},handleAdd(){this.$router.push("/server_group/add")},handleDelete(s){this.$confirm("此操作将永久删除该分组, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(()=>{k(s).then(()=>{this.$message({type:"success",message:"删除成功!"}),this.pageChange()}).catch(n=>{this.$message({type:"error",message:"删除失败!"+n})})}).catch(()=>{this.$message({type:"info",message:"已取消删除"})})}}},S=e("h1",{class:"mb-3"},"分组管理",-1),B={class:"functions mb-3"},w={class:"btn-group",role:"group"},E={class:"shadow-box mb-3"},N={class:"table table-borderless table-hover"},V=e("thead",null,[e("tr",null,[e("th",null,"名称"),e("th",null,"排序"),e("th",null,"操作")])],-1),D={style:{width:"150px"}},G={key:0},L=e("td",{colspan:"3",style:{"text-align":"center"}},"无数据",-1),P=[L],T={key:0,"aria-label":"Page navigation example"};function A(s,n,F,U,a,r){const d=x,u=v;return o(),l("div",null,[S,e("div",B,[e("div",w,[e("button",{class:"btn btn-primary",onClick:n[0]||(n[0]=t=>r.handleAdd())}," 新增 ")])]),e("div",E,[e("table",N,[V,e("tbody",null,[(o(!0),l(y,null,C(a.list,(t,_)=>(o(),l("tr",{key:_},[e("td",null,g(t.group_name),1),e("td",null,g(t.sort),1),e("td",D,[i(d,{class:"btn btn-primary",type:"primary",size:"small",disabled:t.system===1,onClick:m=>r.handleEdit(t.id)},{default:c(()=>[h("编辑")]),_:2},1032,["disabled","onClick"]),i(d,{class:"btn btn-danger",size:"small",type:"danger",disabled:t.system===1,onClick:m=>r.handleDelete(t.id)},{default:c(()=>[h("删除")]),_:2},1032,["disabled","onClick"])])]))),128)),a.list.length===0?(o(),l("tr",G,P)):p("",!0)])]),a.page.total>0?(o(),l("nav",T,[i(u,{"page-size":a.page.limit,"onUpdate:pageSize":n[1]||(n[1]=t=>a.page.limit=t),"current-page":a.page.current,"onUpdate:currentPage":n[2]||(n[2]=t=>a.page.current=t),onSizeChange:r.pageChange,onCurrentChange:r.pageChange,layout:"prev, pager, next",total:a.page.total},null,8,["page-size","current-page","onSizeChange","onCurrentChange","total"])])):p("",!0)])])}const O=b(z,[["render",A]]);export{O as default};
