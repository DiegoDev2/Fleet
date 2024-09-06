package main

// Paket - Dependency manager for .NET with support for NuGet and Git repositories
// Homepage: https://fsprojects.github.io/Paket/

import (
	"fmt"
	
	"os/exec"
)

func installPaket() {
	// Método 1: Descargar y extraer .tar.gz
	paket_tar_url := "https://github.com/fsprojects/Paket/releases/download/6.2.1/paket.exe"
	paket_cmd_tar := exec.Command("curl", "-L", paket_tar_url, "-o", "package.tar.gz")
	err := paket_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	paket_zip_url := "https://github.com/fsprojects/Paket/releases/download/6.2.1/paket.exe"
	paket_cmd_zip := exec.Command("curl", "-L", paket_zip_url, "-o", "package.zip")
	err = paket_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	paket_bin_url := "https://github.com/fsprojects/Paket/releases/download/6.2.1/paket.exe"
	paket_cmd_bin := exec.Command("curl", "-L", paket_bin_url, "-o", "binary.bin")
	err = paket_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	paket_src_url := "https://github.com/fsprojects/Paket/releases/download/6.2.1/paket.exe"
	paket_cmd_src := exec.Command("curl", "-L", paket_src_url, "-o", "source.tar.gz")
	err = paket_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	paket_cmd_direct := exec.Command("./binary")
	err = paket_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mono")
exec.Command("latte", "install", "mono")
}
