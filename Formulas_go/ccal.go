package main

// Ccal - Create Chinese calendars for print or browsing
// Homepage: https://ccal.chinesebay.com/ccal/ccal.htm

import (
	"fmt"
	
	"os/exec"
)

func installCcal() {
	// Método 1: Descargar y extraer .tar.gz
	ccal_tar_url := "https://ccal.chinesebay.com/ccal/ccal-2.5.3.tar.gz"
	ccal_cmd_tar := exec.Command("curl", "-L", ccal_tar_url, "-o", "package.tar.gz")
	err := ccal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccal_zip_url := "https://ccal.chinesebay.com/ccal/ccal-2.5.3.zip"
	ccal_cmd_zip := exec.Command("curl", "-L", ccal_zip_url, "-o", "package.zip")
	err = ccal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccal_bin_url := "https://ccal.chinesebay.com/ccal/ccal-2.5.3.bin"
	ccal_cmd_bin := exec.Command("curl", "-L", ccal_bin_url, "-o", "binary.bin")
	err = ccal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccal_src_url := "https://ccal.chinesebay.com/ccal/ccal-2.5.3.src.tar.gz"
	ccal_cmd_src := exec.Command("curl", "-L", ccal_src_url, "-o", "source.tar.gz")
	err = ccal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccal_cmd_direct := exec.Command("./binary")
	err = ccal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
