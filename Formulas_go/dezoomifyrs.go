package main

// DezoomifyRs - Tiled image downloader
// Homepage: https://github.com/lovasoa/dezoomify-rs

import (
	"fmt"
	
	"os/exec"
)

func installDezoomifyRs() {
	// Método 1: Descargar y extraer .tar.gz
	dezoomifyrs_tar_url := "https://github.com/lovasoa/dezoomify-rs/archive/refs/tags/v2.12.4.tar.gz"
	dezoomifyrs_cmd_tar := exec.Command("curl", "-L", dezoomifyrs_tar_url, "-o", "package.tar.gz")
	err := dezoomifyrs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dezoomifyrs_zip_url := "https://github.com/lovasoa/dezoomify-rs/archive/refs/tags/v2.12.4.zip"
	dezoomifyrs_cmd_zip := exec.Command("curl", "-L", dezoomifyrs_zip_url, "-o", "package.zip")
	err = dezoomifyrs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dezoomifyrs_bin_url := "https://github.com/lovasoa/dezoomify-rs/archive/refs/tags/v2.12.4.bin"
	dezoomifyrs_cmd_bin := exec.Command("curl", "-L", dezoomifyrs_bin_url, "-o", "binary.bin")
	err = dezoomifyrs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dezoomifyrs_src_url := "https://github.com/lovasoa/dezoomify-rs/archive/refs/tags/v2.12.4.src.tar.gz"
	dezoomifyrs_cmd_src := exec.Command("curl", "-L", dezoomifyrs_src_url, "-o", "source.tar.gz")
	err = dezoomifyrs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dezoomifyrs_cmd_direct := exec.Command("./binary")
	err = dezoomifyrs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
}
