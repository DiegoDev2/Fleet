package main

// Smlfmt - Custom parser and code formatter for Standard ML
// Homepage: https://github.com/shwestrick/smlfmt

import (
	"fmt"
	
	"os/exec"
)

func installSmlfmt() {
	// Método 1: Descargar y extraer .tar.gz
	smlfmt_tar_url := "https://github.com/shwestrick/smlfmt/archive/refs/tags/v1.1.0.tar.gz"
	smlfmt_cmd_tar := exec.Command("curl", "-L", smlfmt_tar_url, "-o", "package.tar.gz")
	err := smlfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smlfmt_zip_url := "https://github.com/shwestrick/smlfmt/archive/refs/tags/v1.1.0.zip"
	smlfmt_cmd_zip := exec.Command("curl", "-L", smlfmt_zip_url, "-o", "package.zip")
	err = smlfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smlfmt_bin_url := "https://github.com/shwestrick/smlfmt/archive/refs/tags/v1.1.0.bin"
	smlfmt_cmd_bin := exec.Command("curl", "-L", smlfmt_bin_url, "-o", "binary.bin")
	err = smlfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smlfmt_src_url := "https://github.com/shwestrick/smlfmt/archive/refs/tags/v1.1.0.src.tar.gz"
	smlfmt_cmd_src := exec.Command("curl", "-L", smlfmt_src_url, "-o", "source.tar.gz")
	err = smlfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smlfmt_cmd_direct := exec.Command("./binary")
	err = smlfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mlton")
exec.Command("latte", "install", "mlton")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
