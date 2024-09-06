package main

// Httm - Interactive, file-level Time Machine-like tool for ZFS/btrfs
// Homepage: https://github.com/kimono-koans/httm

import (
	"fmt"
	
	"os/exec"
)

func installHttm() {
	// Método 1: Descargar y extraer .tar.gz
	httm_tar_url := "https://github.com/kimono-koans/httm/archive/refs/tags/0.42.4.tar.gz"
	httm_cmd_tar := exec.Command("curl", "-L", httm_tar_url, "-o", "package.tar.gz")
	err := httm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httm_zip_url := "https://github.com/kimono-koans/httm/archive/refs/tags/0.42.4.zip"
	httm_cmd_zip := exec.Command("curl", "-L", httm_zip_url, "-o", "package.zip")
	err = httm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httm_bin_url := "https://github.com/kimono-koans/httm/archive/refs/tags/0.42.4.bin"
	httm_cmd_bin := exec.Command("curl", "-L", httm_bin_url, "-o", "binary.bin")
	err = httm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httm_src_url := "https://github.com/kimono-koans/httm/archive/refs/tags/0.42.4.src.tar.gz"
	httm_cmd_src := exec.Command("curl", "-L", httm_src_url, "-o", "source.tar.gz")
	err = httm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httm_cmd_direct := exec.Command("./binary")
	err = httm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: acl")
	exec.Command("latte", "install", "acl").Run()
}
