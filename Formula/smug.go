package main

// Smug - Automate your tmux workflow
// Homepage: https://github.com/ivaaaan/smug

import (
	"fmt"
	
	"os/exec"
)

func installSmug() {
	// Método 1: Descargar y extraer .tar.gz
	smug_tar_url := "https://github.com/ivaaaan/smug/archive/refs/tags/v0.3.4.tar.gz"
	smug_cmd_tar := exec.Command("curl", "-L", smug_tar_url, "-o", "package.tar.gz")
	err := smug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smug_zip_url := "https://github.com/ivaaaan/smug/archive/refs/tags/v0.3.4.zip"
	smug_cmd_zip := exec.Command("curl", "-L", smug_zip_url, "-o", "package.zip")
	err = smug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smug_bin_url := "https://github.com/ivaaaan/smug/archive/refs/tags/v0.3.4.bin"
	smug_cmd_bin := exec.Command("curl", "-L", smug_bin_url, "-o", "binary.bin")
	err = smug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smug_src_url := "https://github.com/ivaaaan/smug/archive/refs/tags/v0.3.4.src.tar.gz"
	smug_cmd_src := exec.Command("curl", "-L", smug_src_url, "-o", "source.tar.gz")
	err = smug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smug_cmd_direct := exec.Command("./binary")
	err = smug_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: tmux")
	exec.Command("latte", "install", "tmux").Run()
}
