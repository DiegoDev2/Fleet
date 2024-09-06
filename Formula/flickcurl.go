package main

// Flickcurl - Library for the Flickr API
// Homepage: https://librdf.org/flickcurl/

import (
	"fmt"
	
	"os/exec"
)

func installFlickcurl() {
	// Método 1: Descargar y extraer .tar.gz
	flickcurl_tar_url := "https://download.dajobe.org/flickcurl/flickcurl-1.26.tar.gz"
	flickcurl_cmd_tar := exec.Command("curl", "-L", flickcurl_tar_url, "-o", "package.tar.gz")
	err := flickcurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flickcurl_zip_url := "https://download.dajobe.org/flickcurl/flickcurl-1.26.zip"
	flickcurl_cmd_zip := exec.Command("curl", "-L", flickcurl_zip_url, "-o", "package.zip")
	err = flickcurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flickcurl_bin_url := "https://download.dajobe.org/flickcurl/flickcurl-1.26.bin"
	flickcurl_cmd_bin := exec.Command("curl", "-L", flickcurl_bin_url, "-o", "binary.bin")
	err = flickcurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flickcurl_src_url := "https://download.dajobe.org/flickcurl/flickcurl-1.26.src.tar.gz"
	flickcurl_cmd_src := exec.Command("curl", "-L", flickcurl_src_url, "-o", "source.tar.gz")
	err = flickcurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flickcurl_cmd_direct := exec.Command("./binary")
	err = flickcurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
