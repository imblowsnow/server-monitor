import{b as l,u}from"./server_group-mdiWvcsB.js";import{_ as d,o as m,c as p,a as e,k as n,v as a}from"./index-L-XxcENT.js";const c={name:"index",data(){return{form:{group_name:"",sort:1},id:null}},mounted(){this.id=this.$route.params.id,this.getServerGroup(this.id)},methods:{getServerGroup(t){l(t).then(s=>{this.form=s})},updateServerGroup(){u(this.id,this.form).then(()=>{this.$message({type:"success",message:"更新成功!"}),this.$router.push("/server_group")}).catch(t=>{this.$message({type:"error",message:"更新失败!"+t})})}}},_=e("h1",{class:"mb-3"},"编辑分组",-1),h={class:"shadow-box big-padding"},f={class:"row"},v={class:"mb-3"},b=e("label",{class:"form-label"},"分组名称",-1),g={class:"mb-3"},x=e("label",{class:"form-label"},"分组排序",-1),G={class:"row"};function S(t,s,y,k,o,i){return m(),p("div",null,[_,e("div",h,[e("form",null,[e("div",f,[e("div",v,[b,n(e("input",{type:"text",class:"form-control","onUpdate:modelValue":s[0]||(s[0]=r=>o.form.group_name=r),required:""},null,512),[[a,o.form.group_name]])]),e("div",g,[x,n(e("input",{type:"number",min:1,class:"form-control","onUpdate:modelValue":s[1]||(s[1]=r=>o.form.sort=r),required:""},null,512),[[a,o.form.sort]])])]),e("div",G,[e("button",{type:"button",class:"btn btn-primary",onClick:s[2]||(s[2]=(...r)=>i.updateServerGroup&&i.updateServerGroup(...r))},"更新")])])])])}const V=d(c,[["render",S]]);export{V as default};
