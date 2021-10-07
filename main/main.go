package main

import (
	"flag"
	"fmt"
	"genshinSign/genshin"
	"os"
)

func main() {
	defer handlePanic()

	uid, sToken, cookieToken := getEnv()
	if uid == "" {
		panic("uid can't be empty,")
	}
	if sToken == "" && cookieToken == "" {
		panic("stoken and ctoken can't both be empty")
	}

	var client genshin.Client
	if cookieToken != "" {
		client = genshin.LoginByCookie(uid, cookieToken)
	} else {
		client = genshin.LoginBySToken(uid, sToken)
	}

	roles := client.GetUserGameRolesByCookie()

	fmt.Println()
	for _, role := range roles.Data.List {
		info := client.GetRewardInfo(&role)

		fmt.Println(info.Data.Today)
		fmt.Printf("Already sign in %v days and miss %v days\n", info.Data.TotalSignDay, info.Data.SignCntMissed)

		if info.Data.FirstBind {
			fmt.Printf("Skipped unbinding role %s name:%s lv:%v\n", role.RegionName, role.Nickname, role.Level)
			continue
		}

		if info.Data.IsSign {
			fmt.Printf("Skipped already sign in role %s name:%s lv:%v\n", role.RegionName, role.Nickname, role.Level)
			continue
		}

		sign := client.Sign(&role)
		if sign.Code == 0 {
			fmt.Println("Sign in successfully")
			fmt.Printf("%s name:%s lv:%v\n", role.RegionName, role.Nickname, role.Level)
		} else {
			fmt.Println("Sign in failed: " + sign.Message)
		}
	}

	fmt.Println()
}

func getEnv() (string, string, string) {
	uid := flag.String("uid", "", "BBS uid")
	stoken := flag.String("stoken", "", "stoken")
	ctoken := flag.String("ctoken", "", "cookie token from browser cookie")
	flag.Parse()

	// cover by os env
	if val, ext := os.LookupEnv("uid"); ext {
		uid = &val
	}

	if val, ext := os.LookupEnv("stoken"); ext {
		stoken = &val
	}

	if val, ext := os.LookupEnv("ctoken"); ext {
		ctoken = &val
	}

	return *uid, *stoken, *ctoken
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
