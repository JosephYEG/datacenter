// Code generated by goctl. DO NOT EDIT.
package types

type Actid struct {
	Actid int64 `json:"actid"` //活动id
}

type ActivityResp struct {
	Actid      int64  `json:"actid"`
	Title      string `json:"title"`       //活动名称
	Descr      string `json:"descr"`       //活动描述
	StartDate  int64  `json:"start_date"`  //活动时间
	EnrollDate int64  `json:"enroll_date"` //投票时间
	EndDate    int64  `json:"end_date"`    //活动结束时间
	Votecount  int64  `json:"votecount"`   //当前活动的总票数
	Viewcount  int64  `json:"viewcount"`   //当前活动的总浏览数
	Type       int64  `json:"type"`        //投票方式
	Num        int64  `json:"num"`         //投票几票
}

type AppUser struct {
	Uid      int64  `json:"uid"`
	Auid     int64  `json:"auid"`
	Beid     int64  `json:"beid"`  //应用id
	Ptyid    int64  `json:"ptyid"` //对应平台
	Nickname string `json:"nickname"`
	Openid   string `json:"openid"`
	Avator   string `json:"avator"`
}

type Application struct {
	Sname       string `json:"Sname"`       //名称
	Logo        string `json:"logo"`        // login
	Isclose     int64  `json:"isclose"`     //是否关闭
	Fullwebsite string `json:"fullwebsite"` // 全站名称
}

type Beid struct {
	Beid int64 `json:"beid"`
}

type EnrollReq struct {
	Actid
	Name    string   `json:"name"`    // 名称
	Address string   `json:"address"` //地址
	Images  []string `json:"images"`  //作品图片
	Descr   string   `json:"descr"`   // 作品描述
}

type EnrollResp struct {
	Actid
	Aeid      int64    `json:"aeid"`      // 作品id
	Name      string   `json:"name"`      // 名称
	Address   string   `json:"address"`   //地址
	Images    []string `json:"images"`    //作品图片
	Descr     string   `json:"descr"`     // 作品描述
	Votecount int64    `json:"votecount"` //当前活动的总票数
	Viewcount int64    `json:"viewcount"` //当前活动的总浏览数
}

type JwtToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	AccessExpire int64  `json:"access_expire,omitempty"`
	RefreshAfter int64  `json:"refresh_after,omitempty"`
}

type LoginAppUser struct {
	Uid      int64  `json:"uid"`
	Auid     int64  `json:"auid"`
	Beid     int64  `json:"beid"`  //应用id
	Ptyid    int64  `json:"ptyid"` //对应平台
	Nickname string `json:"nickname"`
	Openid   string `json:"openid"`
	Avator   string `json:"avator"`
	JwtToken
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Type     int64  `json:"type"` //1.密码登陆，2.短信登陆
	Password string `json:"password"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"` //基本一个手机号码就完事
	Password string `json:"password"`
	Smscode  string `json:"smscode"` //短信码
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type SnsReq struct {
	Beid
	Ptyid   int64  `json:"ptyid"`    //对应平台
	BackUrl string `json:"back_url"` //登陆返回的地址
}

type SnsResp struct {
	Beid
	Ptyid    int64  `json:"ptyid"`     //对应平台
	Appid    string `json:"appid"`     //sns 平台的id
	Title    string `json:"title"`     //名称
	LoginUrl string `json:"login_url"` //微信登陆的地址
}

type Token struct {
	Token string `json:"token"`
}

type UserReply struct {
	Auid     int64  `json:"auid"`
	Uid      int64  `json:"uid"`
	Beid     int64  `json:"beid"`  //应用id
	Ptyid    int64  `json:"ptyid"` //对应平台
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Openid   string `json:"openid"`
	Avator   string `json:"avator"`
	JwtToken
}

type UserReq struct {
	Auid  int64 `json:"auid"`
	Uid   int64 `json:"uid"`
	Beid  int64 `json:"beid"`  //应用id
	Ptyid int64 `json:"ptyid"` //对应平台
}

type VoteReq struct {
	Aeid int64 `json:"aeid"` // 作品id
	Actid
}

type VoteResp struct {
	VoteReq
	Votecount int64 `json:"votecount"` //投票票数
	Viewcount int64 `json:"viewcount"` //浏览数
}

type WxLoginReq struct {
	Beid  int64  `json:"beid"`  //应用id
	Code  string `json:"code"`  //微信登陆密钥
	Ptyid int64  `json:"ptyid"` //对应平台
}

type WxShareResp struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
	Noncestr  string `json:"noncestr"`
	Signature string `json:"signature"`
}

type WxTicket struct {
	Ticket string `json:"ticket"`
}
