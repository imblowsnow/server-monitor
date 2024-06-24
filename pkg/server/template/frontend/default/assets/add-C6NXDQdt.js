import{b as c}from"./server-BP7FDecQ.js";import{g as _}from"./server_group-Ddmbvzjs.js";import{_ as u,o as n,c as i,a as s,k as l,l as m,F as p,m as f,v as d,t as h}from"./index-GT17-lK8.js";const v={data(){return{groups:[{id:1,name:"group1"},{id:2,name:"group2"}],form:{group_id:1,name:"",remark:"",sort:1,show_index:1}}},mounted(){this.getGroups()},methods:{getGroups(){_().then(r=>{this.groups=r}).catch(()=>{this.$message.error("获取分组失败!")})},addServer(){c(this.form).then(r=>{this.$bus.emit("serverChanged"),this.$router.go(-1)}).catch(r=>{this.$message.error("添加失败!"+r)})}}},b={class:"server-add"},g=s("h1",{class:"mb-3"},"添加服务器",-1),x={class:"shadow-box big-padding"},k={class:"row"},w={class:"col-md-6"},S={class:"mb-3"},y=s("label",{class:"form-label"},"服务器分组",-1),V={class:"input-group"},U=["value"],q={class:"mb-3"},B=s("label",{class:"form-label"},"服务器名称",-1),G={class:"mb-3"},C=s("label",{class:"form-label"},"服务器备注信息",-1),D={class:"col-md-6"},F={class:"mb-3"},M=s("label",{class:"form-label"},"排序",-1),E={class:"mb-3"},L=s("label",{class:"form-label"},"首页展示",-1),N=s("option",{value:1},"展示",-1),T=s("option",{value:0},"不展示",-1),j=[N,T],z={class:"row"};function A(r,e,H,I,t,a){return n(),i("div",b,[g,s("div",x,[s("form",null,[s("div",k,[s("div",w,[s("div",S,[y,s("div",V,[l(s("select",{class:"form-control","onUpdate:modelValue":e[0]||(e[0]=o=>t.form.group_id=o),required:""},[(n(!0),i(p,null,f(t.groups,o=>(n(),i("option",{value:o.id},h(o.group_name),9,U))),256))],512),[[m,t.form.group_id]])])]),s("div",q,[B,l(s("input",{type:"text",class:"form-control","onUpdate:modelValue":e[1]||(e[1]=o=>t.form.name=o),required:""},null,512),[[d,t.form.name]])]),s("div",G,[C,l(s("input",{type:"text",class:"form-control","onUpdate:modelValue":e[2]||(e[2]=o=>t.form.remark=o),required:""},null,512),[[d,t.form.remark]])])]),s("div",D,[s("div",F,[M,l(s("input",{type:"number",class:"form-control","onUpdate:modelValue":e[3]||(e[3]=o=>t.form.sort=o),min:"0",required:""},null,512),[[d,t.form.sort]])]),s("div",E,[L,l(s("select",{class:"form-control","onUpdate:modelValue":e[4]||(e[4]=o=>t.form.show_index=o)},j,512),[[m,t.form.show_index]])])])]),s("div",z,[s("button",{type:"button",class:"btn btn-primary",onClick:e[5]||(e[5]=(...o)=>a.addServer&&a.addServer(...o))},"添加")])])])])}const P=u(v,[["render",A]]);export{P as default};