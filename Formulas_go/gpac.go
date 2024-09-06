package main

// Gpac - Multimedia framework for research and academic purposes
// Homepage: https://gpac.wp.mines-telecom.fr/

import (
	"fmt"
	
	"os/exec"
)

func installGpac() {
	// Método 1: Descargar y extraer .tar.gz
	gpac_tar_url := "https://github.com/gpac/gpac/archive/refs/tags/v2.4.0.tar.gz"
	gpac_cmd_tar := exec.Command("curl", "-L", gpac_tar_url, "-o", "package.tar.gz")
	err := gpac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpac_zip_url := "https://github.com/gpac/gpac/archive/refs/tags/v2.4.0.zip"
	gpac_cmd_zip := exec.Command("curl", "-L", gpac_zip_url, "-o", "package.zip")
	err = gpac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpac_bin_url := "https://github.com/gpac/gpac/archive/refs/tags/v2.4.0.bin"
	gpac_cmd_bin := exec.Command("curl", "-L", gpac_bin_url, "-o", "binary.bin")
	err = gpac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpac_src_url := "https://github.com/gpac/gpac/archive/refs/tags/v2.4.0.src.tar.gz"
	gpac_cmd_src := exec.Command("curl", "-L", gpac_src_url, "-o", "source.tar.gz")
	err = gpac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpac_cmd_direct := exec.Command("./binary")
	err = gpac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
