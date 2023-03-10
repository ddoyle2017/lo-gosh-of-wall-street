package app

type AuctionLength string

const (
	SHORT     AuctionLength = "SHORT"
	MEDIUM    AuctionLength = "MEDIUM"
	LONG      AuctionLength = "LONG"
	VERY_LONG AuctionLength = "VERY_LONG"
)

type Modifier struct {
	Type_ uint
	Value uint
}

type Item struct {
	Id         uint
	BonusLists []uint
	Context    uint // The context an item was obtained in, e.g. mythic raid, dungeon, quest, etc.
	Modifiers  []Modifier
}

type Auction struct {
	Id        uint64
	Buyout    uint64
	Bid       uint64
	Item      Item
	Quantity  uint
	UnitPrice uint64
	TimeLeft  AuctionLength
}

type AuctionData struct {
	Auctions []Auction
}
