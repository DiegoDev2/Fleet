package main

// Xsv - Fast CSV toolkit written in Rust
// Homepage: https://github.com/BurntSushi/xsv

import (
	"fmt"
	
	"os/exec"
)

func installXsv() {
	// Método 1: Descargar y extraer .tar.gz
	xsv_tar_url := "https://github.com/BurntSushi/xsv/archive/refs/tags/0.13.0.tar.gz"
	xsv_cmd_tar := exec.Command("curl", "-L", xsv_tar_url, "-o", "package.tar.gz")
	err := xsv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xsv_zip_url := "https://github.com/BurntSushi/xsv/archive/refs/tags/0.13.0.zip"
	xsv_cmd_zip := exec.Command("curl", "-L", xsv_zip_url, "-o", "package.zip")
	err = xsv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xsv_bin_url := "https://github.com/BurntSushi/xsv/archive/refs/tags/0.13.0.bin"
	xsv_cmd_bin := exec.Command("curl", "-L", xsv_bin_url, "-o", "binary.bin")
	err = xsv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xsv_src_url := "https://github.com/BurntSushi/xsv/archive/refs/tags/0.13.0.src.tar.gz"
	xsv_cmd_src := exec.Command("curl", "-L", xsv_src_url, "-o", "source.tar.gz")
	err = xsv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xsv_cmd_direct := exec.Command("./binary")
	err = xsv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
