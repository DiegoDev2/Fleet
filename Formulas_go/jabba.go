package main

// Jabba - Cross-platform Java Version Manager
// Homepage: https://github.com/Jabba-Team/jabba

import (
	"fmt"
	
	"os/exec"
)

func installJabba() {
	// Método 1: Descargar y extraer .tar.gz
	jabba_tar_url := "https://github.com/Jabba-Team/jabba/archive/refs/tags/0.14.0.tar.gz"
	jabba_cmd_tar := exec.Command("curl", "-L", jabba_tar_url, "-o", "package.tar.gz")
	err := jabba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jabba_zip_url := "https://github.com/Jabba-Team/jabba/archive/refs/tags/0.14.0.zip"
	jabba_cmd_zip := exec.Command("curl", "-L", jabba_zip_url, "-o", "package.zip")
	err = jabba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jabba_bin_url := "https://github.com/Jabba-Team/jabba/archive/refs/tags/0.14.0.bin"
	jabba_cmd_bin := exec.Command("curl", "-L", jabba_bin_url, "-o", "binary.bin")
	err = jabba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jabba_src_url := "https://github.com/Jabba-Team/jabba/archive/refs/tags/0.14.0.src.tar.gz"
	jabba_cmd_src := exec.Command("curl", "-L", jabba_src_url, "-o", "source.tar.gz")
	err = jabba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jabba_cmd_direct := exec.Command("./binary")
	err = jabba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
