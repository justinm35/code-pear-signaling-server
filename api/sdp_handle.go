package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justinm35/code-pear-signaling-server/models"
	"github.com/teris-io/shortid"
)


func Serve() {
	router := gin.Default()

	router.POST("/offer_provide_sdp", offerProvideSdp)
	router.GET("/offer_poll_for_sdp", offerPollForSdp)
	router.GET("/accept_get_sdp", acceptGetSdp)
	router.POST("/accept_provide_sdp", acceptProvideSdp)

	router.Run("0.0.0.0:4000")
}

func logRandomMessage(c *gin.Context) {
	fmt.Println("Endpoint was hit")
	c.IndentedJSON(http.StatusOK,  gin.H{
		"message": "Hello World",
	} )
}


// Randmly generates a unique identifier
// Takes an encoded SDP Offer and stores it in postgresql under
// Returns the unique identifier
func offerProvideSdp(c *gin.Context) {
	sdp := c.PostForm("sdp")
	if sdp == "" {
		fmt.Println("Nothing provided as offer SDP")
	}
	fmt.Println(sdp)

	// This package is excessive and should be removed
	accessKey, err := shortid.Generate()
	if err != nil {
		fmt.Println("Error generating access key")
	}

	models.InsertConnectionRecord(models.Connection{
		OfferClientSdp: sdp,
		AccessKey: accessKey,
	})

	c.String(http.StatusOK, accessKey)
}

// Offer client can poll this endpoint and retrieve the 
// Acceptor SDP once it is available
func offerPollForSdp(c *gin.Context) {
	access_key := c.Query("access_key")
	if access_key == "" {
		c.String(http.StatusBadRequest, "No access key provided")
	}

	sdpAcceptorKey := models.GetConnectionByAccessKey(access_key).AcceptClientSdp

	if sdpAcceptorKey == "" {
		c.String(http.StatusNotFound, "No Acceptor SDP found yet")
	} else {
		c.String(http.StatusOK, sdpAcceptorKey)
	}
}

// Takes the Access Key and finds the Offer SDP
// If available and sends it back
func acceptGetSdp(c *gin.Context) {
	sdp := c.Query("access_key")
	
	fetchedOfferSdpString := models.GetConnectionByAccessKey(sdp).OfferClientSdp

	c.String(http.StatusOK, fetchedOfferSdpString)
}

// Acceptor can send to this endpoint both the access key in the query params and the 
// SDP in the form which can then be passed to the 
func acceptProvideSdp(c *gin.Context) {
	access_key := c.Query("access_key")
	if access_key == "" {
		c.String(http.StatusBadRequest, "No access key provided")
	}

	sdp := c.PostForm("sdp")
	if sdp == "" {
		c.String(http.StatusBadRequest, "No SDP provided")
	}

	models.UpdateConnectionRecordByAccessKey(access_key, sdp)

	c.String(http.StatusOK, "SDP stored")
}

