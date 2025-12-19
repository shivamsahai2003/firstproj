package templates

const SerpTemplate = `
<!doctype html>
<html lang="en">
<body>
	<div class="ad-item"{{.ad_title_1}}</div>
	<div class="ad-item"{{.ad_desc_1}}</div>

	<div class="ad-item"{{.ad_title_2}}</div>
	<div class="ad-item"{{.ad_desc_2}}</div>

	<div class="ad-item"{{.ad_title_3}}</div>
	<div class="ad-item"{{.ad_desc_3}}</div>
</body>
</html>`

const PublisherTemplatePub1 = `
<!doctype html>
<html lang="en">
<body>
<div class="keyword-pub1" style="color:blue"{{.keyword_1}}</div>

<div class="keyword-pub1"> style="color:blue"{{.keyword_2}}</div>

<div class="keyword-pub1"> style="color:blue"{{.keyword_3}}</div>
</body>
</html>`

const PublisherTemplatePub2 = `
<!doctype html>
<html lang="en">
<body>
<div class="keyword-pub2" style="color:red"{{.keyword_1}}</div>

<div class="keyword-pub2"> style="color:red"{{.keyword_2}}</div>

<div class="keyword-pub2"> style="color:red"{{.keyword_3}}</div>

<div class="keyword-pub2"> style="color:red"{{.keyword_4}}</div>

<div class="keyword-pub2"> style="color:red"{{.keyword_5}}</div>
</body>
</html>`

//const KeywordTemplateBlue=`

//<!doctype html>
//<html lang="en">
//<body>
//	div<

//
//const SerpTemplateBlue = `{{/*MAX_ADS*/}}
//
//<!doctype html>
//<html lang="en">
//<body>
//{{range $i, $ad := .Ads}}
//	{{if lt $i 3}}
//		<div class="ad-item">{{.$ad.TitleHTML}}</div>
//	{{end}}
//{{end}}
//
//</body>
//</html>`
//
//const SerpTemplateRed = `{{/*MAX_ADS*/}}
//<!doctype html>
//<html lang="en">
//<body>
//{{range $i, $ad := .Ads}}
//	{{if lt $i 5}}
//		<div class="ad-item">{{.$ad.TitleHTML}}</div>
//	{{end}}
//{{end}}
//</body>
//</html>`

// KeywordTemplate for the Publisher Page
//const KeywordTemplate = `<!doctype html>
//<html lang="en">
//<head><meta charset="utf-8"><title>{{.Title}}</title><meta name="viewport" content="width=device-width, initial-scale=1"></head>
//<body style="font:15px/1.5 Arial, sans-serif; margin:20px;">
//  <h1>{{.Title}}</h1>
//  <p style="color:#555;">Publisher: {{.PubKey}} | Total fetched: {{.TotalFetched}} | Total shown: {{.TotalShown}}</p>
//  <p style="color:#555;">Slot: {{.Slot}} | CC: {{.CC}} | Domain: {{.D}} | LID: {{.LID}} | PID: {{.PID}} | Size: {{.TSize}}</p>
//  <p style="color:#777;">Referrer: {{.KwRf}}</p>
//  <p style="color:#777;">Page Title: {{.PTitle}}</p>
//  <p style="color:#777;">Page URL: {{.RURL}}</p>
//  <p style="color:#777;">Keyword ID: {{.KID}}</p>
//  {{if .IsBot}}
//    <div style="margin:10px 0; padding:10px; background:#fff7ed; border:1px solid #fed7aa; color:#9a3412; border-radius:6px;">Bot detected.</div>
//  {{end}}
//  <hr>
//  <div style="display:grid; grid-template-columns: 1fr; gap:12px; margin-top:16px;">
//    {{range .Groups}}
//      <section style="border:1px solid #e5e7eb; border-radius:8px; padding:12px;">
//        <div style="font-weight:600; margin-bottom:8px;">{{.Label}} ({{len .Keywords}})</div>
//        {{if .Keywords}}
//          <ul style="margin:0; padding-left:18px;">
//            {{range .Keywords}}
//              <li><span style="color:#0b57d0; font-weight:600;">{{.Name}}</span></li>
//            {{end}}
//          </ul>
//        {{else}}
//          <div style="color:#999;">No keywords in this group</div>
//        {{end}}
//      </section>
//    {{end}}
//  </div>
//  <p style="margin-top:12px;"><a href="javascript:history.back()">Back</a></p>
//</body></html>`

// SerpTemplate is the HTML template for the SERP page
//const SerpTemplate = `<!doctype html>
//<html lang="en">
//<head><meta charset="utf-8"><title>{{.Title}}</title><meta name="viewport" content="width=device-width, initial-scale=1"></head>
//<body style="font:15px/1.5 Arial, sans-serif; margin:20px;">
//  <h1>{{.Title}}</h1>
//  <p style="color:#555;">Slot: {{.Slot}} | CC: {{.CC}} | Domain: {{.D}} | LID: {{.LID}} | PID: {{.PID}} | Size: {{.TSize}}</p>
//  <p style="color:#777;">Referrer: {{.KwRf}}</p>
//  <p style="color:#777;">Page Title: {{.PTitle}}</p>
//  <p style="color:#777;">Page URL: {{.RURL}}</p>
//  <p style="color:#777;">Keyword ID: {{.KID}}</p>
//  {{if .IsBot}}
//    <div style="margin:10px 0; padding:10px; background:#fff7ed; border:1px solid #fed7aa; color:#9a3412; border-radius:6px;">Bot detected: ad clicks are disabled.</div>
//  {{end}}
//  <hr>
//  <div class="sponsored-ads" style="margin-top:16px;">
//    {{if .HasAds}}
//      {{range .Ads}}
//        <div class="ad-item" style="border:1px solid #e5e7eb; border-radius:8px; padding:12px; margin-bottom:10px;">
//          {{if .RenderLinks}}
//            <a href="{{.ClickHref}}" rel="nofollow noopener" target="_blank" style="font-weight:600; color:#0b57d0; text-decoration:none;">{{.TitleHTML}}</a>
//          {{else}}
//            <span style="font-weight:600; color:#0b57d0; text-decoration:none; cursor:not-allowed;">{{.TitleHTML}}</span>
//          {{end}}
//          {{if .DescHTML}}<div class="ad-desc" style="color:#374151; margin-top:6px;">{{.DescHTML}}</div>{{end}}
//          {{if .Host}}<div class="ad-host" style="color:#6b7280; margin-top:6px; font-size:13px;">{{.Host}}</div>{{end}}
//        </div>
//      {{end}}
//    {{else}}
//      <div style="color:#999;">No sponsored ads available</div>
//    {{end}}
//  </div>
//  <p><a href="javascript:history.back()">Back</a></p>
//</body></html>`
