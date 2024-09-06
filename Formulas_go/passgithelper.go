package main

// PassGitHelper - Git credential helper interfacing with pass
// Homepage: https://github.com/languitar/pass-git-helper

import (
	"fmt"
	
	"os/exec"
)

func installPassGitHelper() {
	// Método 1: Descargar y extraer .tar.gz
	passgithelper_tar_url := "https://github.com/languitar/pass-git-helper/archive/refs/tags/v2.0.0.tar.gz"
	passgithelper_cmd_tar := exec.Command("curl", "-L", passgithelper_tar_url, "-o", "package.tar.gz")
	err := passgithelper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	passgithelper_zip_url := "https://github.com/languitar/pass-git-helper/archive/refs/tags/v2.0.0.zip"
	passgithelper_cmd_zip := exec.Command("curl", "-L", passgithelper_zip_url, "-o", "package.zip")
	err = passgithelper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	passgithelper_bin_url := "https://github.com/languitar/pass-git-helper/archive/refs/tags/v2.0.0.bin"
	passgithelper_cmd_bin := exec.Command("curl", "-L", passgithelper_bin_url, "-o", "binary.bin")
	err = passgithelper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	passgithelper_src_url := "https://github.com/languitar/pass-git-helper/archive/refs/tags/v2.0.0.src.tar.gz"
	passgithelper_cmd_src := exec.Command("curl", "-L", passgithelper_src_url, "-o", "source.tar.gz")
	err = passgithelper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	passgithelper_cmd_direct := exec.Command("./binary")
	err = passgithelper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnupg")
exec.Command("latte", "install", "gnupg")
	fmt.Println("Instalando dependencia: pass")
exec.Command("latte", "install", "pass")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
