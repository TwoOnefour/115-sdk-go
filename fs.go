// https://www.yuque.com/115yun/open/qur839kyx9cgxpxi
package sdk

import (
	"context"
	"net/http"
	"strconv"
)

type MkdirResp struct {
	FileName string `json:"file_name"`
	FileID   string `json:"file_id"`
}

func (c *Client) Mkdir(ctx context.Context, pid, filename string) (*MkdirResp, error) {
	var resp MkdirResp
	_, err := c.AuthRequest(ctx, ApiFsMkdir, http.MethodPost, &resp, ReqWithJson(Json{
		"pid":       pid,
		"file_name": filename,
	}))
	if err != nil {
		return nil, err
	}
	return &resp, err
}

type GetFilesReq struct {
	CID         string
	Type        string
	Limit       int64
	Offset      int64
	Suffix      string
	ASC         bool   // 1: asc, 0: desc
	O           string // order by file_name|file_size|user_utime|file_type
	CustomOrder int
	Stdir       int  // 筛选文件时， 是否显示文件夹；1:要展示文件夹 0不展示
	Star        bool // 1: filter star files
	Cur         int  // 是否只显示当前文件夹内文件
	ShowDir     bool // 是否显示目录；0 或 1，默认为0
}

type GetFilesResp struct {
	Resp[[]struct {
		Fid      string   `json:"fid"`  // 文件ID
		Aid      string   `json:"aid"`  // 文件的状态，aid 的别名。1 正常，7 删除(回收站)，120 彻底删除
		Pid      string   `json:"pid"`  // 父文件夹ID
		Fc       string   `json:"fc"`   // 文件分类 0 文件夹 1 文件
		Fn       string   `json:"fn"`   // 文件名
		Fco      string   `json:"fco"`  // 文件夹封面
		Ism      string   `json:"ism"`  // 是否星标，1：星标
		Isp      int      `json:"isp"`  // 是否加密；1：加密
		Pc       string   `json:"pc"`   // 文件提取码
		Upt      int64    `json:"upt"`  // 修改时间
		Uet      int64    `json:"uet"`  // 修改时间
		UpPt     int64    `json:"uppt"` // 上传时间
		Cm       int64    `json:"cm"`
		FDesc    string   `json:"fdesc"`     // 文件备注
		IsPl     int64    `json:"ispl"`      // 是否统计文件夹下视频时长开关
		Fl       []string `json:"fl"`        // 文件标签
		Sha1     string   `json:"sha1"`      // 文件sha1
		FS       int64    `json:"fs"`        // 文件大小
		Fta      string   `json:"fta"`       // 文件状态 0/2 未上传完成，1 已上传完成
		Ico      string   `json:"ico"`       // 文件后缀名
		Fatr     string   `json:"fatr"`      // 音频长度
		IsV      int64    `json:"isv"`       // 是否视频文件
		Def      int64    `json:"def"`       // 视频清晰度；1:标清 2:高清 3:超清 4:1080P 5:4k;100:原画
		Def2     int64    `json:"def2"`      // 视频清晰度；1:标清 2:高清 3:超清 4:1080P 5:4k;100:原画
		PlayLong int64    `json:"play_long"` // 音视频时长
		VImg     string   `json:"v_img"`
		Thumb    string   `json:"thumb"` // 图片缩略图
		Uo       string   `json:"uo"`    // 原图地址
	}]
	Count          int64  `json:"count"`     // 当前目录文件数量
	SysCount       int64  `json:"sys_count"` // 系统文件夹数量
	Offset         int64  `json:"offset"`    // 偏移量
	Limit          int64  `json:"limit"`     // 分页量
	Aid            int    `json:"aid"`       // 文件的状态，aid 的别名。1 正常，7 删除(回收站)，120 彻底删除
	Cid            int64  `json:"cid"`       // 父目录ID
	IsAsc          int    `json:"is_asc"`    // 1: asc, 0: desc
	MinSize        int64  `json:"min_size"`
	MaxSize        int64  `json:"max_size"`
	SysDir         string `json:"sys_dir"`
	HideData       string `json:"hide_data"`        //是否返回文件数据
	RecordOpenTime string `json:"record_open_time"` //是否记录文件夹的打开时间
	Star           int    `json:"star"`             //是否星标；1：星标；0：未星标
	Type           int    `json:"type"`             //一级筛选大分类，1：文档，2：图片，3：音乐，4：视频，5：压缩包，6：应用
	Suffix         string `json:"suffix"`           //一级筛选选其他时填写的后缀名
	Path           []struct {
		Name string `json:"name"` //父目录名
		Aid  int64  `json:"aid"`
		Cid  int64  `json:"cid"`
		Pid  int64  `json:"pid"`
		Isp  int64  `json:"isp"`
		PCid string `json:"p_cid"`
		Fv   string `json:"fv"`
	} `json:"path"` //父目录树
	Cur    int64  `json:"cur"`
	StdDir int    `json:"stdir"`
	Fields string `json:"fields"`
	Order  string `json:"order"`
}

// GetFiles: https://www.yuque.com/115yun/open/kz9ft9a7s57ep868
func (c *Client) GetFiles(ctx context.Context, req *GetFilesReq) (*GetFilesResp, error) {
	var resp GetFilesResp
	_, err := c.AuthRequestRaw(ctx, ApiFsGetFiles, http.MethodGet, &resp, ReqWithQuery(Form{
		"cid":          req.CID,
		"type":         req.Type,
		"limit":        strconv.FormatInt(req.Limit, 10),
		"offset":       strconv.FormatInt(req.Offset, 10),
		"suffix":       req.Suffix,
		"asc":          Ternary(req.ASC, "1", "0"),
		"o":            req.O,
		"custom_order": strconv.Itoa(req.CustomOrder),
		"stdir":        strconv.Itoa(req.Stdir),
		"star":         Ternary(req.Star, "1", "0"),
		"cur":          strconv.Itoa(req.Cur),
		"show_dir":     Ternary(req.ShowDir, "1", "0"),
	}))
	if err != nil {
		return nil, err
	}
	return &resp, err
}

type GetFolderInfoResp struct {
	Count        string `json:"count"`
	Size         string `json:"size"`
	FolderCount  int64  `json:"folder_count"`
	PlayLong     int64  `json:"play_long"`
	ShowPlayLong int64  `json:"show_play_long"`
	PTime        string `json:"ptime"`
	UTime        string `json:"utime"`
	FileName     string `json:"file_name"`
	PickCode     string `json:"pick_code"`
	Sha1         string `json:"sha1"`
	FileID       string `json:"file_id"`
	IsMark       string `json:"is_mark"`
	OpenTime     int64  `json:"open_time"`
	FileCategory string `json:"file_category"`
	Paths        []struct {
		FileID   int64  `json:"file_id"`
		FileName string `json:"file_name"`
	} `json:"paths"`
}

// GetFolderInfo: https://www.yuque.com/115yun/open/rl8zrhe2nag21dfw
func (c *Client) GetFolderInfo(ctx context.Context, fileID string) (*GetFolderInfoResp, error) {
	var resp GetFolderInfoResp
	_, err := c.AuthRequest(ctx, ApiFsGetFolderInfo, http.MethodGet, &resp, ReqWithQuery(Form{
		"file_id": fileID,
	}))
	if err != nil {
		return nil, err
	}
	return &resp, err
}
