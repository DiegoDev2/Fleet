package main

// Chrony - Versatile implementation of the Network Time Protocol (NTP)
// Homepage: https://chrony-project.org

import (
	"fmt"
	
	"os/exec"
)

func installChrony() {
	// Método 1: Descargar y extraer .tar.gz
	chrony_tar_url := "https://chrony-project.org/releases/chrony-4.6.tar.gz"
	chrony_cmd_tar := exec.Command("curl", "-L", chrony_tar_url, "-o", "package.tar.gz")
	err := chrony_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chrony_zip_url := "https://chrony-project.org/releases/chrony-4.6.zip"
	chrony_cmd_zip := exec.Command("curl", "-L", chrony_zip_url, "-o", "package.zip")
	err = chrony_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chrony_bin_url := "https://chrony-project.org/releases/chrony-4.6.bin"
	chrony_cmd_bin := exec.Command("curl", "-L", chrony_bin_url, "-o", "binary.bin")
	err = chrony_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chrony_src_url := "https://chrony-project.org/releases/chrony-4.6.src.tar.gz"
	chrony_cmd_src := exec.Command("curl", "-L", chrony_src_url, "-o", "source.tar.gz")
	err = chrony_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chrony_cmd_direct := exec.Command("./binary")
	err = chrony_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
}
