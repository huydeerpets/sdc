package handler

import (
	"github.com/huydeerpets/sdc/lib"
	"github.com/huydeerpets/sdc/model"
	"strconv"
)

type IgnoreAnswerHandler struct {
	lib.AuthHandler
}

func (self *IgnoreAnswerHandler) Get() {
	aid, _ := self.GetInt64(":aid")
	qid, _ := self.GetInt64(":qid")
	uid, _ := self.GetSession("userid").(int64)
	role, _ := self.GetSession("userrole").(int64)
	if aid > 0 && qid > 0 && uid > 0 {
		if err := model.SetIgnoreAnswer(qid, aid, uid, role); err == nil {
			self.Redirect("/"+strconv.Itoa(int(qid))+"/", 302)
		} else {
			self.Redirect("/", 302)
		}
	} else {
		self.Redirect("/", 302)
	}
}
