package main

// Juise - JUNOS user interface scripting environment
// Homepage: https://github.com/Juniper/juise/wiki

import (
	"fmt"
	
	"os/exec"
)

func installJuise() {
	// Método 1: Descargar y extraer .tar.gz
	juise_tar_url := "https://github.com/Juniper/juise/releases/download/0.9.0/juise-0.9.0.tar.gz"
	juise_cmd_tar := exec.Command("curl", "-L", juise_tar_url, "-o", "package.tar.gz")
	err := juise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	juise_zip_url := "https://github.com/Juniper/juise/releases/download/0.9.0/juise-0.9.0.zip"
	juise_cmd_zip := exec.Command("curl", "-L", juise_zip_url, "-o", "package.zip")
	err = juise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	juise_bin_url := "https://github.com/Juniper/juise/releases/download/0.9.0/juise-0.9.0.bin"
	juise_cmd_bin := exec.Command("curl", "-L", juise_bin_url, "-o", "binary.bin")
	err = juise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	juise_src_url := "https://github.com/Juniper/juise/releases/download/0.9.0/juise-0.9.0.src.tar.gz"
	juise_cmd_src := exec.Command("curl", "-L", juise_src_url, "-o", "source.tar.gz")
	err = juise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	juise_cmd_direct := exec.Command("./binary")
	err = juise_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libslax")
exec.Command("latte", "install", "libslax")
}
