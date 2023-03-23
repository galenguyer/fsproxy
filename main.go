package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	host := flag.String("host", "localhost:6969", "host to bind to")
	upstreamFlag := flag.String("upstream", "", "upstream host to proxy to (required)")
	only404 := flag.Bool("only404", false, "only change status code to 404")
	flag.Parse()

	if *upstreamFlag == "" {
		log.Fatal("--upstream is required")
	}
	upstream, err := url.Parse(*upstreamFlag)
	if err != nil {
		log.Fatal(err)
	}
	if upstream.Host == "" {
		log.Fatal("upstream host is required (you may be missing the protocol, ie http://)")
	}
	if upstream.Scheme == "" {
		upstream.Scheme = "http"
	}

	statusCodes := [...]int{
		http.StatusContinue,
		http.StatusSwitchingProtocols,
		http.StatusProcessing,
		http.StatusEarlyHints,
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed,
		http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusPaymentRequired,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusProxyAuthRequired,
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusGone,
		http.StatusLengthRequired,
		http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,
		http.StatusExpectationFailed,
		http.StatusTeapot,
		http.StatusMisdirectedRequest,
		http.StatusUnprocessableEntity,
		http.StatusLocked,
		http.StatusFailedDependency,
		http.StatusTooEarly,
		http.StatusUpgradeRequired,
		http.StatusPreconditionRequired,
		http.StatusTooManyRequests,
		http.StatusRequestHeaderFieldsTooLarge,
		http.StatusUnavailableForLegalReasons,
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
		http.StatusVariantAlsoNegotiates,
		http.StatusInsufficientStorage,
		http.StatusLoopDetected,
		http.StatusNotExtended,
		http.StatusNetworkAuthenticationRequired,
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(req *httputil.ProxyRequest) {
			req.SetURL(upstream)
		},
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		originStatus := resp.StatusCode
		if *only404 {
			resp.StatusCode = http.StatusNotFound
		} else {
			resp.StatusCode = statusCodes[rand.Intn(len(statusCodes))]
		}
		log.Printf("status code %d (%s) changed to %d (%s) for path %s", originStatus, http.StatusText(originStatus), resp.StatusCode, http.StatusText(resp.StatusCode), resp.Request.URL.Path)
		return nil
	}

	log.Println("listening on", *host)
	log.Println("proxying requests to", upstream.String())
	log.Fatal(http.ListenAndServe(*host, proxy))
}
