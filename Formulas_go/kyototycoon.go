package main

// KyotoTycoon - Database server with interface to Kyoto Cabinet
// Homepage: https://dbmx.net/kyototycoon/

import (
	"fmt"
	
	"os/exec"
)

func installKyotoTycoon() {
	// Método 1: Descargar y extraer .tar.gz
	kyototycoon_tar_url := "https://dbmx.net/kyototycoon/pkg/kyototycoon-0.9.56.tar.gz"
	kyototycoon_cmd_tar := exec.Command("curl", "-L", kyototycoon_tar_url, "-o", "package.tar.gz")
	err := kyototycoon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kyototycoon_zip_url := "https://dbmx.net/kyototycoon/pkg/kyototycoon-0.9.56.zip"
	kyototycoon_cmd_zip := exec.Command("curl", "-L", kyototycoon_zip_url, "-o", "package.zip")
	err = kyototycoon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kyototycoon_bin_url := "https://dbmx.net/kyototycoon/pkg/kyototycoon-0.9.56.bin"
	kyototycoon_cmd_bin := exec.Command("curl", "-L", kyototycoon_bin_url, "-o", "binary.bin")
	err = kyototycoon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kyototycoon_src_url := "https://dbmx.net/kyototycoon/pkg/kyototycoon-0.9.56.src.tar.gz"
	kyototycoon_cmd_src := exec.Command("curl", "-L", kyototycoon_src_url, "-o", "source.tar.gz")
	err = kyototycoon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kyototycoon_cmd_direct := exec.Command("./binary")
	err = kyototycoon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: kyoto-cabinet")
exec.Command("latte", "install", "kyoto-cabinet")
}
