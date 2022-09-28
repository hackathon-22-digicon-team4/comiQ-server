package handler

import (
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

type GetBookUserStampCountsResponse struct {
	BookUserStampCounts []BookUserStampCount `json:"bookUserStampCounts"`
}

type BookUserStampCount struct {
	StampID string `json:"stampId,omitempty"`
	StampName string `json:"stampName,omitempty"`
	StampImageURL string `json:"stampImageUrl,omitempty"`
	Count int `json:"count,omitempty"`
}

// bookIdとbookSeiresIdをクエリパラメータとして受け取る
// クエリパラメータを元に、BookUserStampのスライスを取得する
// BookUserStampのスライスを元に、各スタンプごとに押された数を集計する
// 集計結果を返す
func (h *Handlers) GetBookUserStampCounts(c echo.Context) error {
	ctx := c.Request().Context()
	bookID := c.QueryParam("bookId")
	bookSeriesID := c.QueryParam("bookSeriesId")
	bookUserStamps, err := h.Repository.FindBookUserStampsByQuery(ctx, bookSeriesID, bookID, "", "", "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// stampIDをキーに、スタンプごとに押された数を集計する
	stampCountMap := make(map[string]int)
	stampIDs := make([]string, 0)
	for _, bus := range bookUserStamps {
		stampCountMap[bus.StampID]++
	}
	// stampCountMapに含まれているstampIDをリスト化する
	for stampID := range stampCountMap {
		stampIDs = append(stampIDs, stampID)
	}
	// stampIDを元に、スタンプの情報を取得する
	stamps, err := h.Repository.FindStampsByIDs(ctx, stampIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// stampCountMapを元に、BookUserStampCountを作成する
	bookUserStampCounts := make([]BookUserStampCount, 0)
	for _, stamp := range stamps {
		bookUserStampCounts = append(bookUserStampCounts, BookUserStampCount{
			StampID: stamp.ID,
			StampName: stamp.Name,
			StampImageURL: stamp.ImageURL(h.AssetHost),
			Count: stampCountMap[stamp.ID],
		})
	}
	// BookUserStampCountsをCountが多い順にソートする
	sort.Slice(bookUserStampCounts, func(i, j int) bool {
		return bookUserStampCounts[i].Count > bookUserStampCounts[j].Count
	})

	return c.JSON(http.StatusOK, GetBookUserStampCountsResponse{
		BookUserStampCounts: bookUserStampCounts,
	})
}
