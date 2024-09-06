package main

// Vcprompt - Provide version control info in shell prompts
// Homepage: https://hg.gerg.ca/vcprompt

import (
	"fmt"
	
	"os/exec"
)

func installVcprompt() {
	// Método 1: Descargar y extraer .tar.gz
	vcprompt_tar_url := "https://hg.gerg.ca/vcprompt/archive/1.2.1.tar.gz"
	vcprompt_cmd_tar := exec.Command("curl", "-L", vcprompt_tar_url, "-o", "package.tar.gz")
	err := vcprompt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vcprompt_zip_url := "https://hg.gerg.ca/vcprompt/archive/1.2.1.zip"
	vcprompt_cmd_zip := exec.Command("curl", "-L", vcprompt_zip_url, "-o", "package.zip")
	err = vcprompt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vcprompt_bin_url := "https://hg.gerg.ca/vcprompt/archive/1.2.1.bin"
	vcprompt_cmd_bin := exec.Command("curl", "-L", vcprompt_bin_url, "-o", "binary.bin")
	err = vcprompt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vcprompt_src_url := "https://hg.gerg.ca/vcprompt/archive/1.2.1.src.tar.gz"
	vcprompt_cmd_src := exec.Command("curl", "-L", vcprompt_src_url, "-o", "source.tar.gz")
	err = vcprompt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vcprompt_cmd_direct := exec.Command("./binary")
	err = vcprompt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
}
