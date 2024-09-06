package main

// Hyperestraier - Full-text search system for communities
// Homepage: https://dbmx.net/hyperestraier/

import (
	"fmt"
	
	"os/exec"
)

func installHyperestraier() {
	// Método 1: Descargar y extraer .tar.gz
	hyperestraier_tar_url := "https://dbmx.net/hyperestraier/hyperestraier-1.4.13.tar.gz"
	hyperestraier_cmd_tar := exec.Command("curl", "-L", hyperestraier_tar_url, "-o", "package.tar.gz")
	err := hyperestraier_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hyperestraier_zip_url := "https://dbmx.net/hyperestraier/hyperestraier-1.4.13.zip"
	hyperestraier_cmd_zip := exec.Command("curl", "-L", hyperestraier_zip_url, "-o", "package.zip")
	err = hyperestraier_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hyperestraier_bin_url := "https://dbmx.net/hyperestraier/hyperestraier-1.4.13.bin"
	hyperestraier_cmd_bin := exec.Command("curl", "-L", hyperestraier_bin_url, "-o", "binary.bin")
	err = hyperestraier_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hyperestraier_src_url := "https://dbmx.net/hyperestraier/hyperestraier-1.4.13.src.tar.gz"
	hyperestraier_cmd_src := exec.Command("curl", "-L", hyperestraier_src_url, "-o", "source.tar.gz")
	err = hyperestraier_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hyperestraier_cmd_direct := exec.Command("./binary")
	err = hyperestraier_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: qdbm")
	exec.Command("latte", "install", "qdbm").Run()
}
