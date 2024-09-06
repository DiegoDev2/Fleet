package main

// Bandwhich - Terminal bandwidth utilization tool
// Homepage: https://github.com/imsnif/bandwhich

import (
	"fmt"
	
	"os/exec"
)

func installBandwhich() {
	// Método 1: Descargar y extraer .tar.gz
	bandwhich_tar_url := "https://github.com/imsnif/bandwhich/archive/refs/tags/v0.23.0.tar.gz"
	bandwhich_cmd_tar := exec.Command("curl", "-L", bandwhich_tar_url, "-o", "package.tar.gz")
	err := bandwhich_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bandwhich_zip_url := "https://github.com/imsnif/bandwhich/archive/refs/tags/v0.23.0.zip"
	bandwhich_cmd_zip := exec.Command("curl", "-L", bandwhich_zip_url, "-o", "package.zip")
	err = bandwhich_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bandwhich_bin_url := "https://github.com/imsnif/bandwhich/archive/refs/tags/v0.23.0.bin"
	bandwhich_cmd_bin := exec.Command("curl", "-L", bandwhich_bin_url, "-o", "binary.bin")
	err = bandwhich_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bandwhich_src_url := "https://github.com/imsnif/bandwhich/archive/refs/tags/v0.23.0.src.tar.gz"
	bandwhich_cmd_src := exec.Command("curl", "-L", bandwhich_src_url, "-o", "source.tar.gz")
	err = bandwhich_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bandwhich_cmd_direct := exec.Command("./binary")
	err = bandwhich_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
