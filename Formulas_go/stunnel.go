package main

// Stunnel - SSL tunneling program
// Homepage: https://www.stunnel.org/

import (
	"fmt"
	
	"os/exec"
)

func installStunnel() {
	// Método 1: Descargar y extraer .tar.gz
	stunnel_tar_url := "https://www.stunnel.org/downloads/stunnel-5.72.tar.gz"
	stunnel_cmd_tar := exec.Command("curl", "-L", stunnel_tar_url, "-o", "package.tar.gz")
	err := stunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stunnel_zip_url := "https://www.stunnel.org/downloads/stunnel-5.72.zip"
	stunnel_cmd_zip := exec.Command("curl", "-L", stunnel_zip_url, "-o", "package.zip")
	err = stunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stunnel_bin_url := "https://www.stunnel.org/downloads/stunnel-5.72.bin"
	stunnel_cmd_bin := exec.Command("curl", "-L", stunnel_bin_url, "-o", "binary.bin")
	err = stunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stunnel_src_url := "https://www.stunnel.org/downloads/stunnel-5.72.src.tar.gz"
	stunnel_cmd_src := exec.Command("curl", "-L", stunnel_src_url, "-o", "source.tar.gz")
	err = stunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stunnel_cmd_direct := exec.Command("./binary")
	err = stunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
