package main

// Kcgi - Minimal CGI and FastCGI library for C/C++
// Homepage: https://kristaps.bsd.lv/kcgi/

import (
	"fmt"
	
	"os/exec"
)

func installKcgi() {
	// Método 1: Descargar y extraer .tar.gz
	kcgi_tar_url := "https://kristaps.bsd.lv/kcgi/snapshots/kcgi-0.13.3.tgz"
	kcgi_cmd_tar := exec.Command("curl", "-L", kcgi_tar_url, "-o", "package.tar.gz")
	err := kcgi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kcgi_zip_url := "https://kristaps.bsd.lv/kcgi/snapshots/kcgi-0.13.3.tgz"
	kcgi_cmd_zip := exec.Command("curl", "-L", kcgi_zip_url, "-o", "package.zip")
	err = kcgi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kcgi_bin_url := "https://kristaps.bsd.lv/kcgi/snapshots/kcgi-0.13.3.tgz"
	kcgi_cmd_bin := exec.Command("curl", "-L", kcgi_bin_url, "-o", "binary.bin")
	err = kcgi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kcgi_src_url := "https://kristaps.bsd.lv/kcgi/snapshots/kcgi-0.13.3.tgz"
	kcgi_cmd_src := exec.Command("curl", "-L", kcgi_src_url, "-o", "source.tar.gz")
	err = kcgi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kcgi_cmd_direct := exec.Command("./binary")
	err = kcgi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bmake")
	exec.Command("latte", "install", "bmake").Run()
	fmt.Println("Instalando dependencia: libseccomp")
	exec.Command("latte", "install", "libseccomp").Run()
}
