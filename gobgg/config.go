package gobgg

const (
	BaseURL = "https://www.boardgamegeek.com/xmlapi2/"
	CollectionSuffix = "collection"
	FamilySuffix = "family"
	ForumListSuffix = "forumlist"
	ForumsSuffix = "forum"
	GuildSuffix = "guild"
	HotSuffix = "hot"
	PlaySuffix = "plays"
	SearchSuffix = "search"
	ThingSuffix = "thing"
	ThreadsSuffix = "thread"
	UserSuffix = "user"
)

type Query struct{
	Username 	string
	Domain 		string
	Id 			int
	Page 		int
	Buddies 	bool
	Guilds 		bool
	Hot 		bool
	Top 		bool
	Members		bool
	Sort 		string
	Type		string
}

