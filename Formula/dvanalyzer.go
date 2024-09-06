package main

// Dvanalyzer - Quality control tool for examining tape-to-file DV streams
// Homepage: https://mediaarea.net/DVAnalyzer

import (
	"fmt"
	
	"os/exec"
)

func installDvanalyzer() {
	// Método 1: Descargar y extraer .tar.gz
	dvanalyzer_tar_url := "https://mediaarea.net/download/binary/dvanalyzer/1.4.2/DVAnalyzer_CLI_1.4.2_GNU_FromSource.tar.bz2"
	dvanalyzer_cmd_tar := exec.Command("curl", "-L", dvanalyzer_tar_url, "-o", "package.tar.gz")
	err := dvanalyzer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvanalyzer_zip_url := "https://mediaarea.net/download/binary/dvanalyzer/1.4.2/DVAnalyzer_CLI_1.4.2_GNU_FromSource.tar.bz2"
	dvanalyzer_cmd_zip := exec.Command("curl", "-L", dvanalyzer_zip_url, "-o", "package.zip")
	err = dvanalyzer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvanalyzer_bin_url := "https://mediaarea.net/download/binary/dvanalyzer/1.4.2/DVAnalyzer_CLI_1.4.2_GNU_FromSource.tar.bz2"
	dvanalyzer_cmd_bin := exec.Command("curl", "-L", dvanalyzer_bin_url, "-o", "binary.bin")
	err = dvanalyzer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvanalyzer_src_url := "https://mediaarea.net/download/binary/dvanalyzer/1.4.2/DVAnalyzer_CLI_1.4.2_GNU_FromSource.tar.bz2"
	dvanalyzer_cmd_src := exec.Command("curl", "-L", dvanalyzer_src_url, "-o", "source.tar.gz")
	err = dvanalyzer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvanalyzer_cmd_direct := exec.Command("./binary")
	err = dvanalyzer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
