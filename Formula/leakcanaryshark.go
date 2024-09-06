package main

// LeakcanaryShark - CLI Java memory leak explorer for LeakCanary
// Homepage: https://square.github.io/leakcanary/shark/

import (
	"fmt"
	
	"os/exec"
)

func installLeakcanaryShark() {
	// Método 1: Descargar y extraer .tar.gz
	leakcanaryshark_tar_url := "https://github.com/square/leakcanary/releases/download/v2.14/shark-cli-2.14.zip"
	leakcanaryshark_cmd_tar := exec.Command("curl", "-L", leakcanaryshark_tar_url, "-o", "package.tar.gz")
	err := leakcanaryshark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leakcanaryshark_zip_url := "https://github.com/square/leakcanary/releases/download/v2.14/shark-cli-2.14.zip"
	leakcanaryshark_cmd_zip := exec.Command("curl", "-L", leakcanaryshark_zip_url, "-o", "package.zip")
	err = leakcanaryshark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leakcanaryshark_bin_url := "https://github.com/square/leakcanary/releases/download/v2.14/shark-cli-2.14.zip"
	leakcanaryshark_cmd_bin := exec.Command("curl", "-L", leakcanaryshark_bin_url, "-o", "binary.bin")
	err = leakcanaryshark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leakcanaryshark_src_url := "https://github.com/square/leakcanary/releases/download/v2.14/shark-cli-2.14.zip"
	leakcanaryshark_cmd_src := exec.Command("curl", "-L", leakcanaryshark_src_url, "-o", "source.tar.gz")
	err = leakcanaryshark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leakcanaryshark_cmd_direct := exec.Command("./binary")
	err = leakcanaryshark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
