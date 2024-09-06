package main

// Cgit - Hyperfast web frontend for Git repositories written in C
// Homepage: https://git.zx2c4.com/cgit/

import (
	"fmt"
	
	"os/exec"
)

func installCgit() {
	// Método 1: Descargar y extraer .tar.gz
	cgit_tar_url := "https://git.zx2c4.com/cgit/snapshot/cgit-1.2.3.tar.xz"
	cgit_cmd_tar := exec.Command("curl", "-L", cgit_tar_url, "-o", "package.tar.gz")
	err := cgit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgit_zip_url := "https://git.zx2c4.com/cgit/snapshot/cgit-1.2.3.tar.xz"
	cgit_cmd_zip := exec.Command("curl", "-L", cgit_zip_url, "-o", "package.zip")
	err = cgit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgit_bin_url := "https://git.zx2c4.com/cgit/snapshot/cgit-1.2.3.tar.xz"
	cgit_cmd_bin := exec.Command("curl", "-L", cgit_bin_url, "-o", "binary.bin")
	err = cgit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgit_src_url := "https://git.zx2c4.com/cgit/snapshot/cgit-1.2.3.tar.xz"
	cgit_cmd_src := exec.Command("curl", "-L", cgit_src_url, "-o", "source.tar.gz")
	err = cgit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgit_cmd_direct := exec.Command("./binary")
	err = cgit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
