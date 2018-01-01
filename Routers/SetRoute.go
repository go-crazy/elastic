package Routers

import (
	Gin "github.com/gin-gonic/gin"
	"github.com/go-crazy/elastic/Utils"
	"github.com/go-crazy/elastic/Models"
	"github.com/golang/net/context"
)

func SetHandler(c *Gin.Context)  {
	dat :=  Utils.BodyToJson(c.Request)
	eType := dat["type"].(string)
	Index := dat["index"].(string)
	bodyData := dat["source"]
	id := dat["id"].(string)
	parent_id := dat["parent_id"].(string)
	operation := dat["operation"].(string)
	client :=  Models.GetElasticCon(Utils.ElasticUrl())
	
	var err error
	if operation == "create_index"{
		// 创建索引
		_,err = client.CreateIndex(Index).BodyJson(bodyData).Do(context.Background())
	}else if operation == "del_index"{
		// 删除索引
		_,err = client.DeleteIndex(Index).Do(context.Background())
	}else if operation == "add"{
		indexService := client.Index().Index(Index)
		if parent_id != "" {
			indexService =  indexService.Parent(parent_id)
		}
		_,err = indexService.Id(id).Type(eType).BodyJson(bodyData).Do(context.Background())
	}else if operation == "update"{
		updateSevice := client.Update().Index(Index)
		if parent_id != "" {
			updateSevice = updateSevice.Parent(parent_id)
		}
		_,err = updateSevice.Type(eType).Id(id).Doc(bodyData).DetectNoop(true).Do(context.TODO())
	}else if(operation=="delete"){
		deleteService := client.Delete().Index(Index)
		if parent_id != ""{
			deleteService = deleteService.Parent(parent_id)
		}
		_,err = deleteService.Id(id).Type(eType).Do(context.TODO())
	}
	if err != nil {
		c.AbortWithError(500,err)
		return
	}
   	Api_response(c,Gin.H{"status": "ok"})
}
