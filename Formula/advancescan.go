package main

// Advancescan - Rom manager for AdvanceMAME/MESS
// Homepage: https://www.advancemame.it/scan-readme.html

import (
	"fmt"
	
	"os/exec"
)

func installAdvancescan() {
	// Método 1: Descargar y extraer .tar.gz
	advancescan_tar_url := "https://github.com/amadvance/advancescan/releases/download/v1.18/advancescan-1.18.tar.gz"
	advancescan_cmd_tar := exec.Command("curl", "-L", advancescan_tar_url, "-o", "package.tar.gz")
	err := advancescan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	advancescan_zip_url := "https://github.com/amadvance/advancescan/releases/download/v1.18/advancescan-1.18.zip"
	advancescan_cmd_zip := exec.Command("curl", "-L", advancescan_zip_url, "-o", "package.zip")
	err = advancescan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	advancescan_bin_url := "https://github.com/amadvance/advancescan/releases/download/v1.18/advancescan-1.18.bin"
	advancescan_cmd_bin := exec.Command("curl", "-L", advancescan_bin_url, "-o", "binary.bin")
	err = advancescan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	advancescan_src_url := "https://github.com/amadvance/advancescan/releases/download/v1.18/advancescan-1.18.src.tar.gz"
	advancescan_cmd_src := exec.Command("curl", "-L", advancescan_src_url, "-o", "source.tar.gz")
	err = advancescan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	advancescan_cmd_direct := exec.Command("./binary")
	err = advancescan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
