package main

// Mpdas - C++ client to submit tracks to audioscrobbler
// Homepage: https://www.50hz.ws/mpdas/

import (
	"fmt"
	
	"os/exec"
)

func installMpdas() {
	// Método 1: Descargar y extraer .tar.gz
	mpdas_tar_url := "https://www.50hz.ws/mpdas/mpdas-0.4.5.tar.gz"
	mpdas_cmd_tar := exec.Command("curl", "-L", mpdas_tar_url, "-o", "package.tar.gz")
	err := mpdas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpdas_zip_url := "https://www.50hz.ws/mpdas/mpdas-0.4.5.zip"
	mpdas_cmd_zip := exec.Command("curl", "-L", mpdas_zip_url, "-o", "package.zip")
	err = mpdas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpdas_bin_url := "https://www.50hz.ws/mpdas/mpdas-0.4.5.bin"
	mpdas_cmd_bin := exec.Command("curl", "-L", mpdas_bin_url, "-o", "binary.bin")
	err = mpdas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpdas_src_url := "https://www.50hz.ws/mpdas/mpdas-0.4.5.src.tar.gz"
	mpdas_cmd_src := exec.Command("curl", "-L", mpdas_src_url, "-o", "source.tar.gz")
	err = mpdas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpdas_cmd_direct := exec.Command("./binary")
	err = mpdas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmpdclient")
exec.Command("latte", "install", "libmpdclient")
}
