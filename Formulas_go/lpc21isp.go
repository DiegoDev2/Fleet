package main

// Lpc21isp - In-circuit programming (ISP) tool for several NXP microcontrollers
// Homepage: https://lpc21isp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLpc21isp() {
	// Método 1: Descargar y extraer .tar.gz
	lpc21isp_tar_url := "https://downloads.sourceforge.net/project/lpc21isp/lpc21isp/1.97/lpc21isp_197.tar.gz"
	lpc21isp_cmd_tar := exec.Command("curl", "-L", lpc21isp_tar_url, "-o", "package.tar.gz")
	err := lpc21isp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lpc21isp_zip_url := "https://downloads.sourceforge.net/project/lpc21isp/lpc21isp/1.97/lpc21isp_197.zip"
	lpc21isp_cmd_zip := exec.Command("curl", "-L", lpc21isp_zip_url, "-o", "package.zip")
	err = lpc21isp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lpc21isp_bin_url := "https://downloads.sourceforge.net/project/lpc21isp/lpc21isp/1.97/lpc21isp_197.bin"
	lpc21isp_cmd_bin := exec.Command("curl", "-L", lpc21isp_bin_url, "-o", "binary.bin")
	err = lpc21isp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lpc21isp_src_url := "https://downloads.sourceforge.net/project/lpc21isp/lpc21isp/1.97/lpc21isp_197.src.tar.gz"
	lpc21isp_cmd_src := exec.Command("curl", "-L", lpc21isp_src_url, "-o", "source.tar.gz")
	err = lpc21isp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lpc21isp_cmd_direct := exec.Command("./binary")
	err = lpc21isp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
