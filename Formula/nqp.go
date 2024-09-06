package main

// Nqp - Lightweight Raku-like environment for virtual machines
// Homepage: https://github.com/Raku/nqp

import (
	"fmt"
	
	"os/exec"
)

func installNqp() {
	// Método 1: Descargar y extraer .tar.gz
	nqp_tar_url := "https://github.com/Raku/nqp/releases/download/2024.08/nqp-2024.08.tar.gz"
	nqp_cmd_tar := exec.Command("curl", "-L", nqp_tar_url, "-o", "package.tar.gz")
	err := nqp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nqp_zip_url := "https://github.com/Raku/nqp/releases/download/2024.08/nqp-2024.08.zip"
	nqp_cmd_zip := exec.Command("curl", "-L", nqp_zip_url, "-o", "package.zip")
	err = nqp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nqp_bin_url := "https://github.com/Raku/nqp/releases/download/2024.08/nqp-2024.08.bin"
	nqp_cmd_bin := exec.Command("curl", "-L", nqp_bin_url, "-o", "binary.bin")
	err = nqp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nqp_src_url := "https://github.com/Raku/nqp/releases/download/2024.08/nqp-2024.08.src.tar.gz"
	nqp_cmd_src := exec.Command("curl", "-L", nqp_src_url, "-o", "source.tar.gz")
	err = nqp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nqp_cmd_direct := exec.Command("./binary")
	err = nqp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtommath")
	exec.Command("latte", "install", "libtommath").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: moarvm")
	exec.Command("latte", "install", "moarvm").Run()
}
