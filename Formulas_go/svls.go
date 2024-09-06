package main

// Svls - SystemVerilog language server
// Homepage: https://github.com/dalance/svls

import (
	"fmt"
	
	"os/exec"
)

func installSvls() {
	// Método 1: Descargar y extraer .tar.gz
	svls_tar_url := "https://github.com/dalance/svls/archive/refs/tags/v0.2.12.tar.gz"
	svls_cmd_tar := exec.Command("curl", "-L", svls_tar_url, "-o", "package.tar.gz")
	err := svls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	svls_zip_url := "https://github.com/dalance/svls/archive/refs/tags/v0.2.12.zip"
	svls_cmd_zip := exec.Command("curl", "-L", svls_zip_url, "-o", "package.zip")
	err = svls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	svls_bin_url := "https://github.com/dalance/svls/archive/refs/tags/v0.2.12.bin"
	svls_cmd_bin := exec.Command("curl", "-L", svls_bin_url, "-o", "binary.bin")
	err = svls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	svls_src_url := "https://github.com/dalance/svls/archive/refs/tags/v0.2.12.src.tar.gz"
	svls_cmd_src := exec.Command("curl", "-L", svls_src_url, "-o", "source.tar.gz")
	err = svls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	svls_cmd_direct := exec.Command("./binary")
	err = svls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
