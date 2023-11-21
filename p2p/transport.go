package p2p

// peer represents remote node
type Peer interface {
}

// transport anything that handles the communications between the nodes in the network
// this can be of the form (TCP,UDP,websocket...)
type Transport interface {
}
