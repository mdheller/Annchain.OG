package annsensus

import (
	"github.com/annchain/OG/consensus/bft"
	"github.com/annchain/OG/consensus/dkg"
	"github.com/sirupsen/logrus"
	"sync"
)

// AnnsensusPeerCommunicator routes typess and judge which adapter to use and to receive.
// Do IO work only.
type AnnsensusPeerCommunicator struct {
	incoming          AnnsensusPeerCommunicatorIncoming
	outgoing          AnnsensusPeerCommunicatorOutgoing

	dkgMessageAdapter DkgMessageAdapter
	termHolder        HistoricalTermsHolder
	quit              chan bool
	quitWg            sync.WaitGroup
}

func NewAnnsensusCommunicator(
	incoming AnnsensusPeerCommunicatorIncoming,
	outgoing AnnsensusPeerCommunicatorOutgoing,
	bftMessageAdapter BftMessageAdapter,
	dkgMessageAdapter DkgMessageAdapter,
	termHolder HistoricalTermsHolder) *AnnsensusPeerCommunicator {
	return &AnnsensusPeerCommunicator{
		incoming:          incoming,
		outgoing:          outgoing,
		bftMessageAdapter: bftMessageAdapter,
		dkgMessageAdapter: dkgMessageAdapter,
		termHolder:        termHolder,
		quit:              make(chan bool),
		quitWg:            sync.WaitGroup{},
	}
}

func (r *AnnsensusPeerCommunicator) Run() {
	// keep receiving OG messages and decrypt to incoming channel
	for {
		select {
		case <-r.quit:
			r.quitWg.Done()
			return
		case msg := <-r.incoming.GetPipeIn():
			r.HandleAnnsensusMessage(msg)
		}
	}
}

// HandleConsensusMessage is a sub-router for routing consensus message to either bft,dkg or term.
// As part of Annsensus, bft,dkg and term may not be regarded as a separate component of OG.
// Annsensus itself is also a plugin of OG supporting consensus messages.
// Do not block the pipe for any message processing. Router should not be blocked. Use channel.
func (ap *AnnsensusPeerCommunicator) HandleAnnsensusMessage(annsensusMessage AnnsensusMessage) {
	switch annsensusMessage.GetType() {
	case AnnsensusMessageTypeSigned:
		fallthrough
	case AnnsensusMessageTypePlain:
		// let bft and dkg handle this message since I don't know the content
		bmsg, err := ap.bftMessageAdapter.AdaptAnnsensusMessage(annsensusMessage)
		if err == nil {
			// send to bft
			msgTerm, err := ap.termHolder.GetTermByHeight(bmsg)
			if err != nil {
				logrus.WithError(err).Warn("failed to find appropriate term for msg")
				return
			}
			msgTerm.BftPartner.GetBftPeerCommunicatorIncoming().GetPipeIn() <- bmsg
			return
		}

		dmsg, err := ap.dkgMessageAdapter.AdaptAnnsensusMessage(annsensusMessage)
		if err == nil {
			// send to dkg
			msgTerm, err := ap.termHolder.GetTermByHeight(dmsg)
			if err != nil {
				logrus.WithError(err).Warn("failed to find appropriate term for msg")
				return
			}
			msgTerm.DkgPartner.GetDkgPeerCommunicatorIncoming().GetPipeIn() <- dmsg
			return
		}

		logrus.WithError(err).Warn("error on handling annsensus message")
		return
	default:
		logrus.WithField("msg", annsensusMessage).WithField("IM", ap.termHolder.DebugMyId()).Warn("unsupported annsensus message type")
	}
}

func (ap *AnnsensusPeerCommunicator) BroadcastBft(msg bft.BftMessage, peers []bft.PeerInfo) {
	logrus.WithField("msg", msg).Debug("AnnsensusPeerCommunicator is broadcasting bft message")
	adaptedMessage, err := ap.bftMessageAdapter.AdaptBftMessage(msg)
	if err != nil {
		logrus.WithError(err).Warn("failed to adapt bft message to og message")
		return
	}
	var annsensusPeers []AnnsensusPeer
	for _, peer := range peers {
		annsensusPeers = append(annsensusPeers, AnnsensusPeer{
			Id:             peer.Id,
			PublicKey:      peer.PublicKey,
			Address:        peer.Address,
			PublicKeyBytes: peer.PublicKeyBytes,
		})
	}

	ap.outgoing.Broadcast(adaptedMessage, annsensusPeers)
}

func (ap *AnnsensusPeerCommunicator) UnicastBft(msg bft.BftMessage, peer bft.PeerInfo) {
	logrus.WithField("msg", msg).Debug("AnnsensusPeerCommunicator is unicasting bft message")
	adaptedMessage, err := ap.bftMessageAdapter.AdaptBftMessage(msg)
	if err != nil {
		logrus.WithError(err).Warn("failed to adapt bft message to og message")
		return
	}
	ap.outgoing.Unicast(adaptedMessage, AnnsensusPeer{
		Id:             peer.Id,
		PublicKey:      peer.PublicKey,
		Address:        peer.Address,
		PublicKeyBytes: peer.PublicKeyBytes,
	})
}

func (ap *AnnsensusPeerCommunicator) BroadcastDkg(msg dkg.DkgMessage, peers []dkg.PeerInfo) {
	panic("implement me")
}

func (ap *AnnsensusPeerCommunicator) UnicastDkg(msg dkg.DkgMessage, peer dkg.PeerInfo) {
	panic("implement me")
}

func (ap *AnnsensusPeerCommunicator) Stop() {
	ap.quit <- true
	ap.quitWg.Wait()
}