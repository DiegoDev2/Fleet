package main

// Volt - Meta-level vim package manager
// Homepage: https://github.com/vim-volt/volt

import (
	"fmt"
	
	"os/exec"
)

func installVolt() {
	// Método 1: Descargar y extraer .tar.gz
	volt_tar_url := "https://github.com/vim-volt/volt.git"
	volt_cmd_tar := exec.Command("curl", "-L", volt_tar_url, "-o", "package.tar.gz")
	err := volt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	volt_zip_url := "https://github.com/vim-volt/volt.git"
	volt_cmd_zip := exec.Command("curl", "-L", volt_zip_url, "-o", "package.zip")
	err = volt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	volt_bin_url := "https://github.com/vim-volt/volt.git"
	volt_cmd_bin := exec.Command("curl", "-L", volt_bin_url, "-o", "binary.bin")
	err = volt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	volt_src_url := "https://github.com/vim-volt/volt.git"
	volt_cmd_src := exec.Command("curl", "-L", volt_src_url, "-o", "source.tar.gz")
	err = volt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	volt_cmd_direct := exec.Command("./binary")
	err = volt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
