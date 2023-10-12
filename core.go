package core

import (
	"github.com/defval/di"
	"github.com/ipfs/boxo/blockservice"
	"github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/boxo/fetcher"
	"github.com/ipfs/boxo/filestore"
	"github.com/ipfs/boxo/mfs"
	blocks "github.com/ipfs/go-block-format"
	ipld "github.com/ipfs/go-ipld-format"
	record "github.com/libp2p/go-libp2p-record"
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	identity "github.com/twingdev/go-twing-identity"
	"github.com/twingdev/go-twing/core/repo"
)

type TwingNode struct {
	Identity peer.ID
	Repo     repo.Repo
	twingID  *identity.Identity
	Services NodeServices
}

type NodeServices struct {
	Peerstore            pstore.Peerstore          `optional:"true"` // storage for other Peer instances
	Blockstore           blockstore.Blockstore     // the block store (lower level)
	Filestore            *filestore.Filestore      `optional:"true"` // the filestore blockstore
	BaseBlocks           blocks.BasicBlock         // the raw blockstore, no filestore wrapping
	Blocks               blockservice.BlockService // the block service, get/add blocks.
	DAG                  ipld.DAGService           // the merkle dag service, get/add objects.
	IPLDFetcherFactory   fetcher.Factory           `name:"ipldFetcher"`   // fetcher that paths over the IPLD data model
	UnixFSFetcherFactory fetcher.Factory           `name:"unixfsFetcher"` // fetcher that interprets UnixFS data
	Reporter             *metrics.BandwidthCounter `optional:"true"`
	Discovery            discovery.Service         `optional:"true"`
	FilesRoot            *mfs.Root
	RecordValidator      record.Validator
}

func init() {
	di.SetTracer(&di.StdTracer{})
}

type Module struct {
	Di      []*di.Option
	Name    string
	Version string
}

type IModule interface {
	Version() string
	Name() string
	Load() error
	Destroy() error
}
