import{_ as n}from"./monitor-faults-DDVJoldT.js";import{_ as l,d as c,o as i,c as r,a as s,n as a,t as e,b as _,e as h,p as m,f as p}from"./index-D-XvNmkW.js";/* empty css                  */import"./server-C2qe13OZ.js";const u=c({name:"dashboard",components:{MonitorFaults:n},data(){return{total:{onlineNum:0,offlineNum:0}}},mounted(){this.dashboardTotal()},methods:{dashboardTotal(){h().then(o=>{this.total=o})}}}),t=o=>(m("data-v-4488ea02"),o=o(),p(),o),f={class:"dashboard"},b=t(()=>s("h1",{class:"mb-3"},"状态速览",-1)),v={class:"shadow-box big-padding text-center mb-4"},N={class:"row"},x={class:"col"},w=t(()=>s("h3",null,"在线",-1)),g={class:"col"},y=t(()=>s("h3",null,"离线",-1)),I={class:"shadow-box table-shadow-box",style:{"overflow-x":"hidden"}};function S(o,$,B,T,k,C){const d=n;return i(),r("div",f,[b,s("div",v,[s("div",N,[s("div",x,[w,s("span",{class:a(["num",o.total.onlineNum===0&&"text-secondary"])},e(o.total.onlineNum),3)]),s("div",g,[y,s("span",{class:a(["num",o.total.offlineNum>0?"text-danger":"text-secondary"])},e(o.total.offlineNum),3)])])]),s("div",I,[_(d,{limit:5})])])}const F=l(u,[["render",S],["__scopeId","data-v-4488ea02"]]);export{F as default};
