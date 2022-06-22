package detect

import (
	"context"
	"dbproxy/db/ck"
	"fmt"
	"strings"
	"time"
)

var (
	COLUMNS = []string{
		"t",
		"src_machine_id",
		"src_ip",
		"src_asn",
		"dst_machine_id",
		"dst_ip",
		"dst_asn",
		"dst_eth",
		"ping_loss_rate",
		"ping_max_delay",
		"ping_min_delay",
		"ping_avg_delay",
		"http_code",
		"http_download_speed",
		"http_time_connect",
		"http_time_name_lookup",
		"http_time_start_transfer",
		"http_time_time_redirect",
		"http_time_total",
		"udp_out_succ_rate",
		"udp_out_avg_delay",
		"tcp_out_connect_time",
		"tcp_out_succ_rate",
		"tcp_out_avg_rate",
		"host_retrans_rate",
		"eth_send_err_rate",
		"eth_send_drop_rate",
		"bandwidth_limit",
		"hops",
		"mtr"}
)

type Hop struct {
	RouteNo int     `json:"route_num"`
	Addr    string  `json:"addr"`
	Loss    float32 `json:"loss"`
	Snt     int     `json:"snt"`
	Last    float32 `json:"last"`
	Avg     float32 `json:"avg"`
	Best    float32 `json:"best"`
	Wrst    float32 `json:"wrst"`
	StDev   float64 `json:"stdev"`
}

type Row struct {
	T                     int64   `json:"t"`
	SrcMachineId          string  `json:"src_machine_id"`
	DstMachineId          string  `json:"dst_machine_id"`
	SrcIp                 string  `json:"src_ip"`
	DstIp                 string  `json:"dst_ip"`
	SrcAsn                string  `json:"src_asn"`
	DstAsn                string  `json:"dst_asn"`
	DstEth                string  `json:"dst_eth"`
	Hops                  uint8   `json:"hops"`
	PingLossRate          float64 `json:"ping_loss_rate"`
	PingMaxDelay          float64 `json:"ping_max_delay"`
	PingMinDelay          float64 `json:"ping_min_delay"`
	PingAvgDelay          float64 `json:"ping_avg_delay"`
	HttpCode              uint8   `json:"http_code"`
	HttpDownloadSpeed     float64 `json:"http_download_speed"`
	HttpTimeConnect       float64 `json:"http_time_connect"`
	HttpTimeNameLookUp    float64 `json:"http_time_name_lookup"`
	HttpTimeStartTransfer float64 `json:"http_time_start_transfer"`
	HttpTimeTimeRedirect  float64 `json:"http_time_time_redirect"`
	HttpTimeTotal         float64 `json:"http_time_total"`
	BandwidthLimit        uint8   `json:"bandwidth_limit"`
	EthSendErrRate        float64 `json:"eth_send_err_rate"`
	EthSendDropRate       float64 `json:"eth_send_drop_rate"`
	HostRetransRate       float64 `json:"host_retrans_rate"`
	UdpOutSuccRate        float64 `json:"udp_out_succ_rate"`
	UdpOutAvgDelay        float64 `json:"udp_out_avg_delay"`
	TcpOutConnectTime     float64 `json:"tcp_out_connect_time"`
	TcpOutSuccRate        float64 `json:"tcp_out_succ_rate"`
	TcpOutAvgRate         float64 `json:"tcp_out_avg_rate"`
	Mtr                   string  `json:"mtr"`
}

func Insert(rows []Row) error {
	ctx := context.Background()
	batch, err := ck.DB.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO ipaas.detect(%s)", strings.Join(COLUMNS, ",")))
	if err != nil {
		return err
	}

	for _, row := range rows {
		if err = batch.Append(
			time.Unix(row.T, 0),
			row.SrcMachineId,
			row.SrcIp,
			row.SrcAsn,
			row.DstMachineId,
			row.DstIp,
			row.DstAsn,
			row.DstEth,
			row.PingLossRate,
			row.PingMaxDelay,
			row.PingMinDelay,
			row.PingAvgDelay,
			row.HttpCode,
			row.HttpDownloadSpeed,
			row.HttpTimeConnect,
			row.HttpTimeNameLookUp,
			row.HttpTimeStartTransfer,
			row.HttpTimeTimeRedirect,
			row.HttpTimeTotal,
			row.UdpOutSuccRate,
			row.UdpOutAvgDelay,
			row.TcpOutConnectTime,
			row.TcpOutSuccRate,
			row.TcpOutAvgRate,
			row.HostRetransRate,
			row.EthSendErrRate,
			row.EthSendDropRate,
			row.BandwidthLimit,
			row.Hops,
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
