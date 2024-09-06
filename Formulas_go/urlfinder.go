package main

// Urlfinder - Extracting URLs and subdomains from JS files on a website
// Homepage: https://github.com/pingc0y/URLFinder

import (
	"fmt"
	
	"os/exec"
)

func installUrlfinder() {
	// Método 1: Descargar y extraer .tar.gz
	urlfinder_tar_url := "https://github.com/pingc0y/URLFinder/archive/refs/tags/2023.9.9.tar.gz"
	urlfinder_cmd_tar := exec.Command("curl", "-L", urlfinder_tar_url, "-o", "package.tar.gz")
	err := urlfinder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	urlfinder_zip_url := "https://github.com/pingc0y/URLFinder/archive/refs/tags/2023.9.9.zip"
	urlfinder_cmd_zip := exec.Command("curl", "-L", urlfinder_zip_url, "-o", "package.zip")
	err = urlfinder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	urlfinder_bin_url := "https://github.com/pingc0y/URLFinder/archive/refs/tags/2023.9.9.bin"
	urlfinder_cmd_bin := exec.Command("curl", "-L", urlfinder_bin_url, "-o", "binary.bin")
	err = urlfinder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	urlfinder_src_url := "https://github.com/pingc0y/URLFinder/archive/refs/tags/2023.9.9.src.tar.gz"
	urlfinder_cmd_src := exec.Command("curl", "-L", urlfinder_src_url, "-o", "source.tar.gz")
	err = urlfinder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	urlfinder_cmd_direct := exec.Command("./binary")
	err = urlfinder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
