package v1

import (
	"aio/internal/pb"
	"aio/internal/usecase"
	"aio/pkg/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type extensionRoutes struct {
	e usecase.Extension
	l logger.Interface
}

func newExtensionRouter(e *gin.RouterGroup, l logger.Interface, ext usecase.Extension) {
	r := &extensionRoutes{e: ext, l: l}
	g := e.Group("/ext")
	g.POST("/register", r.register)
	//g.Any("/connect", r.Connect)
	{
		g.GET("/:uid/group", r.pull)
		g.POST("/:uid/group", r.addGroup)
		g.DELETE("/:uid/:group/", r.removeGroup)
		g.DELETE("/:uid/:group/:tab", r.removeTab)
		//g.POST("/uid/swap/group")
		g.POST("/:uid/swap/tab", r.swapTab)
	}
}

// register godoc
// @Summery      register extension client
// @Description  get uid
// @Tags         extension
// @Param        name       query  string  true  "client name"   minlength(1)  maxlength(64)
// @Param        extension  query  string  true  "extension id"  minlength(1)  maxlength(256)
// @Produce      json
// @Success      200  {object}  ExtensionResp{data=object{uid=string}}
// @Router       /ext/register [post]
func (e *extensionRoutes) register(c *gin.Context) {
	name := c.Query("name")
	extension := c.Query("extension")
	obj, err := e.e.Register(c.Request.Context(), name, extension)
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	ExtResp(c, 200, 0, nil, gin.H{
		"uid": obj.ClientUID,
	})
}

// addGroup godoc
// @Summery      store client tab group
// @Description  store one or more tab group
// @Tags         add
// @Param        uid     path  string      true  "client uid"
// @Param        groups  body  []pb.Group  true  "groups"
// @Accept       json
// @Produce      json
// @Response     200  {object}  ExtensionResp{}
// @Router       /ext/{uid}/group [post]
func (e *extensionRoutes) addGroup(c *gin.Context) {
	id := c.Param("uid")
	groups := make([]*pb.Group, 0)
	err := c.ShouldBindJSON(&groups)
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	uid, err := uuid.FromString(id)
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	err = e.e.Add(c.Request.Context(), uid, groups...)
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	ExtResp(c, 200, 0, nil, nil)
}

// pull godoc
// @Summery      pull all tab
// @Description  get client's all groups and tabs
// @Tags         pull
// @Param        uid     path   string  true  "client uid"
// @Accept       json
// @Produce      json
// @Response     200  {object}  ExtensionResp{data=pb.Client}
// @Router       /ext/{uid}/group [get]
func (e *extensionRoutes) pull(c *gin.Context) {
	id := c.Param("uid")
	uid, err := uuid.FromString(id)
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	cli, err := e.e.Pull(c.Request.Context(), uid)
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	ExtResp(c, 200, 0, nil, pb.ClientToPB(cli))
}

// removeGroup godoc
// @Summery      remove a tab group
// @Description  remove tab group and all tabs
// @Tags         remove
// @Param        uid    path  string  true  "client uid"
// @Param        group  path  string  true  "group uid"
// @Accept       json
// @Produce      json
// @Response     200  {object}  ExtensionResp{}
// @Router       /ext/{uid}/{group}/ [delete]
func (e *extensionRoutes) removeGroup(c *gin.Context) {
	uid, err := uuid.FromString(c.Param("uid"))
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	guid, err := uuid.FromString(c.Param("group"))
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	err = e.e.RemoveGroup(c.Request.Context(), uid, guid)
	if err != nil {
		e.l.Errorln(err)
	}
	//ExtResp()
}

// removeTab godoc
// @Summery      remove tab
// @Description  remove tab
// @Tags         remove
// @Param        uid    path  string  true  "client uid"
// @Param        group  path  string  true  "group uid"
// @Param        tab    path  string  true  "tab uid"
// @Accept       json
// @Produce      json
// @Response     200  {object}  ExtensionResp{}
// @Router       /ext/{uid}/{group}/{tab} [delete]
func (e *extensionRoutes) removeTab(c *gin.Context) {
	uid := uuid.FromStringOrNil(c.Param("uid"))
	gid := uuid.FromStringOrNil(c.Param("group"))
	tid := uuid.FromStringOrNil(c.Param("tab"))
	err := e.e.RemoveTab(c.Request.Context(), uid, gid, tid)
	if err != nil {
		e.l.Errorln(err)
	}
	//ExtResp()
}

// swapTab godoc
// @Summery      swap tab a with tab b
// @Description  swap 2 tab
// @Tags         modify
// @Param        uid  path  string  true  "client uid"
// @Param        groupA  query  string  true  "group(a) uid"
// @Param        groupB  query  string  true  "group(b) uid"
// @Param        tabA    query  string  true  "group(a) tab(a) uid"
// @Param        tabB    query  string  true  "group(b) tab(b) uid"
// @Accept       json
// @Produce      json
// @Response     200  {object}  ExtensionResp{}
// @Router       /ext/{uid}/swap/tab [post]
func (e *extensionRoutes) swapTab(c *gin.Context) {
	uid := uuid.FromStringOrNil(c.Param("uid"))
	ga := uuid.FromStringOrNil(c.Query("groupA"))
	gb := uuid.FromStringOrNil(c.Query("groupB"))
	ta := uuid.FromStringOrNil(c.Query("tabA"))
	tb := uuid.FromStringOrNil(c.Query("tabB"))

	err := e.e.SwapTab(c.Request.Context(), uid, ga, ta, gb, tb)
	if err != nil {
		e.l.Errorln(err)
	}
	//ExtResp()
}

func (e *extensionRoutes) addAddress(c *gin.Context) {
	uid := uuid.FromStringOrNil(c.Param("uid"))
	addr := &pb.BarkDevice{}
	err := c.ShouldBind(addr)
	if err != nil {

	}
	err = e.e.AddBarkAddress(c.Request.Context(), uid, addr)
	if err != nil {

	}
}

func (e *extensionRoutes) pullAddress(c *gin.Context) {
	uid := uuid.FromStringOrNil(c.Param("uid"))
	list, err := e.e.BarkAddresses(c.Request.Context(), uid)
	if err != nil {

	}
	result := make([]*pb.BarkDevice, 0, len(list))
	for _, v := range list {
		result = append(result, pb.AddressToPB(v))
	}
	ExtResp(c, 200, 0, nil, result)
}

func (e *extensionRoutes) dropAddress(c *gin.Context) {
	uid := uuid.FromStringOrNil(c.Param("uid"))
	addr := &pb.BarkDevice{}
	err := c.ShouldBind(addr)
	if err != nil {

	}
	err = e.e.DropBarkAddress(c.Request.Context(), uid, addr)
	if err != nil {

	}
}

func (e *extensionRoutes) Connect(c *gin.Context) {
	uid, err := uuid.FromString(c.Query("uid"))
	if err != nil {
		ExtResp(c, 200, 0, err, nil)
		return
	}
	err = e.e.Connect(c.Request.Context(), uid, nil)
	if err != nil {
		ExtResp(c, 200, 0, err, nil)
		return
	}
}
