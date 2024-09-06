package main

// Kubie - Much more powerful alternative to kubectx and kubens
// Homepage: https://blog.sbstp.ca/introducing-kubie/

import (
	"fmt"
	
	"os/exec"
)

func installKubie() {
	// Método 1: Descargar y extraer .tar.gz
	kubie_tar_url := "https://github.com/sbstp/kubie/archive/refs/tags/v0.23.1.tar.gz"
	kubie_cmd_tar := exec.Command("curl", "-L", kubie_tar_url, "-o", "package.tar.gz")
	err := kubie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubie_zip_url := "https://github.com/sbstp/kubie/archive/refs/tags/v0.23.1.zip"
	kubie_cmd_zip := exec.Command("curl", "-L", kubie_zip_url, "-o", "package.zip")
	err = kubie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubie_bin_url := "https://github.com/sbstp/kubie/archive/refs/tags/v0.23.1.bin"
	kubie_cmd_bin := exec.Command("curl", "-L", kubie_bin_url, "-o", "binary.bin")
	err = kubie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubie_src_url := "https://github.com/sbstp/kubie/archive/refs/tags/v0.23.1.src.tar.gz"
	kubie_cmd_src := exec.Command("curl", "-L", kubie_src_url, "-o", "source.tar.gz")
	err = kubie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubie_cmd_direct := exec.Command("./binary")
	err = kubie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: kubernetes-cli")
exec.Command("latte", "install", "kubernetes-cli")
}
