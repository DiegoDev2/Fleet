package main

// Calcurse - Text-based personal organizer
// Homepage: https://calcurse.org/

import (
	"fmt"
	
	"os/exec"
)

func installCalcurse() {
	// Método 1: Descargar y extraer .tar.gz
	calcurse_tar_url := "https://calcurse.org/files/calcurse-4.8.1.tar.gz"
	calcurse_cmd_tar := exec.Command("curl", "-L", calcurse_tar_url, "-o", "package.tar.gz")
	err := calcurse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	calcurse_zip_url := "https://calcurse.org/files/calcurse-4.8.1.zip"
	calcurse_cmd_zip := exec.Command("curl", "-L", calcurse_zip_url, "-o", "package.zip")
	err = calcurse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	calcurse_bin_url := "https://calcurse.org/files/calcurse-4.8.1.bin"
	calcurse_cmd_bin := exec.Command("curl", "-L", calcurse_bin_url, "-o", "binary.bin")
	err = calcurse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	calcurse_src_url := "https://calcurse.org/files/calcurse-4.8.1.src.tar.gz"
	calcurse_cmd_src := exec.Command("curl", "-L", calcurse_src_url, "-o", "source.tar.gz")
	err = calcurse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	calcurse_cmd_direct := exec.Command("./binary")
	err = calcurse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
