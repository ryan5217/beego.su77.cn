package paginator

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/components"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"math"
	"strconv"
)

func Get(path string, params parameter.Parameters, size int, pageSizeList []string) types.PaginatorAttribute {

	paginator := template2.Default().Paginator().(*components.PaginatorAttribute)

	pageInt, _ := strconv.Atoi(params.Page)
	pageSizeInt, _ := strconv.Atoi(params.PageSize)
	totalPage := int(math.Ceil(float64(size) / float64(pageSizeInt)))

	if params.Page == "1" {
		paginator.PreviousClass = "disabled"
		paginator.PreviousUrl = path
	} else {
		paginator.PreviousClass = ""
		paginator.PreviousUrl = path + params.GetLastPageRouteParamStr()
	}

	if pageInt == totalPage {
		paginator.NextClass = "disabled"
		paginator.NextUrl = path
	} else {
		paginator.NextClass = ""
		paginator.NextUrl = path + params.GetNextPageRouteParamStr()
	}
	paginator.Url = path + params.GetRouteParamStrWithoutPageSize()
	paginator.CurPageEndIndex = strconv.Itoa((pageInt) * pageSizeInt)
	paginator.CurPageStartIndex = strconv.Itoa((pageInt - 1) * pageSizeInt)
	paginator.Total = strconv.Itoa(size)

	paginator.Option = make(map[string]template.HTML, len(pageSizeList))
	for i := 0; i < len(pageSizeList); i++ {
		paginator.Option[pageSizeList[i]] = template.HTML("")
	}

	paginator.Option[params.PageSize] = template.HTML("selected")

	paginator.Pages = []map[string]string{}

	if totalPage < 10 {
		var pagesArr []map[string]string
		for i := 1; i < totalPage+1; i++ {
			if i == pageInt {
				pagesArr = append(pagesArr, map[string]string{
					"page":    strconv.Itoa(i),
					"active":  "active",
					"isSplit": "0",
					"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
				})
			} else {
				pagesArr = append(pagesArr, map[string]string{
					"page":    strconv.Itoa(i),
					"active":  "",
					"isSplit": "0",
					"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
				})
			}
		}
		paginator.Pages = pagesArr
	} else {
		var pagesArr []map[string]string
		if pageInt < 6 {
			for i := 1; i < totalPage+1; i++ {

				if i == pageInt {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "active",
						"isSplit": "0",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
				} else {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "",
						"isSplit": "0",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
				}

				if i == 6 {
					pagesArr = append(pagesArr, map[string]string{
						"page":    "",
						"active":  "",
						"isSplit": "1",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
					i = totalPage - 1
				}
			}
		} else if pageInt < totalPage-4 {
			for i := 1; i < totalPage+1; i++ {

				if i == pageInt {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "active",
						"isSplit": "0",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
				} else {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "",
						"isSplit": "0",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
				}

				if i == 2 {
					pagesArr = append(pagesArr, map[string]string{
						"page":    "",
						"active":  "",
						"isSplit": "1",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
					if pageInt < 7 {
						i = 5
					} else {
						i = pageInt - 2
					}
				}

				if pageInt < 7 {
					if i == pageInt+5 {
						pagesArr = append(pagesArr, map[string]string{
							"page":    "",
							"active":  "",
							"isSplit": "1",
							"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
						})
						i = totalPage - 1
					}
				} else {
					if i == pageInt+3 {
						pagesArr = append(pagesArr, map[string]string{
							"page":    "",
							"active":  "",
							"isSplit": "1",
							"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
						})
						i = totalPage - 1
					}
				}
			}
		} else {
			for i := 1; i < totalPage+1; i++ {

				if i == pageInt {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "active",
						"isSplit": "0",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
				} else {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "",
						"isSplit": "0",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
				}

				if i == 2 {
					pagesArr = append(pagesArr, map[string]string{
						"page":    "",
						"active":  "",
						"isSplit": "1",
						"url":     path + params.SetPage(strconv.Itoa(i)).GetRouteParamStr(),
					})
					i = totalPage - 4
				}
			}
		}
		paginator.Pages = pagesArr
	}

	return paginator.SetPageSizeList(pageSizeList)
}
