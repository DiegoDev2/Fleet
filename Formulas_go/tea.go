package main

// Tea - Command-line tool to interact with Gitea servers
// Homepage: https://gitea.com/gitea/tea

import (
	"fmt"
	
	"os/exec"
)

func installTea() {
	// Método 1: Descargar y extraer .tar.gz
	tea_tar_url := "https://gitea.com/gitea/tea/archive/v0.9.2.tar.gz"
	tea_cmd_tar := exec.Command("curl", "-L", tea_tar_url, "-o", "package.tar.gz")
	err := tea_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tea_zip_url := "https://gitea.com/gitea/tea/archive/v0.9.2.zip"
	tea_cmd_zip := exec.Command("curl", "-L", tea_zip_url, "-o", "package.zip")
	err = tea_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tea_bin_url := "https://gitea.com/gitea/tea/archive/v0.9.2.bin"
	tea_cmd_bin := exec.Command("curl", "-L", tea_bin_url, "-o", "binary.bin")
	err = tea_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tea_src_url := "https://gitea.com/gitea/tea/archive/v0.9.2.src.tar.gz"
	tea_cmd_src := exec.Command("curl", "-L", tea_src_url, "-o", "source.tar.gz")
	err = tea_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tea_cmd_direct := exec.Command("./binary")
	err = tea_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
