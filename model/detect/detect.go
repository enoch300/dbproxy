package detect

import (
	"context"
	"dbproxy/db/ck"
	"fmt"
	"strings"
	"time"
)

var COLUMNS = []string{
	"t",
	"src_machine_id",
	"src_ip",
	"src_asn",
	"dst_machine_id",
	"dst_ip",
	"dst_asn",
	"dst_eth",
	"dst_tcp_port",
	"dst_udp_port",
	"dst_http_port",
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
	T                     int64   `json:"t"`                        //开始探测时间, 时间戳
	SrcMachineId          string  `json:"src_machine_id"`           //源节点 id
	DstMachineId          string  `json:"dst_machine_id"`           //目标节点 id
	SrcIp                 string  `json:"src_ip"`                   //源 ip
	DstIp                 string  `json:"dst_ip"`                   //目标 ip
	SrcAsn                string  `json:"src_asn"`                  //源 asn号
	DstAsn                string  `json:"dst_asn"`                  //目标 asn号
	DstEth                string  `json:"dst_eth"`                  //目标 网卡名
	DstTcpPort            string  `json:"dst_tcp_port"`             //目标 tcp 监听ip
	DstUdpPort            string  `json:"dst_udp_port"`             //目标 tcp 监听端口
	DstHttpPort           string  `json:"dst_http_port"`            //目标 http监听端口
	PingLossRate          float64 `json:"ping_loss_rate"`           //ping 丢包率
	PingMaxDelay          float64 `json:"ping_max_delay"`           //ping 最大延时
	PingMinDelay          float64 `json:"ping_min_delay"`           //ping 最小延时
	PingAvgDelay          float64 `json:"ping_avg_delay"`           //ping 平均延时
	HttpCode              uint8   `json:"http_code"`                //http 状态码
	HttpDownloadSpeed     float64 `json:"http_download_speed"`      //http 下载速率
	HttpTimeConnect       float64 `json:"http_time_connect"`        //http 建连耗时
	HttpTimeNameLookUp    float64 `json:"http_time_name_lookup"`    //http 域名解析耗时
	HttpTimeStartTransfer float64 `json:"http_time_start_transfer"` //http 首包时间
	HttpTimeTimeRedirect  float64 `json:"http_time_time_redirect"`  //http 重定向时间
	HttpTimeTotal         float64 `json:"http_time_total"`          //http 总耗时
	UdpOutSuccRate        float64 `json:"udp_out_succ_rate"`        //udp 成功率
	UdpOutAvgDelay        float64 `json:"udp_out_avg_delay"`        //udp 平均延时
	TcpOutConnectTime     float64 `json:"tcp_out_connect_time"`     //tcp 建连耗时
	TcpOutSuccRate        float64 `json:"tcp_out_succ_rate"`        //tcp 成功率
	TcpOutAvgDelay        float64 `json:"tcp_out_avg_delay"`        //tcp 平均延时
	EthSendErrRate        float64 `json:"eth_send_err_rate"`        //源 网卡发包错误率
	EthSendDropRate       float64 `json:"eth_send_drop_rate"`       //源 网卡发包丢包率
	BandwidthLimit        uint8   `json:"bandwidth_limit"`          //源 网卡是否限速
	HostRetransRate       float64 `json:"host_retrans_rate"`        //源 整机重传率
	Hops                  uint8   `json:"hops"`                     //源到目标 路由跳数
	Mtr                   string  `json:"mtr"`                      //源到目标 路由详情
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
			row.DstTcpPort,
			row.DstUdpPort,
			row.DstHttpPort,
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
			row.TcpOutAvgDelay,
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
