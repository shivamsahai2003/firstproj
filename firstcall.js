(function(){
  var src = (document.currentScript && document.currentScript.src) || '';   // todo remove
  var ORIGIN;

  // todo why capital case and not small or camel case
  var PUBLISHER_ID = ''; // Publisher identifier from script src
  try {
    var srcUrl = new URL(src, location.href);
    ORIGIN = srcUrl.origin;
    PUBLISHER_ID = srcUrl.searchParams.get('pub_id') || ''; // todo remember to change publisher page contract
  } catch(e){ ORIGIN = location.origin; }

  // Publisher configuration by pub param or host
  var PUB_CONFIG = {
    'blue': { pid: 100, lid: 224, actno: 5, maxno: 5, cc: 'US', tsize: '300x250', pubKey: 'blue' },
    'red':  { pid: 200, lid: 224, actno: 5, maxno: 5, cc: 'US', tsize: '300x250', pubKey: 'red' }
  };

  var HOST_CONFIG = {
    'blue.localhost': PUB_CONFIG['blue'],
    'red.localhost':  PUB_CONFIG['red'],
    'localhost':      { pid: 0, lid: 224, actno: 5, maxno: 5, cc: 'US', tsize: '300x250', pubKey: 'default' }
  };

  var DEFAULT_CONFIG = { pid: 0, lid: 224, actno: 5, maxno: 5, cc: 'US', tsize: '300x250', pubKey: 'default' };

  function getConfig() {
    // First check pub param from script src
    if (PUBLISHER_ID && PUB_CONFIG[PUBLISHER_ID]) {
      return PUB_CONFIG[PUBLISHER_ID];
    }
    // Fallback to host-based config
    var host = (location.hostname || '').replace(/^www\./, '').split(':')[0];
    if (HOST_CONFIG[host]) return HOST_CONFIG[host];
    for (var h in HOST_CONFIG) {
      if (host === h || host.indexOf(h) === 0) return HOST_CONFIG[h];
    }
    return DEFAULT_CONFIG;
  }

  function getPageDefaults(){
    var config = getConfig();
    var qs = new URLSearchParams(window.location.search);
    return {
      actno: config.actno,
      maxno: config.maxno,
      cc: config.cc,
      lid: config.lid,
      pid: config.pid,
      pub: config.pubKey, // Pass publisher key to backend
      d: location.hostname || '',
      rurl: location.href,
      ptitle: qs.get('ptitle') || (document.title || ''),
      tsize: config.tsize,
      kwrf: document.referrer || ''
    };
  }

  function slotIdFromEl(el){
    if (!el) return '';
    var dataSlot = el.getAttribute('data-kw-slot');
    if (dataSlot) return String(dataSlot);
    var id = el.id || '';
    if (id) return id;
    return '';
  }

  function injectForSlot(el){
    if (!el || el.__kwInjected) return;
    var defaults = getPageDefaults();
    var slotId = slotIdFromEl(el);
    if (!slotId) return;
    var p = [];
    p.push('slot=' + encodeURIComponent(slotId));
    p.push('actno=' + encodeURIComponent(String(defaults.actno)));
    p.push('maxno=' + encodeURIComponent(String(defaults.maxno)));
    p.push('cc=' + encodeURIComponent(String(defaults.cc)));
    p.push('lid=' + encodeURIComponent(String(defaults.lid)));
    if (defaults.d) p.push('d=' + encodeURIComponent(String(defaults.d)));
    if (defaults.rurl) p.push('rurl=' + encodeURIComponent(String(defaults.rurl)));
    if (defaults.ptitle) p.push('ptitle=' + encodeURIComponent(String(defaults.ptitle)));
    if (defaults.tsize) p.push('tsize=' + encodeURIComponent(String(defaults.tsize)));
    if (defaults.kwrf) p.push('kwrf=' + encodeURIComponent(String(defaults.kwrf)));
    p.push('pid=' + encodeURIComponent(String(defaults.pid))); // Always send PID (even if 0)
    if (defaults.pub) p.push('pub=' + encodeURIComponent(String(defaults.pub))); // Publisher key
    var s = document.createElement('script');
    s.async = true;
    s.src = ORIGIN + '/keyword_render?' + p.join('&');   // todo change path here
    (document.head || document.documentElement || document.body).appendChild(s);
    el.__kwInjected = true;
  }

  function findSlots(){
    var list = [];
    try {
      var a = document.querySelectorAll('[data-kw-slot]');
      for (var i=0; i<a.length; i++) list.push(a[i]);
      var b = document.querySelectorAll('[id^="kw-slot-"]');
      for (var j=0; j<b.length; j++) { if (list.indexOf(b[j]) === -1) list.push(b[j]); }
    } catch(e){}
    return list;
  }

  function run(){ var slots = findSlots(); for (var i=0; i<slots.length; i++){ injectForSlot(slots[i]); } }
  if (document.readyState === 'complete' || document.readyState === 'interactive') { run(); }
  else { document.addEventListener('DOMContentLoaded', run); }
})();// todo-> what problem this fucntion is solving
