package main

// Hss - Interactive parallel SSH client
// Homepage: https://github.com/six-ddc/hss

import (
	"fmt"
	
	"os/exec"
)

func installHss() {
	// Método 1: Descargar y extraer .tar.gz
	hss_tar_url := "https://github.com/six-ddc/hss/archive/refs/tags/1.9.tar.gz"
	hss_cmd_tar := exec.Command("curl", "-L", hss_tar_url, "-o", "package.tar.gz")
	err := hss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hss_zip_url := "https://github.com/six-ddc/hss/archive/refs/tags/1.9.zip"
	hss_cmd_zip := exec.Command("curl", "-L", hss_zip_url, "-o", "package.zip")
	err = hss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hss_bin_url := "https://github.com/six-ddc/hss/archive/refs/tags/1.9.bin"
	hss_cmd_bin := exec.Command("curl", "-L", hss_bin_url, "-o", "binary.bin")
	err = hss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hss_src_url := "https://github.com/six-ddc/hss/archive/refs/tags/1.9.src.tar.gz"
	hss_cmd_src := exec.Command("curl", "-L", hss_src_url, "-o", "source.tar.gz")
	err = hss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hss_cmd_direct := exec.Command("./binary")
	err = hss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
