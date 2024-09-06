package main

// ColorCode - Free advanced MasterMind clone
// Homepage: http://colorcode.laebisch.com/

import (
	"fmt"
	
	"os/exec"
)

func installColorCode() {
	// Método 1: Descargar y extraer .tar.gz
	colorcode_tar_url := "http://colorcode.laebisch.com/download/ColorCode-0.8.7.tar.gz"
	colorcode_cmd_tar := exec.Command("curl", "-L", colorcode_tar_url, "-o", "package.tar.gz")
	err := colorcode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	colorcode_zip_url := "http://colorcode.laebisch.com/download/ColorCode-0.8.7.zip"
	colorcode_cmd_zip := exec.Command("curl", "-L", colorcode_zip_url, "-o", "package.zip")
	err = colorcode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	colorcode_bin_url := "http://colorcode.laebisch.com/download/ColorCode-0.8.7.bin"
	colorcode_cmd_bin := exec.Command("curl", "-L", colorcode_bin_url, "-o", "binary.bin")
	err = colorcode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	colorcode_src_url := "http://colorcode.laebisch.com/download/ColorCode-0.8.7.src.tar.gz"
	colorcode_cmd_src := exec.Command("curl", "-L", colorcode_src_url, "-o", "source.tar.gz")
	err = colorcode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	colorcode_cmd_direct := exec.Command("./binary")
	err = colorcode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
}
