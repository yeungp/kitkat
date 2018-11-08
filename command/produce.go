package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/waltzofpearls/kitkat/producer"
)

func init() {
	p := producer.New()
	cmd := &cobra.Command{
		Use:     "produce",
		Short:   "Producer mode. Writes messages to given kinesis stream(s)",
		Aliases: []string{"p"},
		Run:     produce(p),
	}
	rootCmd.AddCommand(cmd)
	cmd.Flags().StringVarP(&p.Stream, "stream", "t", "", "Kinesis stream name")
	cmd.Flags().StringVarP(&p.Region, "region", "r", "us-west-2", "AWS region")
	cmd.Flags().StringVarP(&p.PartitionKey, "key", "k", "", "Partition key or random key if left empty")
	cmd.Flags().BoolVarP(&p.Aggregated, "aggregated", "a", false, "Produce in KPL aggregated record format")
}

func produce(p *producer.Producer) runFunc {
	return func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().Changed("stream") {
			cmd.Help()
			os.Exit(1)
		}
		p.Verbose = verbose
		if err := p.Write(); err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	}
}
