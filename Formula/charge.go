package main

// Charge - Opinionated, zero-config static site generator
// Homepage: https://charge.js.org

import (
	"fmt"
	
	"os/exec"
)

func installCharge() {
	// Método 1: Descargar y extraer .tar.gz
	charge_tar_url := "https://registry.npmjs.org/@static/charge/-/charge-1.7.0.tgz"
	charge_cmd_tar := exec.Command("curl", "-L", charge_tar_url, "-o", "package.tar.gz")
	err := charge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	charge_zip_url := "https://registry.npmjs.org/@static/charge/-/charge-1.7.0.tgz"
	charge_cmd_zip := exec.Command("curl", "-L", charge_zip_url, "-o", "package.zip")
	err = charge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	charge_bin_url := "https://registry.npmjs.org/@static/charge/-/charge-1.7.0.tgz"
	charge_cmd_bin := exec.Command("curl", "-L", charge_bin_url, "-o", "binary.bin")
	err = charge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	charge_src_url := "https://registry.npmjs.org/@static/charge/-/charge-1.7.0.tgz"
	charge_cmd_src := exec.Command("curl", "-L", charge_src_url, "-o", "source.tar.gz")
	err = charge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	charge_cmd_direct := exec.Command("./binary")
	err = charge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
