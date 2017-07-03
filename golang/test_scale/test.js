function genHMAC() {
    console.log("### Generating HMAC ###"); 
    var    e = customTextbox.getValue(document.getElementById("ENVVAR")), 
            t = customTextbox.getValue(document.getElementById("application")), 
            o = customTextbox.getValue(document.getElementById("application-version")), 
            n = [e, t, o].join("/"), 
            a = customTextbox.getValue(document.getElementById("GPVAR")), 
            s = a.split("/"), 
            r = s[0], u = s[1], 
            c = customTextbox.getValue(document.getElementById("CUSTID")), 
            m = Date.now(), 
            l = customTextbox.getValue(document.getElementById("TXID")), 
            d = customTextbox.getValue(document.getElementById("APIKEY")), 
            E = "<br>", 
            i = !1, 
            g = !0; 
            
    switch (u) { 
        case "get-report": g = validateInputs(u) 
    }
    
    var T = r + ":" + u + ":" + c + ":" + m + ":" + l, 
        p = CryptoJS.HmacSHA256(T, d), 
        I = CryptoJS.enc.Base64.stringify(p); 
    
    "post-market-transactions-internal" === u && (i = !0), 
    console.log("apikey: ", d), 
    console.log("data2hash: "), 
    console.log("service: ", r), 
    console.log("sub service: ", u), 
    console.log("customer Id: ", c), 
    console.log("timestamp: ", m), 
    console.log("transaction Id: ", l), 
    console.log(T), 
    console.log("hashBase64: ", I); 

    var x = "Header Data" + E + E; 
    x = x + "SCALE_TS: " + m + E, 
    x = x + "SCALE_TX_ID: " + l + E, 
    x = x + "SCALE_HMAC: " + I + E + E, 
    x = x + "URL" + E + E; 
    
    var y = document.getElementById("PROXY_SELECTED").checked, S = null; 
    y && (S = customTextbox.getValue(document.getElementById("PROXY"))); 
    
    var v = customTextbox.getValue(document.getElementById("CERT")), 
        C = [n, a, genUrlParams(u)].join("/"); 
    x = x + C + E + E; 
    
    var L = ""; 
    switch (SELECTED_ENV) { case "DEMO": case "PROD": L = "--tlsv1"; break; default: L = "" }
    
    var f = { timestamp: m, txId: l, hmac: I, proxy: S, cert: v, tlsv1: L, url: C }, 
        V = genCurl(u, f); 
    
    x = x + "cURL" + E + E, x = x + V + E; 
    
    var b = document.querySelector(".output-container"), 
        D = document.getElementById("OUTPUT"); 
    
    return b.classList.add("show"), 
        customTextbox.setValue(D, x), 
        D.scrollIntoView(), 
        CURL_COMMAND = V, { ts: m, txid: l, hmac: I, finalURL: C } 
    } 
        
    function createCORSRequest(e, t) { var o = new XMLHttpRequest; return "withCredentials" in o ? o.open(e, t, !0) : "undefined" != typeof XDomainRequest ? (o = new XDomainRequest, o.open(e, t)) : o = null, o } function appendGETParams(e, t) { var o = []; for (var n in t) o.push(encodeURIComponent(n) + "=" + encodeURIComponent(t[n])); return e + (o.length ? "?" + o.join("&") : "") } function emptyDom(e) { for (; e.firstChild;)e.removeChild(e.firstChild) } function userInputOnFocusOut(e) { var t = e.target || e.toElement; if ("div" === t.nodeName.toLowerCase() && t.classList.contains("custom-textbox") && "OUTPUT" !== t.id) { var o = t.innerHTML.replace(/<(?:.|\n)*?>/gm, "").trim(); if ("CUSTID" === t.id) { var n = APIKEY_LIST[SELECTED_ENV].apiKey || null; APIKEY_LIST[SELECTED_ENV] && APIKEY_LIST[SELECTED_ENV].customer && APIKEY_LIST[SELECTED_ENV].customer[o] && (n = APIKEY_LIST[SELECTED_ENV].customer[o] || null), n && customTextbox.setValue(document.getElementById("APIKEY"), n.trim()) } t.innerHTML = o } e.stopPropagation() } var ENV = { "https://www.demo.fx-scale.com": "DEMO" }, APIKEY_LIST = { DEMO: { customer: { EZL: "a1fc360d-fe95-4182-a9f3-a7180887771a" } } }, CERT_PWD_LIST = { "default": "", DEMO: "GEU00429-1364209.pem:1234" }, PROXY_LIST = { "default": "", DEMO: "10.65.128.43:8080" }, PROPS = { DEMO: {} }, CURL_COMMAND = "", SELECTED_ENV = "", SELECTED_APPLICATION = ""; String.prototype.trim || (String.prototype.trim = function () { return this.replace(/^\s+|\s+$/g, "") }); var addLeadingZeros = function (e, t) { return t > e.toString().length ? (Array(t).join("0") + e).slice(-t) : e }, formatDate = function (e, t) { var o = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]; try { var n = new Date(e), a = addLeadingZeros(n.getDate(), 2), s = n.getMonth() + 1, r = o[s], u = addLeadingZeros(s, 2), c = n.getFullYear(), m = addLeadingZeros(n.getHours(), 2), l = addLeadingZeros(n.getMinutes(), 2), d = addLeadingZeros(n.getSeconds(), 2), E = ""; if (t) switch (t) { case "YYYY-MM-DD": E = c + "-" + u + "-" + a; break; case "DD-MM-YYYY": E = a + "-" + u + "-" + c; break; case "THH-TMM": E = m + ":" + l } else E = a + " " + r + " " + c + " @ " + m + ":" + l + ":" + d; return E } catch (i) { return "" } }, customTextbox = { getValue: function (e) { return e.innerHTML.replace(/<(?:.|\n)*?>/gm, "").trim() }, setValue: function (e, t) { e.innerHTML = t } }; domReady(function (e) { document.getElementById("user-inputs").addEventListener("focusout", userInputOnFocusOut, !1), document.getElementById("ENV").dispatchEvent(new Event("change")) })