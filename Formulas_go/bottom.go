package main

// Bottom - Yet another cross-platform graphical process/system monitor
// Homepage: https://clementtsang.github.io/bottom/

import (
	"fmt"
	
	"os/exec"
)

func installBottom() {
	// Método 1: Descargar y extraer .tar.gz
	bottom_tar_url := "https://github.com/ClementTsang/bottom/archive/refs/tags/0.10.2.tar.gz"
	bottom_cmd_tar := exec.Command("curl", "-L", bottom_tar_url, "-o", "package.tar.gz")
	err := bottom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bottom_zip_url := "https://github.com/ClementTsang/bottom/archive/refs/tags/0.10.2.zip"
	bottom_cmd_zip := exec.Command("curl", "-L", bottom_zip_url, "-o", "package.zip")
	err = bottom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bottom_bin_url := "https://github.com/ClementTsang/bottom/archive/refs/tags/0.10.2.bin"
	bottom_cmd_bin := exec.Command("curl", "-L", bottom_bin_url, "-o", "binary.bin")
	err = bottom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bottom_src_url := "https://github.com/ClementTsang/bottom/archive/refs/tags/0.10.2.src.tar.gz"
	bottom_cmd_src := exec.Command("curl", "-L", bottom_src_url, "-o", "source.tar.gz")
	err = bottom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bottom_cmd_direct := exec.Command("./binary")
	err = bottom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
