import{q as r}from"./index-L-XxcENT.js";async function s(e){return r({url:"/admin-api/v1/server/"+e,method:"get"})}async function i(e){return r({url:"/admin-api/v1/server/info/"+e,method:"get"})}async function u(e){return r({url:"/admin-api/v1/server",method:"post",data:e})}async function d(e,t){return r({url:"/admin-api/v1/server/"+e,method:"put",data:t})}async function v(e,t=1){return r({url:`/admin-api/v1/server/statistics/${e}?type=${t}`,method:"get"})}async function o(e,t=1,a=10){return r({url:`/admin-api/v1/server/fault/page?page=${t}&size=${a}&serverId=${e||""}`,method:"get"})}async function c(e,t){return r({url:`/admin-api/v1/server/fault/${e}`,method:"put",data:{remark:t}})}async function m(e){return r({url:"/admin-api/v1/server/"+e,method:"delete"})}export{v as a,u as b,s as c,m as d,c as e,i as g,o as p,d as u};
