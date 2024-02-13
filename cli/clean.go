package cli

import (
	"github.com/elliotchance/pie/v2"
	"github.com/lazygophers/log"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean log",
	Long:  `clean log`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		files, err := os.ReadDir(filepath.Dir(filepath.Join(os.TempDir(), "lazygophers", "log") + "/"))
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}

		log.Info("files:", files)

		pie.Each(
			pie.DropTop(
				pie.SortUsing(
					pie.Map(
						files,
						func(file os.DirEntry) string {
							return file.Name()
						},
					),
					func(a, b string) bool {
						return a > b
					},
				),
				12,
			),
			func(s string) {
				log.Warnf("remove:%s", s)
				err = os.Remove(filepath.Join(filepath.Dir(filepath.Join(os.TempDir(), "lazygophers", "log")+"/"), s))
				if err != nil {
					log.Errorf("err:%v", err)
				}
			})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
