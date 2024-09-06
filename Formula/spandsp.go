package main

// Spandsp - DSP functions library for telephony
// Homepage: https://www.soft-switch.org/

import (
	"fmt"
	
	"os/exec"
)

func installSpandsp() {
	// Método 1: Descargar y extraer .tar.gz
	spandsp_tar_url := "https://www.soft-switch.org/downloads/spandsp/spandsp-0.0.6.tar.gz"
	spandsp_cmd_tar := exec.Command("curl", "-L", spandsp_tar_url, "-o", "package.tar.gz")
	err := spandsp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spandsp_zip_url := "https://www.soft-switch.org/downloads/spandsp/spandsp-0.0.6.zip"
	spandsp_cmd_zip := exec.Command("curl", "-L", spandsp_zip_url, "-o", "package.zip")
	err = spandsp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spandsp_bin_url := "https://www.soft-switch.org/downloads/spandsp/spandsp-0.0.6.bin"
	spandsp_cmd_bin := exec.Command("curl", "-L", spandsp_bin_url, "-o", "binary.bin")
	err = spandsp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spandsp_src_url := "https://www.soft-switch.org/downloads/spandsp/spandsp-0.0.6.src.tar.gz"
	spandsp_cmd_src := exec.Command("curl", "-L", spandsp_src_url, "-o", "source.tar.gz")
	err = spandsp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spandsp_cmd_direct := exec.Command("./binary")
	err = spandsp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
}
