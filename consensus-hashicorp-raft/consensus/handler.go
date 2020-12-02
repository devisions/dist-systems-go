package consensus

import (
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	// When communicating with a Raft cluster, clients should communicate with the leader node.
	// If stale reads are acceptable, queries can also be performed on the follower nodes.
	for k, rf := range rafts {
		if k == rf.Leader() {
			state := r.FormValue("next")
			// We run apply on the leader node.
			// Internally, the leader runs apply on all the other nodes.
			result := rf.Apply([]byte(state), 1*time.Second)
			if result.Error() != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			newState, ok := result.Response().(string)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if newState != state {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte("invalid transition"))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(result.Response().(string)))
			return
		}
	}
}
