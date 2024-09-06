package main

// Patchelf - Modify dynamic ELF executables
// Homepage: https://github.com/NixOS/patchelf

import (
	"fmt"
	
	"os/exec"
)

func installPatchelf() {
	// Método 1: Descargar y extraer .tar.gz
	patchelf_tar_url := "https://github.com/NixOS/patchelf/releases/download/0.18.0/patchelf-0.18.0.tar.bz2"
	patchelf_cmd_tar := exec.Command("curl", "-L", patchelf_tar_url, "-o", "package.tar.gz")
	err := patchelf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	patchelf_zip_url := "https://github.com/NixOS/patchelf/releases/download/0.18.0/patchelf-0.18.0.tar.bz2"
	patchelf_cmd_zip := exec.Command("curl", "-L", patchelf_zip_url, "-o", "package.zip")
	err = patchelf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	patchelf_bin_url := "https://github.com/NixOS/patchelf/releases/download/0.18.0/patchelf-0.18.0.tar.bz2"
	patchelf_cmd_bin := exec.Command("curl", "-L", patchelf_bin_url, "-o", "binary.bin")
	err = patchelf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	patchelf_src_url := "https://github.com/NixOS/patchelf/releases/download/0.18.0/patchelf-0.18.0.tar.bz2"
	patchelf_cmd_src := exec.Command("curl", "-L", patchelf_src_url, "-o", "source.tar.gz")
	err = patchelf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	patchelf_cmd_direct := exec.Command("./binary")
	err = patchelf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
