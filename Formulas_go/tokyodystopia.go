package main

// TokyoDystopia - Lightweight full-text search system
// Homepage: https://dbmx.net/tokyodystopia/

import (
	"fmt"
	
	"os/exec"
)

func installTokyoDystopia() {
	// Método 1: Descargar y extraer .tar.gz
	tokyodystopia_tar_url := "https://dbmx.net/tokyodystopia/tokyodystopia-0.9.15.tar.gz"
	tokyodystopia_cmd_tar := exec.Command("curl", "-L", tokyodystopia_tar_url, "-o", "package.tar.gz")
	err := tokyodystopia_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tokyodystopia_zip_url := "https://dbmx.net/tokyodystopia/tokyodystopia-0.9.15.zip"
	tokyodystopia_cmd_zip := exec.Command("curl", "-L", tokyodystopia_zip_url, "-o", "package.zip")
	err = tokyodystopia_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tokyodystopia_bin_url := "https://dbmx.net/tokyodystopia/tokyodystopia-0.9.15.bin"
	tokyodystopia_cmd_bin := exec.Command("curl", "-L", tokyodystopia_bin_url, "-o", "binary.bin")
	err = tokyodystopia_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tokyodystopia_src_url := "https://dbmx.net/tokyodystopia/tokyodystopia-0.9.15.src.tar.gz"
	tokyodystopia_cmd_src := exec.Command("curl", "-L", tokyodystopia_src_url, "-o", "source.tar.gz")
	err = tokyodystopia_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tokyodystopia_cmd_direct := exec.Command("./binary")
	err = tokyodystopia_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tokyo-cabinet")
exec.Command("latte", "install", "tokyo-cabinet")
}
