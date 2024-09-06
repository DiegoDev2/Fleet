package main

// Buku - Powerful command-line bookmark manager
// Homepage: https://github.com/jarun/buku

import (
	"fmt"
	
	"os/exec"
)

func installBuku() {
	// Método 1: Descargar y extraer .tar.gz
	buku_tar_url := "https://github.com/jarun/buku/archive/refs/tags/v4.9.tar.gz"
	buku_cmd_tar := exec.Command("curl", "-L", buku_tar_url, "-o", "package.tar.gz")
	err := buku_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	buku_zip_url := "https://github.com/jarun/buku/archive/refs/tags/v4.9.zip"
	buku_cmd_zip := exec.Command("curl", "-L", buku_zip_url, "-o", "package.zip")
	err = buku_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	buku_bin_url := "https://github.com/jarun/buku/archive/refs/tags/v4.9.bin"
	buku_cmd_bin := exec.Command("curl", "-L", buku_bin_url, "-o", "binary.bin")
	err = buku_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	buku_src_url := "https://github.com/jarun/buku/archive/refs/tags/v4.9.src.tar.gz"
	buku_cmd_src := exec.Command("curl", "-L", buku_src_url, "-o", "source.tar.gz")
	err = buku_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	buku_cmd_direct := exec.Command("./binary")
	err = buku_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
