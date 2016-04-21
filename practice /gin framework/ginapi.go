package main

import (
    "github.com/gin-gonic/gin"
    "database/sql"
    //"github.com/coopernurse/gorp"
  _ "github.com/mattn/go-sqlite3"
    "log"
    //"time"
    "strconv"
)

//var dbmap = initDb()

func main(){

    //defer dbmap.Db.Close()

    router := gin.Default()
    router.GET("/articles", ArticlesList)
   // router.POST("/articles", ArticlePost)
  //  router.GET("/articles/:id", ArticlesDetail)
    router.Run(":8111")
}

type stu struct {
    Id int64 `db:"id"`
    Name int64
    Dep string
    //Content string
}

/*func createArticle(title, body string) Article {
    article := Article{
        Created:    time.Now().UnixNano(),
        Title:      title,
        Content:    body,
    }

    err := dbmap.Insert(&article)
    checkErr(err, "Insert failed")
    return article
}*/

/*func getArticle(article_id int) Article {
    article := Article{}
    err := dbmap.SelectOne(&stu, "select * from student where id = $1", id)
    checkErr(err, "SelectOne failed")
    return article
}*/

func ArticlesList(c *gin.Context) {
     db, err := sql.Open("postgres", "user=postgres password='root' dbname=aaron sslmode=disable")
    checkErr(err, "sql.Open failed")
    var articles []stu
    err:= db.Select(&articles, "select * from student")
    checkErr(err, "Select failed")
    content := gin.H{}
    for k, v := range articles {
        content[strconv.Itoa(k)] = v
    }
    c.JSON(200, content)
}

/*func ArticlesDetail(c *gin.Context) {
    article_id := c.Params.ByName("id")
    a_id, _ := strconv.Atoi(article_id)
    article := getArticle(a_id)
    content := gin.H{"title": article.Title, "content": article.Content}
    c.JSON(200, content)
}*/

/*func ArticlePost(c *gin.Context) {
    var json Article

    c.Bind(&json) // This will infer what binder to use depending on the content-type header.
    article := createArticle(json.Title, json.Content)
    if article.Title == json.Title {
        content := gin.H{
            "result": "Success",
            "title": article.Title,
            "content": article.Content,
        }
        c.JSON(201, content)
    } else {
        c.JSON(500, gin.H{"result": "An error occured"})
    }
}*/

//func initDb() *gorp.DbMap {
  //  db, err := sql.Open("postgres", "user=postgres password='root' dbname=aaron sslmode=disable")
    //checkErr(err, "sql.Open failed")
    
   // dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

   // dbmap.AddTableWithName(stu{}, "articles").SetKeys(true, "Id")

   // err = dbmap.CreateTablesIfNotExists()
   // checkErr(err, "Create tables failed")

    //return dbmap
//}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}