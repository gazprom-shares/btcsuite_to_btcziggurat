package pfcregtest_spawn

import (
	"github.com/jfixby/pin"
	"github.com/btcziggurat/btcsuite_to_btcziggurat/proc"
	"github.com/picfight/coin_knife/fileproc"
	"github.com/picfight/coin_knife/injector"
	"github.com/picfight/coin_knife/projectops"
	"path/filepath"
)

type Settings struct {
	PathToBtcsuiteBTCDRepo string
	PathToPicfightPFCDRepo string

	DoNotProcessAnyFiles bool
	FileContentProcessor fileproc.StringProcessor
	FileNameProcessor    fileproc.StringProcessor
	IsFileProcessable    fileproc.FileToProcessDetector
}

func Build() {

	set := &Settings{
		PathToBtcsuiteBTCDRepo: `D:\PICFIGHT\src\github.com\jfixby\btcregtest`,
		PathToPicfightPFCDRepo: `D:\PICFIGHT\src\github.com\picfight\pfcregtest`,
		DoNotProcessAnyFiles:   false,
		FileNameProcessor:      proc.PicfightCoinFileNameGenerator,
		IsFileProcessable:      proc.ProcessableFiles,
		FileContentProcessor:   proc.PicfightCoinFileGenerator,
	}

	pin.D(" Input", set.PathToBtcsuiteBTCDRepo)
	pin.D("Output", set.PathToPicfightPFCDRepo)
	pin.D("")

	projectops.ClearProject(set.PathToPicfightPFCDRepo)

	fileproc.TransferFiles(
		set.IsFileProcessable,
		set.FileNameProcessor,
		set.DoNotProcessAnyFiles,
		set.FileContentProcessor,
		set.PathToBtcsuiteBTCDRepo,
		set.PathToPicfightPFCDRepo,
	)

	injector.PerformInjections(set.PathToPicfightPFCDRepo, filepath.Join("", "code_injections", "regtest"))

	//FixNetworkParams(set.PathToPfcdRepo, set)
	//FixFullBlockTests(set)

	projectops.AppendGitIgnore(set.PathToPicfightPFCDRepo)

	projectops.GoFmt(set.PathToPicfightPFCDRepo)

}
