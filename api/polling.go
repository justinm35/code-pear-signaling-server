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
	router.POST("/offer", sdpOffer)
	router.GET("/accept", sdpAccept)
	
	router.Run("localhost:8080")
}

func logRandomMessage(c *gin.Context) {
	fmt.Println("Endpoint was hit")
	c.IndentedJSON(http.StatusOK,  gin.H{
		"message": "Hello World",
	} )
}

func sdpOffer(c *gin.Context) {
	// Randmly generates a unique identifier
	// Takes an encoded SDP Offer and stores it in postgresql under
	// Returns the unique identifier

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

func sdpAccept(c *gin.Context) {
	sdp := c.Query("access_key")
	
	fetchedOfferSdpString := models.GetConnectionByAccessKey(sdp).OfferClientSdp

	c.String(http.StatusOK, fetchedOfferSdpString)
}

	// var myConnection models.Connection
	// if os.Args[1] == "insert" {
	////  models.InsertConnectionRecord(models.Connection{
		//	OfferClientSdp: "This Is a test offer",
		//	AccessKey: "Random Key wowwhah",
		//})
	//} else {
	//	myConnection = models.GetConnectionByAccessKey(os.Args[1])
		//fmt.Println(myConnection)
	//}
