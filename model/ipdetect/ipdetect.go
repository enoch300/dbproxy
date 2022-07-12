package ipdetect

import (
	"context"
	"dbproxy/db/ck"
	"dbproxy/utils"
	"fmt"
	"strings"
	"time"
)

var columns = make([]string, 0)

type Row struct {
	T            int64   `json:"t"`
	Id           string  `json:"id"`
	SrcMachineId string  `json:"src_machine_id"`
	DstMachineId string  `json:"dst_machine_id"`
	Biz          string  `json:"biz"`
	Bd           string  `json:"bd"`
	Bid          string  `json:"bid"`
	SrcIp        string  `json:"src_ip"`
	SrcCountry   string  `json:"src_country"`
	SrcProvince  string  `json:"src_province"`
	SrcCity      string  `json:"src_city"`
	SrcIsp       string  `json:"src_isp"`
	DstDevice    string  `json:"dst_device"`
	DstIp        string  `json:"dst_ip"`
	DstCountry   string  `json:"dst_country"`
	DstProvince  string  `json:"dst_province"`
	DstCity      string  `json:"dst_city"`
	DstIsp       string  `json:"dst_isp"`
	DstPort      string  `json:"dst_port"`
	DstPortAlive uint8   `json:"dst_port_alive"`
	PingAvg      float64 `json:"ping_avg"`
	PingMax      float64 `json:"ping_max"`
	PingMin      float64 `json:"ping_min"`
	PingLoss     float64 `json:"ping_loss"`
	Mtr          float64 `json:"mtr"`
}

func init() {
	columns = utils.StructTags(&Row{}, "json")
}

func Insert(rows []Row) error {
	ctx := context.Background()
	batch, err := ck.DB.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO ipaas.ipdetect(%s)", strings.Join(columns, ",")))
	if err != nil {
		return err
	}

	for _, row := range rows {
		if err = batch.Append(
			time.Unix(row.T, 0),
			row.Id,
			row.SrcMachineId,
			row.DstMachineId,
			row.Biz,
			row.Bd,
			row.Bid,
			row.SrcIp,
			row.SrcCountry,
			row.SrcProvince,
			row.SrcCity,
			row.SrcIsp,
			row.DstDevice,
			row.DstIp,
			row.DstCountry,
			row.DstProvince,
			row.DstCity,
			row.DstIsp,
			row.DstPort,
			row.DstPortAlive,
			row.PingAvg,
			row.PingMax,
			row.PingMin,
			row.PingLoss,
			row.Mtr,
		); err != nil {
			return err
		}
	}

	if err = batch.Send(); err != nil {
		return err
	}

	return nil
}
