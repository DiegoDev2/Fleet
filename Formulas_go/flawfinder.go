package main

// Flawfinder - Examines code and reports possible security weaknesses
// Homepage: https://dwheeler.com/flawfinder/

import (
	"fmt"
	
	"os/exec"
)

func installFlawfinder() {
	// Método 1: Descargar y extraer .tar.gz
	flawfinder_tar_url := "https://dwheeler.com/flawfinder/flawfinder-2.0.19.tar.gz"
	flawfinder_cmd_tar := exec.Command("curl", "-L", flawfinder_tar_url, "-o", "package.tar.gz")
	err := flawfinder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flawfinder_zip_url := "https://dwheeler.com/flawfinder/flawfinder-2.0.19.zip"
	flawfinder_cmd_zip := exec.Command("curl", "-L", flawfinder_zip_url, "-o", "package.zip")
	err = flawfinder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flawfinder_bin_url := "https://dwheeler.com/flawfinder/flawfinder-2.0.19.bin"
	flawfinder_cmd_bin := exec.Command("curl", "-L", flawfinder_bin_url, "-o", "binary.bin")
	err = flawfinder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flawfinder_src_url := "https://dwheeler.com/flawfinder/flawfinder-2.0.19.src.tar.gz"
	flawfinder_cmd_src := exec.Command("curl", "-L", flawfinder_src_url, "-o", "source.tar.gz")
	err = flawfinder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flawfinder_cmd_direct := exec.Command("./binary")
	err = flawfinder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
