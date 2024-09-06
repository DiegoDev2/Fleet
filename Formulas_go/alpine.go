package main

// Alpine - News and email agent
// Homepage: https://alpineapp.email

import (
	"fmt"
	
	"os/exec"
)

func installAlpine() {
	// Método 1: Descargar y extraer .tar.gz
	alpine_tar_url := "https://alpineapp.email/alpine/release/src/alpine-2.26.tar.xz"
	alpine_cmd_tar := exec.Command("curl", "-L", alpine_tar_url, "-o", "package.tar.gz")
	err := alpine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alpine_zip_url := "https://alpineapp.email/alpine/release/src/alpine-2.26.tar.xz"
	alpine_cmd_zip := exec.Command("curl", "-L", alpine_zip_url, "-o", "package.zip")
	err = alpine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alpine_bin_url := "https://alpineapp.email/alpine/release/src/alpine-2.26.tar.xz"
	alpine_cmd_bin := exec.Command("curl", "-L", alpine_bin_url, "-o", "binary.bin")
	err = alpine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alpine_src_url := "https://alpineapp.email/alpine/release/src/alpine-2.26.tar.xz"
	alpine_cmd_src := exec.Command("curl", "-L", alpine_src_url, "-o", "source.tar.gz")
	err = alpine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alpine_cmd_direct := exec.Command("./binary")
	err = alpine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
