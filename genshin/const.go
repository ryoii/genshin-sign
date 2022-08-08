package genshin

const (

	/* Maybe changed */

	// SIGN_UA Mhy bbs app user-agent
	webUA = "Mozilla/5.0 (Linux; Android 6.0.1; MuMu Build/V417IR; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.100 Mobile Safari/537.36 miHoYoBBS/" + appVersion
	appUA = "okhttp/4.8.0"

	appVersion = "2.33.1"
	//appVersion = "2.13.1"

	// 2.33.1
	salt = "1OUn34iIy84ypu9cpXyun2VaQ2zuFeLm"

	// 2.13.1
	//salt = "6zT9berkIjLBimVKLeQiyYCN0tatGDpP"

	activityId = "e202009291139501"

	/* Unchanged */

	webUrl = "https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html"
	apiUrl = "https://api-takumi.mihoyo.com/"

	webReferer = webUrl + "?bbs_auth_required=true&act_id=" + activityId + "&utm_source=bbs&utm_medium=mys&utm_campaign=icon"
	appReferer = "https://app.mihoyo.com"

	appClientType = "2"
	webClientType = "4"

	acceptEncoding = "gzip"

	/* api path */

	// Get action ticket
	getActionTicket = apiUrl + "auth/api/getActionTicketBySToken"

	// Get binding user info
	getUserGameRoles = apiUrl + "binding/api/getUserGameRoles"

	// Get binding user info by cookie
	getUserGameRolesByCookie = apiUrl + "binding/api/getUserGameRolesByCookie"

	// Get cookie for web api
	getCookieAccount = apiUrl + "auth/api/getCookieAccountInfoBySToken"

	// Check sign status
	signRewardInfo = apiUrl + "event/bbs_sign_reward/info"

	// Actually sign
	signRewardSign = apiUrl + "event/bbs_sign_reward/sign"
)
