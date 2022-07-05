package httpquality

import (
	"context"
	"dbproxy/db/ck"
	"fmt"
	"strings"
	"time"
)

var COLUMNS = []string{
	"t",
	"name",
	"src_machine_id",
	"src_region",
	"src_province",
	"src_isp",
	"size_download",
	"connect_timeout",
	"http_code",
	"download_speed",
	"time_total",
	"time_connect",
	"time_name_lookup",
	"time_pre_transfer",
	"time_start_transfer",
}

type Row struct {
	T                 int64   `json:"t"`
	Name              string  `json:"name"`
	SrcMachineId      string  `json:"src_machine_id"`
	SrcRegion         string  `json:"src_region"`
	SrcProvince       string  `json:"src_province"`
	SrcIsp            string  `json:"src_isp"`
	SizeDownload      float64 `json:"size_download"`
	ConnectTimeout    uint8   `json:"connect_timeout"`
	HttpCode          uint8   `json:"http_code"`
	DownloadSpeed     float64 `json:"download_speed"`
	TimeTotal         float64 `json:"time_total"`
	TimeConnect       float64 `json:"time_connect"`
	TimeNameLookup    float64 `json:"time_name_lookup"`
	TimePreTransfer   float64 `json:"time_pre_transfer"`
	TimeStartTransfer float64 `json:"time_start_transfer"`
}

func Insert(rows []Row) error {
	ctx := context.Background()
	batch, err := ck.DB.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO ipaas.httpquality(%s)", strings.Join(COLUMNS, ",")))
	if err != nil {
		return err
	}

	for _, row := range rows {
		if err = batch.Append(
			time.Unix(row.T, 0),
			row.Name,
			row.SrcMachineId,
			row.SrcRegion,
			row.SrcProvince,
			row.SrcIsp,
			row.SizeDownload,
			row.ConnectTimeout,
			row.HttpCode,
			row.DownloadSpeed,
			row.TimeTotal,
			row.TimeConnect,
			row.TimeNameLookup,
			row.TimePreTransfer,
			row.TimeStartTransfer,
		); err != nil {
			return err
		}
	}

	if err = batch.Send(); err != nil {
		return err
	}

	return nil
}
