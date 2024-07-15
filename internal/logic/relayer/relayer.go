package relayer

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	v1 "relayer_ton/api/relayer/v1"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

type sRelayer struct {
	cli    *liteclient.ConnectionPool
	wallet *wallet.Wallet
}

func NewRelayer() *sRelayer {
	client := liteclient.NewConnectionPool()

	err := client.AddConnection(context.Background(), "65.21.238.183:16351", "Mf/JGvcWAvcrN3oheze8RF/ps6p7oL6ifrIzFmGQFQ8=")
	if err != nil {
		panic(err)
	}
	api := ton.NewAPIClient(client).WithRetry()
	w, err := openWallet(api)
	if err != nil {
		panic(err)
	}

	return &sRelayer{
		cli:    client,
		wallet: w,
	}
}

func (s *sRelayer) Deploy(ctx context.Context, req *v1.DeployReq) (*v1.DeployRes, error) {
	msgBody := cell.BeginCell().EndCell()
	addr, tx, block, err := s.wallet.DeployContractWaitTransaction(ctx,
		tlb.MustFromTON("0.02"),
		msgBody, getCode(), getData(req.Address))
	if err != nil {
		g.Log().Warning(ctx, "tx:", tx, "block:", block)
	}
	return &v1.DeployRes{
		Address: addr.String(),
	}, nil
}
func openWallet(api ton.APIClientWrapped) (*wallet.Wallet, error) {
	words := strings.Split("august jungle large onion disorder chase measure soft great history bid ceiling embark utility magnet response fade copper once setup need vivid above between", " ")

	w, err := wallet.FromSeed(api, words, wallet.V4R2)
	if err != nil {
		return nil, err
	}
	return w, nil
}
func getCode() *cell.Cell {
	var hexBOC = "B5EE9C7241021601000304000114FF00F4A413F4BCF2C80B010201200203020148040504F8F2810208D71820D31FD31FD31F02F823BBF264ED44D0D31FD31FD3FFF404D15143BAF2A15151BAF2A205F9015006F00123BAF2A3F80024A4C8CB1F5240CB1F5230CBFF5210F400C9ED54F80F01D30721C0009F6C519320D74A96D307D402FB00E830E021C001E30021C002E30001C0039130E30D03A4C8CB1F12CB1F121314150202CF06070201200A0B02E5007434C0C85C6C2497C13808B5D270482497C13800B4C7C860841C1B1D59EF48A084191CDD1CAF6C2497C17800FE900C083E9100723281F2FFF2743B513420405035C87D010C172040423D029BE84C6CE497C1F80174CFF20960841C1B1D59EEA48E0C38C340E084191CDD1CAEA497C1B8C3600809005934FFF4FFF4C1CC083086A4A9B977967E449BE85BE94CC8ACE456F80CB232FFF2FFC83F880C1C7E4100E127EC20007201FA00F40430F8276F2230500AA121BEF2E0508210F06C75677080185004CB0526CF1658FA0219F400CB6917CB1F5260CB3F20C98040FB000600845004810108F45930ED44D0810140D720C801CF16F400C9ED540172B08E208210E47374727080185005CB055003CF1623FA0213CB6ACB1FCB3FC98040FB00925F03E20201200C0D0059BD242B6F6A2684080A06B90FA0218470D4080847A4937D29910CE6903E9FF9837812801B7810148987159F31840201580E0F0011B8C97ED44D0D70B1F8003DB29DFB513420405035C87D010C00B23281F2FFF274006040423D029BE84C6002012010110019ADCE76A26840206B90EB85FFC00019AF1DF6A26840106B90EB858FC0006ED207FA00D4D422F90005C8CA0715CBFFC9D077748018C8CB05CB0222CF165005FA0214CB6B12CCCCC973FB00C84014810108F451F2A7020070810108D718FA00D33FC8542047810108F451F2A782106E6F746570748018C8CB05CB025006CF165004FA0214CB6A12CB1FCB3FC973FB0002006C810108D718FA00D33F305224810108F459F2A782106473747270748018C8CB05CB025005CF165003FA0213CB6ACB1F12CB3FC973FB00000ECBFFF400C9ED5405F7415B"
	codeCellBytes, _ := hex.DecodeString(hexBOC)

	codeCell, err := cell.FromBOC(codeCellBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(common.Bytes2Hex(codeCellBytes))

	return codeCell
}
func getData(addr string) *cell.Cell {
	data := cell.BeginCell().
		MustStoreBigUInt(big.NewInt(0), 32).                                 //chainid
		MustStoreBigUInt(big.NewInt(100), 32).                               //min fee
		MustStoreBigUInt(big.NewInt(0).SetBytes(common.FromHex(addr)), 256). //verify addr
		MustStoreDict(nil).
		EndCell()

	return data
}
