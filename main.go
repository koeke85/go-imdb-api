package main
import(
    "net/http"
    "gopkg.in/gin-gonic/gin.v1"
    "github.com/eefret/gomdb"
    _"log"
    )

func main(){

    router := gin.Default()
    router.GET("/kago", func(c *gin.Context){
            movieName := c.DefaultQuery("movieName", "Mars")
            //lastname := c.DefaultQuery("movie","Jeremy")
            target := &gomdb.QueryData{Title: movieName, SearchType: gomdb.MovieSearch}
            res, _ := gomdb.Search(target)
            c.String(http.StatusOK, "Hello %s %s", movieName, res.Search)
            })
    router.Run(":9487")


}

