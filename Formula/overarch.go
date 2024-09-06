package main

// Overarch - Data driven description of software architecture
// Homepage: https://github.com/soulspace-org/overarch

import (
	"fmt"
	
	"os/exec"
)

func installOverarch() {
	// Método 1: Descargar y extraer .tar.gz
	overarch_tar_url := "https://github.com/soulspace-org/overarch/releases/download/v0.30.0/overarch.jar"
	overarch_cmd_tar := exec.Command("curl", "-L", overarch_tar_url, "-o", "package.tar.gz")
	err := overarch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	overarch_zip_url := "https://github.com/soulspace-org/overarch/releases/download/v0.30.0/overarch.jar"
	overarch_cmd_zip := exec.Command("curl", "-L", overarch_zip_url, "-o", "package.zip")
	err = overarch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	overarch_bin_url := "https://github.com/soulspace-org/overarch/releases/download/v0.30.0/overarch.jar"
	overarch_cmd_bin := exec.Command("curl", "-L", overarch_bin_url, "-o", "binary.bin")
	err = overarch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	overarch_src_url := "https://github.com/soulspace-org/overarch/releases/download/v0.30.0/overarch.jar"
	overarch_cmd_src := exec.Command("curl", "-L", overarch_src_url, "-o", "source.tar.gz")
	err = overarch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	overarch_cmd_direct := exec.Command("./binary")
	err = overarch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: leiningen")
	exec.Command("latte", "install", "leiningen").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
