package main

// Eiffelstudio - Development environment for the Eiffel language
// Homepage: https://www.eiffel.com

import (
	"fmt"
	
	"os/exec"
)

func installEiffelstudio() {
	// Método 1: Descargar y extraer .tar.gz
	eiffelstudio_tar_url := "https://ftp.eiffel.com/pub/download/22.05/pp/PorterPackage_std_106302.tar"
	eiffelstudio_cmd_tar := exec.Command("curl", "-L", eiffelstudio_tar_url, "-o", "package.tar.gz")
	err := eiffelstudio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eiffelstudio_zip_url := "https://ftp.eiffel.com/pub/download/22.05/pp/PorterPackage_std_106302.tar"
	eiffelstudio_cmd_zip := exec.Command("curl", "-L", eiffelstudio_zip_url, "-o", "package.zip")
	err = eiffelstudio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eiffelstudio_bin_url := "https://ftp.eiffel.com/pub/download/22.05/pp/PorterPackage_std_106302.tar"
	eiffelstudio_cmd_bin := exec.Command("curl", "-L", eiffelstudio_bin_url, "-o", "binary.bin")
	err = eiffelstudio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eiffelstudio_src_url := "https://ftp.eiffel.com/pub/download/22.05/pp/PorterPackage_std_106302.tar"
	eiffelstudio_cmd_src := exec.Command("curl", "-L", eiffelstudio_src_url, "-o", "source.tar.gz")
	err = eiffelstudio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eiffelstudio_cmd_direct := exec.Command("./binary")
	err = eiffelstudio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
}
