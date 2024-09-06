package main

// Wcurl - Wrapper around curl to easily download files
// Homepage: https://github.com/curl/wcurl

import (
	"fmt"
	
	"os/exec"
)

func installWcurl() {
	// Método 1: Descargar y extraer .tar.gz
	wcurl_tar_url := "https://github.com/curl/wcurl/archive/refs/tags/2024.07.10.tar.gz"
	wcurl_cmd_tar := exec.Command("curl", "-L", wcurl_tar_url, "-o", "package.tar.gz")
	err := wcurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wcurl_zip_url := "https://github.com/curl/wcurl/archive/refs/tags/2024.07.10.zip"
	wcurl_cmd_zip := exec.Command("curl", "-L", wcurl_zip_url, "-o", "package.zip")
	err = wcurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wcurl_bin_url := "https://github.com/curl/wcurl/archive/refs/tags/2024.07.10.bin"
	wcurl_cmd_bin := exec.Command("curl", "-L", wcurl_bin_url, "-o", "binary.bin")
	err = wcurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wcurl_src_url := "https://github.com/curl/wcurl/archive/refs/tags/2024.07.10.src.tar.gz"
	wcurl_cmd_src := exec.Command("curl", "-L", wcurl_src_url, "-o", "source.tar.gz")
	err = wcurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wcurl_cmd_direct := exec.Command("./binary")
	err = wcurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
}
