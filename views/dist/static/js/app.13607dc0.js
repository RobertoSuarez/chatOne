(function(t){function e(e){for(var s,r,i=e[0],c=e[1],u=e[2],d=0,m=[];d<i.length;d++)r=i[d],Object.prototype.hasOwnProperty.call(n,r)&&n[r]&&m.push(n[r][0]),n[r]=0;for(s in c)Object.prototype.hasOwnProperty.call(c,s)&&(t[s]=c[s]);l&&l(e);while(m.length)m.shift()();return o.push.apply(o,u||[]),a()}function a(){for(var t,e=0;e<o.length;e++){for(var a=o[e],s=!0,i=1;i<a.length;i++){var c=a[i];0!==n[c]&&(s=!1)}s&&(o.splice(e--,1),t=r(r.s=a[0]))}return t}var s={},n={app:0},o=[];function r(e){if(s[e])return s[e].exports;var a=s[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,r),a.l=!0,a.exports}r.m=t,r.c=s,r.d=function(t,e,a){r.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},r.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},r.t=function(t,e){if(1&e&&(t=r(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(r.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var s in t)r.d(a,s,function(e){return t[e]}.bind(null,s));return a},r.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return r.d(e,"a",e),e},r.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},r.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],c=i.push.bind(i);i.push=e,i=i.slice();for(var u=0;u<i.length;u++)e(i[u]);var l=c;o.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},"034f":function(t,e,a){"use strict";a("85ec")},"0c7c":function(t,e,a){"use strict";var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticClass:"row",attrs:{id:"header"}},[a("h3",[t._v("Chat")]),a("button",{staticClass:"btn waves-effect waves-light blue darken-1",attrs:{type:"submit",name:"action"},on:{click:t.closeSession}},[t._v(" Cerrar Sessión ")])]),a("div",{staticClass:"row"},[a("div",{staticClass:"col s3",attrs:{id:"Contactos"}},[a("h2",[t._v("Contactos")]),a("Contactos",{attrs:{contactos:t.contactos},on:{"click-set-chat":t.setChat}}),a("button",{staticClass:"btn waves-effect waves-light blue darken-1",on:{click:t.NewContact}},[t._v("Nuevo contato")])],1),a("div",{staticClass:"col s9",attrs:{id:"chat"}},[a("h2",[t._v("chat")]),t.isSelected?a("div",[a("h2",[t._v(t._s(t.chat.name))]),t._l(t.chat.mensajes,(function(e){return a("div",{key:e.id},[t._v(" "+t._s(e))])})),a("input",{directives:[{name:"model",rawName:"v-model",value:t.mensaje,expression:"mensaje"}],attrs:{type:"text"},domProps:{value:t.mensaje},on:{keydown:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.insertarmsg(e)},input:function(e){e.target.composing||(t.mensaje=e.target.value)}}})],2):a("div",[a("h2",[t._v("Chatea con tus contactos")])])])])])},n=[],o=(a("c740"),a("9612")),r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",t._l(t.contactos,(function(e){return a("div",{key:e.id},[a("Contacto",{attrs:{contacto:e},on:{"click-set-chat":function(a){return t.$emit("click-set-chat",e.id)}}})],1)})),0)},i=[],c=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticClass:"contacto",on:{click:function(e){return t.$emit("click-set-chat",t.contacto.id)}}},[t._v(" "+t._s(t.contacto.name)+" ")])])},u=[],l=(a("b0c0"),{methods:{chatear:function(){console.log(this.contacto.name)}},props:["contacto"]}),d=l,m=(a("9021"),a("2877")),p=Object(m["a"])(d,c,u,!1,null,null,null),v=p.exports,f={name:"Contactos",components:{Contacto:v},data:function(){return{}},methods:{},props:["contactos"]},h=f,b=Object(m["a"])(h,r,i,!1,null,null,null),g=b.exports,w={name:"Dashboard",components:{Contactos:g},data:function(){return{user:{},contactos:[],chat:{},isSelected:!1,mensaje:""}},methods:{closeSession:function(){this.$cookies.remove("Token"),this.$router.push("/")},setChat:function(t){var e=this;this.contactos.findIndex((function(a){if(a.id==t)return e.chat=a,!0})),this.isSelected=!0},insertarmsg:function(){var t={id:o["a"].v4(),mensaje:this.mensaje};this.chat.mensajes.push(t),this.mensaje=""},NewContact:function(){var t=this;console.log("Creando un nuevo contacto");var e=prompt("Nombre del contacto?");this.$chat.post("/users/".concat(this.user.ID,"/contacts"),{number:"0997454001",name:e,nickname:"pepe rata"}).then((function(e){console.log(e),t.GetContactos()}))},GetContactos:function(){var t=this;this.$chat.get("/users/".concat(this.user.ID,"/contacts")).then((function(e){return t.contactos=e.data}))}},created:function(){var t=this,e=this.$cookies.get("Token");null!==e?(this.$chat.defaults.headers.common["Token"]=e,this.$chat.get("/users/me").then((function(e){t.user=e.data,t.GetContactos()}))):this.$router.push("/")}},_=w,C=Object(m["a"])(_,s,n,!1,null,null,null);e["a"]=C.exports},2531:function(t,e,a){},"3dfd":function(t,e,a){"use strict";var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("router-view")],1)},n=[],o={name:"App",components:{},data:function(){return{}},methods:{}},r=o,i=(a("034f"),a("2877")),c=Object(i["a"])(r,s,n,!1,null,null,null);e["a"]=c.exports},"56d7":function(t,e,a){"use strict";a.r(e),function(t){a("e260"),a("e6cf"),a("cca6"),a("a79d");var e=a("2b0e"),s=a("3dfd"),n=(a("8266"),a("619b"),a("4d5c"),a("8c4f")),o=a("bc3a"),r=a.n(o),i=a("2b27"),c=a.n(i),u=a("578a"),l=a("57da"),d=a("fac8"),m=a("0c7c");e["a"].prototype.$http=r.a,e["a"].prototype.$chat=r.a.create({baseURL:"https://chatone1.herokuapp.com/api/v1",timeout:5e3}),e["a"].use(n["a"]),e["a"].use(c.a),e["a"].config.productionTip=!1;var p=new n["a"]({mode:"history",base:t,routes:[{path:"/",component:l["a"]},{path:"/login",component:u["a"]},{path:"/Registrar",component:d["a"]},{path:"/dashboard",component:m["a"]}]});new e["a"]({router:p,render:function(t){return t(s["a"])}}).$mount("#app")}.call(this,"/")},"578a":function(t,e,a){"use strict";var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[a("h2",[t._v("Iniciar Sessión")]),a("form",{on:{submit:function(e){return e.preventDefault(),t.login(e)}}},[a("div",{staticClass:"row"},[a("div",{staticClass:"input-field col s12"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.auth.email,expression:"auth.email"}],attrs:{type:"email",id:"email"},domProps:{value:t.auth.email},on:{input:function(e){e.target.composing||t.$set(t.auth,"email",e.target.value)}}}),a("label",{attrs:{for:"email"}},[t._v("Email")])]),a("div",{staticClass:"input-field col s12"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.auth.password,expression:"auth.password"}],attrs:{type:"password",id:"password"},domProps:{value:t.auth.password},on:{input:function(e){e.target.composing||t.$set(t.auth,"password",e.target.value)}}}),a("label",{attrs:{for:"password"}},[t._v("Password")])]),a("button",{staticClass:"btn waves-effect waves-light blue darken-1",class:t.disable?"disabled":"",attrs:{type:"submit",name:"action"}},[t._v("Registrate ")])])])])},n=[],o={name:"Login",data:function(){return{auth:{email:"",password:""},disable:!1}},methods:{login:function(){var t=this;console.log(this.auth),this.disable=!0,this.$chat.post("/login",this.auth).then((function(e){e.data.iserror?alert(e.data.message):(alert(e.data.message),t.$cookies.set("Token",e.data.data.token),t.$router.push("/dashboard")),t.disable=!1}))}}},r=o,i=a("2877"),c=Object(i["a"])(r,s,n,!1,null,null,null);e["a"]=c.exports},"57da":function(t,e,a){"use strict";var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"row"},[t._m(0),a("div",{staticClass:"col s8 maincontainer"},[a("h1",{staticClass:"title"},[t._v("Comunicate con chat")]),a("router-link",{staticClass:"link",attrs:{to:"/login"}},[a("button",{staticClass:"btn waves-effect waves-light blue darken-1",attrs:{type:"submit",name:"action"}},[t._v("Iniciar sessión ")])]),a("router-link",{staticClass:"link",attrs:{to:"/registrar"}},[a("button",{staticClass:"btn waves-effect waves-light blue darken-1",attrs:{type:"submit",name:"action"}},[t._v("Registrate ")])])],1)])},n=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"col s4 headercontainer"},[a("h3",{staticClass:"title"},[t._v("Comunicate")])])}],o={name:"Home",created:function(){console.log("creado")}},r=o,i=(a("a224"),a("2877")),c=Object(i["a"])(r,s,n,!1,null,"6abb514b",null);e["a"]=c.exports},"85ec":function(t,e,a){},9021:function(t,e,a){"use strict";a("c636")},a224:function(t,e,a){"use strict";a("2531")},c636:function(t,e,a){},fac8:function(t,e,a){"use strict";var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[a("h2",[t._v("Crea tu cuenta")]),a("form",{on:{submit:function(e){return e.preventDefault(),t.Register(e)}}},[a("div",{staticClass:"row"},[a("div",{staticClass:"input-field col s6"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.user.name,expression:"user.name"}],attrs:{type:"text",id:"nombre"},domProps:{value:t.user.name},on:{input:function(e){e.target.composing||t.$set(t.user,"name",e.target.value)}}}),a("label",{attrs:{for:"nombre"}},[t._v("Nombre")])]),a("div",{staticClass:"input-field col s6"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.user.password,expression:"user.password"}],attrs:{type:"password",id:"password"},domProps:{value:t.user.password},on:{input:function(e){e.target.composing||t.$set(t.user,"password",e.target.value)}}}),a("label",{attrs:{for:"password"}},[t._v("Password")])]),a("div",{staticClass:"input-field col s12"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.user.email,expression:"user.email"}],attrs:{type:"email",id:"email"},domProps:{value:t.user.email},on:{input:function(e){e.target.composing||t.$set(t.user,"email",e.target.value)}}}),a("label",{attrs:{for:"email"}},[t._v("Email")])]),a("div",{staticClass:"input-field col s6"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.user.number,expression:"user.number"}],attrs:{type:"text",id:"number"},domProps:{value:t.user.number},on:{input:function(e){e.target.composing||t.$set(t.user,"number",e.target.value)}}}),a("label",{attrs:{for:"number"}},[t._v("Number")])]),a("div",{staticClass:"input-field col s6"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.user.role,expression:"user.role"}],attrs:{type:"text",id:"role"},domProps:{value:t.user.role},on:{input:function(e){e.target.composing||t.$set(t.user,"role",e.target.value)}}}),a("label",{attrs:{for:"role"}},[t._v("Role")])]),a("button",{staticClass:"btn waves-effect waves-light blue darken-1",class:t.disable?"disabled":"",attrs:{type:"submit",name:"action"}},[t._v("Registrate ")])])])])},n=[],o={name:"Registrar",data:function(){return{user:{name:"",password:"",number:"",email:"",role:""},disable:!1}},methods:{Register:function(){var t=this;this.disable||(this.disable=!0,this.$chat.post("/users",this.user).then((function(e){console.log(e.data),200==e.status&&(e.data.iserror?alert(e.data.message):(alert(e.data.message),t.$router.push("/"))),t.disable=!1})).catch((function(t){return console.log(t)})))}}},r=o,i=a("2877"),c=Object(i["a"])(r,s,n,!1,null,null,null);e["a"]=c.exports}});
//# sourceMappingURL=app.13607dc0.js.map