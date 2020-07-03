package pirsch

// Contains all substrings (in lowercase) used to filter the User-Agent header.
// Please add the reference in case you copy an existing list.
// - https://github.com/gorangajic/isbot/blob/master/list.json
var userAgentBlacklist = []string{
	"://", // urls
	"bot",
	"crawler",
	"spider",
	"go-http-client",
	"twingly",
	"evc-batch",
	"mailto",
	"newspaper",
	"feedreader",
	"magic browser",
	"ruby",
	"weavr",
	"hackernews",
	"cfnetwork",
	"socialbeeagent",
	"php",
	"daum",
	"deusu",
	"muckrack",
	"splash",
	"sysomos",
	"um-ln",
	"site",
	"vbseo",
	"12345",
	"acoon",
	"activebookmark",
	"activerefresh",
	"activeworlds",
	"ad muncher",
	"addthis",
	"ahc",
	"amazon cloudfront",
	"amigavoyager",
	"anonymous_agent",
	"apache",
	"asafaweb",
	"asynchttp",
	"axios",
	"azureus",
	"biglotron",
	"binlar",
	"blackboard",
	"safeassign",
	"bloglines",
	"bloglovin",
	"browsershots",
	"btwebclient",
	"bubing",
	"cakephp",
	"camo asset proxy",
	"catexplorador",
	"clamav",
	"client",
	"custom",
	"dap ",
	"davclnt",
	"digg",
	"dispatch",
	"disqus",
	"docomo",
	"download",
	"drupact",
	"drupal",
	"duckduckgo",
	"ecatch",
	"evernote clip resolver",
	"facebook$",
	"facebookexternalhit",
	"facebookplatform",
	"faraday",
	"fasthttp",
	"fdm",
	"findlink",
	"flashget",
	"friendica",
	"gigablastopensource",
	"gooblog",
	"googal",
	"goose",
	"grammarly",
	"greatnews",
	"greenbrowser",
	"gregarius",
	"guzzlehttp",
	"hatena",
	"hexometer",
	"hobbit",
	"hotzonu",
	"http",
	"hubspot",
	"hwcdn",
	"ibrowse",
	"ice browser",
	"iframely",
	"infox-wisg",
	"ingrid",
	"java",
	"jetbrains",
	"jetty",
	"jigsaw",
	"libtorrent",
	"libwww",
	"liferea",
	"linkdex",
	"linkwalker",
	"ltx71",
	"lua-resty-http",
	"lwp-",
	"lwp::simple",
	"magpierss",
	"mailchimp",
	"mechanize",
	"meltwaternews",
	"metainspector",
	"metauri",
	"microsoft bits",
	"microsoft data",
	"microsoft office existence",
	"microsoft office protocol discovery",
	"microsoft windows network diagnostics",
	"microsoft-cryptoapi",
	"microsoft-webdav-miniredir",
	"mixmax-linkpreview",
	"mixnodecache",
	"monit",
	"movabletype",
	"duckduckgo",
	"mucommander",
	"my browser",
	"navermailapp",
	"netsurf",
	"nettrack anonymous web statistics",
	"netvibes",
	"newsgator",
	"newspaper",
	"nextcloud-news",
	"ning",
	"node-superagent",
	"nokiac3",
	"notetextview",
	"nuzzel",
	"offline explorer",
	"okhttp",
	"omgili",
	"ossproxy",
	"panscient",
	"pcore-http",
	"pear http_request",
	"photon",
	"postman",
	"postrank",
	"python",
	"ramblermail",
	"raynette_httprequest",
	"realdownload",
	"rebelmouse",
	"riddler",
	"rssbandit",
	"rssowl",
	"ruby$",
	"scrapy",
	"selenium",
	"sentry",
	"seostats",
	"set:",
	"shareaza",
	"shortlinktranslate",
	"snap$",
	"snapchat",
	"space bison",
	"spring ",
	"sprinklr",
	"summify",
	"svn",
	"t-online browser",
	"taringa",
	"test certificate info",
	"the knowledge ai",
	"thinklab",
	"tiny tiny rss",
	"traackr",
	"transmission",
	"tumblr",
	"typhoeus",
	"ubuntu apt-http",
	"upflow",
	"user_agent",
	"utorrent",
	"vbulletin",
	"venus/fedoraplanet",
	"vse",
	"w3c",
	"webcopier",
	"wget",
	"whatsapp",
	"whatweb",
	"windows-rss-platform",
	"www-mechanize",
	"xenu link sleuth",
	"xymon",
	"yahoo",
	"yandex",
	"zabbix",
	"zdm",
	"zend_http_client",
	"zjavascript",
	"zmeu$",
	"abonti",
	"adbeat",
	"amiga-aweb",
	"analyz",
	"anyevent-http",
	"appinsights",
	"archive",
	"ask jeeves/teoma",
	"auto",
	"bingpreview",
	"bluecoat drtr",
	"bordermanager",
	"brandverity",
	"browsex",
	"burpcollaborator",
	"capture",
	"catchpoint",
	"check",
	"chrome-lighthouse",
	"cloudflare",
	"coccoc",
	"collect",
	"commons-httpclient",
	"crawl",
	"daemon",
	"dareboost",
	"datanyze",
	"dataprovider",
	"dejaclick",
	"detector",
	"dmbrowser",
	"domains project",
	"embedly",
	"feed",
	"fetch",
	"finder",
	"flipboardproxy",
	"freesafeip",
	"genieo",
	"getlinkinfo",
	"gomezagent",
	"google",
	"grouphigh",
	"gtmetrix",
	"headlesschrome",
	"heritrix",
	"httrack",
	"hubspot marketing grader",
	"hydra",
	"ichiro",
	"images",
	"index",
	"ips-agent",
	"java",
	"javafx",
	"jorgee",
	"kulturarw3",
	"library",
	"link preview",
	"lipperhey",
	"lucidworks-anda",
	"mail",
	"miniflux",
	"monitor",
	"netcraftsurveyagent",
	"netnewswire",
	"newsfox",
	"nmap scripting engine",
	"nutch",
	"offbyone",
	"outbrain",
	"page2rss",
	"parse",
	"perl",
	"phantom",
	"pingadmin",
	"pingdom",
	"powermarks",
	"powerpc amigaos",
	"pr-cy.ru",
	"prlog",
	"proximic",
	"ptst",
	"qqdownload",
	"qwantify",
	"ranksonicsiteauditor",
	"reader",
	"rivva",
	"scan",
	"scoutjet",
	"scrape",
	"search",
	"server",
	"seznamemailproxy",
	"shrinktheweb",
	"skypeuripreview",
	"slurp",
	"snacktory",
	"sogou",
	"speedmode",
	"spider",
	"spy",
	"statuscake",
	"stumbleupon",
	"synapse",
	"synthetic",
	"thumb",
	"tineye",
	"tracemyfile",
	"trendsmapresolver",
	"tweetedtimes",
	"twingly recon",
	"url",
	"vagabondo",
	"validator",
	"vkshare",
	"wapchoi",
	"wappalyzer",
	"webcorp",
	"webdatastats",
	"webglance",
	"webkit2png",
	"winhttp",
	"wmtips.com/\\d",
	"woorankreview",
	"wordpress",
	"ackerm",
	"alertra",
	"iskanie",
	"megaproxy",
	"moreover",
	"mowser",
	"nearsoftware",
	"ssllabs",
	"yeti",
	"zgrab",
	"validator",
	"vkshare",
	"WAPCHOI",
	"Wappalyzer",
	"webcorp",
	"WebDataStats",
	"Webglance",
	"webkit2png",
	"WinHTTP",
	"wmtips.com/\\d",
	"woorankreview",
	"WordPress",
	"ackerm",
	"alertra",
	"iskanie",
	"megaproxy",
	"moreover",
	"mowser",
	"nearsoftware",
	"ssllabs",
	"yeti",
}
