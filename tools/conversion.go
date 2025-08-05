package tools

import (
	"China_Telecom_Monitor/models"
	"fmt" // 导入 fmt 包，用于错误处理
	"github.com/golang-module/carbon/v2"
	"strconv"
	"strings"
)

// ToSummary 函数将重要数据转换为摘要信息。
// 优化了错误处理，并简化了代码结构。
func ToSummary(qryImportantData *models.Result[models.ImportantData], username string, time carbon.Carbon) models.Summary {
	var ds models.Summary

	// 检查输入数据是否有效。
	if qryImportantData == nil || qryImportantData.HeaderInfos.Code != "0000" || qryImportantData.ResponseData.ResultCode != "0000" {
		return ds // 如果数据无效，则返回空摘要。
	}

	data := qryImportantData.ResponseData.Data

	// 使用闭包处理流量转换，避免重复代码
	parseInt64 := func(s string) int64 {
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing int64: %s, error: %v\n", s, err) // 打印错误信息，方便调试
			return 0                                                        // 出错时返回 0，避免程序崩溃
		}
		return val
	}

	parseFloat64 := func(s string) float64 {
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("Error parsing float64: %s, error: %v\n", s, err) // 打印错误信息，方便调试
			return 0                                                          // 出错时返回 0，避免程序崩溃
		}
		return val
	}

	useFlow := parseInt64(data.FlowInfo.TotalAmount.Used)
	balanceFlow := parseInt64(data.FlowInfo.TotalAmount.Balance)
	totalFlow := useFlow + balanceFlow

	generalUse := parseInt64(data.FlowInfo.CommonFlow.Used)
	generalBalance := parseInt64(data.FlowInfo.CommonFlow.Balance)
	generalTotal := generalUse + generalBalance

	specialUse := parseInt64(data.FlowInfo.SpecialAmount.Used)
	specialBalance := parseInt64(data.FlowInfo.SpecialAmount.Balance)
	specialTotal := specialUse + specialBalance

	voiceUsage := parseInt64(data.VoiceInfo.VoiceDataInfo.Used)
	voiceAmount := parseInt64(data.VoiceInfo.VoiceDataInfo.Total)

	balanceFloat := parseFloat64(data.BalanceInfo.IndexBalanceDataInfo.Balance)
	balance := int64(balanceFloat * 100)

	var items []models.SummaryItems
	flowLists := data.FlowInfo.FlowList
	if flowLists != nil && len(flowLists) > 0 {
		items = make([]models.SummaryItems, 0, len(flowLists)) // 初始化 items 切片，预分配容量
		for _, flowList := range flowLists {
			if !strings.Contains(flowList.Title, "流量") {
				continue // 如果标题不包含 "流量"，则跳过
			}

			var use, balanceF int64
			if strings.Contains(flowList.LeftTitle, "已用") {
				use, _ = ToInt64(flowList.LeftTitleHh) // 这里假设 ToInt64 函数存在且可用
			}
			if strings.Contains(flowList.RightTitle, "剩余") {
				balanceF, _ = ToInt64(flowList.RightTitleHh) // 这里假设 ToInt64 函数存在且可用
			}

			items = append(items, models.SummaryItems{ // 使用 append 添加元素
				Name:  flowList.Title,
				Use:   use,
				Total: use + balanceF,
			})
		}
	}

	// 使用闭包处理 itemsMB 的转换
	convertItems := func(items []models.SummaryItems) []models.SummaryItems {
		itemsMB := make([]models.SummaryItems, len(items))
		for i, item := range items {
			itemsMB[i] = models.SummaryItems{
				Name:  item.Name,
				Use:   float64(item.Use) / 1024,
				Total: float64(item.Total) / 1024,
			}
		}
		return itemsMB
	}

	return models.Summary{
		Username:     username,
		Use:          float64(useFlow) / 1024,
		Total:        float64(totalFlow) / 1024,
		Balance:      balance,
		VoiceUsage:   voiceUsage,
		VoiceAmount:  voiceAmount,
		GeneralUse:   float64(generalUse) / 1024,
		GeneralTotal: float64(generalTotal) / 1024,
		SpecialUse:   float64(specialUse) / 1024,
		SpecialTotal: float64(specialTotal) / 1024,
		CreateTime: carbon.DateTime{
			Carbon: time,
		},
		Items: convertItems(items), // 直接调用转换函数
	}
}