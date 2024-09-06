package main

// Vcsh - Config manager based on git
// Homepage: https://github.com/RichiH/vcsh

import (
	"fmt"
	
	"os/exec"
)

func installVcsh() {
	// Método 1: Descargar y extraer .tar.gz
	vcsh_tar_url := "https://github.com/RichiH/vcsh/releases/download/v2.0.10/vcsh-2.0.10.tar.zst"
	vcsh_cmd_tar := exec.Command("curl", "-L", vcsh_tar_url, "-o", "package.tar.gz")
	err := vcsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vcsh_zip_url := "https://github.com/RichiH/vcsh/releases/download/v2.0.10/vcsh-2.0.10.tar.zst"
	vcsh_cmd_zip := exec.Command("curl", "-L", vcsh_zip_url, "-o", "package.zip")
	err = vcsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vcsh_bin_url := "https://github.com/RichiH/vcsh/releases/download/v2.0.10/vcsh-2.0.10.tar.zst"
	vcsh_cmd_bin := exec.Command("curl", "-L", vcsh_bin_url, "-o", "binary.bin")
	err = vcsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vcsh_src_url := "https://github.com/RichiH/vcsh/releases/download/v2.0.10/vcsh-2.0.10.tar.zst"
	vcsh_cmd_src := exec.Command("curl", "-L", vcsh_src_url, "-o", "source.tar.gz")
	err = vcsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vcsh_cmd_direct := exec.Command("./binary")
	err = vcsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
