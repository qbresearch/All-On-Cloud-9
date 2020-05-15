package common

const (
	LAMBDA_BLOCK = "XXX"

	LOCAL_TXN  = "LOCAL_TXN"
	GLOBAL_TXN = "GLOBAL_TXN"

	LOCAL_BLOCK_NUM  = "LOCAL-BLOCK-%d%d" // LOCAL_BLOCK-<APP_ID><LOCAL_TXN_NUM>
	GLOBAL_BLOCK_NUM = "GLOBAL-BLOCK-%d%d-%d"

	// Orderer Message Types
	O_REQUEST = "REQUEST"
	O_ORDER   = "ORDER"
	O_SYNC    = "SYNC"

	// -------------- inter application messages --------------
	// Message from primary agent of the sender application to the receiver application
	NATS_ORD_REQUEST = "NATS_ORDERER_REQUEST"

	// NATS inbox messages
	// ORDERER MESSAGES
	NATS_ORD_ORDER = "NATS_ORDERER_ORDER"
	NATS_ORD_SYNC  = "NATS_ORDERER_SYNC"

	// NATS start/stop messages
	NATS_CONSENSUS_INITIATE_MSG = "NATS_CONSENSUS_START"
	NATS_CONSENSUS_DONE_MSG     = "NATS_CONSENSUS_DONE"

	// BPaxos NATS message subjects
	LEADER_TO_DEPS        = "LEADER_TO_DEPS"
	DEPS_TO_LEADER        = "DEPS_TO_LEADER"
	LEADER_TO_PROPOSER    = "LEADER_TO_PROPOSER"
	PROPOSER_TO_CONSENSUS = "PROPOSER_TO_CONSENSUS"
	CONSENSUS_TO_PROPOSER = "CONSENSUS_TO_PROPOSER"
	PROPOSER_TO_REPLICA   = "PROPOSER_TO_REPLICA"
	CLIENT_TO_LEADER      = "CLIENT_TO_LEADER"

	// Number of tolerable failures
	F = 1

	// Number of BPaxos Proposer nodes
	NUM_PROPOSERS                  = 1
	CONSENSUS_TIMEOUT_MILLISECONDS = 1000

	NATS_MANUFACTURER_INBOX = "NATS_MANUFACTURER_INBOX"
	NATS_CARRIER_INBOX      = "NATS_CARRIER_INBOX"
	NATS_BUYER_INBOX        = "NATS_BUYER_INBOX"
	NATS_SUPPLIER_INBOX     = "NATS_SUPPLIER_INBOX"

	NATS_ADD_TO_BC = "NATS_ADD_TO_BLOCKCHAIN"

)

var (
	NATS_ORDERER_SUBJECTS = [...]string{NATS_ORD_ORDER, NATS_ORD_SYNC}
)
