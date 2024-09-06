package main

// Prqlc - Simple, powerful, pipelined SQL replacement
// Homepage: https://prql-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installPrqlc() {
	// Método 1: Descargar y extraer .tar.gz
	prqlc_tar_url := "https://github.com/PRQL/prql/archive/refs/tags/0.13.0.tar.gz"
	prqlc_cmd_tar := exec.Command("curl", "-L", prqlc_tar_url, "-o", "package.tar.gz")
	err := prqlc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prqlc_zip_url := "https://github.com/PRQL/prql/archive/refs/tags/0.13.0.zip"
	prqlc_cmd_zip := exec.Command("curl", "-L", prqlc_zip_url, "-o", "package.zip")
	err = prqlc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prqlc_bin_url := "https://github.com/PRQL/prql/archive/refs/tags/0.13.0.bin"
	prqlc_cmd_bin := exec.Command("curl", "-L", prqlc_bin_url, "-o", "binary.bin")
	err = prqlc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prqlc_src_url := "https://github.com/PRQL/prql/archive/refs/tags/0.13.0.src.tar.gz"
	prqlc_cmd_src := exec.Command("curl", "-L", prqlc_src_url, "-o", "source.tar.gz")
	err = prqlc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prqlc_cmd_direct := exec.Command("./binary")
	err = prqlc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
