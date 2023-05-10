package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"maan/common"
	"maan/common/api"
	"maan/common/copier"
	"maan/common/errs"
	"maan/common/tms"
	validate2 "maan/common/validate"
	"maan/pkg/connections/database/transaction"
	"maan/pkg/domain"
	"maan/pkg/dto"
	"maan/pkg/fuckqr"
	"maan/pkg/model"
	"maan/pkg/public"
	"net/http"
)

type HandlerScanRecord struct {
	scanRecordDomain *domain.ScanRecordDomain
	tx               *transaction.Transaction
}

func NewHandlerScanRecord() *HandlerScanRecord {
	return &HandlerScanRecord{
		scanRecordDomain: domain.NewScanRecordDomain(),
		tx:               transaction.NewTransaction(),
	}
}

func (h *HandlerScanRecord) AnalysisQrScan(ctx *gin.Context) {
	var result = &common.Result{}
	var req = dto.QrScanReq{}
	var resp = dto.QrScanResp{}

	// mvss初始分数
	var mvss = 100
	// 风险类型
	var risk = public.InitType

	// 参数绑定
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(errs.ParameterError))
		return
	}

	// 参数校验
	validate := validator.New()
	err2 := validate.Struct(req)
	if err2 != nil {
		ctx.JSON(http.StatusOK, result.Fail(errs.ParameterError))
		return
	}

	// 判断扫描结果是否是一个url类型的地址，如果是，则进行地址风险检测，否则对内容进行检测
	if !validate2.VerifyUrlFormat(req.Content) {

		// 检测扫描内容风险
		subMvss, isPassDfa := fuckqr.ContentCheck(req.Content)
		if subMvss >= 41 {
			mvss -= subMvss
			risk = public.UnSafeContent
		} else {
			mvss -= subMvss
		}

		if isPassDfa {
			resp.IsPassDfa = 1
		} else {
			resp.IsPassDfa = 0
		}

		resp.Mvss = mvss
		resp.RiskType = risk

		ctx.JSON(http.StatusOK, result.Success(resp))
		return
	}

	ip := tms.UrlToIp(req.Content)

	// 得到ip归属地信息
	ipInfo, err := api.GetIpInfo(ip)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
	}

	// 检测html文档内容风险
	subMvss, urlTitle, isPassDfa := fuckqr.UrlContentCheck(req.Content)
	if subMvss >= 21 {
		mvss -= subMvss
		risk = public.UnSafeDocContent
	} else {
		mvss -= subMvss
		risk = public.UnSafeFormOrJs
	}

	// 验证ssl证书的有效性
	sslCheck := fuckqr.SslCheck(req.Content)
	if !sslCheck {
		mvss -= public.UnSafeSslMvss
		risk = public.UnSafeSsl
	}

	// ip归属地威胁分析
	check := fuckqr.IpAddrCheck(ipInfo.Adcode.O)
	if !check {
		mvss -= public.UnSafeIpAddrMvss
		risk = public.UnSafeIpAddr
	}

	// 尝试命中恶意ip库
	hackerIpCheck := fuckqr.HackerIpCheck(ip)
	if !hackerIpCheck {
		mvss -= public.RelHackerIpMvss
		risk = public.RelHackerIp
	}

	if isPassDfa {
		resp.IsPassDfa = 1
	} else {
		resp.IsPassDfa = 0
	}

	resp.Mvss = mvss
	resp.IpAddr = ipInfo.Adcode.O
	resp.RiskType = risk
	resp.UrlTitle = urlTitle

	// 进行更细粒度的风险检测
	riskAnalyzer := fuckqr.RiskAnalyzer(resp)
	if riskAnalyzer != "" {
		resp.RiskType = riskAnalyzer
	}

	// 保存扫描结果到数据库
	_ = h.SaveScanRecord(resp, req)

	ctx.JSON(http.StatusOK, result.Success(resp))
	return
}

// 保存扫描结果到数据库

func (h *HandlerScanRecord) SaveScanRecord(resp dto.QrScanResp, resp2 dto.QrScanReq) *errs.BError {

	isSafeFlag := 1
	if resp.Mvss >= 60 && resp.Mvss < 80 {
		isSafeFlag = 2
	}

	if resp.Mvss < 60 {
		isSafeFlag = 3
	}

	scanRecord := &model.ScanRecord{
		Content:      resp2.Content,
		Mvss:         int32(resp.Mvss),
		IpAddr:       resp.IpAddr,
		RiskType:     resp.RiskType,
		ContentTitle: resp.UrlTitle,
		IsDfa:        int32(resp.IsPassDfa),
		IsSafe:       int32(isSafeFlag),
	}

	err := h.scanRecordDomain.SaveScanRecord(scanRecord)
	return err
}

// 返回情报信息

func (h *HandlerScanRecord) FindScanRecord(ctx *gin.Context) {
	var result = &common.Result{}
	var resp = dto.IntelligenceResp{}

	// 得到最近20条扫描结果
	scanRecords, err := h.scanRecordDomain.FindScanRecord()
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
	}

	for _, scanRecord := range scanRecords {
		if len(scanRecord.Content) > 50 {
			scanRecord.Content = scanRecord.Content[0:50] + "..."
		}
	}

	if err := copier.Copy(&resp.IntelligenceList, scanRecords); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		return
	}

	// 防止返回 null 值
	if resp.IntelligenceList == nil {
		resp.IntelligenceList = make([]*dto.Intelligence, 0)
	}

	ctx.JSON(http.StatusOK, result.Success(resp))
	return
}
