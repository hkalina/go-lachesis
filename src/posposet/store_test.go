package posposet

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"go.etcd.io/bbolt"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/src/inter/pos"
	"github.com/Fantom-foundation/go-lachesis/src/kvdb"
	"github.com/Fantom-foundation/go-lachesis/src/logger"
)

/*
 * bench:
 */

func BenchmarkStore(b *testing.B) {
	logger.SetTestMode(b)

	benchmarkStore(b)
}

func benchmarkStore(b *testing.B) {
	dir, err := ioutil.TempDir("", "poset-bench")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			panic(err)
		}
	}()

	// open history DB
	pathDb := filepath.Join(dir, "lachesis.bolt")
	db, err := bbolt.Open(pathDb, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer os.Remove(pathDb)
	defer db.Close()

	historyCache := kvdb.NewCacheWrapper(kvdb.NewBoltDatabase(db, nil, nil))

	var (
		tempCount int
		tempCache *kvdb.CacheWrapper
	)
	newTempDb := func(name string) kvdb.Database {
		pathTemp := filepath.Join(dir, fmt.Sprintf("lachesis.%d.tmp.bolt", tempCount))
		tempDb, err := bbolt.Open(pathTemp, 0600, nil)
		if err != nil {
			panic(err)
		}
		// counter
		tempCount++

		tempCache = kvdb.NewCacheWrapper(
			kvdb.NewBoltDatabase(
				tempDb,
				tempDb.Close,
				func() error {
					return os.Remove(pathTemp)
				}))
		return tempCache
	}

	input := NewEventStore(historyCache)
	defer input.Close()

	store := NewStore(historyCache, newTempDb)
	defer store.Close()

	nodes := inter.GenNodes(5)

	p := benchPoset(nodes, input, store)

	// flushes both epoch DB and history DB
	flushAll := func() {
		err := historyCache.Flush()
		if err != nil {
			b.Fatal(err)
		}
		err = tempCache.Flush()
		if err != nil {
			b.Fatal(err)
		}
	}

	// run test with random DAG, N + 1 epochs long
	b.ResetTimer()
	maxEpoch := idx.SuperFrame(b.N) + 1
	for epoch := idx.SuperFrame(1); epoch <= maxEpoch; epoch++ {
		buildEvent := func(e *inter.Event) *inter.Event {
			if e.Seq == 1 && e.Creator == nodes[0] {
				// move stake from node0 to node1, to test that it doesn't break anything
				e.InternalTransactions = append(e.InternalTransactions,
					&inter.InternalTransaction{
						Nonce:    0,
						Amount:   1,
						Receiver: nodes[1],
					})
			}
			e.Epoch = epoch
			return p.Prepare(e)
		}
		onNewEvent := func(e *inter.Event) {
			input.SetEvent(e)
			_ = p.ProcessEvent(e)

			if (historyCache.NotFlushedSizeEst() + tempCache.NotFlushedSizeEst()) >= 1024*1024 {
				flushAll()
			}
		}

		r := rand.New(rand.NewSource(int64((epoch))))
		_ = inter.GenEventsByNode(nodes, int(SuperFrameLen*3), 3, buildEvent, onNewEvent, r)
	}

	flushAll()
}

func benchPoset(nodes []hash.Peer, input EventSource, store *Store) *Poset {
	balances := make(map[hash.Peer]pos.Stake, len(nodes))
	for _, addr := range nodes {
		balances[addr] = pos.Stake(1)
	}

	if err := store.ApplyGenesis(balances, genesisTestTime); err != nil {
		panic(err)
	}

	poset := New(store, input)
	poset.Bootstrap()

	return poset
}
