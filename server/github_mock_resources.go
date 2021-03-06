// +build mock

package server

var releasesPage = [][]byte{
	[]byte(`[
		{
			"url": "https://api.github.com/repos/getlantern/lantern/releases/3533390",
			"assets_url": "https://api.github.com/repos/getlantern/lantern/releases/3533390/assets",
			"upload_url": "https://uploads.github.com/repos/getlantern/lantern/releases/3533390/assets{?name,label}",
			"html_url": "https://github.com/getlantern/lantern/releases/tag/2.2.5",
			"id": 3533390,
			"tag_name": "2.2.5",
			"target_commitish": "devel",
			"name": "2.2.5",
			"draft": false,
			"author": {
				"login": "myleshorton",
				"id": 1143966,
				"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
				"gravatar_id": "",
				"url": "https://api.github.com/users/myleshorton",
				"html_url": "https://github.com/myleshorton",
				"followers_url": "https://api.github.com/users/myleshorton/followers",
				"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
				"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
				"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
				"organizations_url": "https://api.github.com/users/myleshorton/orgs",
				"repos_url": "https://api.github.com/users/myleshorton/repos",
				"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
				"received_events_url": "https://api.github.com/users/myleshorton/received_events",
				"type": "User",
				"site_admin": false
			},
			"prerelease": false,
			"created_at": "2016-06-27T20:11:42Z",
			"published_at": "2016-06-27T20:23:39Z",
			"assets": [
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1907977",
					"id": 1907977,
					"name": "update_darwin_amd64.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5490268,
					"download_count": 328,
					"created_at": "2016-06-27T20:23:53Z",
					"updated_at": "2016-06-27T20:24:25Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.5/update_darwin_amd64.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1907978",
					"id": 1907978,
					"name": "update_linux_386.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5081950,
					"download_count": 52,
					"created_at": "2016-06-27T20:24:33Z",
					"updated_at": "2016-06-27T20:24:57Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.5/update_linux_386.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1907979",
					"id": 1907979,
					"name": "update_linux_amd64.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5194961,
					"download_count": 54,
					"created_at": "2016-06-27T20:25:08Z",
					"updated_at": "2016-06-27T20:25:34Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.5/update_linux_amd64.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1907981",
					"id": 1907981,
					"name": "update_windows_386.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5160278,
					"download_count": 885,
					"created_at": "2016-06-27T20:25:42Z",
					"updated_at": "2016-06-27T20:26:13Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.5/update_windows_386.bz2"
				}
			],
			"tarball_url": "https://api.github.com/repos/getlantern/lantern/tarball/2.2.5",
			"zipball_url": "https://api.github.com/repos/getlantern/lantern/zipball/2.2.5",
			"body": ""
		},
		{
			"url": "https://api.github.com/repos/getlantern/lantern/releases/3245059",
			"assets_url": "https://api.github.com/repos/getlantern/lantern/releases/3245059/assets",
			"upload_url": "https://uploads.github.com/repos/getlantern/lantern/releases/3245059/assets{?name,label}",
			"html_url": "https://github.com/getlantern/lantern/releases/tag/2.2.4",
			"id": 3245059,
			"tag_name": "2.2.4",
			"target_commitish": "devel",
			"name": "2.2.4",
			"draft": false,
			"author": {
				"login": "myleshorton",
				"id": 1143966,
				"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
				"gravatar_id": "",
				"url": "https://api.github.com/users/myleshorton",
				"html_url": "https://github.com/myleshorton",
				"followers_url": "https://api.github.com/users/myleshorton/followers",
				"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
				"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
				"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
				"organizations_url": "https://api.github.com/users/myleshorton/orgs",
				"repos_url": "https://api.github.com/users/myleshorton/repos",
				"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
				"received_events_url": "https://api.github.com/users/myleshorton/received_events",
				"type": "User",
				"site_admin": false
			},
			"prerelease": false,
			"created_at": "2016-05-17T19:05:38Z",
			"published_at": "2016-05-17T19:32:38Z",
			"assets": [
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1697374",
					"id": 1697374,
					"name": "update_darwin_amd64.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5490162,
					"download_count": 3567,
					"created_at": "2016-05-17T19:32:51Z",
					"updated_at": "2016-05-17T19:33:20Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.4/update_darwin_amd64.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1697378",
					"id": 1697378,
					"name": "update_linux_386.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5073999,
					"download_count": 479,
					"created_at": "2016-05-17T19:33:34Z",
					"updated_at": "2016-05-17T19:33:55Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.4/update_linux_386.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1697384",
					"id": 1697384,
					"name": "update_linux_amd64.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5196158,
					"download_count": 3694,
					"created_at": "2016-05-17T19:34:05Z",
					"updated_at": "2016-05-17T19:34:30Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.4/update_linux_amd64.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1697389",
					"id": 1697389,
					"name": "update_windows_386.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5166401,
					"download_count": 12230,
					"created_at": "2016-05-17T19:34:38Z",
					"updated_at": "2016-05-17T19:35:01Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.4/update_windows_386.bz2"
				}
			],
			"tarball_url": "https://api.github.com/repos/getlantern/lantern/tarball/2.2.4",
			"zipball_url": "https://api.github.com/repos/getlantern/lantern/zipball/2.2.4",
			"body": ""
		},
		{
			"url": "https://api.github.com/repos/getlantern/lantern/releases/3211864",
			"assets_url": "https://api.github.com/repos/getlantern/lantern/releases/3211864/assets",
			"upload_url": "https://uploads.github.com/repos/getlantern/lantern/releases/3211864/assets{?name,label}",
			"html_url": "https://github.com/getlantern/lantern/releases/tag/2.2.3",
			"id": 3211864,
			"tag_name": "2.2.3",
			"target_commitish": "devel",
			"name": "2.2.3",
			"draft": false,
			"author": {
				"login": "myleshorton",
				"id": 1143966,
				"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
				"gravatar_id": "",
				"url": "https://api.github.com/users/myleshorton",
				"html_url": "https://github.com/myleshorton",
				"followers_url": "https://api.github.com/users/myleshorton/followers",
				"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
				"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
				"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
				"organizations_url": "https://api.github.com/users/myleshorton/orgs",
				"repos_url": "https://api.github.com/users/myleshorton/repos",
				"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
				"received_events_url": "https://api.github.com/users/myleshorton/received_events",
				"type": "User",
				"site_admin": false
			},
			"prerelease": false,
			"created_at": "2016-05-12T14:13:18Z",
			"published_at": "2016-05-12T14:22:22Z",
			"assets": [
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1676475",
					"id": 1676475,
					"name": "update_android_arm.bz2",
					"label": null,
					"uploader": {
						"login": "atavism",
						"id": 6346213,
						"avatar_url": "https://avatars.githubusercontent.com/u/6346213?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/atavism",
						"html_url": "https://github.com/atavism",
						"followers_url": "https://api.github.com/users/atavism/followers",
						"following_url": "https://api.github.com/users/atavism/following{/other_user}",
						"gists_url": "https://api.github.com/users/atavism/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/atavism/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/atavism/subscriptions",
						"organizations_url": "https://api.github.com/users/atavism/orgs",
						"repos_url": "https://api.github.com/users/atavism/repos",
						"events_url": "https://api.github.com/users/atavism/events{/privacy}",
						"received_events_url": "https://api.github.com/users/atavism/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 4974460,
					"download_count": 1911,
					"created_at": "2016-05-12T16:10:07Z",
					"updated_at": "2016-05-12T16:10:15Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.3/update_android_arm.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1675792",
					"id": 1675792,
					"name": "update_darwin_amd64.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5493956,
					"download_count": 38497,
					"created_at": "2016-05-12T14:22:30Z",
					"updated_at": "2016-05-12T14:22:50Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.3/update_darwin_amd64.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1675795",
					"id": 1675795,
					"name": "update_linux_386.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5080211,
					"download_count": 148,
					"created_at": "2016-05-12T14:22:57Z",
					"updated_at": "2016-05-12T14:23:16Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.3/update_linux_386.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1675799",
					"id": 1675799,
					"name": "update_linux_amd64.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5196941,
					"download_count": 1040,
					"created_at": "2016-05-12T14:23:23Z",
					"updated_at": "2016-05-12T14:23:42Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.3/update_linux_amd64.bz2"
				},
				{
					"url": "https://api.github.com/repos/getlantern/lantern/releases/assets/1675801",
					"id": 1675801,
					"name": "update_windows_386.bz2",
					"label": "",
					"uploader": {
						"login": "myleshorton",
						"id": 1143966,
						"avatar_url": "https://avatars.githubusercontent.com/u/1143966?v=3",
						"gravatar_id": "",
						"url": "https://api.github.com/users/myleshorton",
						"html_url": "https://github.com/myleshorton",
						"followers_url": "https://api.github.com/users/myleshorton/followers",
						"following_url": "https://api.github.com/users/myleshorton/following{/other_user}",
						"gists_url": "https://api.github.com/users/myleshorton/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/myleshorton/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/myleshorton/subscriptions",
						"organizations_url": "https://api.github.com/users/myleshorton/orgs",
						"repos_url": "https://api.github.com/users/myleshorton/repos",
						"events_url": "https://api.github.com/users/myleshorton/events{/privacy}",
						"received_events_url": "https://api.github.com/users/myleshorton/received_events",
						"type": "User",
						"site_admin": false
					},
					"content_type": "application/x-bzip2",
					"state": "uploaded",
					"size": 5167505,
					"download_count": 496370,
					"created_at": "2016-05-12T14:23:48Z",
					"updated_at": "2016-05-12T14:24:07Z",
					"browser_download_url": "https://github.com/getlantern/lantern/releases/download/2.2.3/update_windows_386.bz2"
				}
			],
			"tarball_url": "https://api.github.com/repos/getlantern/lantern/tarball/2.2.3",
			"zipball_url": "https://api.github.com/repos/getlantern/lantern/zipball/2.2.3",
			"body": ""
		}
	]`),
	[]byte(`[
	]`),
}
