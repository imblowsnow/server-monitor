import{_ as d,r as h,o as v,c as f,b as e,w as l,h as u,t as b,a as S,L as k,M as C,N as x,O as y}from"./index-L-XxcENT.js";/* empty css                  */import{p as F,e as w}from"./server-CCX-AtOc.js";const z={name:"monitor-faults",props:{serverId:{type:Number,default:null},limit:{type:Number,default:10}},data(){return{faults:[{id:1,name:"fault1",status:"up"}],page:{current:1,limit:10,total:0}}},watch:{serverId(){this.pageServerFaults()}},mounted(){this.page.limit=this.limit,this.pageServerFaults()},methods:{pageServerFaults(){F(this.serverId,this.page.current,this.page.limit).then(s=>{this.faults=s.list,this.page.total=s.total})},editRemark(s,a){this.$prompt("请输入备注","提示",{confirmButtonText:"确定",cancelButtonText:"取消",inputPattern:/\S/,inputErrorMessage:"备注不能为空",inputValue:a}).then(({value:i})=>{w(s,i).then(()=>{this.$message({type:"success",message:"修改成功!"}),this.pageServerFaults()}).catch(p=>{this.$message({type:"error",message:"修改失败!"+p})})})}}},B={class:"monitor-faults"},E={"aria-label":"Page navigation example"};function N(s,a,i,p,n,o){const m=h("router-link"),r=k,_=C,c=x,g=y;return v(),f("div",B,[e(c,{data:n.faults,class:"table table-borderless table-hover"},{default:l(()=>[e(r,{prop:"server_name",label:"名称"},{default:l(t=>[e(m,{to:"/server/"+t.row.server_id},{default:l(()=>[u(b(t.row.server_name),1)]),_:2},1032,["to"])]),_:1}),e(r,{prop:"start_time",label:"开始时间"}),e(r,{prop:"end_time",label:"结束时间"}),e(r,{prop:"duration",label:"耗时(秒)"}),e(r,{prop:"remark",label:"备注"}),e(r,{label:"操作"},{default:l(t=>[e(_,{type:"primary",size:"small",onClick:T=>o.editRemark(t.row.id,t.row.remark)},{default:l(()=>[u("修改备注")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"]),S("nav",E,[e(g,{"page-size":n.page.limit,"onUpdate:pageSize":a[0]||(a[0]=t=>n.page.limit=t),"current-page":n.page.current,"onUpdate:currentPage":a[1]||(a[1]=t=>n.page.current=t),onSizeChange:o.pageServerFaults,onCurrentChange:o.pageServerFaults,layout:"prev, pager, next",total:this.page.total},null,8,["page-size","current-page","onSizeChange","onCurrentChange","total"])])])}const R=d(z,[["render",N]]);export{R as _};
