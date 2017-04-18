package bitbucket

type Payload struct {
	//	repository struct {
	//	}
	//
	//	refChanges []struct {

	//	}

	Changesets struct {
		Size       int
		Limit      int
		IsLastPage bool
		Values     []struct {
			FromCommit struct {
				Id        string
				DisplayId string
			}
			ToCommit struct {
				Id        string
				DisplayId string
				Author    struct {
					Name         string
					EmailAddress string
				}
				AuthorTimestamp uint64
				Message         string
				Parents         []struct {
					Id        string
					DisplayId string
				}
			}
		}
	}
}

//type refChanges struct {
//}

//type changesets struct {
//}
