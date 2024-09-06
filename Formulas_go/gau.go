package main

// Gau - Open Threat Exchange, Wayback Machine, and Common Crawl URL fetcher
// Homepage: https://github.com/lc/gau

import (
	"fmt"
	
	"os/exec"
)

func installGau() {
	// Método 1: Descargar y extraer .tar.gz
	gau_tar_url := "https://github.com/lc/gau/archive/refs/tags/v2.2.3.tar.gz"
	gau_cmd_tar := exec.Command("curl", "-L", gau_tar_url, "-o", "package.tar.gz")
	err := gau_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gau_zip_url := "https://github.com/lc/gau/archive/refs/tags/v2.2.3.zip"
	gau_cmd_zip := exec.Command("curl", "-L", gau_zip_url, "-o", "package.zip")
	err = gau_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gau_bin_url := "https://github.com/lc/gau/archive/refs/tags/v2.2.3.bin"
	gau_cmd_bin := exec.Command("curl", "-L", gau_bin_url, "-o", "binary.bin")
	err = gau_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gau_src_url := "https://github.com/lc/gau/archive/refs/tags/v2.2.3.src.tar.gz"
	gau_cmd_src := exec.Command("curl", "-L", gau_src_url, "-o", "source.tar.gz")
	err = gau_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gau_cmd_direct := exec.Command("./binary")
	err = gau_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
