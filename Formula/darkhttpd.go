package main

// Darkhttpd - Small static webserver without CGI
// Homepage: https://unix4lyfe.org/darkhttpd/

import (
	"fmt"
	
	"os/exec"
)

func installDarkhttpd() {
	// Método 1: Descargar y extraer .tar.gz
	darkhttpd_tar_url := "https://github.com/emikulic/darkhttpd/archive/refs/tags/v1.16.tar.gz"
	darkhttpd_cmd_tar := exec.Command("curl", "-L", darkhttpd_tar_url, "-o", "package.tar.gz")
	err := darkhttpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	darkhttpd_zip_url := "https://github.com/emikulic/darkhttpd/archive/refs/tags/v1.16.zip"
	darkhttpd_cmd_zip := exec.Command("curl", "-L", darkhttpd_zip_url, "-o", "package.zip")
	err = darkhttpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	darkhttpd_bin_url := "https://github.com/emikulic/darkhttpd/archive/refs/tags/v1.16.bin"
	darkhttpd_cmd_bin := exec.Command("curl", "-L", darkhttpd_bin_url, "-o", "binary.bin")
	err = darkhttpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	darkhttpd_src_url := "https://github.com/emikulic/darkhttpd/archive/refs/tags/v1.16.src.tar.gz"
	darkhttpd_cmd_src := exec.Command("curl", "-L", darkhttpd_src_url, "-o", "source.tar.gz")
	err = darkhttpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	darkhttpd_cmd_direct := exec.Command("./binary")
	err = darkhttpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
